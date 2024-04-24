package types

import "github.com/graphql-go/graphql"

var ObligationsForAccountingOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ObligationsForAccountingOverviewType",
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
			Type: graphql.NewList(ObligationsForAccountingType),
		},
	},
})

var ObligationsForAccountingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ObligationsForAccountingType",
	Fields: graphql.Fields{
		"invoice_id": &graphql.Field{
			Type: graphql.Int,
		},
		"salary_id": &graphql.Field{
			Type: graphql.Int,
		},
		"price": &graphql.Field{
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
