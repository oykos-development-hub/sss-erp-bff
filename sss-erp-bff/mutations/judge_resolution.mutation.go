package mutations

import "github.com/graphql-go/graphql"

var JudgeResolutionItemMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "JudgeResolutionItemMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"number_of_judges": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"number_of_presidents": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var JudgeResolutionInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "JudgeResolutionInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(JudgeResolutionItemMutation),
		},
	},
})
