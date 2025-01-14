// Package store provides a distributed SQLite instance.
//
// Distributed consensus is provided via the Raft algorithm.
package store

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/binary"
	"expvar"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/hashicorp/raft"
	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	"github.com/pastelnetwork/gonode/metadb/rqlite/command"
	sql "github.com/pastelnetwork/gonode/metadb/rqlite/db"
	rlog "github.com/pastelnetwork/gonode/metadb/rqlite/log"
)

const (
	raftDBPath          = "raft.db" // Changing this will break backwards compatibility.
	retainSnapshotCount = 2
	applyTimeout        = 10 * time.Second
	openTimeout         = 120 * time.Second
	sqliteFile          = "db.sqlite"
	leaderWaitDelay     = 100 * time.Millisecond
	appliedWaitDelay    = 100 * time.Millisecond
	connectionPoolCount = 5
	connectionTimeout   = 10 * time.Second
	raftLogCacheSize    = 512
	trailingScale       = 1.25
)

const (
	numSnaphots             = "num_snapshots"
	numBackups              = "num_backups"
	numRestores             = "num_restores"
	numUncompressedCommands = "num_uncompressed_commands"
	numCompressedCommands   = "num_compressed_commands"
)

// BackupFormat represents the format of database backup.
type BackupFormat int

const (
	// BackupSQL is the plaintext SQL command format.
	BackupSQL BackupFormat = iota

	// BackupBinary is a SQLite file backup format.
	BackupBinary
)

// stats captures stats for the Store.
var stats *expvar.Map

func init() {
	stats = expvar.NewMap("store")
	stats.Add(numSnaphots, 0)
	stats.Add(numBackups, 0)
	stats.Add(numRestores, 0)
	stats.Add(numUncompressedCommands, 0)
	stats.Add(numCompressedCommands, 0)
}

// ClusterState defines the possible Raft states the current node can be in
type ClusterState int

// Represents the Raft cluster states
const (
	Leader ClusterState = iota
	Follower
	Candidate
	Shutdown
	Unknown
)

// Store is a SQLite database, where all changes are made via Raft consensus.
type Store struct {
	raftDir string

	raft   *raft.Raft // The consensus mechanism.
	ln     Listener
	raftTn *raft.NetworkTransport
	raftID string    // Node ID.
	dbConf *DBConfig // SQLite database config.
	dbPath string    // Path to underlying SQLite file, if not in-memory.
	db     *sql.DB   // The underlying SQLite store.

	reqMarshaller *command.RequestMarshaler // Request marshaler for writing to log.
	raftLog       raft.LogStore             // Persistent log store.
	raftStable    raft.StableStore          // Persistent k-v store.
	boltStore     *rlog.Log                 // Physical store.

	onDiskCreated        bool      // On disk database actually created?
	snapsExistOnOpen     bool      // Any snaps present when store opens?
	firstIdxOnOpen       uint64    // First index on log when Store opens.
	lastIdxOnOpen        uint64    // Last index on log when Store opens.
	lastCommandIdxOnOpen uint64    // Last command index on log when Store opens.
	firstLogAppliedT     time.Time // Time first log is applied
	appliedOnOpen        uint64    // Number of logs applied at open.
	openT                time.Time // Timestamp when Store opens.

	numNoops int // For whitebox testing

	txMu    sync.RWMutex // Sync between snapshots and query-level transactions.
	queryMu sync.RWMutex // Sync queries generally with other operations.

	ShutdownOnRemove   bool
	SnapshotThreshold  uint64
	SnapshotInterval   time.Duration
	LeaderLeaseTimeout time.Duration
	HeartbeatTimeout   time.Duration
	ElectionTimeout    time.Duration
	ApplyTimeout       time.Duration

	numTrailingLogs uint64

	ctx context.Context
}

// IsNewNode returns whether a node using raftDir would be a brand new node.
// It also means that the window this node joining a different cluster has passed.
func IsNewNode(raftDir string) bool {
	// If there is any pre-existing Raft state, then this node
	// has already been created.
	return !pathExists(filepath.Join(raftDir, raftDBPath))
}

// Config represents the configuration of the underlying Store.
type Config struct {
	DBConf *DBConfig // The DBConfig object for this Store.
	Dir    string    // The working directory for raft.
	Tn     Transport // The underlying Transport for raft.
	ID     string    // Node ID.
}

// New returns a new Store.
func New(ctx context.Context, ln Listener, c *Config) *Store {
	return &Store{
		ctx:           ctx,
		ln:            ln,
		raftDir:       c.Dir,
		raftID:        c.ID,
		dbConf:        c.DBConf,
		dbPath:        filepath.Join(c.Dir, sqliteFile),
		reqMarshaller: command.NewRequestMarshaler(),
		ApplyTimeout:  applyTimeout,
	}
}

