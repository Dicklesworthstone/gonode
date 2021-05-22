package node

// NatsUsecase contains methods for Send Ticket to SuperNode.
type NatsUsecase interface {
	// UploadSignedTicket for Publish Signed Ticket to Supernode
	UploadSignedTicket(msg interface{}, destSubject string) error
}
