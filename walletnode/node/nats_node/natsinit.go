package nats_node

import (
	"github.com/pastelnetwork/gonode/common/nats"
)

type natsClient struct {
	NatsConn nats.NatsConnection
}

func NewClient(host string, port int) (natsClient, error) {
	// Init Nats Connection
	NatsConn, err := nats.Connect(host, port)
	if err != nil {
		return natsClient{NatsConn}, err
	}
	return natsClient{NatsConn}, nil
}
