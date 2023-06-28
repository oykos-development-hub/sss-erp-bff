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
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.LoginResolver,
}
