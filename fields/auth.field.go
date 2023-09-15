package fields

import (
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var LoginField = &graphql.Field{
	Type:        types.LoginType,
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

var RefreshField = &graphql.Field{
	Type:        types.RefreshTokenType,
	Description: "Returns a basic data for logged in user",
	Resolve:     resolvers.RefreshTokenResolver,
}
