package mutations

import "github.com/graphql-go/graphql"

var UserProfileAbsentInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileAbsentInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"vacation_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"location": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"target_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
