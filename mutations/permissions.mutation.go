package mutations

import "github.com/graphql-go/graphql"

var PermissionsMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PermissionsMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"role_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"permission_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"create": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"read": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"update": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"delete": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
})

var PermissionsInputMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PermissionsInputMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"permissions_list": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(PermissionsMutation),
		},
		"role_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})
