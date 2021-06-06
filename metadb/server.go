package metadb

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	"github.com/pastelnetwork/gonode/metadb/rqlite/cluster"
	"github.com/pastelnetwork/gonode/metadb/rqlite/disco"
	httpd "github.com/pastelnetwork/gonode/metadb/rqlite/http"
	"github.com/pastelnetwork/gonode/metadb/rqlite/store"
	"github.com/pastelnetwork/gonode/metadb/rqlite/tcp"
)

const (
	defaultJoinAttempts           = 5
	defaultJoinInterval           = 5 * time.Second
	defaultRaftHeartbeatTimeout   = time.Second
	defaultRaftElectionTimeout    = time.Second
	defaultRaftApplyTimeout       = 10 * time.Second
	defaultRaftOpenTimeout        = 120 * time.Second
	defaultRaftWaitForLeader      = true
	defaultRaftSnapThreshold      = 8192
	defaultRaftSnapInterval       = 30 * time.Second
	defaultRaftLeaderLeaseTimeout = 0 * time.Second
	defaultCompressionSize        = 150
	defaultCompressionBatch       = 5
)

// the order is: node id, raft advertise address vs raft address
func (s *service) idOrRaftAddr() string {
	if s.config.NodeID != "" {
		return s.config.NodeID
	}
	return s.config.RaftAddress
}

// determine the join addresses
func (s *service) determineJoinAddresses(ctx context.Context) ([]string, error) {
	var addrs []string
	if s.config.JoinAddress != "" {
		// explicit join addresses are first priority.
		addrs = strings.Split(s.config.JoinAddress, ",")
	}

	if s.config.DiscoveryID != "" {
		log.WithContext(ctx).Infof("register with discovery service at %s with ID %s", s.config.DiscoveryURL, s.config.DiscoveryID)

		c := disco.New(ctx, s.config.DiscoveryURL)
		r, err := c.Register(s.config.DiscoveryID, s.config.HTTPAddress)
		if err != nil {
			return nil, errors.Errorf("discovery register: %w", err)
		}
		log.WithContext(ctx).Infof("discovery service responded with nodes: %v", r.Nodes)

		for _, a := range r.Nodes {
			if a != s.config.HTTPAddress {
				addrs = append(addrs, a)
			}
		}
	}

	return addrs, nil
}

// wait until the store is in full consensus
func (s *service) waitForConsensus(ctx context.Context, dbStore *store.Store) error {
	if _, err := dbStore.WaitForLeader(defaultRaftOpenTimeout); err != nil {
		if defaultRaftWaitForLeader {
			return errors.Errorf("leader did not appear within timeout: %w", err)
		}
		log.WithContext(ctx).Infof("ignoring error while waiting for leader")
	}
	if err := dbStore.WaitForApplied(defaultRaftOpenTimeout); err != nil {
		return errors.Errorf("store log not applied within timeout: %w", err)
	}

	return nil
}

// start the http server
func (s *service) startHTTPServer(ctx context.Context, dbStore *store.Store, cs *cluster.Service) error {
	// create http server
	server := httpd.New(ctx, s.config.HTTPAddress, dbStore, cs, nil)

	// start the http server
	return server.Start()
}

// start a mux for rqlite node
func (s *service) startNodeMux(ctx context.Context, ln net.Listener) (*tcp.Mux, error) {
	mux, err := tcp.NewMux(ctx, ln, nil)
	if err != nil {
		return nil, errors.Errorf("create node-to-node mux: %w", err)
	}

	go mux.Serve()

	return mux, nil
}

// start the cluster server
func (s *service) startClusterService(ctx context.Context, tn cluster.Transport) (*cluster.Service, error) {
	c := cluster.New(ctx, tn)

	// set the api address
	c.SetAPIAddr(s.config.HTTPAddress)
	// open the cluster service
	if err := c.Open(); err != nil {
		return nil, err
	}
	return c, nil
}

