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
		"activity": &graphql.Field{
			Type: DropdownItemType,
		},
		"month": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_calculation": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"registred": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deletable": &graphql.Field{
			Type: graphql.Boolean,
		},
		"number_of_employees": &graphql.Field{
			Type: graphql.Int,
		},
		"obligations_price": &graphql.Field{
			Type: graphql.Float,
		},
		"gross_price": &graphql.Field{
			Type: graphql.Float,
		},
		"vat_price": &graphql.Field{
			Type: graphql.Float,
		},
		"net_price": &graphql.Field{
			Type: graphql.Float,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"salary_additional_expenses": &graphql.Field{
			Type: graphql.NewList(SalaryAdditionalExpenses),
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var SalaryAdditionalExpenses = graphql.NewObject(graphql.ObjectConfig{
	Name: "SalaryAdditionalExpenses",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"debtor": &graphql.Field{
			Type: DropdownItemType,
		},
		"activity": &graphql.Field{
			Type: DropdownItemType,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"amount": &graphql.Field{
			Type: graphql.Float,
		},
		"subject": &graphql.Field{
			Type: DropdownItemType,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"bank_account": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"identificator_number": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
