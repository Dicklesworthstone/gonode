package node

// NatsUsecase contains methods for Send Ticket to SuperNode.
type NatsUsecase interface {
	UploadSignedTicket()
}
