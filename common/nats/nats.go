package nats

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
)

// Connection for saving *nats.Conn object to used for Subscribing or Publishing
type Connection struct {
	Natsconn *nats.Conn
}

// Publish for publish message to Destination Subject / Topic
func Publish(message []byte, destSubject string, nc Connection) error {
	err := nc.Natsconn.Publish(destSubject, message)
	if err != nil {
		return err
	}
	return nil
}

// Connect for start connection to Nats and save object nats connection to struct Connection
func Connect(host string, port int) (Connection, error) {
	// Init Nats Connection
	natsServer := fmt.Sprintf("%s:%d", host, port)
	nc, err := nats.Connect(natsServer)
	if err != nil {
		return Connection{}, err
	}
	return Connection{nc}, nil
}
