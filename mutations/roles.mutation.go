package mutations

import "github.com/graphql-go/graphql"

var RolesInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RolesInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
})
