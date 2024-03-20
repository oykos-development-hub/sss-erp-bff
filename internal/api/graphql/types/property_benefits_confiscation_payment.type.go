package types

import "github.com/graphql-go/graphql"

var PropBenConfPaymentOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PropBenConfPaymentOverview",
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
			Type: graphql.NewList(PropBenConfPaymentType),
		},
	},
})

var PropBenConfPaymentInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PropBenConfPaymentInsert",
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
			Type: PropBenConfPaymentType,
		},
	},
})

var PropBenConfPaymentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PropBenConfPaymentType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"property_benefits_confiscation_id": &graphql.Field{
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

var PropBenConfPaymentDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PropBenConfPaymentDelete",
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
