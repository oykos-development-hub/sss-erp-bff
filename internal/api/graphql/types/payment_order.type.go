package types

import "github.com/graphql-go/graphql"

var PaymentOrderOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentOrderOverviewType",
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
			Type: graphql.NewList(PaymentOrderType),
		},
	},
})

var PaymentOrderInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentOrderInsertType",
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
			Type: PaymentOrderType,
		},
	},
})

var PaymentOrderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentOrderType",
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
			Type: graphql.String,
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
		"items": &graphql.Field{
			Type: graphql.NewList(PaymentOrderItemsType),
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var PaymentOrderItemsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentOrderItemsType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"invoice_id": &graphql.Field{
			Type: graphql.Int,
		},
		"additional_expenses_id": &graphql.Field{
			Type: graphql.Int,
		},
		"salary_additional_expenses_id": &graphql.Field{
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
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var ObligationsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ObligationsOverviewType",
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
		"items": &graphql.Field{
			Type: graphql.NewList(ObligationsType),
		},
	},
})

var ObligationsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ObligationsType",
	Fields: graphql.Fields{
		"invoice_id": &graphql.Field{
			Type: graphql.Int,
		},
		"additional_expense_id": &graphql.Field{
			Type: graphql.Int,
		},
		"salary_additional_expense_id": &graphql.Field{
			Type: graphql.Int,
		},
		"total_price": &graphql.Field{
			Type: graphql.Float,
		},
		"remain_price": &graphql.Field{
			Type: graphql.Float,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})
