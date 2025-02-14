package broker

import (
	"fmt"
	"log"
	"net"
)

type Broker struct {
	Address     string
	Topics      map[string][]string
	Connections []net.Conn
}

func NewBroker(address string) *Broker {
	topics := make(map[string][]string)
	connections := make([]net.Conn, 3)
	return &Broker{
		Address:     address,
		Topics:      topics,
		Connections: connections,
	}
}

func (broker *Broker) Start() error {
	listener, err := net.Listen("tcp", broker.Address)
	if err != nil {
		log.Fatalf("Error starting listener %v", err)
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
		go broker.HandleConnection(conn)
	}
}

func (broker *Broker) HandleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error reading from connection %s: %v", conn.RemoteAddr(), err)
			return
		}
		fmt.Printf("Received data: %s\n", string(buffer[:n]))
		topic := string(buffer[:4])
		message := string(buffer[5:n])
		fmt.Printf("Topic: %s, Message: %s\n", topic, message)
		broker.Topics[topic] = append(broker.Topics[topic], message)

		// Write back to Connection
		_, err = conn.Write(
			[]byte(fmt.Sprintf("Received data with topic: %s, message: %s", topic, message)),
		)
		if err != nil {
			log.Printf("Error writing to connection %s: %v", conn.RemoteAddr(), err)
			return
		}
	}
}
