package broker

import (
	"fmt"
	"log"
	"net"
)

func (broker *Broker) SendAck(conn net.Conn) error {
	_, err := conn.Write([]byte("ACK"))
	if err != nil {
		log.Fatalf("Error in sending ack: %v", err)
	}
	return err
}

func (broker *Broker) ReceiveAck(conn net.Conn) error {
	buffer := make([]byte, 3)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error in receiving ack: %v", err)
	}
	ack := string(buffer)
	if ack != "ACK" {
		return fmt.Errorf("invalid ack: %s", ack)
	}
	log.Println("Received ack")
	return nil
}
