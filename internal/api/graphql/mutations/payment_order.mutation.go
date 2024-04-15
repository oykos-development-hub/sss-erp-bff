package mutations

import "github.com/graphql-go/graphql"

var PaymentOrderMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PaymentOrderMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_payment": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"date_of_sap": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"id_of_statement": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"sap_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(PaymentItems),
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var PaymentItems = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PaymentItems",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"invoice_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"additional_expense_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"salary_additional_expense_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
