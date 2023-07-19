package types

import "github.com/graphql-go/graphql"

var SettingsDropdownItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SettingsDropdownItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"abbreviation": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SettingsDropdownType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SettingsDropdownTypes",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(SettingsDropdownItemType),
		},
	},
})

var SettingsDropdownInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SettingsDropdownInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: SettingsDropdownItemType,
		},
	},
})

var SettingsDropdownDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SettingsDropdownDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
