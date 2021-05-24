package nats_node

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/panjf2000/ants"
	"github.com/pastelnetwork/gonode/supernode/model"

	nats "github.com/nats-io/nats.go"
)

// StartSubscribe to Start Subscribe Subject / Topic to Nats
func (n *natsClient) StartSubscribe(services []func()) {
	// what Subscribe need
	// - subscribe many subject

	var (
		topic = "Pastel Mail Messaging" // example
		wg    *sync.WaitGroup
	)
	poolingConnection, _ := ants.NewPool(20) // example 20

	wg.Add(1)
	go func() {
		// Simple Async Subscriber
		fmt.Printf("Application is Subscribing Pastel Mail Message %v", topic)
		nc.NatsConn.Natsconn.Subscribe(topic, func(m *nats.Msg) {
			task := func(m *nats.Msg) func() {
				return func() {
					ReceivedMessage(m)
				}
			}
			poolingConnection.Submit(task(m))
		})
	}()

	wg.Wait()
}

// ReceivedMessage to processing input message from Nats from byte to Model
func ReceivedMessage(m *nats.Msg) {
	// If Message is encrypted first before send to nats
	// decrypt in here

	// Unmarshall
	inp := model.PastelMailMessaging{}
	err := json.Unmarshal(m.Data, &inp)
	if err != nil {
		return
	}
}
