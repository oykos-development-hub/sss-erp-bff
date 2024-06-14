package types

import "github.com/graphql-go/graphql"

var EnforcedPaymentOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EnforcedPaymentOverviewType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(EnforcedPaymentType),
		},
	},
})

var EnforcedPaymentInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EnforcedPaymentInsertType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: EnforcedPaymentType,
		},
	},
})

var EnforcedPaymentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EnforcedPaymentType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"amount": &graphql.Field{
			Type: graphql.Float,
		},
		"amount_for_lawyer": &graphql.Field{
			Type: graphql.Float,
		},
		"amount_for_agent": &graphql.Field{
			Type: graphql.Float,
		},
		"amount_for_bank": &graphql.Field{
			Type: graphql.Float,
		},
		"account_for_expenses": &graphql.Field{
			Type: graphql.Int,
		},
		"execution_number": &graphql.Field{
			Type: graphql.String,
		},
		"agent": &graphql.Field{
			Type: DropdownItemType,
		},
		"bank_account": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_order": &graphql.Field{
			Type: graphql.DateTime,
		},
		"date_of_payment": &graphql.Field{
			Type: graphql.DateTime,
		},
		"date_of_sap": &graphql.Field{
			Type: graphql.DateTime,
		},
		"id_of_statement": &graphql.Field{
			Type: graphql.Int,
		},
		"return_date": &graphql.Field{
			Type: graphql.String,
		},
		"return_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"sap_id": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"return_file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(EnforcedPaymentItemsType),
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var EnforcedPaymentItemsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EnforcedPaymentItemsType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"invoice_id": &graphql.Field{
			Type: graphql.Int,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"amount": &graphql.Field{
			Type: graphql.Float,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})
