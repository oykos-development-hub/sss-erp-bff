package types

import "github.com/graphql-go/graphql"

var JobPositionItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobPositionItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"abbreviation": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"requirements": &graphql.Field{
			Type: graphql.String,
		},
		"is_judge": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_judge_president": &graphql.Field{
			Type: graphql.Boolean,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var JobPositionOrganizationUnitItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobPositionOrganizationUnitItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var JobPositionsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobPositions",
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
			Type: graphql.NewList(JobPositionItemType),
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var JobPositionsOrganizationUnitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobPositionsOrganizationUnitType",
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
			Type: graphql.NewList(JobPositionOrganizationUnitItemType),
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var JobPositionInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobPositionInsert",
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
			Type: JobPositionItemType,
		},
	},
})

var JobPositionDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobPositionDelete",
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
	},
})

var JobPositionInOrganizationUnitItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobPositionInOrganizationUnitItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"systematization_id": &graphql.Field{
			Type: graphql.Int,
		},
		"parent_organization_unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"job_position_id": &graphql.Field{
			Type: graphql.Int,
		},
		"parent_job_position_id": &graphql.Field{
			Type: graphql.Int,
		},
		"system_permission_id": &graphql.Field{
			Type: graphql.Int,
		},
		"available_slots": &graphql.Field{
			Type: graphql.Int,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"requirements": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var JobPositionInOrganizationUnitInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobPositionInOrganizationUnitInsert",
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
			Type: JobPositionInOrganizationUnitItemType,
		},
	},
})

var JobPositionInOrganizationUnitDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobPositionInOrganizationUnitDelete",
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
	},
})

var EmployeeInOrganizationUnitItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EmployeeInOrganizationUnitItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_account_id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"position_in_organization_unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var EmployeeInOrganizationUnitInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EmployeeInOrganizationUnitInsert",
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
			Type: EmployeeInOrganizationUnitItemType,
		},
	},
})

var EmployeeInOrganizationUnitDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EmployeeInOrganizationUnitDelete",
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
	},
})
