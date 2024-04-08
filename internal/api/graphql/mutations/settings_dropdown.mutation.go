package mutations

import "github.com/graphql-go/graphql"

var SettingsDropdownInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SettingsDropdownInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"parent_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"entity": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"color": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"icon": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
