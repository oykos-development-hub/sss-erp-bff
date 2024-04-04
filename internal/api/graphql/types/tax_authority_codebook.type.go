package types

import "github.com/graphql-go/graphql"

var TaxAuthorityCodebooksType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaxAuthorityCodebooks",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"tax_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"pio_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"release_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"previous_income_percentage_less_than_700": &graphql.Field{
			Type: graphql.Float,
		},
		"previous_income_percentage_less_than_1000": &graphql.Field{
			Type: graphql.Float,
		},
		"previous_income_percentage_more_than_1000": &graphql.Field{
			Type: graphql.Float,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var TaxAuthorityCodebooksOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaxAuthorityCodebooksOverview",
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
			Type: graphql.NewList(TaxAuthorityCodebooksType),
		},
	},
})

var TaxAuthorityCodebooksInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaxAuthorityCodebooksInsert",
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
			Type: TaxAuthorityCodebooksType,
		},
	},
})

var TaxAuthorityCodebooksDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaxAuthorityCodebooksDelete",
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
