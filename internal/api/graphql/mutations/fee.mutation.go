package mutations

import "github.com/graphql-go/graphql"

var FeeMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FeeMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"fee_type": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"fee_subcategory": &graphql.InputObjectFieldConfig{
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
		"court_account": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"fee_fee_details": &graphql.InputObjectFieldConfig{
			Type: FeeFeeDetailsInputType,
		},
		"file": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
})

var FeeFeeDetailsInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FeeFeeDetailsInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"fee_all_payments_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"fee_left_to_pay_amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
