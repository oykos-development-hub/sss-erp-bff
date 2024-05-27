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

var SpendingDynamicMonthEntryItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingDynamicMonthEntryItem",
	Fields: graphql.Fields{
		"value": &graphql.Field{
			Type: graphql.String,
		},
		"savings": &graphql.Field{
			Type: graphql.String,
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
		"current_budget_id": &graphql.Field{
			Type: graphql.Int,
		},
		"actual": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"january": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"february": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"march": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"april": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"may": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"june": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"july": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"august": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"september": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"october": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"november": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
		},
		"december": &graphql.Field{
			Type: SpendingDynamicMonthEntryItemType,
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
		"version": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})
