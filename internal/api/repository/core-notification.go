package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetNotification(id int) (*structs.Notifications, error) {
	res := &dto.GetNotificationResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.NOTIFICATIONS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateNotification(notificationID int, notification *structs.Notifications) error {
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.NOTIFICATIONS+"/"+strconv.Itoa(notificationID), notification, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) DeleteNotification(notificationId int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Core.NOTIFICATIONS+"/"+strconv.Itoa(notificationId), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) FetchNotifications(userId int) ([]*structs.Notifications, error) {
	input := dto.GetNotificationInputMS{}
	input.ToUserID = &userId

	res := &dto.GetNotificationListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.NOTIFICATIONS, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateNotification(notification *structs.Notifications) (*structs.Notifications, error) {
	res := &dto.GetNotificationResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.NOTIFICATIONS, notification, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) MarkNotificationRead(notificationID int) error {
	notification, err := repo.GetNotification(notificationID)
	if err != nil {
		return err
	}

	notification.IsRead = true

	err = repo.UpdateNotification(notificationID, notification)
	if err != nil {
		return err
	}

	return nil
}
