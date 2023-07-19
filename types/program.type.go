package types

import "github.com/graphql-go/graphql"

var ProgramType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProgramType",
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
	},
})
