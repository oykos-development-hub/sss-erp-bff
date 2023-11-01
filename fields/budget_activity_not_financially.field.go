package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var BudgetActivityNotFinanciallyOverviewField = &graphql.Field{
	Type:        types.BudgetActivityNotFinanciallyOverviewType,
	Description: "Returns a data of Not Financially Budget Activity item",
	Args: graphql.FieldConfigArgument{
		"request_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BudgetActivityNotFinanciallyOverviewResolver,
}

var BudgetActivityNotFinanciallyInsertField = &graphql.Field{
	Type:        types.BudgetActivityNotFinanciallyOverviewType,
	Description: "Insert Not Financially Budget Activity item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.BudgetActivityNotFinanciallyInsertMutation),
		},
	},
	Resolve: resolvers.BudgetActivityNotFinanciallyInsertResolver,
}

var ProgramNotFinanciallyInsertField = &graphql.Field{
	Type:        types.ProgramNotFinanciallyOverviewType,
	Description: "Returns a data of Not Financially Program item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.ProgramNotFinanciallyInsertMutation),
		},
	},
	Resolve: resolvers.BudgetActivityNotFinanciallyProgramInsertResolver,
}

var GoalsNotFinanciallyInsertField = &graphql.Field{
	Type:        types.GoalsNotFinanciallyOverviewType,
	Description: "Insert Not Financially Goals item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.GoalsNotFinanciallyInsertMutation),
		},
	},
	Resolve: resolvers.BudgetActivityNotFinanciallyGoalsInsertResolver,
}

var InductorNotFinanciallyOverviewField = &graphql.Field{
	Type:        types.IndicatorNotFinanciallyOverviewType,
	Description: "Returns a data of Inductor items",
	Args: graphql.FieldConfigArgument{
		"goals_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BudgetActivityNotFinanciallyInductorResolver,
}

var InductorNotFinanciallyInsertField = &graphql.Field{
	Type:        types.IndicatorNotFinanciallyInsertType,
	Description: "Insert Inductor item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.IndicatorNotFinanciallyInsertMutation),
		},
	},
	Resolve: resolvers.BudgetActivityNotFinanciallyInductorInsertResolver,
}

var CheckBudgetActivityNotFinanciallyIsDoneField = &graphql.Field{
	Type:        types.CheckBudgetActivityNotFinanciallyIsDoneType,
	Description: "Returns a data of Inductor items",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.CheckBudgetActivityNotFinanciallyIsDoneResolver,
}
