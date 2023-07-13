package mutations

import "github.com/graphql-go/graphql"

var UserProfileContractInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ContractInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"contract_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"net_salary": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"gross_salary": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"bank_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_signature": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_eligibility": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
