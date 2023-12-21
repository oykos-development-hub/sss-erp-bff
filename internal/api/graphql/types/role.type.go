package types

import "github.com/graphql-go/graphql"

var RoleDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RoleDetails",
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
			Type: RoleItemDetailsType,
		},
	},
})

var RoleOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RoleOverview",
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
			Type: graphql.NewList(RoleItemType),
		},
	},
})

var RoleItemDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RoleItemDetails",
	Fields: graphql.Fields{
		"role": &graphql.Field{
			Type: RoleItemType,
		},
		"users": &graphql.Field{
			Type: graphql.NewList(UserAccountsOverviewItemType),
		},
	},
})

var RoleItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RoleItem",
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
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var RoleInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RoleInsert",
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
			Type: RoleItemType,
		},
	},
})
