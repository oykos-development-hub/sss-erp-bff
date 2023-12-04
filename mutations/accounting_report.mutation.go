package mutations

import "github.com/graphql-go/graphql"

var OverallSpendingMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OverallSpendingMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"start_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"end_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"office_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"articles": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.String),
		},
		"exception": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
	},
})
