package nats_node

type natsClient struct{}

func NewClient() *nats.Conn {
	// Init Nats Connection
	nc, err := nats.Connect(host, port)
	if err != nil {

	}
	return nc
}
