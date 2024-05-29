package types

import "github.com/graphql-go/graphql"

var NonFinancialBudgetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NonFinancialBudgetType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"request_id": &graphql.Field{
			Type: graphql.Int,
		},
		"status": &graphql.Field{
			Type: DropdownItemType,
		},
		"impl_contact_fullname": &graphql.Field{
			Type: graphql.String,
		},
		"impl_contact_working_place": &graphql.Field{
			Type: graphql.String,
		},
		"impl_contact_phone": &graphql.Field{
			Type: graphql.String,
		},
		"impl_contact_email": &graphql.Field{
			Type: graphql.String,
		},
		"contact_fullname": &graphql.Field{
			Type: graphql.String,
		},
		"contact_working_place": &graphql.Field{
			Type: graphql.String,
		},
		"contact_phone": &graphql.Field{
			Type: graphql.String,
		},
		"contact_email": &graphql.Field{
			Type: graphql.String,
		},
		"statement": &graphql.Field{
			Type: graphql.String,
		},
		"activity": &graphql.Field{
			Type: ActivityType,
		},
	},
})

var BudgetRequest = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetRequestType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"budget_id": &graphql.Field{
			Type: graphql.String,
		},
		"request_type": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})

var ActivityType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActivityType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"sub_program": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"goals": &graphql.Field{
			Type: graphql.NewList(ActivityGoalType),
		},
	},
})

var ActivityGoalType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActivityGoalType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"indicators": &graphql.Field{
			Type: graphql.NewList(IndicatorNotFinanciallyType),
		},
	},
})

var IndicatorNotFinanciallyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "IndicatorNotFinanciallyType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"performance_indicator_code": &graphql.Field{
			Type: graphql.String,
		},
		"indicator_source": &graphql.Field{
			Type: graphql.String,
		},
		"base_year": &graphql.Field{
			Type: graphql.String,
		},
		"gender_equality": &graphql.Field{
			Type: graphql.String,
		},
		"base_value": &graphql.Field{
			Type: graphql.String,
		},
		"source_of_information": &graphql.Field{
			Type: graphql.String,
		},
		"unit_of_measure": &graphql.Field{
			Type: graphql.String,
		},
		"indicator_description": &graphql.Field{
			Type: graphql.String,
		},
		"planned_value_1": &graphql.Field{
			Type: graphql.String,
		},
		"revised_value_1": &graphql.Field{
			Type: graphql.String,
		},
		"achieved_value_1": &graphql.Field{
			Type: graphql.String,
		},
		"planned_value_2": &graphql.Field{
			Type: graphql.String,
		},
		"revised_value_2": &graphql.Field{
			Type: graphql.String,
		},
		"achieved_value_2": &graphql.Field{
			Type: graphql.String,
		},
		"planned_value_3": &graphql.Field{
			Type: graphql.String,
		},
		"revised_value_3": &graphql.Field{
			Type: graphql.String,
		},
		"achieved_value_3": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var NonFinancialBudgetInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NonFinancialBudgetInsertType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: NonFinancialBudgetType,
		},
	},
})

var NonFinancialBudgetOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NonFinancialBudgetOverviewType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(NonFinancialBudgetType),
		},
	},
})

var NonFinancialBudgetGoalInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NonFinancialBudgetGoalInsertType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"item": &graphql.Field{
			Type: ActivityGoalType,
		},
	},
})

var NonFinancialGoalIndicatorInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NonFinancialGoalIndicatorInsertType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"item": &graphql.Field{
			Type: IndicatorNotFinanciallyType,
		},
	},
})

var ActivityInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActivityInsertType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: ActivityType,
		},
	},
})

var IndicatorNotFinanciallyInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "IndicatorNotFinanciallyInsertType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: IndicatorNotFinanciallyType,
		},
	},
})

var IndicatorNotFinanciallyOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "IndicatorNotFinanciallyOverviewType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(IndicatorNotFinanciallyType),
		},
	},
})

var CheckBudgetActivityNotFinanciallyIsDoneType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CheckBudgetActivityNotFinanciallyIsDoneType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
