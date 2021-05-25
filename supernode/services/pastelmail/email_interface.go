package pastelmail

import "github.com/pastelnetwork/gonode/supernode/model"

type PastelMailUsecase interface {
	DoSomething(msg model.PastelMailMessaging)
}
