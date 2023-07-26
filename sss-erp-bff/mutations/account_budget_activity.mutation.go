package mutations

import "github.com/graphql-go/graphql"

var BudgetAccountActivityMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BudgetAccountActivityMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"activity_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"value_current_year": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"value_next_year": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"value_after_next_year": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
