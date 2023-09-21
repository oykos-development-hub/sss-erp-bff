package mutations

import "github.com/graphql-go/graphql"

var JobPositionInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "JobPositionInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"requirements": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"is_judge": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"is_judge_president": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"color": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"icon": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var JobPositionInOrganizationUnitInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "JobPositionInOrganizationUnitInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"systematization_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"parent_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"job_position_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"available_slots": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"requirements": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"employees": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
})

var EmployeeInOrganizationUnitInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "EmployeeInOrganizationUnitInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"position_in_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
})
