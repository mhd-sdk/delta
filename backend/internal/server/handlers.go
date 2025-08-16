package server

import (
	"log"

	"github.com/delta/internal/client"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func initHandlers(s *Server) {
	s.fiberServer.Get("/ws/status", s.subscribeToStatus())
	s.fiberServer.Post("/api/market-data/bars", s.GetCandlesticks())
}

func (s *Server) subscribeToStatus() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		s.AddClient(&client.Client{Conn: c})
		defer s.RemoveClient(&client.Client{Conn: c})

		for {
			if _, _, err := c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
		}

	})
}
