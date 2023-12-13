package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var RoleDetailsField = &graphql.Field{
	Type:        types.RoleDetailsType,
	Description: "Returns details for role",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.RoleDetailsResolver,
}

var RoleOverviewField = &graphql.Field{
	Type:        types.RoleOverviewType,
	Description: "Returns roles",
	Args:        graphql.FieldConfigArgument{},
	Resolve:     resolvers.RoleOverviewResolver,
}

var RoleInsertField = &graphql.Field{
	Type:        types.RoleInsertType,
	Description: "Inserts a data of Role",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.RolesInsertMutation),
		},
	},
	Resolve: resolvers.RolesInsertResolver,
}
