package types

import "github.com/graphql-go/graphql"

var ActivitiesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActivitiesType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"subroutine_id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"organization_unit_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
