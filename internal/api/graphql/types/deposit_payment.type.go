package types

import "github.com/graphql-go/graphql"

var DepositPaymentOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositPaymentOverviewType",
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
			Type: graphql.NewList(DepositPaymentType),
		},
	},
})

var DepositPaymentInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositPaymentInsertType",
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
			Type: DepositPaymentType,
		},
	},
})

var DepositPaymentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositPaymentType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"payer": &graphql.Field{
			Type: graphql.String,
		},
		"case_number": &graphql.Field{
			Type: graphql.String,
		},
		"party_name": &graphql.Field{
			Type: graphql.String,
		},
		"current_bank_account": &graphql.Field{
			Type: graphql.String,
		},
		"number_of_bank_statement": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_bank_statement": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"amount": &graphql.Field{
			Type: graphql.Float,
		},
		"main_bank_account": &graphql.Field{
			Type: graphql.Boolean,
		},
		"date_of_transfer_main_account": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var DepositPaymentAdditionalExpensesOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositPaymentAdditionalExpensesOverviewType",
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
			Type: graphql.NewList(AdditionalExpensesType),
		},
	},
})

var DepositPaymentAdditionalExpensesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositPaymentAdditionalExpensesType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"case_number": &graphql.Field{
			Type: graphql.String,
		},
		"subject": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"bank_account": &graphql.Field{
			Type: graphql.String,
		},
		"payment_order": &graphql.Field{
			Type: DropdownItemType,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
	},
})
