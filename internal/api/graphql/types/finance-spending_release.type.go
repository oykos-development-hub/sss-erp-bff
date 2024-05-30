package types

import (
	"github.com/graphql-go/graphql"
)

var SpendingReleaseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingReleaseType",
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
			Type: SpendingReleaseItemType,
		},
	},
})

var SpendingReleaseItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingReleaseItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"current_budget_id": &graphql.Field{
			Type: graphql.Int,
		},
		"budget_id": &graphql.Field{
			Type: graphql.Int,
		},
		"unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"account_id": &graphql.Field{
			Type: graphql.Int,
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"month": &graphql.Field{
			Type: graphql.Int,
		},
		"value": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SpendingReleaseOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingReleaseOverview",
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
			Type: graphql.NewList(SpendingReleaseItemType),
		},
	},
})
