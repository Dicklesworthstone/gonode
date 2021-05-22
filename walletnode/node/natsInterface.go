package node

// NatsUsecase contains methods for Send Ticket to SuperNode.
type NatsUsecase interface {
	UploadSignedTicket(msg interface{}, destSubject string) error
}
