package mutations

import "github.com/graphql-go/graphql"

var BudgeMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BudgeMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"activity_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"source": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
