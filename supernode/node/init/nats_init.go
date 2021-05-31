package init

import (
	"github.com/pastelnetwork/gonode/common/nats"
	natscomm "github.com/pastelnetwork/gonode/common/nats"
)

// NewNatsConnection return Interface of publish instance
func NewNatsConnection() nats.Connection {
	// Init Nats Connection
	NatsConn, err := natscomm.Connect("localhost", 4222, false, false, nil)
	if err != nil {
		return NatsConn
	}
	return NatsConn
}
