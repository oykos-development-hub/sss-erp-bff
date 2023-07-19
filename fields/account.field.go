package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var AccountInsertField = &graphql.Field{
	Type:        types.AccountInsertType,
	Description: "Creates new or alter existing Account",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.AccountMutation),
		},
	},
	Resolve: resolvers.AccountInsertResolver,
}

var AccountDeleteField = &graphql.Field{
	Type:        types.AccountDeleteType,
	Description: "Deleted Account",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.AccountDeleteResolver,
}

var AccountOverviewField = &graphql.Field{
	Type:        types.AccountOverviewType,
	Description: "Returns a data of Account items",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"tree": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: resolvers.AccountOverviewResolver,
}
