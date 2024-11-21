package fields

import (
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) NotificationOverview() *graphql.Field {
	return &graphql.Field{
		Type:        types.NotificationsGetType,
		Description: "Get notifications for authenticated user",
		Resolve:     f.Resolvers.NotificationOverviewResolver,
	}
}

func (f *Field) NotificationRead() *graphql.Field {
	return &graphql.Field{
		Type:        types.NotificationReadType,
		Description: "Mark notification as read",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.NotificationReadResolver,
	}
}

func (f *Field) NotificationDelete() *graphql.Field {
	return &graphql.Field{
		Type:        types.NotificationDeleteType,
		Description: "Delete notification by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.NotificationDeleteResolver,
	}
}
