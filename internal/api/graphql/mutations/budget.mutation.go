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

var FinancialBudgetFillMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FinancialBudgetFillMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_request_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"current_year": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"next_year": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"year_after_next": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
