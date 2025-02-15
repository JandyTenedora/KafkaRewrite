package broker

import (
	"fmt"
	"log"
	"net"
)

func (broker *Broker) SendAck(conn net.Conn) error{
	_, err := conn.Write([]byte(fmt.Sprintf("Message received from address: %v", broker.Address)))
	if err != nil {
		log.Fatalf("Error in sending ack: %v", err)
	}
	return err
}
