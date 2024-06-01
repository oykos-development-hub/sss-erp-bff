package types

import (
	"bff/internal/api/dto"
	"sync"

	"github.com/graphql-go/graphql"
)

var (
	spendingDynamicType                   *graphql.Object
	onceInitializeSpendingDynamic         sync.Once
	spendingDynamicOverviewType           *graphql.Object
	onceInitializeSpendingDynamicOverview sync.Once
)

func GetSpendingDynamicOverviewType() *graphql.Object {
	onceInitializeSpendingDynamicOverview.Do(func() {
		initSpendingDynamicOverviewType()
	})
	return spendingDynamicOverviewType
}

func GetSpendingDynamicType() *graphql.Object {
	onceInitializeSpendingDynamic.Do(func() {
		initSpendingDynamicType()
	})
	return spendingDynamicType
}

func initSpendingDynamicOverviewType() {
	spendingDynamicOverviewType = graphql.NewObject(graphql.ObjectConfig{
		Name: "SpendingDynamicOverviewType",
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
				Type: graphql.NewList(GetSpendingDynamicType()),
			},
		},
	})
}

func initSpendingDynamicType() {
	spendingDynamicType = graphql.NewObject(graphql.ObjectConfig{
		Name: "SpendingDynamicType",
		Fields: (graphql.FieldsThunk)(func() graphql.Fields {
			return graphql.Fields{
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
				"account_serial_number": &graphql.Field{
					Type: graphql.String,
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
				"total_savings": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
				"children": &graphql.Field{
					Type: graphql.NewList(spendingDynamicType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if dynamicItem, ok := p.Source.(*dto.SpendingDynamicDTO); ok {
							return dynamicItem.Children, nil
						}
						return nil, nil
					},
				},
			}
		}),
	})
}

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
