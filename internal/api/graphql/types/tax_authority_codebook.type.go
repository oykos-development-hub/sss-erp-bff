package types

import "github.com/graphql-go/graphql"

var TaxAuthorityCodebooksType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaxAuthorityCodebooks",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"tax_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"tax_supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"pio_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"release_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"previous_income_percentage_less_than_700": &graphql.Field{
			Type: graphql.Float,
		},
		"previous_income_percentage_less_than_1000": &graphql.Field{
			Type: graphql.Float,
		},
		"previous_income_percentage_more_than_1000": &graphql.Field{
			Type: graphql.Float,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"pio_supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"pio_percentage_employer_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"pio_employer_supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"pio_percentage_employee_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"pio_employee_supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"unemployment_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"unemployment_supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"unemployment_employer_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"unemployment_employer_supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"unemployment_employee_percentage": &graphql.Field{
			Type: graphql.Float,
		},
		"unemployment_employee_supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"labor_fund": &graphql.Field{
			Type: graphql.Float,
		},
		"labor_fund_supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"coefficient": &graphql.Field{
			Type: graphql.Float,
		},
		"coefficient_less_700": &graphql.Field{
			Type: graphql.Float,
		},
		"coefficient_less_1000": &graphql.Field{
			Type: graphql.Float,
		},
		"coefficient_more_1000": &graphql.Field{
			Type: graphql.Float,
		},
		"amount_less_700": &graphql.Field{
			Type: graphql.Float,
		},
		"amount_less_1000": &graphql.Field{
			Type: graphql.Float,
		},
		"amount_more_1000": &graphql.Field{
			Type: graphql.Float,
		},
		"include_subtax": &graphql.Field{
			Type: graphql.Boolean,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var TaxAuthorityCodebooksOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaxAuthorityCodebooksOverview",
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
			Type: graphql.NewList(TaxAuthorityCodebooksType),
		},
	},
})

var TaxAuthorityCodebooksInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaxAuthorityCodebooksInsert",
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
			Type: TaxAuthorityCodebooksType,
		},
	},
})

var TaxAuthorityCodebooksDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaxAuthorityCodebooksDelete",
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
