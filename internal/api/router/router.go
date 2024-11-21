package router

import (
	"bff/config"
	"bff/internal/api/graphql"
	"bff/internal/api/http/files"
	"bff/internal/api/middleware"
	"bff/internal/api/repository"
	"bff/internal/api/sse"
	"bff/internal/api/sse/notifications"

	"github.com/go-chi/chi/v5"
)

// Setup creates the main router with all the routes.
func Setup(cfg *config.Config) (*chi.Mux, error) {
	mainRouter := chi.NewRouter()

	repo := repository.NewMicroserviceRepository(cfg)

	middleware := middleware.NewMiddleware(repo, cfg)

	sseManager := sse.NewServerSentEvent()
	mainRouter.Mount("/sse", sseManager.Router(middleware))

	fileHandler := files.NewHandler(repo, cfg)
	mainRouter.Mount("/files", files.SetupFileHandler(fileHandler, middleware))

	notificationService := notifications.NewNotificationService(repo, sseManager)
	graphqlHandler, err := graphql.SetupGraphQLHandler(notificationService, repo, middleware, cfg)
	if err != nil {
		return nil, err
	}
	mainRouter.Handle("/", *graphqlHandler)

	return mainRouter, nil
}
