package types

import "github.com/graphql-go/graphql"

var FlatRatePaymentOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlatRatePaymentOverview",
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
			Type: graphql.NewList(FlatRatePaymentType),
		},
	},
})

var FlatRatePaymentInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlatRatePaymentInsert",
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
			Type: FlatRatePaymentType,
		},
	},
})

var FlatRatePaymentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlatRatePaymentType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"flat_rate_id": &graphql.Field{
			Type: graphql.Int,
		},
		"payment_method": &graphql.Field{
			Type: DropdownItemType,
		},
		"amount": &graphql.Field{
			Type: graphql.Float,
		},
		"payment_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"payment_due_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"receipt_number": &graphql.Field{
			Type: graphql.String,
		},
		"payment_reference_number": &graphql.Field{
			Type: graphql.String,
		},
		"debit_reference_number": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: DropdownItemType,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var FlatRatePaymentDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FlatRatePaymentDelete",
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
	},
})
