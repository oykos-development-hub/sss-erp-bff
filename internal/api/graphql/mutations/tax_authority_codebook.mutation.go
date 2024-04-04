package mutations

import "github.com/graphql-go/graphql"

var TaxAuthorityCodebookInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "TaxAuthorityCodebookInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"code": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pio_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"tax_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"release_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"previous_income_percentage_less_than_700": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"previous_income_percentage_less_than_1000": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"previous_income_percentage_more_than_1000": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
