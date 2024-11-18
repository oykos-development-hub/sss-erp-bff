package types

import "github.com/graphql-go/graphql"

var ContractItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ContractItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"contract_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"department": &graphql.Field{
			Type: DropdownItemType,
		},
		"job_position_in_organization_unit": &graphql.Field{
			Type: DropdownItemType,
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
		"number_of_conference": &graphql.Field{
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
		"files": &graphql.Field{
			Type: graphql.NewList(FileDropdownItemType),
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})
