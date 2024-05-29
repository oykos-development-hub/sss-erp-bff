package mutations

import "github.com/graphql-go/graphql"

var TaxAuthorityCodebookInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "TaxAuthorityCodebookInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"code": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pio_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"tax_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"release_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"release_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"tax_supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"pio_supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"pio_percentage_employer_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"pio_employer_supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"pio_percentage_employee_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"pio_employee_supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"unemployment_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"unemployment_supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"unemployment_employer_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"unemployment_employer_supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"unemployment_employee_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"unemployment_employee_supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"labor_fund": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"labor_fund_supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"previous_income_percentage_less_than_700": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"previous_income_percentage_less_than_1000": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"previous_income_percentage_more_than_1000": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"coefficient": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"coefficient_less_700": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"coefficient_less_1000": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"coefficient_more_1000": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"amount_less_700": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"amount_less_1000": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"amount_more_1000": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"include_subtax": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
	},
})
