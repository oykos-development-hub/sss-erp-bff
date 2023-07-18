package mutations

import "github.com/graphql-go/graphql"

var JudgeNormInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "JudgeNormInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"area": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"norm": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"percentage_of_norm_decrease": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"number_of_items": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"number_of_solved_items": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"start_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"end_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"evaluation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"evaluation_valid_to": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"relocation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
