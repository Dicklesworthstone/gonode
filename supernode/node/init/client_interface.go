package init

import (
	"github.com/pastelnetwork/gonode/supernode/model"
)

// PastelMail represents interaction supernode with Nats for feature PastelMail
type PastelMail interface {
	MailToSuperNode(msg model.PastelMailMessaging, destSubject string) error
}
