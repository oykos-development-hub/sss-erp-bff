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

func (f *Field) BudgetSendOnReviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetSendOnReviewType,
		Description: "Send Budget",
		Args: graphql.FieldConfigArgument{
			"request_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BudgetSendOnReviewResolver,
	}
}

func (f *Field) BudgetRequestAcceptField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetRequestAcceptType,
		Description: "Official accepts request",
		Args: graphql.FieldConfigArgument{
			"request_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BudgetRequestAcceptResolver,
	}
}

func (f *Field) BudgetRequestRejectField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetRequestRejectType,
		Description: "Official accepts request",
		Args: graphql.FieldConfigArgument{
			"request_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"comment": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: f.Resolvers.BudgetRequestRejectResolver,
	}
}

func (f *Field) BudgetRequestsDetailsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetRequestsDetailsType,
		Description: "Send Budget",
		Args: graphql.FieldConfigArgument{
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"unit_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BudgetRequestsDetailsResolver,
	}
}

func (f *Field) FinancialBudgetSummary() *graphql.Field {
	return &graphql.Field{
		Type:        types.FinancialBudgetSummaryType,
		Description: "Send Budget",
		Args: graphql.FieldConfigArgument{
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"unit_id": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: 0,
			},
		},
		Resolve: f.Resolvers.FinancialBudgetSummary,
	}
}

func (f *Field) BudgetRequestsOfficialField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OfficialBudgetRequestOverviewType,
		Description: "Budget requests overview",
		Args: graphql.FieldConfigArgument{
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BudgetRequestsOfficialResolver,
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
				Type: graphql.Int,
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

func (f *Field) FilledFinancialBudgetOverview() *graphql.Field {
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

func (f *Field) FinancialBudgetDetails() *graphql.Field {
	return &graphql.Field{
		Type:        types.FinancialBudgetDetailsType,
		Description: "Returns a data of Financial Budget",
		Args: graphql.FieldConfigArgument{
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.FinancialBudgetDetails,
	}
}

func (f *Field) FinancialBudgetFillField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FinancialBudgetFillType,
		Description: "Fill Financially Budget item",
		Args: graphql.FieldConfigArgument{
			"request_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.FinancialBudgetFillMutation),
			},
			"comment": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
		},
		Resolve: f.Resolvers.FinancialBudgetFillResolver,
	}
}

func (f *Field) FinancialBudgetFillActualField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FinancialBudgetFillType,
		Description: "Fill Financially Budget item",
		Args: graphql.FieldConfigArgument{
			"request_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.FinancialBudgetFillActualMutation),
			},
		},
		Resolve: f.Resolvers.FinancialBudgetFillActualResolver,
	}
}

func (f *Field) FinancialBudgetVersionUpdateField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FinancialBudgetDetailsType,
		Description: "Updates version of financial budget",
		Args: graphql.FieldConfigArgument{
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.FinancialBudgetVersionUpdate,
	}
}

func (f *Field) SpendingDynamicOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SpendingDynamicType,
		Description: "Spending dynamic overview",
		Args: graphql.FieldConfigArgument{
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"unit_id": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: 0,
			},
			"history": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
		},
		Resolve: f.Resolvers.SpendingDynamicOverview,
	}
}
