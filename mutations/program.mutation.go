package mutations

import "github.com/graphql-go/graphql"

var ProgramMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ProgramMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"parent_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
