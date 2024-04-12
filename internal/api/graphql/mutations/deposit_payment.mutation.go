package mutations

import "github.com/graphql-go/graphql"

var DepositPaymentMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "DepositPaymentMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"payer": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"case_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"party_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"number_of_bank_statement": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_bank_statement": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"main_bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"date_of_transfer_main_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"updated_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
