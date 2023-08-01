package types

import "github.com/graphql-go/graphql"

var JudgeNormItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgeNormItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"topic": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"percentage_of_norm_decrease": &graphql.Field{
			Type: graphql.Int,
		},
		"percentage_of_norm_fulfilment": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_norm_decrease": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_items": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_items_solved": &graphql.Field{
			Type: graphql.Int,
		},
		"evaluation": &graphql.Field{
			Type: UserProfileEvaluationItemType,
		},
		"date_of_evaluation_validity": &graphql.Field{
			Type: graphql.String,
		},
		"relocation": &graphql.Field{
			Type: UserProfileAbsentItemType,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var JudgesOverviewItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgesOverviewItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"job_position": &graphql.Field{
			Type: DropdownItemType,
		},
		"is_judge_president": &graphql.Field{
			Type: graphql.Boolean,
		},
		"first_name": &graphql.Field{
			Type: graphql.String,
		},
		"last_name": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"folder_id": &graphql.Field{
			Type: graphql.Int,
		},
		"norms": &graphql.Field{
			Type: graphql.NewList(JudgeNormItemType),
		},
	},
})

var JudgesOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgesOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(JudgesOverviewItemType),
		},
	},
})

var JudgeNormInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgeNormInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: JudgeNormItemType,
		},
	},
})

var JudgeNormDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgeNormDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var JudgeResolutionItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgeResolutionItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"available_slots_presidents": &graphql.Field{
			Type: graphql.Int,
		},
		"available_slots_judges": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_presidents": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_judges": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_employees": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_relocated_judges": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_suspended_judges": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var JudgeResolutionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgeResolution",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"available_slots_judges": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_judges": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(JudgeResolutionItemType),
		},
	},
})

var JudgeResolutionsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgeResolutionsOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(JudgeResolutionType),
		},
	},
})

var JudgeResolutionsInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgeResolutionsInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: JudgeResolutionType,
		},
	},
})

var JudgeResolutionsDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JudgeResolutionDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