// Open opens the Store. If enableBootstrap is set, then this node becomes a
// standalone node. If not set, then the calling layer must know that this
// node has pre-existing state, or the calling layer will trigger a join
// operation after opening the Store.
func (s *Store) Open(enableBootstrap bool) error {
	s.openT = time.Now()
	log.WithContext(s.ctx).Debugf("opening store with node ID %s", s.raftID)

	log.WithContext(s.ctx).Debugf("ensuring directory at %s exists", s.raftDir)
	err := os.MkdirAll(s.raftDir, 0755)
	if err != nil {
		return err
	}

	// Create Raft-compatible network layer.
	s.raftTn = raft.NewNetworkTransport(NewTransport(s.ln), connectionPoolCount, connectionTimeout, nil)

	// Don't allow control over trailing logs directly, just implement a policy.
	s.numTrailingLogs = uint64(float64(s.SnapshotThreshold) * trailingScale)

	config := s.raftConfig()
	config.Logger = NewLogger(s.ctx)
	config.LocalID = raft.ServerID(s.raftID)

	// Create the snapshot store. This allows Raft to truncate the log.
	snapshots, err := raft.NewFileSnapshotStore(s.raftDir, retainSnapshotCount, os.Stderr)
	if err != nil {
		return fmt.Errorf("file snapshot store: %s", err)
	}
	snaps, err := snapshots.List()
	if err != nil {
		return fmt.Errorf("list snapshots: %s", err)
	}
	log.WithContext(s.ctx).Debugf("%d preexisting snapshots present", len(snaps))
	s.snapsExistOnOpen = len(snaps) > 0

	// Create the log store and stable store.
	s.boltStore, err = rlog.NewLog(filepath.Join(s.raftDir, raftDBPath))
	if err != nil {
		return fmt.Errorf("new log store: %s", err)
	}
	s.raftStable = s.boltStore
	s.raftLog, err = raft.NewLogCache(raftLogCacheSize, s.boltStore)
	if err != nil {
		return fmt.Errorf("new cached store: %s", err)
	}

	// Get some info about the log, before any more entries are committed.
	if err := s.setLogInfo(); err != nil {
		return fmt.Errorf("set log info: %s", err)
	}
	log.WithContext(s.ctx).Debugf("first log index: %d, last log index: %d, last command log index: %d:",
		s.firstIdxOnOpen, s.lastIdxOnOpen, s.lastCommandIdxOnOpen)

	// If an on-disk database has been requested, and there are no snapshots, and
	// there are no commands in the log, then this is the only opportunity to
	// create that on-disk database file before Raft initializes.
	if !s.dbConf.Memory && !s.snapsExistOnOpen && s.lastCommandIdxOnOpen == 0 {
		s.db, err = s.openOnDisk(nil)
		if err != nil {
			return fmt.Errorf("failed to open on-disk database")
		}
		s.onDiskCreated = true
	} else {
		// We need an in-memory database, at least for bootstrapping purposes.
		s.db, err = s.openInMemory(nil)
		if err != nil {
			return fmt.Errorf("failed to open in-memory database")
		}
	}

	// Instantiate the Raft system.
	ra, err := raft.NewRaft(config, s, s.raftLog, s.raftStable, snapshots, s.raftTn)
	if err != nil {
		return fmt.Errorf("new raft: %s", err)
	}

	if enableBootstrap {
		log.WithContext(s.ctx).Debugf("executing new cluster bootstrap")
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: s.raftTn.LocalAddr(),
				},
			},
		}
		ra.BootstrapCluster(configuration)
	} else {
		log.WithContext(s.ctx).Debugf("no cluster bootstrap requested")
	}

	s.raft = ra

	return nil
}

// Close closes the store. If wait is true, waits for a graceful shutdown.
func (s *Store) Close(wait bool) error {
	f := s.raft.Shutdown()
	if wait {
		if f.Error() != nil {
			return f.Error()
		}
	}
	// Only shutdown Bolt and SQLite when Raft is done.
	if err := s.db.Close(); err != nil {
		return err
	}
	return s.boltStore.Close()
}

// WaitForApplied waits for all Raft log entries to to be applied to the
// underlying database.
func (s *Store) WaitForApplied(ctx context.Context, timeout time.Duration) error {
	if timeout == 0 {
		return nil
	}
	log.WithContext(ctx).Debugf("waiting for up to %s for application of initial logs", timeout)
	if err := s.WaitForAppliedIndex(ctx, s.raft.LastIndex(), timeout); err != nil {
		return errors.New(ErrOpenTimeout)
	}
	return nil
}

