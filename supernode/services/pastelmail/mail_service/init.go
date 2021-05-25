package mail_service

import (
	"github.com/pastelnetwork/gonode/common/nats"
	"github.com/pastelnetwork/gonode/common/storage"
	"github.com/pastelnetwork/gonode/pastel"
	"github.com/pastelnetwork/gonode/supernode/node"
	"github.com/pastelnetwork/gonode/supernode/node/nats_node"
	"github.com/pastelnetwork/gonode/supernode/node/nats_node/nats_client"
	"github.com/pastelnetwork/gonode/supernode/services/pastelmail"
)

// Service represent artwork service.
type Service struct {
	db           storage.KeyValue
	pastelClient pastel.Client
	nodeClient   node.Client
	pastelMail   nats_node.PastelMail
}

// NewService returns a new Service instance.
func NewService(db storage.KeyValue, pastelClient pastel.Client, nodeClient node.Client, nc nats.Connection) pastelmail.PastelMailUsecase {
	return &Service{
		db:           db,
		pastelClient: pastelClient,
		nodeClient:   nodeClient,
		pastelMail:   nats_client.NewPublishService(nc),
	}
}
