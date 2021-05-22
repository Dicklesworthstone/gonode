package nats

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
)

type NatsConnection struct {
	Natsconn *nats.Conn
}

func Publish(message []byte, destSubject string, nc NatsConnection) error {
	err := nc.Natsconn.Publish(destSubject, message)
	if err != nil {
		return err
	}
	return nil
}

func Connect(host string, port int) (NatsConnection, error) {
	// Init Nats Connection
	natsServer := fmt.Sprintf("%s:%d", host, port)
	nc, err := nats.Connect(natsServer)
	if err != nil {
		return NatsConnection{}, err
	}
	return NatsConnection{nc}, nil
}
