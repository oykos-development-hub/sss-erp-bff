package types

import "github.com/graphql-go/graphql"

var BudgetActivityNotFinanciallyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetActivityNotFinanciallyType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name_of_organization_unit": &graphql.Field{
			Type: graphql.String,
		},
		"organization_code": &graphql.Field{
			Type: graphql.String,
		},
		"mission_statement": &graphql.Field{
			Type: graphql.String,
		},
		"person_responsible_name_surname": &graphql.Field{
			Type: graphql.String,
		},
		"person_responsible_working_place": &graphql.Field{
			Type: graphql.String,
		},
		"person_responsible_telephone_number": &graphql.Field{
			Type: graphql.String,
		},
		"person_responsible_email": &graphql.Field{
			Type: graphql.String,
		},
		"contact_person_name_surname": &graphql.Field{
			Type: graphql.String,
		},
		"contact_person_working_place": &graphql.Field{
			Type: graphql.String,
		},
		"contact_person_telephone_number": &graphql.Field{
			Type: graphql.String,
		},
		"contact_person_email": &graphql.Field{
			Type: graphql.String,
		},
		"program": &graphql.Field{
			Type: ProgramNotFinanciallyType,
		},
		"subprogram": &graphql.Field{
			Type: ProgramNotFinanciallyType,
		},
		"activity": &graphql.Field{
			Type: ProgramNotFinanciallyType,
		},
	},
})

var ProgramNotFinanciallyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProgramNotFinanciallyType",
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
		"goals": &graphql.Field{
			Type: graphql.NewList(GoalsNotFinanciallyType),
		},
	},
})

var GoalsNotFinanciallyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GoalsNotFinanciallyType",
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
	},
})

var IndicatorNotFinanciallyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "IndicatorNotFinanciallyType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"source": &graphql.Field{
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
		"source_information": &graphql.Field{
			Type: graphql.String,
		},
		"unit": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"planned_current_year": &graphql.Field{
			Type: graphql.String,
		},
		"revised_current_year": &graphql.Field{
			Type: graphql.String,
		},
		"value_current_year": &graphql.Field{
			Type: graphql.String,
		},
		"planned_next_year": &graphql.Field{
			Type: graphql.String,
		},
		"revised_next_year": &graphql.Field{
			Type: graphql.String,
		},
		"value_next_year": &graphql.Field{
			Type: graphql.String,
		},
		"planned_after_next_year": &graphql.Field{
			Type: graphql.String,
		},
		"revised_after_next_year": &graphql.Field{
			Type: graphql.String,
		},
		"value_after_next_year": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var BudgetActivityNotFinanciallyOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BudgetActivityNotFinanciallyOverviewType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: BudgetActivityNotFinanciallyType,
		},
	},
})

var ProgramNotFinanciallyOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProgramNotFinanciallyOverviewType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: ProgramNotFinanciallyType,
		},
	},
})

var GoalsNotFinanciallyOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GoalsNotFinanciallyOverviewType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: GoalsNotFinanciallyType,
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
