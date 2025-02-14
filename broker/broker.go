package broker

import (
	"fmt"
	"log"
	"net"
)

type Broker struct {
	Address     string
	Topics      map[string][]byte
	Connections []net.Conn
}

func NewBroker(address string) *Broker {
	topics := make(map[string][]byte)
	connections := make([]net.Conn, 3)
	return &Broker{
		Address:     address,
		Topics:      topics,
		Connections: connections,
	}
}

func (broker *Broker) Start() {
	listener, error := net.Listen("tcp", broker.Address)
	if error != nil {
		log.Fatalf("Error starting listener %v", error)
	}
	defer listener.Close()

	fmt.Printf("Broker started on on %s\n", broker.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		broker.Connections = append(broker.Connections, conn)
		fmt.Printf("New connection from %s\n", conn.RemoteAddr())
	}
}