// IsLeader is used to determine if the current node is cluster leader
func (s *Store) IsLeader() bool {
	return s.raft.State() == raft.Leader
}

// State returns the current node's Raft state
func (s *Store) State() ClusterState {
	state := s.raft.State()
	switch state {
	case raft.Leader:
		return Leader
	case raft.Candidate:
		return Candidate
	case raft.Follower:
		return Follower
	case raft.Shutdown:
		return Shutdown
	default:
		return Unknown
	}
}

// Path returns the path to the store's storage directory.
func (s *Store) Path() string {
	return s.raftDir
}

// Addr returns the address of the store.
func (s *Store) Addr() string {
	return string(s.raftTn.LocalAddr())
}

// ID returns the Raft ID of the store.
func (s *Store) ID() string {
	return s.raftID
}

// LeaderAddr returns the address of the current leader. Returns a
// blank string if there is no leader.
func (s *Store) LeaderAddr() (string, error) {
	return string(s.raft.Leader()), nil
}

// LeaderID returns the node ID of the Raft leader. Returns a
// blank string if there is no leader, or an error.
func (s *Store) LeaderID() (string, error) {
	addr, err := s.LeaderAddr()
	if err != nil {
		return "", nil
	}
	configFuture := s.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		log.WithContext(s.ctx).WithError(err).Errorf("failed to get raft configuration")
		return "", err
	}

	for _, srv := range configFuture.Configuration().Servers {
		if srv.Address == raft.ServerAddress(addr) {
			return string(srv.ID), nil
		}
	}
	return "", nil
}

// Nodes returns the slice of nodes in the cluster, sorted by ID ascending.
func (s *Store) Nodes() ([]*Server, error) {
	f := s.raft.GetConfiguration()
	if f.Error() != nil {
		return nil, f.Error()
	}

	rs := f.Configuration().Servers
	servers := make([]*Server, len(rs))
	for i := range rs {
		servers[i] = &Server{
			ID:       string(rs[i].ID),
			Addr:     string(rs[i].Address),
			Suffrage: rs[i].Suffrage.String(),
		}
	}

	sort.Sort(Servers(servers))
	return servers, nil
}

// WaitForLeader blocks until a leader is detected, or the timeout expires.
func (s *Store) WaitForLeader(ctx context.Context, timeout time.Duration) (string, error) {
	tck := time.NewTicker(leaderWaitDelay)
	defer tck.Stop()
	tmr := time.NewTimer(timeout)
	defer tmr.Stop()

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-tck.C:
			l, err := s.LeaderAddr()
			if err != nil {
				return "", nil
			}
			if l != "" {
				return l, nil
			}
		case <-tmr.C:
			return "", fmt.Errorf("timeout expired")
		}
	}
}

// SetRequestCompression allows low-level control over the compression threshold
// for the request marshaler.
func (s *Store) SetRequestCompression(batch, size int) {
	s.reqMarshaller.BatchThreshold = batch
	s.reqMarshaller.SizeThreshold = size
}

// WaitForAppliedIndex blocks until a given log index has been applied,
// or the timeout expires.
func (s *Store) WaitForAppliedIndex(ctx context.Context, idx uint64, timeout time.Duration) error {
	tck := time.NewTicker(appliedWaitDelay)
	defer tck.Stop()
	tmr := time.NewTimer(timeout)
	defer tmr.Stop()

	for {
		select {
		case <-ctx.Done():
			return s.ctx.Err()
		case <-tck.C:
			if s.raft.AppliedIndex() >= idx {
				return nil
			}
		case <-tmr.C:
			return fmt.Errorf("timeout expired")
		}
	}
}

