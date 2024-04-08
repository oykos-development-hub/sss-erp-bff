package mutations

import "github.com/graphql-go/graphql"

var SalaryMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SalaryMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"activity_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"month": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_calculation": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"salary_additional_expenses": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(SalaryAdditionalExpensesMutation),
		},
	},
})

var SalaryAdditionalExpensesMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SalaryAdditionalExpensesMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"salary_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"subject_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
