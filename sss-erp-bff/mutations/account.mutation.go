package mutations

import "github.com/graphql-go/graphql"

var AccountMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AccountMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"parent_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
