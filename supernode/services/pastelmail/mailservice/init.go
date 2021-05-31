package mailservice

import (
	"github.com/pastelnetwork/gonode/common/nats"
	"github.com/pastelnetwork/gonode/common/storage"
	"github.com/pastelnetwork/gonode/pastel"
	"github.com/pastelnetwork/gonode/supernode/node"
	"github.com/pastelnetwork/gonode/supernode/node/init"
	"github.com/pastelnetwork/gonode/supernode/node/init/client"
	"github.com/pastelnetwork/gonode/supernode/services/pastelmail"
)

// Service represent artwork service.
type Service struct {
	db           storage.KeyValue
	pastelClient pastel.Client
	nodeClient   node.Client
	pastelMail   init.PastelMail
}

// NewService returns a new Service instance.
func NewService(db storage.KeyValue, pastelClient pastel.Client, nodeClient node.Client, nc nats.Connection) pastelmail.PastelMailUsecase {
	return &Service{
		db:           db,
		pastelClient: pastelClient,
		nodeClient:   nodeClient,
		pastelMail:   client.NewPublishService(nc),
	}
}
