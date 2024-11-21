package resolvers

import (
	"bff/config"
	"bff/internal/api/repository"
	"bff/internal/api/sse/notifications"
)

type Resolver struct {
	Config               *config.Config
	NotificationsService *notifications.NotificationService
	Repo                 repository.MicroserviceRepositoryInterface
}

func NewResolver(cfg *config.Config, notificationService *notifications.NotificationService, repo repository.MicroserviceRepositoryInterface) *Resolver {
	return &Resolver{
		Config:               cfg,
		NotificationsService: notificationService,
		Repo:                 repo,
	}
}
