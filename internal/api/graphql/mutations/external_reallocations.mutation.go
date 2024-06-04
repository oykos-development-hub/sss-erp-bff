package mutations

import "github.com/graphql-go/graphql"

var ExternalReallocationMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ExternalReallocationMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"source_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"destination_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_request": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_action_dest_org_unit": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_action_sss": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"requested_by": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"accepted_by": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"destination_org_unit_file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"sss_file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ExternalReallocationItemMutation),
		},
	},
})

var ExternalReallocationItemMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ExternalReallocationItemMutation",
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
