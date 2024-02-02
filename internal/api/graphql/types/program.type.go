package types

import "github.com/graphql-go/graphql"

var ProgramType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProgramType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"parent": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var ProgramsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProgramsOverview",
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
			Type: graphql.NewList(ProgramType),
		},
	},
})

var ProgramDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProgramsDelete",
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

var ProgramInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProgramsInsert",
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
			Type: ProgramType,
		},
	},
})
