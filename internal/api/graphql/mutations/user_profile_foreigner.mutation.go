package mutations

import "github.com/graphql-go/graphql"

var UserProfileForeignerInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileForeignerInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"work_permit_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"work_permit_issuer": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"work_permit_date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"work_permit_date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"work_permit_indefinite_length": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"residence_permit_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"residence_permit_date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"residence_permit_indefinite_length": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"country_of_origin": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"work_permit_file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"residence_permit_file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_ids": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
})
