package types

import "github.com/graphql-go/graphql"

var JobTenderTypeItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderTypeItem",
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
		"value": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"is_judge": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_judge_president": &graphql.Field{
			Type: graphql.Boolean,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var JobTenderTypesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderTypes",
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
			Type: graphql.NewList(JobTenderTypeItemType),
		},
	},
})

var JobTenderTypeInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderTypeInsert",
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
			Type: JobTenderTypeItemType,
		},
	},
})

var JobTenderTypeDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderTypeDelete",
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
