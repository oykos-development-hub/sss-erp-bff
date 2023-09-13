package mutations

import "github.com/graphql-go/graphql"

var BasicInventoryDispatchMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BasicInventoryDispatchMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"source_user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"target_user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"target_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"source_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"office_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"dispatch_description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"inventory_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
