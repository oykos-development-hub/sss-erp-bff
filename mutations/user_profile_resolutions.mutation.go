package mutations

import "github.com/graphql-go/graphql"

var UserProfileResolutionInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileResolutionInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"resolution_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"resolution_purpose": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
