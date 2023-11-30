package websocketmanager

import "github.com/gorilla/websocket"

// Client represents a WebSocket client connection.
type Client struct {
	Conn   *websocket.Conn
	UserID int
}
