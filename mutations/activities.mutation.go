package mutations

import "github.com/graphql-go/graphql"

var ActivitiesMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ActivitiesMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"subroutine_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
