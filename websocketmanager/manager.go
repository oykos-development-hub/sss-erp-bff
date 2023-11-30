package websocketmanager

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var Manager = struct {
	sync.RWMutex
	Clients map[*Client]bool
}{
	Clients: make(map[*Client]bool),
}

func BroadcastNotification(message []byte) {
	Manager.Lock()
	defer Manager.Unlock()

	for client := range Manager.Clients {
		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error writing message:", err)
			client.Conn.Close()
			delete(Manager.Clients, client)
		}
	}
}