// Stats returns stats for the store.
func (s *Store) Stats() (map[string]interface{}, error) {
	fkEnabled, err := s.db.FKConstraints()
	if err != nil {
		return nil, err
	}

	dbSz, err := s.db.Size()
	if err != nil {
		return nil, err
	}
	dbStatus := map[string]interface{}{
		"dsn":            s.dbConf.DSN,
		"fk_constraints": enabledFromBool(fkEnabled),
		"version":        sql.DBVersion,
		"db_size":        dbSz,
	}
	if s.dbConf.Memory {
		dbStatus["path"] = ":memory:"
	} else {
		dbStatus["path"] = s.dbPath
		if s.onDiskCreated {
			if dbStatus["size"], err = s.db.FileSize(); err != nil {
				return nil, err
			}
		}
	}

	nodes, err := s.Nodes()
	if err != nil {
		return nil, err
	}
	leaderID, err := s.LeaderID()
	if err != nil {
		return nil, err
	}

	// Perform type-conversion to actual numbers where possible.
	raftStats := make(map[string]interface{})
	for k, v := range s.raft.Stats() {
		if s, err := strconv.ParseInt(v, 10, 64); err != nil {
			raftStats[k] = v
		} else {
			raftStats[k] = s
		}
	}
	raftStats["log_size"], err = s.logSize()
	if err != nil {
		return nil, err
	}

	dirSz, err := dirSize(s.raftDir)
	if err != nil {
		return nil, err
	}

	leaderAddr, err := s.LeaderAddr()
	if err != nil {
		return nil, err
	}
	status := map[string]interface{}{
		"node_id": s.raftID,
		"raft":    raftStats,
		"addr":    s.Addr(),
		"leader": map[string]string{
			"node_id": leaderID,
			"addr":    leaderAddr,
		},
		"apply_timeout":      s.ApplyTimeout.String(),
		"heartbeat_timeout":  s.HeartbeatTimeout.String(),
		"election_timeout":   s.ElectionTimeout.String(),
		"snapshot_threshold": s.SnapshotThreshold,
		"snapshot_interval":  s.SnapshotInterval,
		"trailing_logs":      s.numTrailingLogs,
		"request_marshaler":  s.reqMarshaller.Stats(),
		"nodes":              nodes,
		"dir":                s.raftDir,
		"dir_size":           dirSz,
		"sqlite3":            dbStatus,
		"db_conf":            s.dbConf,
	}
	return status, nil
}

// Execute executes queries that return no rows, but do modify the database.
func (s *Store) Execute(ex *command.ExecuteRequest) ([]*sql.Result, error) {
	if s.raft.State() != raft.Leader {
		return nil, errors.New(ErrNotLeader)
	}
	return s.execute(ex)
}

// ExecuteOrAbort executes the requests, but aborts any active transaction
// on the underlying database in the case of any error.
func (s *Store) ExecuteOrAbort(ex *command.ExecuteRequest) (results []*sql.Result, retErr error) {
	defer func() {
		var errored bool
		for i := range results {
			if results[i].Error != "" {
				errored = true
				break
			}
		}
		if retErr != nil || errored {
			if err := s.db.AbortTransaction(); err != nil {
				log.WithContext(s.ctx).WithError(err).Error("failed to abort transaction")
			}
		}
	}()
	return s.execute(ex)
}

func (s *Store) execute(ex *command.ExecuteRequest) ([]*sql.Result, error) {
	b, compressed, err := s.reqMarshaller.Marshal(ex)
	if err != nil {
		return nil, err
	}
	if compressed {
		stats.Add(numCompressedCommands, 1)
	} else {
		stats.Add(numUncompressedCommands, 1)
	}

	c := &command.Command{
		Type:       command.Command_COMMAND_TYPE_EXECUTE,
		SubCommand: b,
		Compressed: compressed,
	}

	b, err = command.Marshal(c)
	if err != nil {
		return nil, err
	}

	f := s.raft.Apply(b, s.ApplyTimeout)
	if e := f.(raft.Future); e.Error() != nil {
		if e.Error() == raft.ErrNotLeader {
			return nil, errors.New(ErrNotLeader)
		}
		return nil, e.Error()
	}

	r := f.Response().(*fsmExecuteResponse)
	return r.results, r.error
}

// Backup writes a snapshot of the underlying database to dst
//
// If leader is true, this operation is performed with a read consistency
// level equivalent to "weak". Otherwise no guarantees are made about the
// read consistency level.
func (s *Store) Backup(leader bool, fmt BackupFormat, dst io.Writer) error {
	if leader && s.raft.State() != raft.Leader {
		return errors.New(ErrNotLeader)
	}

	if fmt == BackupBinary {
		if err := s.database(leader, dst); err != nil {
			return err
		}
	} else if fmt == BackupSQL {
		if err := s.db.Dump(dst); err != nil {
			return err
		}
	} else {
		return errors.New(ErrInvalidBackupFormat)
	}

	stats.Add(numBackups, 1)
	return nil
}

