package types

import "github.com/graphql-go/graphql"

var BasicInventoryAssessmentsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryAssessmentsItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"inventory_id": &graphql.Field{
			Type: graphql.Int,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"estimated_duration": &graphql.Field{
			Type: graphql.Int,
		},
		"depreciation_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"gross_price_new": &graphql.Field{
			Type: graphql.Float,
		},
		"gross_price_difference": &graphql.Field{
			Type: graphql.Float,
		},
		"date_of_assessment": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var BasicInventoryAssessmentsInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryAssessmentsInsert",
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
			Type: BasicInventoryAssessmentsItemType,
		},
	},
})

var BasicInventoryAssessmentsDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryAssessmentsDelete",
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
