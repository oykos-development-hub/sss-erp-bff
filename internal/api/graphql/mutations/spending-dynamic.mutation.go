package mutations

import "github.com/graphql-go/graphql"

var SpendingDynamicMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SpendingDynamicMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"budget_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"january": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"february": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"march": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"april": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"may": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"june": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"july": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"august": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"september": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"october": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"november": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"december": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