// Query executes queries that return rows, and do not modify the database.
func (s *Store) Query(qr *command.QueryRequest) ([]*sql.Rows, error) {
	s.queryMu.RLock()
	defer s.queryMu.RUnlock()

	if qr.Level == command.QueryRequest_QUERY_REQUEST_LEVEL_STRONG {
		b, compressed, err := s.reqMarshaller.Marshal(qr)
		if err != nil {
			return nil, err
		}
		if compressed {
			stats.Add(numCompressedCommands, 1)
		} else {
			stats.Add(numUncompressedCommands, 1)
		}

		c := &command.Command{
			Type:       command.Command_COMMAND_TYPE_QUERY,
			SubCommand: b,
			Compressed: compressed,
		}

		b, err = command.Marshal(c)
		if err != nil {
			return nil, err
		}

		f := s.raft.Apply(b, s.ApplyTimeout)
		if e := f.(raft.Future); e.Error() != nil {
			if e.Error() == raft.ErrNotLeader {
				return nil, errors.New(ErrNotLeader)
			}
			return nil, e.Error()
		}

		r := f.Response().(*fsmQueryResponse)
		return r.rows, r.error
	}

	if qr.Level == command.QueryRequest_QUERY_REQUEST_LEVEL_WEAK && s.raft.State() != raft.Leader {
		return nil, errors.New(ErrNotLeader)
	}

	if qr.Level == command.QueryRequest_QUERY_REQUEST_LEVEL_NONE && qr.Freshness > 0 &&
		time.Since(s.raft.LastContact()).Nanoseconds() > qr.Freshness {
		return nil, errors.New(ErrStaleRead)
	}

	// Read straight from database. If a transaction is requested, we must block
	// certain other database operations.
	if qr.Request.Transaction {
		s.txMu.Lock()
		defer s.txMu.Unlock()
	}
	return s.db.Query(qr.Request, qr.Timings)
}

// Join joins a node, identified by id and located at addr, to this store.
// The node must be ready to respond to Raft communications at that address.
func (s *Store) Join(id, addr string, voter bool) error {
	log.WithContext(s.ctx).Debugf("received request to join node at %s", addr)
	if s.raft.State() != raft.Leader {
		return errors.New(ErrNotLeader)
	}

	configFuture := s.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		log.WithContext(s.ctx).WithError(err).Error("failed to get raft configuration")
		return err
	}

	for _, srv := range configFuture.Configuration().Servers {
		// If a node already exists with either the joining node's ID or address,
		// that node may need to be removed from the config first.
		if srv.ID == raft.ServerID(id) || srv.Address == raft.ServerAddress(addr) {
			// However if *both* the ID and the address are the same, the no
			// join is actually needed.
			if srv.Address == raft.ServerAddress(addr) && srv.ID == raft.ServerID(id) {
				log.WithContext(s.ctx).Debugf("node %s at %s already member of cluster, ignoring join request", id, addr)
				return nil
			}

			if err := s.remove(id); err != nil {
				log.WithContext(s.ctx).WithError(err).Error("failed to remove node")
				return err
			}
		}
	}

	var f raft.IndexFuture
	if voter {
		f = s.raft.AddVoter(raft.ServerID(id), raft.ServerAddress(addr), 0, 0)
	} else {

		f = s.raft.AddNonvoter(raft.ServerID(id), raft.ServerAddress(addr), 0, 0)
	}
	if e := f.(raft.Future); e.Error() != nil {
		if e.Error() == raft.ErrNotLeader {
			return errors.New(ErrNotLeader)
		}
		return e.Error()
	}

	log.WithContext(s.ctx).Debugf("node at %s joined successfully as %s", addr, prettyVoter(voter))
	return nil
}

// Remove removes a node from the store, specified by ID.
func (s *Store) Remove(id string) error {
	log.WithContext(s.ctx).Debugf("received request to remove node %s", id)
	if err := s.remove(id); err != nil {
		log.WithContext(s.ctx).WithError(err).Errorf("failed to remove node %s", id)
		return err
	}

	log.WithContext(s.ctx).Debugf("node %s removed successfully", id)
	return nil
}

// Noop writes a noop command to the Raft log. A noop command simply
// consumes a slot in the Raft log, but has no other affect on the
// system.
func (s *Store) Noop(id string) error {
	n := &command.Noop{
		Id: id,
	}
	b, err := command.MarshalNoop(n)
	if err != nil {
		return err
	}

	c := &command.Command{
		Type:       command.Command_COMMAND_TYPE_NOOP,
		SubCommand: b,
	}
	bc, err := command.Marshal(c)
	if err != nil {
		return err
	}

	f := s.raft.Apply(bc, s.ApplyTimeout)
	if e := f.(raft.Future); e.Error() != nil {
		if e.Error() == raft.ErrNotLeader {
			return errors.New(ErrNotLeader)
		}
		return e.Error()
	}
	return nil
}

// openInMemory returns an in-memory database. If b is non-nil, then the
// database will be initialized with the contents of b.
func (s *Store) openInMemory(b []byte) (db *sql.DB, err error) {
	if b == nil {
		db, err = sql.OpenInMemoryWithDSN(s.dbConf.DSN)
	} else {
		db, err = sql.DeserializeInMemoryWithDSN(b, s.dbConf.DSN)
	}
	return
}

