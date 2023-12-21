package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) AccountBudgetActivityOverviewField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.BudgetAccountOverviewResolver,
	}
}
func (f *Field) AccountBudgetActivityInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.AccountBudgetActivityInsertType,
		Description: "Creates new or alter existing Account Budget Activity",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.BudgetAccountActivityMutation),
			},
		},
		Resolve: f.Resolvers.BudgetAccountInsertResolver,
	}
}
