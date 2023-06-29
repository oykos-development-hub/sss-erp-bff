package types

import "github.com/graphql-go/graphql"

var BasicInventoryRealEstatesItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryRealEstatesItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"square_area": &graphql.Field{
			Type: graphql.Int,
		},
		"land_serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"estate_serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"ownership_type": &graphql.Field{
			Type: graphql.String,
		},
		"ownership_scope": &graphql.Field{
			Type: graphql.String,
		},
		"ownership_investment_scope": &graphql.Field{
			Type: graphql.String,
		},
		"limitations_description": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var BasicInventoryRealEstatesOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryRealEstatesOverview",
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
			Type: graphql.NewList(BasicInventoryRealEstatesItemType),
		},
	},
})
