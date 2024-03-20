package mutations

import "github.com/graphql-go/graphql"

var PropBenConfPaymentMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PropBenConfPaymentMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"property_benefits_confiscation_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"payment_method": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"payment_date": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"payment_due_date": &graphql.InputObjectFieldConfig{
			Type: graphql.DateTime,
		},
		"receipt_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"payment_reference_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"debit_reference_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