// create and open the store of rqlite cluster
func (s *service) initStore(ctx context.Context, raftTn *tcp.Layer) (*store.Store, error) {
	// create and open the store, which is on disk
	dbConf := store.NewDBConfig("", false)
	db := store.New(ctx, raftTn, &store.Config{
		DBConf: dbConf,
		Dir:    s.config.DataDir,
		ID:     s.idOrRaftAddr(),
	})

	// set optional parameters on store
	db.SetRequestCompression(defaultCompressionBatch, defaultCompressionSize)
	db.ShutdownOnRemove = false
	db.SnapshotThreshold = defaultRaftSnapThreshold
	db.SnapshotInterval = defaultRaftSnapInterval
	db.LeaderLeaseTimeout = defaultRaftLeaderLeaseTimeout
	db.HeartbeatTimeout = defaultRaftHeartbeatTimeout
	db.ElectionTimeout = defaultRaftElectionTimeout
	db.ApplyTimeout = defaultRaftApplyTimeout

	// a pre-existing node
	bootstrap := false
	isNew := store.IsNewNode(s.config.DataDir)
	if isNew {
		bootstrap = true // new node, it needs to bootstrap
	} else {
		log.WithContext(ctx).Infof("node is detected in: %v", s.config.DataDir)
	}

	// determine the join addresses
	joins, err := s.determineJoinAddresses(ctx)
	if err != nil {
		return nil, errors.Errorf("determine join addresses: %w", err)
	}
	// supplying join addresses means bootstrapping a new cluster won't be required.
	if len(joins) > 0 {
		bootstrap = false
		log.WithContext(ctx).Info("join addresses specified, node is not bootstrap")
	} else {
		log.WithContext(ctx).Info("no join addresses")
	}
	// join address supplied, but we don't need them
	if !isNew && len(joins) > 0 {
		log.WithContext(ctx).Info("node is already member of cluster")
	}

	// open store
	if err := db.Open(bootstrap); err != nil {
		return nil, errors.Errorf("open store: %w", err)
	}
	s.db = db

	// execute any requested join operation
	if len(joins) > 0 && isNew {
		log.WithContext(ctx).Infof("join addresses are: %v", joins)

		// join rqlite cluster
		joinAddr, err := cluster.Join(
			ctx,
			"",
			joins,
			db.ID(),
			s.config.RaftAddress,
			true,
			defaultJoinAttempts,
			defaultJoinInterval,
			nil,
		)
		if err != nil {
			return nil, errors.Errorf("join cluster at %v: %w", joins, err)
		}
		log.WithContext(ctx).Infof("successfully joined cluster at %v", joinAddr)
	}

	return db, nil
}

// start the rqlite server, and try to join rqlite cluster if the join addresses is not empty
func (s *service) startServer(ctx context.Context) error {
	ctx = log.ContextWithPrefix(ctx, logPrefix)

	// create internode network mux and configure.
	muxListener, err := net.Listen("tcp", s.config.RaftAddress)
	if err != nil {
		return errors.Errorf("listen on %s: %w", s.config.RaftAddress, err)
	}
	mux, err := s.startNodeMux(ctx, muxListener)
	if err != nil {
		return errors.Errorf("start node mux: %w", err)
	}
	raftTn := mux.Listen(cluster.MuxRaftHeader)

	// create cluster service, so nodes can learn information about each other.
	// This can be started now since it doesn't require a functioning Store yet.
	cs, err := s.startClusterService(ctx, mux.Listen(cluster.MuxClusterHeader))
	if err != nil {
		return errors.Errorf("start create cluster service: %w", err)
	}

	// create and open the store
	db, err := s.initStore(ctx, raftTn)
	if err != nil {
		return errors.Errorf("create and open store: %w", err)
	}
	s.db = db

	// wait until the store is in full consensus
	if err := s.waitForConsensus(ctx, db); err != nil {
		return errors.Errorf("wait for consensus: %w", err)
	}
	log.WithContext(ctx).Info("store has reached consensus")

	// start the HTTP API server
	if err := s.startHTTPServer(ctx, db, cs); err != nil {
		return errors.Errorf("start http server: %w", err)
	}
	log.WithContext(ctx).Info("node is ready, block until context is done")

	// mark the rqlite node is ready
	s.ready <- struct{}{}

	// block until context is done
	<-ctx.Done()

	// close the rqlite store
	if err := db.Close(true); err != nil {
		log.WithContext(ctx).Errorf("close store: %v", err)
	}

	// close the mux listener
	if err := muxListener.Close(); err != nil {
		log.WithContext(ctx).Errorf("close mux listener: %v", err)
	}

	log.WithContext(ctx).Info("rqlite server is stopped")
	return nil
}
