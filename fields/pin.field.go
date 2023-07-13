package fields

import (
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var PinField = &graphql.Field{
	Type:        types.PinType,
	Description: "Validates user pin",
	Args: graphql.FieldConfigArgument{
		"pin": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.PinResolver,
}
