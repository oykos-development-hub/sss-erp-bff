package mutations

import "github.com/graphql-go/graphql"

var UserProfileBasicInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileBasicInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"first_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"last_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_birth": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"birth_last_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"country_of_birth": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"city_of_birth": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"nationality": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"citizenship": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"address": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"father_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"mother_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"mother_birth_last_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"bank_account": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"bank_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"official_personal_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"official_personal_document_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"official_personal_document_issuer": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"gender": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"single_parent": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"housing_done": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"revisor_role": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"housing_description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"marital_status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_taking_oath": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_becoming_judge": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"phone": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"job_position_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"contract_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"updated_at": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"national_minority": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"private_email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pin": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
