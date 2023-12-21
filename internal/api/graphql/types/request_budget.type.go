package types

import "github.com/graphql-go/graphql"

var RequestBudgetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RequestBudgetType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"date_create": &graphql.Field{
			Type: graphql.String,
		},
		"activity": &graphql.Field{
			Type: DropdownItemType,
		},
		"budget": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})
