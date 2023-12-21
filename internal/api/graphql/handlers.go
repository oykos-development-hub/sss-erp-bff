package graphql

import (
	"bff/config"
	"bff/internal/api/graphql/schema"
	"bff/internal/api/middleware"
	"bff/internal/api/repository"
	"bff/internal/api/websockets/notifications"
	"net/http"

	"github.com/graphql-go/handler"
)

func SetupGraphQLHandler(notificationsService *notifications.Websockets, repo repository.MicroserviceRepositoryInterface, m *middleware.Middleware, cfg *config.Config) (*http.Handler, error) {
	schema, err := schema.SetupGraphQLSchema(notificationsService, repo, cfg)
	if err != nil {
		return nil, err
	}

	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	})

	graphqlHandler := m.ErrorHandlerMiddleware(
		m.GetCorsMiddleware(
			m.AuthMiddleware(
				m.AddResponseWriterToContext(
					m.RequestContextMiddleware(h),
				),
			),
		),
	)

	return &graphqlHandler, nil
}
