package types

import "github.com/graphql-go/graphql"

var BudgetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"year": &graphql.Field{
			Type: graphql.String,
		},
		"activity": &graphql.Field{
			Type: DropdownItemType,
		},
		"source": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
	},
})
