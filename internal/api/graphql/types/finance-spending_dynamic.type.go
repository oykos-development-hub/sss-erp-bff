package types

import (
	"github.com/graphql-go/graphql"
)

var SpendingDynamicType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingDynamic",
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
			Type: graphql.NewList(SpendingDynamicItemType),
		},
	},
})

var SpendingDynamicItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingDynamicItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
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
		"actual": &graphql.Field{
			Type: graphql.String,
		},
		"spending_dynamic_id": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"january": &graphql.Field{
			Type: graphql.String,
		},
		"february": &graphql.Field{
			Type: graphql.String,
		},
		"march": &graphql.Field{
			Type: graphql.String,
		},
		"april": &graphql.Field{
			Type: graphql.String,
		},
		"may": &graphql.Field{
			Type: graphql.String,
		},
		"june": &graphql.Field{
			Type: graphql.String,
		},
		"july": &graphql.Field{
			Type: graphql.String,
		},
		"august": &graphql.Field{
			Type: graphql.String,
		},
		"september": &graphql.Field{
			Type: graphql.String,
		},
		"october": &graphql.Field{
			Type: graphql.String,
		},
		"november": &graphql.Field{
			Type: graphql.String,
		},
		"december": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SpendingDynamicHistoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingDynamicHistory",
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
			Type: graphql.NewList(SpendingDynamicHistoryItemType),
		},
	},
})

var SpendingDynamicHistoryItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingDynamicHistoryItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"budget_id": &graphql.Field{
			Type: graphql.Int,
		},
		"unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})
