package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) BudgetActivityNotFinanciallyOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetActivityNotFinanciallyOverviewType,
		Description: "Returns a data of Not Financially Budget Activity item",
		Args: graphql.FieldConfigArgument{
			"request_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BudgetActivityNotFinanciallyOverviewResolver,
	}
}
func (f *Field) BudgetActivityNotFinanciallyInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BudgetActivityNotFinanciallyOverviewType,
		Description: "Insert Not Financially Budget Activity item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.BudgetActivityNotFinanciallyInsertMutation),
			},
		},
		Resolve: f.Resolvers.BudgetActivityNotFinanciallyInsertResolver,
	}
}
func (f *Field) ProgramNotFinanciallyInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProgramNotFinanciallyOverviewType,
		Description: "Returns a data of Not Financially Program item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.ProgramNotFinanciallyInsertMutation),
			},
		},
		Resolve: f.Resolvers.BudgetActivityNotFinanciallyProgramInsertResolver,
	}
}
func (f *Field) GoalsNotFinanciallyInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.GoalsNotFinanciallyOverviewType,
		Description: "Insert Not Financially Goals item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.GoalsNotFinanciallyInsertMutation),
			},
		},
		Resolve: f.Resolvers.BudgetActivityNotFinanciallyGoalsInsertResolver,
	}
}
func (f *Field) InductorNotFinanciallyOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.IndicatorNotFinanciallyOverviewType,
		Description: "Returns a data of Inductor items",
		Args: graphql.FieldConfigArgument{
			"goals_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BudgetActivityNotFinanciallyInductorResolver,
	}
}
func (f *Field) InductorNotFinanciallyInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.IndicatorNotFinanciallyInsertType,
		Description: "Insert Inductor item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.IndicatorNotFinanciallyInsertMutation),
			},
		},
		Resolve: f.Resolvers.BudgetActivityNotFinanciallyInductorInsertResolver,
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
