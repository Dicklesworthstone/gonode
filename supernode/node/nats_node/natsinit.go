package nats_node

import (
	"github.com/pastelnetwork/gonode/common/nats"
)

type natsClient struct {
	NatsConn nats.Connection
}

// NewClient return natsClient instance
func NewClient(host string, port int) *natsClient {
	// Init Nats Connection
	NatsConn, err := nats.Connect(host, port)
	if err != nil {
		return &natsClient{}
	}
	return &natsClient{NatsConn}
}
