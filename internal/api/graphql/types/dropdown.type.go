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

var FileDropdownItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FileDropdownItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var DropdownBudgetIndentItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DropdownBudgetIndentItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
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

var DropdownOUItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DropdownItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
	},
})
