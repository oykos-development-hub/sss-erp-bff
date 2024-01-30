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
	},
})
