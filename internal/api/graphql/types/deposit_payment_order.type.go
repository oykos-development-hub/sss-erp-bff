package types

import "github.com/graphql-go/graphql"

var DepositPaymentOrderOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositPaymentOrderOverviewType",
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
			Type: graphql.NewList(DepositPaymentOrderType),
		},
	},
})

var DepositPaymentOrderInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositPaymentOrderInsertType",
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
			Type: DepositPaymentOrderType,
		},
	},
})

var DepositPaymentOrderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DepositPaymentOrderType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"case_number": &graphql.Field{
			Type: graphql.String,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"net_amount": &graphql.Field{
			Type: graphql.Float,
		},
		"bank_account": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_payment": &graphql.Field{
			Type: graphql.DateTime,
		},
		"date_of_statement": &graphql.Field{
			Type: graphql.DateTime,
		},
		"id_of_statement": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"additional_expenses": &graphql.Field{
			Type: graphql.NewList(DepositPaymentAdditionalExpensesType),
		},
		"additional_expenses_for_paying": &graphql.Field{
			Type: graphql.NewList(DepositPaymentAdditionalExpensesType),
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
