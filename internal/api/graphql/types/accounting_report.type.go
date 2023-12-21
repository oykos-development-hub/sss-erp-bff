package types

import "github.com/graphql-go/graphql"

var OverallSpendingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OverallSpendingType",
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
			Type: graphql.NewList(ArticleReport),
		},
	},
})

var ArticleReport = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArticleReport",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
