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

func NewClient() *Client {
	return &Client{
		Conn: nil, // Connection will be set later
	}
}

func (c *Client) SetConnection(conn *websocket.Conn) {
	c.Conn = conn
}
