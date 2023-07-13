package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"
	"github.com/graphql-go/graphql"
)

var UserProfileAbsentField = &graphql.Field{
	Type:        types.UserProfileAbsentType,
	Description: "Returns a data of User Profile for displaying inside Absent tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"user_account_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileAbsentResolver,
}

var UserProfileAbsentInsertField = &graphql.Field{
	Type:        types.UserProfileAbsentInsertType,
	Description: "Creates new or alter existing User Profile's Absent item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileAbsentInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileAbsentInsertResolver,
}

var UserProfileAbsentDeleteField = &graphql.Field{
	Type:        types.UserProfileAbsentDeleteType,
	Description: "Deletes existing User Profile's Absent",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"vacation_type_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileAbsentDeleteResolver,
}
