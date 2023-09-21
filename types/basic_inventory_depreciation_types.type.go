package types

import "github.com/graphql-go/graphql"

var BasicInventoryDepreciationTypesItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDepreciationTypesItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"abbreviation": &graphql.Field{
			Type: graphql.String,
		},
		"lifetime_in_months": &graphql.Field{
			Type: graphql.Int,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var BasicInventoryDepreciationTypesOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDepreciationTypesOverview",
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
			Type: graphql.NewList(BasicInventoryDepreciationTypesItemType),
		},
	},
})

var BasicInventoryDepreciationTypesInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDepreciationTypesInsert",
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
			Type: BasicInventoryDepreciationTypesItemType,
		},
	},
})

var BasicInventoryDepreciationTypesDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDepreciationTypesDelete",
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
