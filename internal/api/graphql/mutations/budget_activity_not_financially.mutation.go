package mutations

import "github.com/graphql-go/graphql"

var BudgetActivityNotFinanciallyInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BudgetActivityNotFinanciallyInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
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
