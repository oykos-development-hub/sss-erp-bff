package websocketmanager

import (
	"encoding/json"
)

type ActionType string

const (
	ActionRead   ActionType = "read"
	ActionDelete ActionType = "delete"
)

func processMessage(client *Client, msg []byte) {
	var message struct {
		Action         ActionType `json:"action"`
		NotificationId int        `json:"notification_id"`
	}
	_ = json.Unmarshal(msg, &message)

	switch message.Action {
	case ActionRead:
		_ = markNotificationRead(message.NotificationId)
	case ActionDelete:
		_ = deleteNotification(message.NotificationId)
	}
}
