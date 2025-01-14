package producer

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

var netDial = net.Dial

type Connection struct {
	conn net.Conn
}

/*
New Connection establishes a TCP connection to the Kafka brokerAddress
Parameters:
  - brokerAddress: the address of a kafka broker (e.g. "localhost: 9092")

Returns:
  - Pointer to a connection instance if the connection is successful
  - An error if the connection can not be established
*/
func NewConnection(brokerAddress string) (*Connection, error) {
	conn, err := netDial("tcp", brokerAddress)
	if err != nil {
		return nil, errors.New("failed to establish connection")
	}
	connection := &Connection{conn}
	return connection, err
}

// WriteMessage sends a message to the specified Kafka topic.
// Parameters:
// - topic: The name of the Kafka topic to which the message should be sent.
// - message: The serialized message payload to send.
// Returns:
// - An error if the message could not be sent successfully.
//
// Topic to be stored as 4 byte header of message
func (c *Connection) WriteMessage(topic int32, message []byte) error {
	// TODO: Implement the logic to send a message over the connection.
	topicBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(topicBuffer, uint32(topic))
	bufferedMessage := append(topicBuffer, message...)
	_, err := c.conn.Write([]byte(bufferedMessage))
	if err != nil {
		return fmt.Errorf("Error in writing message: %v", err)
	}
	return err // Should always be nil
}

// Close terminates the TCP connection to the Kafka broker.
// Returns:
// - An error if the connection could not be closed successfully.
func (c *Connection) Close() error {
	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("Error in closing connection: %v", err)
	}
	return err
}
