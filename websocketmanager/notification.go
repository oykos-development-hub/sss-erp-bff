package websocketmanager

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"log"
	"strconv"
)

type NotificationMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func fetchNotifications(userId int) ([]*structs.Notifications, error) {
	input := dto.GetNotificationInputMS{}
	input.ToUserID = &userId

	res := &dto.GetNotificationListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.NOTIFICATIONS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func CreateNotification(notification *structs.Notifications) (*structs.Notifications, error) {
	res := &dto.GetNotificationResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.NOTIFICATIONS_ENDPOINT, notification, res)
	if err != nil {
		return nil, err
	}

	notificationData := NotificationMessage{
		Data: res.Data,
		Type: "new_notification",
	}

	notificationJSON, err := json.Marshal(notificationData)
	if err != nil {
		log.Println("Error marshaling notification:", err)
		return nil, err
	}

	BroadcastNotification(notificationJSON)

	return &res.Data, nil
}

func getNotification(id int) (*structs.Notifications, error) {
	res := &dto.GetNotificationResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.NOTIFICATIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func markNotificationRead(notificationID int) error {
	notification, err := getNotification(notificationID)
	if err != nil {
		return err
	}

	notification.IsRead = true

	err = updateNotification(notificationID, notification)
	if err != nil {
		return err
	}

	return nil
}

func updateNotification(notificationID int, notification *structs.Notifications) error {
	_, err := shared.MakeAPIRequest("PUT", config.NOTIFICATIONS_ENDPOINT+"/"+strconv.Itoa(notificationID), notification, nil)
	if err != nil {
		return err
	}

	return nil
}

func deleteNotification(notificationId int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.NOTIFICATIONS_ENDPOINT+"/"+strconv.Itoa(notificationId), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
