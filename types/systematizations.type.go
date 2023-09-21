package types

import "github.com/graphql-go/graphql"

var SystematizationsOverviewItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SystematizationsOverviewItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Int,
		},
		"date_of_activation": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SystematizationsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SystematizationsOverview",
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
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(SystematizationsOverviewItemType),
		},
	},
})

var SystematizationEmployeesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SystematizationEmployeesType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SystematizationJobPositionsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SystematizationJobPositionsType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"job_positions": &graphql.Field{
			Type: DropdownItemType,
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
		"available_slots": &graphql.Field{
			Type: graphql.Int,
		},
		"employees": &graphql.Field{
			Type: graphql.NewList(DropdownItemType),
		},
	},
})

var SystematizationSectorsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SystematizationSectors",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"parent_id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"abbreviation": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"job_positions_organization_units": &graphql.Field{
			Type: graphql.NewList(SystematizationJobPositionsType),
		},
	},
})

var SystematizationDetailsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SystematizationsDetailsItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Int,
		},
		"date_of_activation": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"sectors": &graphql.Field{
			Type: graphql.NewList(SystematizationSectorsType),
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"active_employees": &graphql.Field{
			Type: graphql.NewList(ActiveEmployeesItemType),
		},
	},
})

var ActiveEmployeesItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActiveEmployeesItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"full_name": &graphql.Field{
			Type: graphql.String,
		},
		"job_position": &graphql.Field{
			Type: DropdownItemType,
		},
		"sector": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SystematizationDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SystematizationDetailsType",
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
			Type: SystematizationDetailsItemType,
		},
	},
})

var SystematizationDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SystematizationDelete",
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
