package broker

import "net"

type Broker struct {
	Address     string
	Topics      map[string][]byte
	Connections []net.Conn
}
