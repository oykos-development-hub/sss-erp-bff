package mutations

import "github.com/graphql-go/graphql"

var ModelsOfAccountingMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ModelsOfAccountingMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"items": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ModelOfAccountingItems),
		},
	},
})

var ModelOfAccountingItems = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ModelOfAccountingItems",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"debit_account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"credit_account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
