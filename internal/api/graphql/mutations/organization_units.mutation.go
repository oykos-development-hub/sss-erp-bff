package mutations

import "github.com/graphql-go/graphql"

var OrganizationUnitInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OrganizationUnitInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"parent_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"number_of_judges": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"order_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"city": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pib": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"address": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"color": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"icon": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"code": &graphql.InputObjectFieldConfig{
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
