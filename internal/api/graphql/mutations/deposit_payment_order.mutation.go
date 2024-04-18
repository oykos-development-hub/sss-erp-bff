package mutations

import "github.com/graphql-go/graphql"

var DepositPaymentOrderMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "DepositPaymentOrderMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"case_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"subject_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"municipality_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"tax_authority_codebook_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"net_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"source_bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_payment": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"date_of_statement": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"id_of_statement": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"additional_expenses": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(DepositPaymentAdditionalExpenses),
		},
		"additional_expenses_for_paying": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(DepositPaymentAdditionalExpenses),
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var DepositPaymentAdditionalExpenses = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "DepositPaymentAdditionalExpenses",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"subject_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"source_bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"payment_order_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
