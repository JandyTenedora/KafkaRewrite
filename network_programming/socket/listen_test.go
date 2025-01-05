package ch03

import (
	"fmt"
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	defer func() { _ = listener.Close() }()
	t.Logf("bound to %q", listener.Addr())
	for {
		conn, err_2 := listener.Accept()
		if err_2 != nil {
			t.Fatal(err_2)
		}

		go func(c net.Conn) {
			defer c.Close()
			fmt.Println("Placeholder for connection handling")
		}(conn)
	}
}
