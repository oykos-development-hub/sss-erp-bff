package mutations

import "github.com/graphql-go/graphql"

var AccountingOrderForObligationsMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AccountingOrderForObligationsMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_booking": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"invoice_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"salary_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
})
