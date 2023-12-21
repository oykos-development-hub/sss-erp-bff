package mutations

import "github.com/graphql-go/graphql"

var JudgeNormInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "JudgeNormInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"topic": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"number_of_norm_decrease": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"number_of_items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"number_of_items_solved": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"evaluation_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_evaluation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_evaluation_validity": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"relocation_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
