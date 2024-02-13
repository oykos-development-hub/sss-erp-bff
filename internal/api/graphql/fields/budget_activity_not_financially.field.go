package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) NonFinancialBudgetOverviewType() *graphql.Field {
	return &graphql.Field{
		Type:        types.NonFinancialBudgetOverviewType,
		Description: "Returns a data of Not Financially Budget Activity item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.NonFinancialBudgetOverviewResolver,
	}
}

func (f *Field) NonFinancialBudgetInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.NonFinancialBudgetInsertType,
		Description: "Insert Non Financially Budget item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.BudgetActivityNotFinanciallyInsertMutation),
			},
		},
		Resolve: f.Resolvers.BudgetActivityNotFinanciallyInsertResolver,
	}
}

func (f *Field) NonFinacialBudgetGoalInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.NonFinancialBudgetGoalInsertType,
		Description: "Insert Not Financially Goals item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.NonFinancialGoalInsertMutation),
			},
		},
		Resolve: f.Resolvers.NonFinancialGoalInsertResolver,
	}
}

func (f *Field) NonFinacialGoalIndicatorInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.NonFinancialGoalIndicatorInsertType,
		Description: "Insert Not Financially Goal Indicator item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.NonFinancialGoalIndicatorInsertMutation),
			},
		},
		Resolve: f.Resolvers.NonFinancialGoalIndicatorInsertResolver,
	}
}

func (f *Field) CheckBudgetActivityNotFinanciallyIsDoneField() *graphql.Field {
	return &graphql.Field{
		Type:        types.CheckBudgetActivityNotFinanciallyIsDoneType,
		Description: "Returns a data of Inductor items",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.CheckBudgetActivityNotFinanciallyIsDoneResolver,
	}
}
