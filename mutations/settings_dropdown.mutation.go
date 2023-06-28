package mutations

import "github.com/graphql-go/graphql"

var SettingsDropdownInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SettingsDropdownInsertMutation",
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
		"color": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"icon": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
