package mutations

import (
	"github.com/graphql-go/graphql"
)

var UserProfileBasicInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileBasicInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"first_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"middle_name": &graphql.InputObjectFieldConfig{
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
		"contract": &graphql.InputObjectFieldConfig{
			Type: UserProfileContractInput,
		},
		"date_of_becoming_judge": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"secondary_email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"phone": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"position_in_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"role_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"national_minority": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pin": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var UserProfileUpdateMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileUpdateMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"user_account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"first_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"middle_name": &graphql.InputObjectFieldConfig{
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
		"date_of_becoming_judge": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"position_in_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"contract": &graphql.InputObjectFieldConfig{
			Type: UserProfileContractInsertMutation,
		},
	},
})

var UserProfileContractInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileContractInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"contract_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"department_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"job_position_in_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
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
