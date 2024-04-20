package mutations

import "github.com/graphql-go/graphql"

var NonFinancialBudgetInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NonFinancialBudgetInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"request_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"impl_contact_fullname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"impl_contact_working_place": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"impl_contact_phone": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"impl_contact_email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"contact_fullname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"contact_working_place": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"contact_phone": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"contact_email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"goals": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(NonFinancialGoalInsertMutation),
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

var NonFinancialGoalInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NonFinancialGoalInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"non_financial_budget_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"indicators": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(NonFinancialGoalIndicatorInsertMutation),
		},
	},
})

var NonFinancialGoalIndicatorInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "NonFinancialGoalIndicatorInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"goal_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"performance_indicator_code": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"indicator_source": &graphql.InputObjectFieldConfig{
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
		"source_of_information": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"unit_of_measure": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"indicator_description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"planned_value_1": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revised_value_1": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"achieved_value_1": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"planned_value_2": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revised_value_2": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"achieved_value_2": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"planned_value_3": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revised_value_3": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"achieved_value_3": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
