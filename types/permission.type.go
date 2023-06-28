package types

import "github.com/graphql-go/graphql"

var PermissionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Permission",
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
		"route": &graphql.Field{
			Type: graphql.String,
		},
		"create": &graphql.Field{
			Type: graphql.Boolean,
		},
		"read": &graphql.Field{
			Type: graphql.Boolean,
		},
		"update": &graphql.Field{
			Type: graphql.Boolean,
		},
		"delete": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