// openOnDisk opens an on-disk database file at the Store's configured path. If
// b is non-nil, any preexisting file will first be overwritten with those contents.
// Otherwise any pre-existing file will be removed before the database is opened.
func (s *Store) openOnDisk(b []byte) (*sql.DB, error) {
	if err := os.Remove(s.dbPath); err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if b != nil {
		if err := ioutil.WriteFile(s.dbPath, b, 0660); err != nil {
			return nil, err
		}
	}
	return sql.OpenWithDSN(s.dbPath, s.dbConf.DSN)
}

// setLogInfo records some key indexs about the log.
func (s *Store) setLogInfo() error {
	var err error
	s.firstIdxOnOpen, err = s.boltStore.FirstIndex()
	if err != nil {
		return fmt.Errorf("failed to get last index: %s", err)
	}
	s.lastIdxOnOpen, err = s.boltStore.LastIndex()
	if err != nil {
		return fmt.Errorf("failed to get last index: %s", err)
	}
	s.lastCommandIdxOnOpen, err = s.boltStore.LastCommandIndex()
	if err != nil {
		return fmt.Errorf("failed to get last command index: %s", err)
	}
	return nil
}

// remove removes the node, with the given ID, from the cluster.
func (s *Store) remove(id string) error {
	if s.raft.State() != raft.Leader {
		return errors.New(ErrNotLeader)
	}

	f := s.raft.RemoveServer(raft.ServerID(id), 0, 0)
	if f.Error() != nil {
		if f.Error() == raft.ErrNotLeader {
			return errors.New(ErrNotLeader)
		}
		return f.Error()
	}

	return nil
}

// raftConfig returns a new Raft config for the store.
func (s *Store) raftConfig() *raft.Config {
	config := raft.DefaultConfig()
	config.ShutdownOnRemove = s.ShutdownOnRemove
	if s.SnapshotThreshold != 0 {
		config.SnapshotThreshold = s.SnapshotThreshold
		config.TrailingLogs = s.numTrailingLogs
	}
	if s.SnapshotInterval != 0 {
		config.SnapshotInterval = s.SnapshotInterval
	}
	if s.LeaderLeaseTimeout != 0 {
		config.LeaderLeaseTimeout = s.LeaderLeaseTimeout
	}
	if s.HeartbeatTimeout != 0 {
		config.HeartbeatTimeout = s.HeartbeatTimeout
	}
	if s.ElectionTimeout != 0 {
		config.ElectionTimeout = s.ElectionTimeout
	}
	return config
}

type fsmExecuteResponse struct {
	results []*sql.Result
	error   error
}

type fsmQueryResponse struct {
	rows  []*sql.Rows
	error error
}

type fsmGenericResponse struct {
	error error
}

// Apply applies a Raft log entry to the database.
func (s *Store) Apply(l *raft.Log) (e interface{}) {
	defer func() {
		if l.Index <= s.lastCommandIdxOnOpen {
			// In here means at least one command entry was in the log when the Store
			// opened.
			s.appliedOnOpen++
			if l.Index == s.lastCommandIdxOnOpen {
				log.WithContext(s.ctx).Debugf("%d committed log entries applied in %s, took %s since open",
					s.appliedOnOpen, time.Since(s.firstLogAppliedT), time.Since(s.openT))

				// Last command log applied. Time to switch to on-disk database?
				if s.dbConf.Memory {
					log.WithContext(s.ctx).Debug("continuing use of in-memory database")
				} else {
					// Since we're here, it means that a) an on-disk database was requested
					// *and* there were commands in the log. A snapshot may or may not have
					// been applied, but it wouldn't have created the on-disk database in that
					// case since there were commands in the log. This is the very last chance
					// to do convert from in-memory to on-disk.
					b, _ := s.db.Serialize()
					err := s.db.Close()
					if err != nil {
						e = &fsmGenericResponse{error: fmt.Errorf("close failed: %s", err)}
						return
					}
					// Open a new on-disk database.
					s.db, err = s.openOnDisk(b)
					if err != nil {
						e = &fsmGenericResponse{error: fmt.Errorf("open on-disk failed: %s", err)}
						return
					}
					s.onDiskCreated = true
					log.WithContext(s.ctx).Debug("successfully switched to on-disk database")
				}
			}
		}
	}()

	if s.firstLogAppliedT.IsZero() {
		s.firstLogAppliedT = time.Now()
	}

	var c command.Command

	if err := command.Unmarshal(l.Data, &c); err != nil {
		panic(fmt.Sprintf("failed to unmarshal cluster command: %s", err.Error()))
	}

	switch c.Type {
	case command.Command_COMMAND_TYPE_QUERY:
		var qr command.QueryRequest
		if err := command.UnmarshalSubCommand(&c, &qr); err != nil {
			panic(fmt.Sprintf("failed to unmarshal query subcommand: %s", err.Error()))
		}
		// Read from database. If a transaction is requested, we must block
		// certain other database operations.
		if qr.Request.Transaction {
			s.txMu.Lock()
			defer s.txMu.Unlock()
		}
		r, err := s.db.Query(qr.Request, qr.Timings)
		return &fsmQueryResponse{rows: r, error: err}
	case command.Command_COMMAND_TYPE_EXECUTE:
		var er command.ExecuteRequest
		if err := command.UnmarshalSubCommand(&c, &er); err != nil {
			panic(fmt.Sprintf("failed to unmarshal execute subcommand: %s", err.Error()))
		}
		r, err := s.db.Execute(er.Request, er.Timings)
		return &fsmExecuteResponse{results: r, error: err}
	case command.Command_COMMAND_TYPE_NOOP:
		s.numNoops++
		return &fsmGenericResponse{}
	default:
		return &fsmGenericResponse{error: fmt.Errorf("unhandled command: %v", c.Type)}
	}
}

