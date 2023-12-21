package router

import (
	"bff/config"
	"bff/internal/api/graphql"
	"bff/internal/api/http/files"
	"bff/internal/api/middleware"
	"bff/internal/api/repository"
	"bff/internal/api/websockets"
	"bff/internal/api/websockets/notifications"

	"github.com/go-chi/chi/v5"
)

// Setup creates the main router with all the routes.
func Setup(cfg *config.Config) (*chi.Mux, error) {
	mainRouter := chi.NewRouter()

	wsManager := websockets.NewClientManager()
	repo := repository.NewMicroserviceRepository(cfg)
	notificationsService := notifications.NewWebsockets(wsManager, repo)
	middleware := middleware.NewMiddleware(repo, cfg)

	mainRouter.HandleFunc("/ws", notificationsService.Handler)

	fileHandler := files.NewHandler(repo, cfg)
	mainRouter.Mount("/files", files.SetupFileHandler(fileHandler, middleware))

	graphqlHandler, err := graphql.SetupGraphQLHandler(notificationsService, repo, middleware, cfg)
	if err != nil {
		return nil, err
	}
	mainRouter.Handle("/", *graphqlHandler)

	return mainRouter, nil
}
