package nats_server

import (
	"encoding/json"
	"sync"

	natgo "github.com/nats-io/nats.go"

	"github.com/panjf2000/ants"
	"github.com/pastelnetwork/gonode/common/nats"
	"github.com/pastelnetwork/gonode/supernode/model"
	"github.com/pastelnetwork/gonode/supernode/services/pastelmail"
)

type NatsServer struct {
	pastelemail pastelmail.PastelMailUsecase
}

// StartSubscribe to Start Subscribe Subject / Topic to Nats
func StartSubscribe(service pastelmail.PastelMailUsecase, n nats.Connection) {
	handler := NatsServer{
		pastelemail: service,
	}

	var (
		topic = "pastel.pastelmail" // example
		wg    *sync.WaitGroup
	)
	poolingConnection, _ := ants.NewPool(20) // example 20

	wg.Add(1)
	go func() {
		// Simple Async Subscriber
		n.Natsconn.Subscribe(topic, func(m *natgo.Msg) {
			task := func(m *natgo.Msg) func() {
				return func() {
					handler.ReceivedMessage(m)
				}
			}
			poolingConnection.Submit(task(m))
		})
	}()

	wg.Wait()
}

// ReceivedMessage to processing input message from Nats from byte to Model
func (natserver *NatsServer) ReceivedMessage(m *natgo.Msg) {
	// If Message is encrypted first before send to nats
	// decrypt in here

	//

	// Unmarshall
	inp := model.PastelMailMessaging{}
	err := json.Unmarshal(m.Data, &inp)
	if err != nil {
		return
	}

	natserver.pastelemail.DoSomething(inp)
}
