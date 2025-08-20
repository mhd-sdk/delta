package client

import (
	"testing"

	"github.com/gofiber/contrib/websocket"
)

func TestNewClient(t *testing.T) {
	client := &Client{}
	if client.Conn != nil {
		t.Error("New client should have nil connection")
	}
}

func TestSetConnection(t *testing.T) {
	client := &Client{}

	// We can't easily test with a real websocket connection in unit tests
	// but we can test the struct field assignment
	if client.Conn != nil {
		t.Error("Connection should be nil initially")
	}

	// Mock connection assignment (in real usage, this would be a real websocket.Conn)
	var mockConn *websocket.Conn
	client.Conn = mockConn

	if client.Conn != mockConn {
		t.Error("Connection assignment failed")
	}
}
