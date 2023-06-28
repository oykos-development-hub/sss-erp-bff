package mutations

import "github.com/graphql-go/graphql"

var UserAccountInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserAccountInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"role_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"secondary_email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"phone": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pin": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"verified_email": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"verified_phone": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"updated_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"folder_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
