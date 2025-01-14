package metadb

import (
	"context"

	"github.com/pastelnetwork/gonode/metadb/rqlite/store"
)

const (
	logPrefix = "metadb"
)

// MetaDB represents the metadb service
type MetaDB interface {
	// Run starts the rqlite service
	Run(ctx context.Context) error
	// Query execute a query, not support multple statements
	// level can be 'none', 'weak', and 'strong'
	Query(ctx context.Context, statement string, level string) (*QueryResult, error)
	// Write execute a statement, not support multple statements
	Write(ctx context.Context, statement string) (*WriteResult, error)
}

type service struct {
	nodeID string        // the node id for rqlite cluster
	config *Config       // the service configuration
	db     *store.Store  // the store for accessing the rqlite cluster
	ready  chan struct{} // mark the rqlite node is started
}

// New returns a new service for metadb
func New(config *Config, nodeID string) MetaDB {
	return &service{
		nodeID: nodeID,
		config: config,
		ready:  make(chan struct{}, 5),
	}
}

// Run starts the rqlite server
func (s *service) Run(ctx context.Context) error {
	return s.startServer(ctx)
}
