package notifications

import (
	"bff/internal/api/repository"
	"bff/internal/api/sse"
	"bff/structs"
	"encoding/json"
	"errors"
)

type NotificationService struct {
	Repo repository.MicroserviceRepositoryInterface
	SSE  *sse.ServerSentEvent
}

func NewNotificationService(repo repository.MicroserviceRepositoryInterface, sse *sse.ServerSentEvent) *NotificationService {
	return &NotificationService{
		Repo: repo,
		SSE:  sse,
	}
}

type NotificationMessage struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

func (svc *NotificationService) sendNotificationToUser(userID int, notification *structs.Notifications) error {
	messageJSON, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	svc.SSE.Broadcast(userID, string(messageJSON))
	return nil
}

func (svc *NotificationService) CreateNotification(notification *structs.Notifications) (*structs.Notifications, error) {
	newNotification, err := svc.Repo.CreateNotification(notification)
	if err != nil {
		return nil, err
	}

	user, err := svc.Repo.GetUserAccountByID(notification.ToUserID)
	if err != nil || user == nil {
		return nil, errors.New("invalid user")
	}

	role, err := svc.Repo.GetRole(*user.RoleID)
	if err != nil || !role.Active {
		return nil, nil
	}

	return newNotification, svc.sendNotificationToUser(notification.ToUserID, newNotification)
}
