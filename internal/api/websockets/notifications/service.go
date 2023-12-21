package notifications

import (
	"bff/internal/api/repository"
	"bff/internal/api/websockets"
	"bff/structs"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

type Websockets struct {
	Wsmanager WebSocketManager
	Repo      repository.MicroserviceRepositoryInterface
}

// NewService creates a new instance of the notification service with the necessary dependencies.
func NewWebsockets(wsmanager WebSocketManager, repo repository.MicroserviceRepositoryInterface) *Websockets {
	return &Websockets{
		Wsmanager: wsmanager,
		Repo:      repo,
	}
}

type ActionType string

const (
	ActionRead   ActionType = "read"
	ActionDelete ActionType = "delete"
)

type NotificationMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketManager interface {
	BroadcastMessage(message []byte, userID int)
	AddClient(client *websockets.Client)
	RemoveClient(client *websockets.Client)
	RemoveClientByUserID(userID int)
}

// Service is responsible for managing notifications.

func (s *Websockets) CreateNotification(notification *structs.Notifications) (*structs.Notifications, error) {
	res, err := s.Repo.CreateNotification(notification)
	if err != nil {
		return nil, err
	}

	notificationData := NotificationMessage{
		Data: res.Data,
		Type: "new_notification",
	}

	notificationJSON, err := json.Marshal(notificationData)
	if err != nil {
		return nil, err
	}

	s.Wsmanager.BroadcastMessage(notificationJSON, res.ToUserID)

	return res, nil
}

func processNotificationMessage(repo repository.MicroserviceRepositoryInterface, msg []byte) {
	var message struct {
		Action         ActionType `json:"action"`
		NotificationId int        `json:"notification_id"`
	}
	_ = json.Unmarshal(msg, &message)

	switch message.Action {
	case ActionRead:
		_ = repo.MarkNotificationRead(message.NotificationId)
	case ActionDelete:
		_ = repo.DeleteNotification(message.NotificationId)
	}
}
