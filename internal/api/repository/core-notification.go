package repository

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetNotification(id int) (*structs.Notifications, error) {
	res := &dto.GetNotificationResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Notifications+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateNotification(notificationID int, notification *structs.Notifications) error {
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.Notifications+"/"+strconv.Itoa(notificationID), notification, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) DeleteNotification(notificationID int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Core.Notifications+"/"+strconv.Itoa(notificationID), nil, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) FetchNotifications(userID int) ([]*structs.Notifications, error) {
	input := dto.GetNotificationInputMS{}
	input.ToUserID = &userID

	res := &dto.GetNotificationListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Notifications, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateNotification(notification *structs.Notifications) (*structs.Notifications, error) {
	res := &dto.GetNotificationResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.Notifications, notification, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) MarkNotificationRead(notificationID int) error {
	notification, err := repo.GetNotification(notificationID)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	notification.IsRead = true

	err = repo.UpdateNotification(notificationID, notification)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}
