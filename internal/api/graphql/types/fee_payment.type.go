package types

import "github.com/graphql-go/graphql"

var FeePaymentOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeePaymentOverview",
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
			Type: graphql.NewList(FeePaymentType),
		},
	},
})

var FeePaymentInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeePaymentInsert",
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
			Type: FeePaymentType,
		},
	},
})

var FeePaymentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeePaymentType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"fee_id": &graphql.Field{
			Type: graphql.Int,
		},
		"payment_method": &graphql.Field{
			Type: graphql.String,
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
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var FeePaymentDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeePaymentDelete",
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
