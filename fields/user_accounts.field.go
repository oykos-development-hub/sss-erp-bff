package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"
	"github.com/graphql-go/graphql"
)

var UserAccountField = &graphql.Field{
	Type:        types.UserAccountsOverviewType,
	Description: "Returns a data of User Accounts for displaying on Settings screen",
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
		"is_active": &graphql.ArgumentConfig{
			Type:         graphql.Boolean,
			DefaultValue: nil,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.UserAccountsOverviewResolver,
}

var UserAccountInsertField = &graphql.Field{
	Type:        types.UserAccountInsertType,
	Description: "Inserts a data of User Account for displaying on Settings screen",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserAccountInsertMutation),
		},
	},
	Resolve: resolvers.UserAccountBasicInsertResolver,
}

var UserAccountDeleteField = &graphql.Field{
	Type:        types.UserAccountDeleteType,
	Description: "Deletes existing User Account's data",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserAccountDeleteResolver,
}
