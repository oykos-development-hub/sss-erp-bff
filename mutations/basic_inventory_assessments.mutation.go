package mutations

import "github.com/graphql-go/graphql"

var BasicInventoryAssessmentsMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BasicInventoryAssessmentsMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"inventory_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"depreciation_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"gross_price_new": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"gross_price_difference": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_assessment": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
