package client

import (
	"github.com/delta/internal/worker/scanner"
	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

func (c *Client) NotifyScan(results scanner.ScanResults) {
	c.Conn.WriteJSON(results)
}
