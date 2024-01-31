package mutations

import "github.com/graphql-go/graphql"

var BudgetMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BudgetMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"year": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_type": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"limits": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(FinancialBudgetLimitMutation),
		},
	},
})

var FinancialBudgetLimitMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FinancialBudgetLimitMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"limit": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
