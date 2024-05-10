package mutations

import "github.com/graphql-go/graphql"

var AccountingOrderForObligationsMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AccountingOrderForObligationsMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"date_of_booking": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"invoice_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"salary_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"payment_order_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"enforced_payment_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"return_enforced_payment_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

var AccountingEntryMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AccountingEntryMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_booking": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(AccountingEntryItemMutation),
		},
	},
})

var AccountingEntryItemMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AccountingEntryItemMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"invoice_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"salary_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"payment_order_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"enforced_payment_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"return_enforced_payment_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"credit_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"debit_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
