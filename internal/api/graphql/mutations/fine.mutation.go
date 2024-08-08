package mutations

import "github.com/graphql-go/graphql"

var FineMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FineMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"act_type": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"decision_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"decision_date": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"subject": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
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
		"fine_fee_details": &graphql.InputObjectFieldConfig{
			Type: FineFeeDetailsInputType, // Assuming you have defined this GraphQL type
		},
		"file": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
})

var FineFeeDetailsInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FineFeeDetailsInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"fee_all_payments_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"fee_amount_grace_period": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"fee_amount_grace_period_due_date": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"fee_amount_grace_period_available": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"fee_left_to_pay_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"fee_court_costs_paid": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"fee_court_costs_left_to_pay_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
