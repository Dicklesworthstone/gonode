package model

type (

	// PastelMailMessaging is Struct for Send Message to other Supernodes through Nats
	PastelMailMessaging struct {
		TrxId         string
		Message       string
		SourceAddress string
	}
)
