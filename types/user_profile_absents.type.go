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
		"vacation_type": &graphql.Field{
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
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
