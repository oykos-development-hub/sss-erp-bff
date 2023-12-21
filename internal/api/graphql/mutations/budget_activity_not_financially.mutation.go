package mutations

import "github.com/graphql-go/graphql"

var BudgetActivityNotFinanciallyInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BudgetActivityNotFinanciallyInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"request_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"person_responsible_name_surname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"person_responsible_working_place": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"person_responsible_telephone_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"person_responsible_email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"contact_person_name_surname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"contact_person_working_place": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"contact_person_telephone_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"contact_person_email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var ProgramNotFinanciallyInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ProgramNotFinanciallyInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"program_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var GoalsNotFinanciallyInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GoalsNotFinanciallyInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_program_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var IndicatorNotFinanciallyInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "IndicatorNotFinanciallyInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"goals_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"code": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"source": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"base_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"gender_equality": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"base_value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"source_information": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"unit": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"planned_current_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revised_current_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value_current_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"planned_next_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revised_next_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value_next_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"planned_after_next_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revised_after_next_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value_after_next_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
