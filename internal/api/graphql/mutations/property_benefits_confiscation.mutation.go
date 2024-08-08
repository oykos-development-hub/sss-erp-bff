package mutations

import "github.com/graphql-go/graphql"

var PropBenConfMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PropBenConfMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"property_benefits_confiscation_type": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"decision_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"decision_date": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"subject": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"jmbg": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"residence": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"payment_reference_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"debit_reference_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"execution_date": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"payment_deadline_date": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"court_costs": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"court_account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"property_benefits_confiscation_details": &graphql.InputObjectFieldConfig{
			Type: PropBenConfDetailsInputType,
		},
		"file": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
})

var PropBenConfDetailsInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PropBenConfDetailsInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"all_payments_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"amount_grace_period": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"amount_grace_period_due_date": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"amount_grace_period_available": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"left_to_pay_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"court_costs_paid": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"court_costs_left_to_pay_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
