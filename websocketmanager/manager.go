package websocketmanager

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientManager struct {
	sync.RWMutex
	Clients map[int]map[string]*Client // Map user IDs to Clients
}

var Manager = ClientManager{
	Clients: make(map[int]map[string]*Client),
}

func BroadcastNotification(message []byte, userID int) {
	Manager.RLock()
	clients, exists := Manager.Clients[userID]
	Manager.RUnlock()

	if exists {
		for _, client := range clients {
			err := client.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error writing message:", err)
				removeClient(client)
			}
		}
	}
}

func removeClient(client *Client) {
	Manager.Lock()
	defer Manager.Unlock()

	err := client.Conn.Close()
	if err != nil {
		log.Printf("Error closing connection for user %d: %s\n", client.UserID, err)
	}

	if _, exists := Manager.Clients[client.UserID]; exists {
		delete(Manager.Clients[client.UserID], client.ID)

		// Optionally, clean up the userID entry if no clients are left
		if len(Manager.Clients[client.UserID]) == 0 {
			delete(Manager.Clients, client.UserID)
		}
	}
}

func RemoveClientByUserID(userID int) {
	if clients, exists := Manager.Clients[userID]; exists {
		for _, client := range clients {
			removeClient(client)
		}
	}
}
