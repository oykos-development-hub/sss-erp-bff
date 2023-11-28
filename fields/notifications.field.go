package fields

import (
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var NotificationsOverviewField = &graphql.Field{
	Type:        types.NotificationsOverviewType,
	Description: "Returns a data of Notification items",
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.NotificationsOverviewResolver,
}

var NotificationsReadField = &graphql.Field{
	Type:        types.NotificationsInsertType,
	Description: "Read or unread notification item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"is_read": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
	Resolve: resolvers.NotificationsReadResolver,
}

var NotificationsDeleteField = &graphql.Field{
	Type:        types.NotificationsDeleteType,
	Description: "Deletes existing Notification item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.NotificationsDeleteResolver,
}
