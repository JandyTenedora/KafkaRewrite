package producer

import (
	"errors"
	"net"
	"testing"
	"time"
)

// Mock net.Conn implementation for testing
type mockConn struct {
	closeErr error
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	return len(b), nil
}

func (m *mockConn) Close() error {
	return m.closeErr
}

func (m *mockConn) LocalAddr() net.Addr {
	return nil
}

func (m *mockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

// Test NewConnection with valid brokerAddress
func TestNewConnection_Success(t *testing.T) {
	// Mock net.Dial to return a mockConn instance
	originalDial := netDial
	netDial = func(network, address string) (net.Conn, error) {
		return &mockConn{}, nil
	}
	defer func() { netDial = originalDial }() // Restore the original function after the test

	conn, err := NewConnection("localhost:9092")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if conn == nil {
		t.Fatalf("expected a connection instance, got nil")
	}
}

// Test NewConnection with invalid brokerAddress
func TestNewConnection_Failure(t *testing.T) {
	// Mock net.Dial to simulate an error
	originalDial := netDial
	netDial = func(network, address string) (net.Conn, error) {
		return nil, errors.New("connection failed")
	}
	defer func() { netDial = originalDial }() // Restore the original function after the test

	conn, err := NewConnection("invalid-broker:9092")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if conn != nil {
		t.Fatalf("expected nil connection, got %v", conn)
	}
}

// Test WriteMessage 
func TestWriteMessage(t *testing.T) {
	conn := &Connection{conn: &mockConn{}}
	err := conn.WriteMessage(int32(10), []byte("message"))
	if err != nil {
		t.Fatalf("expected success, got %v", err)
	}
}

//// Test Close (method not implemented)
func TestClose(t *testing.T) {
	conn := &Connection{conn: &mockConn{}}
	err := conn.Close()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "Close not implemented" {
		t.Fatalf("expected 'Close not implemented', got %v", err)
	}
}

