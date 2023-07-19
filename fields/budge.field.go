package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var BudgeInsertField = &graphql.Field{
	Type:        types.BudgetType,
	Description: "Creates new or alter existing Budge",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.BudgeMutation),
		},
	},
	Resolve: resolvers.BudgeInsertResolver,
}

var BudgeDeleteField = &graphql.Field{
	Type:        types.BudgetType,
	Description: "Deleted Budge",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BudgeDeleteResolver,
}

var BudgeOverviewField = &graphql.Field{
	Type:        types.BudgetType,
	Description: "Returns a data of Budge items",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BudgeOverviewResolver,
}
