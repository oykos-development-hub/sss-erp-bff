package mutations

import "github.com/graphql-go/graphql"

var EnforcedPaymentMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "EnforcedPaymentMutation",
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
		"date_of_order": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"id_of_statement": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"sap_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_sap": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"amount_for_lawyer": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"amount_for_agent": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(PaymentItems),
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"return_file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"return_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"return_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})

var EnforcedPaymentItemMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "EnforcedPaymentItemMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"invoice_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
