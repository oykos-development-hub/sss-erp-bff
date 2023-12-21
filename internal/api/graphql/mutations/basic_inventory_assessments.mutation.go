package mutations

import "github.com/graphql-go/graphql"

var BasicInventoryAssessmentsMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BasicInventoryAssessmentsMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"inventory_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"estimated_duration": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"depreciation_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"purchase_gross_price": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"gross_price_new": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"gross_price_difference": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"residual_price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"date_of_assessment": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
