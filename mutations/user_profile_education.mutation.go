package mutations

import "github.com/graphql-go/graphql"

var UserProfileEducationInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileEducationInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_certification": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"academic_title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"expertise_level": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"score": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"certificate_issuer": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
