package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"
	"github.com/graphql-go/graphql"
)

var UserProfileForeignerField = &graphql.Field{
	Type:        types.UserProfileForeignerType,
	Description: "Returns a data of User Profile for displaying inside Foreigner tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"user_account_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileForeignerResolver,
}

var UserProfileForeignerInsertField = &graphql.Field{
	Type:        types.UserProfileForeignerInsertType,
	Description: "Creates new or alter existing User Profile's Foreigner item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileForeignerInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileForeignerInsertResolver,
}

var UserProfileForeignerDeleteField = &graphql.Field{
	Type:        types.UserProfileForeignerDeleteType,
	Description: "Deletes existing User Profile's Foreigner",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileForeignerDeleteResolver,
}
