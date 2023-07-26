package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var UserProfileResolutionField = &graphql.Field{
	Type:        types.UserProfileResolutionType,
	Description: "Returns a data of User Profile for displaying inside Resolution tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileResolutionResolver,
}

var UserProfileResolutionInsertField = &graphql.Field{
	Type:        types.UserProfileResolutionInsertType,
	Description: "Creates new or alter existing User Profile's Resolution item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileResolutionInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileResolutionInsertResolver,
}

var UserProfileResolutionDeleteField = &graphql.Field{
	Type:        types.UserProfileResolutionDeleteType,
	Description: "Deletes existing User Profile's Resolution",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileResolutionDeleteResolver,
}
