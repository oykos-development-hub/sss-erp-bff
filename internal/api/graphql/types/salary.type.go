package types

import "github.com/graphql-go/graphql"

var SalaryOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SalaryOverviewType",
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
			Type: graphql.NewList(SalaryType),
		},
	},
})

var SalaryInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SalaryInsertType",
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
			Type: SalaryType,
		},
	},
})

var SalaryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SalaryType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"subject": &graphql.Field{
			Type: graphql.String,
		},
		"judge": &graphql.Field{
			Type: DropdownItemType,
		},
		"case_number": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_recipiet": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_case": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_finality": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_enforceability": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_end": &graphql.Field{
			Type: graphql.String,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
