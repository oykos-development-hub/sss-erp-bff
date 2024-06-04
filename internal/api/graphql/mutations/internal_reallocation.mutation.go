package mutations

import "github.com/graphql-go/graphql"

var InternalReallocationMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InternalReallocationMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_request": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"request_by": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(InternalReallocationItemMutation),
		},
	},
})

var InternalReallocationItemMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InternalReallocationItemMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"reallocation_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"source_account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"destination_account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
