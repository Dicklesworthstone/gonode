package nats

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
)

func Publish(message []byte, destSubject string, nc *nats.Conn) error {
	err := nc.Publish(destSubject, message)
	if err != nil {
		return err
	}
	return nil
}

func Connect(host string, port int) (*nats.Conn, error) {
	// Init Nats Connection
	natsServer := fmt.Sprintf("%s:%d", host, port)
	nc, err := nats.Connect(natsServer)
	if err != nil {
		return nil, err
	}
	return nc, nil
}