// Database returns a copy of the underlying database. The caller MUST
// ensure that no transaction is taking place during this call, or an error may
// be returned. If leader is true, this operation is performed with a read
// consistency level equivalent to "weak". Otherwise no guarantees are made
// about the read consistency level.
//
// http://sqlite.org/howtocorrupt.html states it is safe to do this
// as long as no transaction is in progress.
func (s *Store) Database(leader bool) ([]byte, error) {
	if leader && s.raft.State() != raft.Leader {
		return nil, errors.New(ErrNotLeader)
	}
	return s.db.Serialize()
}

// Snapshot returns a snapshot of the database. The caller must ensure that
// no transaction is taking place during this call. Hashicorp Raft guarantees
// that this function will not be called concurrently with Apply, as it states
// Apply and Snapshot are always called from the same thread. This means there
// is no need to synchronize this function with Execute(). However queries that
// involve a transaction must be blocked.
//
// http://sqlite.org/howtocorrupt.html states it is safe to do this
// as long as no transaction is in progress.
func (s *Store) Snapshot() (raft.FSMSnapshot, error) {
	fsm := &fsmSnapshot{
		startT: time.Now(),
		ctx:    s.ctx,
	}

	s.txMu.Lock()
	defer s.txMu.Unlock()

	fsm.database, _ = s.db.Serialize()
	// The error code is not meaningful from Serialize(). The code needs to be able
	// handle a nil byte slice being returned.

	stats.Add(numSnaphots, 1)
	log.WithContext(s.ctx).Debugf("node snapshot created in %s", time.Since(fsm.startT))
	return fsm, nil
}

// Restore restores the node to a previous state. The Hashicorp docs state this
// will not be called concurrently with Apply(), so synchronization with Execute()
// is not necessary.To prevent problems during queries, which may not go through
// the log, it blocks all query requests.
func (s *Store) Restore(rc io.ReadCloser) error {
	startT := time.Now()

	s.queryMu.Lock()
	defer s.queryMu.Unlock()

	var uint64Size uint64
	inc := int64(unsafe.Sizeof(uint64Size))

	// Read all the data into RAM, since we have to decode known-length
	// chunks of various forms.
	var offset int64
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		return fmt.Errorf("readall: %s", err)
	}

	// Get size of database, checking for compression.
	compressed := false
	sz, err := readUint64(b[offset : offset+inc])
	if err != nil {
		return fmt.Errorf("read compression check: %s", err)
	}
	offset = offset + inc

	if sz == math.MaxUint64 {
		compressed = true
		// Database is actually compressed, read actual size next.
		sz, err = readUint64(b[offset : offset+inc])
		if err != nil {
			return fmt.Errorf("read compressed size: %s", err)
		}
		offset = offset + inc
	}

	// Now read in the database file data, decompress if necessary, and restore.
	var database []byte
	if sz > 0 {
		if compressed {
			buf := new(bytes.Buffer)
			gz, err := gzip.NewReader(bytes.NewReader(b[offset : offset+int64(sz)]))
			if err != nil {
				return err
			}

			if _, err := io.Copy(buf, gz); err != nil {
				return fmt.Errorf("SQLite database decompress: %s", err)
			}

			if err := gz.Close(); err != nil {
				return err
			}
			database = buf.Bytes()
		} else {
			database = b[offset : offset+int64(sz)]
		}
	} else {
		log.WithContext(s.ctx).Debug("no database data present in restored snapshot")
		database = nil
	}

	if err := s.db.Close(); err != nil {
		return fmt.Errorf("failed to close pre-restore database: %s", err)
	}

	var db *sql.DB
	if !s.dbConf.Memory && s.lastCommandIdxOnOpen == 0 {
		// A snapshot clearly exists (this function has been called) but there
		// are no command entries in the log -- so Apply will not be called.
		// Therefore this is the last opportunity to create the on-disk database
		// before Raft starts.
		db, err = s.openOnDisk(database)
		if err != nil {
			return fmt.Errorf("open on-disk file during restore: %s", err)
		}
		s.onDiskCreated = true
		log.WithContext(s.ctx).Debug("successfully switched to on-disk database due to restore")
	} else {
		// Deserialize into an in-memory database because a) an in-memory database
		// has been requested, or b) while there was a snapshot, there are also
		// command entries in the log. So by sticking with an in-memory database
		// those entries will be applied in the fastest possible manner. We will
		// defer creation of any database on disk until the Apply function.
		db, err = s.openInMemory(database)
		if err != nil {
			return fmt.Errorf("openInMemory: %s", err)
		}
	}
	s.db = db

	stats.Add(numRestores, 1)
	log.WithContext(s.ctx).Debugf("node restored in %s", time.Since(startT))
	return nil
}

