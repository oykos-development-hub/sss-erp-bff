package types

import "github.com/graphql-go/graphql"

var DropdownItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DropdownItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var DropdownItemWithValueType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DropdownWithValueItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"value": &graphql.Field{
			Type: graphql.String,
		},
	},
})
