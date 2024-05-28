package mutations

import "github.com/graphql-go/graphql"

var SpendingReleaseMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SpendingReleaseMutation",
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
		"month": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
