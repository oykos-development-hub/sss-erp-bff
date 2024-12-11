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
		"number_of_conference": &graphql.Field{
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
		"is_judge": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_president": &graphql.Field{
			Type: graphql.Boolean,
		},
		"judge_application_submission_date": &graphql.Field{
			Type: graphql.String,
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
