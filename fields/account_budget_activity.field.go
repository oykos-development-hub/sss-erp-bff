package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var AccountBudgetActivityOverviewField = &graphql.Field{
	Type:        types.AccountBudgetActivityOverviewType,
	Description: "Returns a data of AccountBudgetActivity items",
	Args: graphql.FieldConfigArgument{
		"budget_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"activity_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BudgetAccountOverviewResolver,
}

var AccountBudgetActivityInsertField = &graphql.Field{
	Type:        types.AccountBudgetActivityInsertType,
	Description: "Creates new or alter existing Account Budget Activity",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewList(mutations.BudgetAccountActivityMutation),
		},
	},
	Resolve: resolvers.BudgetAccountInsertResolver,
}
