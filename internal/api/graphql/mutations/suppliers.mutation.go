package mutations

import "github.com/graphql-go/graphql"

var SuppliersInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SuppliersInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"entity": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"official_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"address": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"folder_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"bank_accounts": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.String),
		},
	},
})
