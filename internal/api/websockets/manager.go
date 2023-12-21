package websockets

import (
	"bff/log"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientManager struct {
	sync.RWMutex
	Clients map[int]map[string]*Client // Map user IDs to Clients
}

type Client struct {
	Conn   *websocket.Conn
	UserID int
	ID     string
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients: make(map[int]map[string]*Client),
	}
}

func (m *ClientManager) AddClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, exists := m.Clients[client.UserID]; !exists {
		m.Clients[client.UserID] = make(map[string]*Client)
	}
	m.Clients[client.UserID][client.ID] = client
}

func (m *ClientManager) RemoveClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, exists := m.Clients[client.UserID]; exists {
		delete(m.Clients[client.UserID], client.ID)
		if len(m.Clients[client.UserID]) == 0 {
			delete(m.Clients, client.UserID)
		}
	}

	err := client.Conn.Close()
	if err != nil {
		log.Logger.Printf("Error closing connection for user %d: %s\n", client.UserID, err)
	}
}

func (m *ClientManager) RemoveClientByUserID(userID int) {
	if clients, exists := m.Clients[userID]; exists {
		for _, client := range clients {
			m.RemoveClient(client)
		}
	}
}

func (m *ClientManager) BroadcastMessage(message []byte, userID int) {
	m.RLock()
	clients, exists := m.Clients[userID]
	m.RUnlock()

	if exists {
		for _, client := range clients {
			err := client.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Logger.Println("Error writing message:", err)
				m.RemoveClient(client)
			}
		}
	}
}
