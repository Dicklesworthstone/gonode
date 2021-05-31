package nats

import (
	"fmt"
	"strings"
	"time"

	nats "github.com/nats-io/nats.go"
)

// Connection for saving *nats.Conn object to used for Subscribing or Publishing
type Connection struct {
	Natsconn *nats.Conn
}

// Connect for start connection to Nats and save object nats connection to struct Connection
// For connecting with Ping Interval the withPing variable should be set 'true'
// For starting cluster connection isCluster should be set to 'true' and listConnection []string parameter should be sent
func Connect(host string, port int, withPing bool, isCluster bool, listConnection []string) (Connection, error) {
	natsServer := fmt.Sprintf("%s:%d", host, port)
	if withPing {
		nc, err := nats.Connect(natsServer, nats.PingInterval(20*time.Second), nats.MaxPingsOutstanding(5))
		if err != nil {
			return Connection{}, err
		}
		return Connection{nc}, nil
	}
	if isCluster {
		nc, err := nats.Connect(strings.Join(listConnection, ","))
		if err != nil {
			return Connection{}, err
		}
		return Connection{nc}, nil
	}
	nc, err := nats.Connect(natsServer)
	if err != nil {
		return Connection{}, err
	}
	return Connection{nc}, nil
}

// Publish for publish message to Destination Subject
func Publish(message []byte, destSubject string, nc Connection) error {
	err := nc.Natsconn.Publish(destSubject, message)
	if err != nil {
		return err
	}
	return nil
}

// Request for Request a message to Destination Subject within timeout duration
func Request(message []byte, destSubject string, timeout time.Duration, nc Connection) ([]byte, error) {
	resp, err := nc.Natsconn.Request(destSubject, message, timeout*time.Millisecond)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
