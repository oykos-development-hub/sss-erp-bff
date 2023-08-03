package types

import "github.com/graphql-go/graphql"

var UserProfileResolutionItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResolutionItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"resolution_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"resolution_purpose": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_start": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_end": &graphql.Field{
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

var UserProfileResolutionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResolution",
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
			Type: graphql.NewList(UserProfileResolutionItemType),
		},
	},
})

var UserProfileResolutionInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileResolutionInsert",
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
			Type: UserProfileResolutionItemType,
		},
	},
})

var UserProfileResolutionDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ResolutionDelete",
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
