package types

import "github.com/graphql-go/graphql"

var UserAccountsOverviewItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserAccountsOverviewItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"role": &graphql.Field{
			Type: DropdownItemType,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"secondary_email": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"pin": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"verified_email": &graphql.Field{
			Type: graphql.Boolean,
		},
		"verified_phone": &graphql.Field{
			Type: graphql.Boolean,
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
	},
})

var UserAccountsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserAccountsOverview",
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
			Type: graphql.NewList(UserAccountsOverviewItemType),
		},
	},
})

var UserAccountInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserAccountInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(UserAccountsOverviewItemType),
		},
	},
})

var UserAccountDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserAccountDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
