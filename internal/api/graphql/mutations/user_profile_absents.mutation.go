package mutations

import "github.com/graphql-go/graphql"

var UserProfileAbsentInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileAbsentInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"absent_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"location": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"target_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_start": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_end": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var UserProfileVacationInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileVacationInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"year": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"number_of_days": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"resolution_purpose": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var UserProfileVacationsInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileVacationsInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"year": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"data": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(UserProfileVacationsItemMutation),
		},
	},
})

var UserProfileVacationsItemMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserProfileVacationsItemMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"number_of_days": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var AbsentTypeInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AbsentTypeInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"abbreviation": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"accounting_days_off": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"relocation": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"color": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"icon": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
