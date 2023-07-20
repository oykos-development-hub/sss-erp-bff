package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var BudgetInsertField = &graphql.Field{
	Type:        types.BudgetInsertType,
	Description: "Creates new or alter existing Budget",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.BudgetMutation),
		},
	},
	Resolve: resolvers.BudgetInsertResolver,
}

var BudgetDeleteField = &graphql.Field{
	Type:        types.BudgetDeleteType,
	Description: "Deleted Budget",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BudgetDeleteResolver,
}

var BudgetOverviewField = &graphql.Field{
	Type:        types.BudgetOverviewType,
	Description: "Returns a data of Budget items",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"year": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"type_budget": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.BudgetOverviewResolver,
}
