package types

import "github.com/graphql-go/graphql"

var ContractItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ContractItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"contract_type_id": &graphql.Field{
			Type: graphql.Int,
		},
		"contract_type": &graphql.Field{
			Type: ContractTypeItemType,
		},
		"abbreviation": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"net_salary": &graphql.Field{
			Type: graphql.String,
		},
		"gross_salary": &graphql.Field{
			Type: graphql.String,
		},
		"bank_account": &graphql.Field{
			Type: graphql.String,
		},
		"bank_name": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_signature": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_eligibility": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_start": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_end": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})
