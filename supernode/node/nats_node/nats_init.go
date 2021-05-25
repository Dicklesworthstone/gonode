package nats_node

import (
	"github.com/pastelnetwork/gonode/common/nats"
	natscomm "github.com/pastelnetwork/gonode/common/nats"
)

// NewNatsConnection return Interface of publish instance
func NewNatsConnection() nats.Connection {
	// Init Nats Connection
	NatsConn, err := natscomm.Connect("localhost", 4222)
	if err != nil {
		return NatsConn
	}
	return NatsConn
}
