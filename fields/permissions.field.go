package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var PermissionsForRoleField = &graphql.Field{
	Type:        types.PermissionsForRoleOverviewType,
	Description: "Returns permissions for role",
	Args: graphql.FieldConfigArgument{
		"role_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.PermissionsForRoleResolver,
}

var PermissionsUpdate = &graphql.Field{
	Type:        types.PermissionsForRoleOverviewType,
	Description: "Sync Permissions data",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.PermissionsInputMutation),
		},
	},
	Resolve: resolvers.PermissionsUpdateResolver,
}
