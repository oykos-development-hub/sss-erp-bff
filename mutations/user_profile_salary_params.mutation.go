package mutations

import "github.com/graphql-go/graphql"

var UserProfileSalaryParamsInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileSalaryParamsInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"benefited_track": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"without_raise": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"insurance_basis": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"salary_rank": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"daily_work_hours": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"weekly_work_hours": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"education_rank": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"education_naming": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"user_resolution_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
