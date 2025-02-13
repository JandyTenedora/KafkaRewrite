package broker

import "net"

type Broker struct {
	Address     string
	Topics      map[string][]byte
	Connections []net.Conn
}

func NewBroker(address string) *Broker{
				topics := make(map[string][]byte)
				connections := make([]net.Conn, 3)
				return &Broker{
								Address: address,
								Topics: topics,
								Connections: connections,
				}
}
