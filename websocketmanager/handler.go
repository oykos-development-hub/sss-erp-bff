package websocketmanager

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	loggedInAccount, err := authenticateUser(r)
	if err != nil {
		log.Println("Authentication failed:", err)
		return
	}

	client := &Client{
		Conn:   conn,
		UserID: loggedInAccount.Id,
		ID:     uuid.New().String(),
	}

	if _, exists := Manager.Clients[loggedInAccount.Id]; !exists {
		Manager.Clients[loggedInAccount.Id] = make(map[string]*Client)
	}
	Manager.Lock()
	Manager.Clients[loggedInAccount.Id][client.ID] = client
	Manager.Unlock()

	notificiations, err := fetchNotifications(loggedInAccount.Id)
	if err != nil {
		log.Println("Error fetching initial data:", err)
		return
	}

	message := NotificationMessage{
		Type: "initial_data",
		Data: notificiations,
	}

	notificationsJSON, _ := json.Marshal(message)

	err = client.Conn.WriteMessage(websocket.TextMessage, notificationsJSON)
	if err != nil {
		log.Println("Error sending initial data:", err)
		return
	}

	handleMessages(client)
}

func handleMessages(client *Client) {
	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			removeClient(client)
			break
		}
		log.Printf("Received: %s\n", msg)

		processMessage(client, msg)
	}
}
