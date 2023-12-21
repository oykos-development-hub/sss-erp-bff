package notifications

import (
	"bff/internal/api/websockets"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (ws *Websockets) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	loggedInAccount, err := ws.Repo.AuthenticateUser(r)
	if err != nil {
		log.Println("Authentication failed:", err)
		return
	}

	client := &websockets.Client{
		Conn:   conn,
		UserID: loggedInAccount.Id,
		ID:     uuid.New().String(),
	}

	ws.Wsmanager.AddClient(client)

	notificiations, err := ws.Repo.FetchNotifications(loggedInAccount.Id)
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

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			ws.Wsmanager.RemoveClient(client)
			break
		}
		log.Printf("Received: %s\n", msg)

		processNotificationMessage(ws.Repo, msg)
	}
}