// RegisterObserver registers an observer of Raft events
func (s *Store) RegisterObserver(o *raft.Observer) {
	s.raft.RegisterObserver(o)
}

// DeregisterObserver deregisters an observer of Raft events
func (s *Store) DeregisterObserver(o *raft.Observer) {
	s.raft.DeregisterObserver(o)
}

// logSize returns the size of the Raft log on disk.
func (s *Store) logSize() (int64, error) {
	fi, err := os.Stat(filepath.Join(s.raftDir, raftDBPath))
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

type fsmSnapshot struct {
	startT   time.Time
	database []byte
	ctx      context.Context
}

// Persist writes the snapshot to the given sink.
func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	defer func() {
		log.WithContext(f.ctx).Debugf("snapshot and persist took %s", time.Since(f.startT))
	}()

	err := func() error {
		b := new(bytes.Buffer)

		// Flag compressed database by writing max uint64 value first.
		// No SQLite database written by earlier versions will have this
		// as a size. *Surely*.
		err := writeUint64(b, math.MaxUint64)
		if err != nil {
			return err
		}
		if _, err := sink.Write(b.Bytes()); err != nil {
			return err
		}
		b.Reset() // Clear state of buffer for future use.

		// Get compressed copy of database.
		cdb, err := f.compressedDatabase()
		if err != nil {
			return err
		}

		if cdb != nil {
			// Write size of compressed database.
			err = writeUint64(b, uint64(len(cdb)))
			if err != nil {
				return err
			}
			if _, err := sink.Write(b.Bytes()); err != nil {
				return err
			}

			// Write compressed database to sink.
			if _, err := sink.Write(cdb); err != nil {
				return err
			}
		} else {
			log.WithContext(f.ctx).Debug("no database data available for snapshot")
			err = writeUint64(b, uint64(0))
			if err != nil {
				return err
			}
			if _, err := sink.Write(b.Bytes()); err != nil {
				return err
			}
		}

		// Close the sink.
		return sink.Close()
	}()

	if err != nil {
		sink.Cancel()
		return err
	}

	return nil
}

func (f *fsmSnapshot) compressedDatabase() ([]byte, error) {
	if f.database == nil {
		return nil, nil
	}

	var buf bytes.Buffer
	gz, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	if _, err := gz.Write(f.database); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Database copies contents of the underlying SQLite database to dst
func (s *Store) database(leader bool, dst io.Writer) error {
	if leader && s.raft.State() != raft.Leader {
		return errors.New(ErrNotLeader)
	}

	f, err := ioutil.TempFile("", "rqlilte-snap-")
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	if err := s.db.Backup(f.Name()); err != nil {
		return err
	}

	of, err := os.Open(f.Name())
	if err != nil {
		return err
	}
	defer of.Close()

	_, err = io.Copy(dst, of)
	return err
}

// Release is a no-op.
func (f *fsmSnapshot) Release() {}

func readUint64(b []byte) (uint64, error) {
	var sz uint64
	if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &sz); err != nil {
		return 0, err
	}
	return sz, nil
}

func writeUint64(w io.Writer, v uint64) error {
	return binary.Write(w, binary.LittleEndian, v)
}

// enabledFromBool converts bool to "enabled" or "disabled".
func enabledFromBool(b bool) string {
	if b {
		return "enabled"
	}
	return "disabled"
}

// prettyVoter converts bool to "voter" or "non-voter"
func prettyVoter(v bool) string {
	if v {
		return "voter"
	}
	return "non-voter"
}

// pathExists returns true if the given path exists.
func pathExists(p string) bool {
	if _, err := os.Lstat(p); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// dirSize returns the total size of all files in the given directory
func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
