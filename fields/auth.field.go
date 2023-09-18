package fields

import (
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var LoginField = &graphql.Field{
	Type:        types.LogoutType,
	Description: "Returns a basic data for logged in user",
	Args: graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.LoginResolver,
}

var LogoutField = &graphql.Field{
	Type:        types.LoginType,
	Description: "Logout the user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.LogoutResolver,
}

var RefreshField = &graphql.Field{
	Type:        types.RefreshTokenType,
	Description: "Returns a basic data for logged in user",
	Resolve:     resolvers.RefreshTokenResolver,
}
