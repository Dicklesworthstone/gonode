package client

import (
	"encoding/json"

	"github.com/pastelnetwork/gonode/common/nats"

	"github.com/pastelnetwork/gonode/supernode/model"
	"github.com/pastelnetwork/gonode/supernode/node/init"
)

type natsClient struct {
	NatsConn nats.Connection
}

// NewPublishService Create Repository / Service for send Message via Publish Nats
func NewPublishService(NatsConn nats.Connection) init.PastelMail {
	return &natsClient{
		NatsConn: NatsConn,
	}
}

// MailToSuperNode implements interface node.UploadSignedTicket()
func (nc *natsClient) MailToSuperNode(msg model.PastelMailMessaging, destSubject string) error {
	outMessage, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// After marshalling object JSON to String, then Encrypt in here

	// -- end here

	// Publish message to Super Nodes who are subscribed to Destination Subject
	err = nats.Publish(outMessage, destSubject, nc.NatsConn)
	if err != nil {
		return err
	}

	return nil
}
