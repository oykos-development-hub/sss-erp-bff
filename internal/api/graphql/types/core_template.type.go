package types

import "github.com/graphql-go/graphql"

var TemplateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TemplateTypes",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(TemplateItemType),
		},
	},
})

var TemplateItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TemplateItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"template": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})

var TemplateInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TemplateInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: TemplateItemType,
		},
	},
})

var TemplateDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TemplateDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
