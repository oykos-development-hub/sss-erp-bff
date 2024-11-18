package mutations

import "github.com/graphql-go/graphql"

var UserProfileEvaluationInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileEvaluationInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"evaluation_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_evaluation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"score": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"evaluator": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"is_relevant": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"reason_for_evaluation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"evaluation_period": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"decision_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_ids": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
})
