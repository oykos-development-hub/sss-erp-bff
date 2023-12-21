package mutations

import "github.com/graphql-go/graphql"

var JobTenderInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "JobTenderInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"position_in_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"number_of_vacant_seats": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var JobTenderApplicationInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "JobTenderApplicationInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"job_tender_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"first_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"last_name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"citizenship": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_birth": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_application": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"official_personal_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"evaluation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
