package types

import "github.com/graphql-go/graphql"

var SuppliersItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SuppliersItem",
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
		"official_id": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"folder_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var SuppliersOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SuppliersOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(SuppliersItemType),
		},
	},
})

var SuppliersInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SuppliersInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(SuppliersItemType),
		},
	},
})

var SuppliersDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SuppliersDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
