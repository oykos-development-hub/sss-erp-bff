package types

import "github.com/graphql-go/graphql"

var UserProfileAbsentItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileAbsentItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"absent_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"location": &graphql.Field{
			Type: graphql.String,
		},
		"target_organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"date_of_start": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_end": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var UserProfileAbsentSummaryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileAbsentSummary",
	Fields: graphql.Fields{
		"current_available_days": &graphql.Field{
			Type: graphql.Int,
		},
		"past_available_days": &graphql.Field{
			Type: graphql.Int,
		},
		"used_days": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var UserProfileAbsentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileAbsent",
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
		"summary": &graphql.Field{
			Type: UserProfileAbsentSummaryType,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(UserProfileAbsentItemType),
		},
	},
})

var UserProfileAbsentInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileAbsentInsert",
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
			Type: UserProfileAbsentItemType,
		},
	},
})

var UserProfileAbsentDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AbsentDelete",
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

var AbsentTypeInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AbsentTypeInsert",
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
			Type: AbsentTypeItem,
		},
	},
})
var AbsentTypeChildren = graphql.NewObject(graphql.ObjectConfig{
	Name: "AbsentTypeChildren",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"parent_id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"abbreviation": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"accounting_days_off": &graphql.Field{
			Type: graphql.Boolean,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
	}})
var AbsentTypeItem = graphql.NewObject(graphql.ObjectConfig{
	Name: "AbsentTypeItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"abbreviation": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"accounting_days_off": &graphql.Field{
			Type: graphql.Boolean,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"children": {
			Type: graphql.NewList(AbsentTypeChildren),
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})
var AbsentTypeItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AbsentTypeItemType",
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
			Type: graphql.NewList(AbsentTypeItem),
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
