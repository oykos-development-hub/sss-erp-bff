package websocketmanager

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientManager struct {
	sync.RWMutex
	Clients map[int]*Client // Map user IDs to Clients
}

var Manager = ClientManager{
	Clients: make(map[int]*Client),
}

func BroadcastNotification(message []byte, userID int) {
	Manager.RLock() // Use RLock for read-only access
	client, exists := Manager.Clients[userID]
	Manager.RUnlock()

	if exists {
		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error writing message:", err)
			removeClient(client)
		}
	}

}
