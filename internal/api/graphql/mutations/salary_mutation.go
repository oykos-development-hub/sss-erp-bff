package mutations

import "github.com/graphql-go/graphql"

var SalaryMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SalaryMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"judge_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"subject": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"case_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_recipiet": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_case": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_finality": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_enforceability": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var SalaryItemMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SalaryItemMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"deposit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"category_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"unit": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"currency": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_confiscation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"case_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"judge_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
