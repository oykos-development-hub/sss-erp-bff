package mutations

import "github.com/graphql-go/graphql"

var InvoiceMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InvoiceMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"passed_to_inventory": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"passed_to_accounting": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"invoice_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"type_of_subject": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type_of_contract": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"source_of_funding": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"supplier_title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"activity_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"tax_authority_codebook_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"order_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_invoice": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"receipt_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_payment": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"sss_invoice_receipt_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"pro_forma_invoice_file_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pro_forma_invoice_file_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pro_forma_invoice_file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"articles": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(InvoiceArticleMutation),
		},
		"additional_expenses": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(AdditionalExpenses),
		},
	},
})

var InvoiceArticleMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InvoiceArticleMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"net_price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"vat_price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"vat_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"cost_account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var AdditionalExpenses = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AdditionalExpensesMutation",
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
		"invoice_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
