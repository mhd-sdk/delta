package server

import (
	"log"

	"github.com/delta/internal/client"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func initHandlers(s *Server) {
	s.fiberServer.Get("/scan", getScans(s))
	s.fiberServer.Get("/ws/scan", subscribeToScans(s))
	// alert will receive parameters to filter the alerts
	// s.fiberServer.Get("/ws/alert", subscribeToScans(s))
}

func getScans(s *Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(s.scanner.GetResults())
	}
}

func subscribeToScans(s *Server) fiber.Handler {
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
