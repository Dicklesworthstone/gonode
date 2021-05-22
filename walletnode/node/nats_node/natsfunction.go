package nats_node

import (
	"encoding/json"
	"github.com/pastelnetwork/gonode/common/nats"
)

func (nc *natsClient) UploadSignedTicket(msg interface{}, destSubject string) error {
	outMessage, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = nats.Publish(outMessage, destSubject, nc.NatsConn)
	if err != nil {
		return err
	}

	return nil
}
