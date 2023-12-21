package resolvers

import (
	"bff/config"
	"bff/internal/api/repository"
	"bff/internal/api/websockets/notifications"
)

type Resolver struct {
	Config               *config.Config
	NotificationsService *notifications.Websockets
	Repo                 repository.MicroserviceRepositoryInterface
}

func NewResolver(cfg *config.Config, notificationService *notifications.Websockets, repo repository.MicroserviceRepositoryInterface) *Resolver {
	return &Resolver{
		Config:               cfg,
		NotificationsService: notificationService,
		Repo:                 repo,
	}
}
