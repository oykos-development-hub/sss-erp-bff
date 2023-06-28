package mutations

import "github.com/graphql-go/graphql"

var UserProfileExperienceInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileExperienceInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"relevant": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"organization_unit": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount_of_experience": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount_of_insured_experience": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_signature": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"reference_file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
