module github.com/pastelnetwork/gonode/supernode

go 1.16

require (
	github.com/nats-io/nats.go v1.11.0 // indirect
	github.com/panjf2000/ants v1.3.0 // indirect
	github.com/panjf2000/ants/v2 v2.4.5 // indirect
	github.com/pastelnetwork/gonode/common v0.0.0
	github.com/pastelnetwork/gonode/pastel v0.0.0
	github.com/pastelnetwork/gonode/proto v0.0.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.37.0
)

replace github.com/pastelnetwork/gonode/common => ../common

replace github.com/pastelnetwork/gonode/proto => ../proto

replace github.com/pastelnetwork/gonode/pastel => ../pastel
