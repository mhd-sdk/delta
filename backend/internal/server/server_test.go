package server

import (
	"testing"

	"github.com/delta/internal/client"
)

func TestGetSlaves(t *testing.T) {
	discover := false
	s := New(&discover)
	slaves := s.GetSlaves()
	if len(slaves) != 0 {
		t.Errorf("Expected 0 slaves, got %d", len(slaves))
	}
}

func TestAddClient(t *testing.T) {
	discover := false
	s := New(&discover)

	c1 := &client.Client{}
	c2 := &client.Client{}

	s.AddClient(c1)
	s.AddClient(c2)

	if len(s.clients) != 2 {
		t.Errorf("expected 2 clients, got %d", len(s.clients))
	}

}

func TestRemoveClient(t *testing.T) {
	discover := false
	s := New(&discover)

	c1 := &client.Client{}
	c2 := &client.Client{}

	s.AddClient(c1)
	s.AddClient(c2)

	s.RemoveClient(c1)
	if len(s.clients) != 1 {
		t.Errorf("expected 1 client after removal, got %d", len(s.clients))
	}

	if s.clients[0] != c2 {
		t.Errorf("expected remaining client to be c2")
	}
}
