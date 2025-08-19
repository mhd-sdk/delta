package client

import (
	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

// func (c *Client) NotifyChange(logs Logs) {
// 	c.Conn.WriteJSON(results)
// }
