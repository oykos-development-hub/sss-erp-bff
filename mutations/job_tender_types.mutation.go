package mutations

import "github.com/graphql-go/graphql"

var JobTenderTypeInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "JobTenderTypeInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"is_judge": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"is_judge_president": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"color": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"icon": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
