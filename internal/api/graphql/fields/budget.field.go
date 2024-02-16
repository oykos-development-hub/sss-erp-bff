package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) BudgetInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetInsertType,
		Description: "Creates new or alter existing Budget",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.BudgetMutation),
			},
		},
		Resolve: f.Resolvers.BudgetInsertResolver,
	}
}
func (f *Field) BudgetDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetDeleteType,
		Description: "Deleted Budget",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.BudgetDeleteResolver,
	}
}
func (f *Field) BudgetSendField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetSendType,
		Description: "Send Budget",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BudgetSendResolver,
	}
}

func (f *Field) BudgetOverviewField() *graphql.Field {
	return &graphql.Field{
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
				Type: graphql.Int,
			},
			"budget_type": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.BudgetOverviewResolver,
	}
}

func (f *Field) BudgetDetailsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetOverviewType,
		Description: "Returns a data of Budget item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BudgetOverviewResolver,
	}
}

func (f *Field) FinancialBudgetOverview() *graphql.Field {
	return &graphql.Field{
		Type:        types.AccountWithFilledDataOverviewType,
		Description: "Returns a data of Financial Budget",
		Args: graphql.FieldConfigArgument{
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.FinancialBudgetOverview,
	}
}

func (f *Field) FinancialBudgetFillField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FinancialBudgetFillType,
		Description: "Fill Financially Budget item",
		Args: graphql.FieldConfigArgument{
			"request_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.FinancialBudgetFillMutation),
			},
		},
		Resolve: f.Resolvers.FinancialBudgetFillResolver,
	}
}
