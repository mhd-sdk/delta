package server

import (
	"context"
	"log"

	"github.com/delta/internal/client"
	"github.com/delta/internal/pb"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func initHandlers(s *Server) {
	s.fiberServer.Get("/ws/logs", s.subscribeToLogs())
	s.fiberServer.Post("/api/market-data/bars", s.GetCandlesticks())
	s.fiberServer.Post("/api/algorithms/state", s.UpdateState())
}

func (s *Server) subscribeToLogs() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		client := &client.Client{Conn: c}
		s.AddClient(client)
		defer s.RemoveClient(client)

		// Channel to receive log events
		logChan := make(chan *pb.LogEvent, 100)

		// Start log subscription for all slaves
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go s.subscribeToAllSlaveLogs(ctx, logChan)

		// Handle WebSocket messages and forward log events
		go func() {
			for {
				if _, _, err := c.ReadMessage(); err != nil {
					log.Println("WebSocket read error:", err)
					cancel()
					return
				}
			}
		}()

		// Forward log events to WebSocket client
		for {
			select {

			case <-ctx.Done():
				return
			}
		}
	})
}
