package types

import (
	"bff/internal/api/dto"
	"sync"

	"github.com/graphql-go/graphql"
)

var SpendingReleaseListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingReleaseListType",
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

var SpendingReleaseOverviwItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingReleaseOverviewItem",
	Fields: graphql.Fields{
		"month": &graphql.Field{
			Type: graphql.Int,
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"value": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
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
			Type: graphql.NewList(SpendingReleaseOverviwItemType),
		},
	},
})

var SpendingReleaseDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpendingReleaseDelete",
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

var (
	spendingReleaseType              *graphql.Object
	onceInitializeSpendingRelease    sync.Once
	spendingReleaseGetType           *graphql.Object
	onceInitializeSpendingReleaseGet sync.Once
)

func GetSpendingReleaseGetType() *graphql.Object {
	onceInitializeSpendingReleaseGet.Do(func() {
		initSpendingReleaseGetType()
	})
	return spendingReleaseGetType
}

func GetSpendingReleaseType() *graphql.Object {
	onceInitializeSpendingRelease.Do(func() {
		initSpendingReleaseType()
	})
	return spendingReleaseType
}

func initSpendingReleaseGetType() {
	spendingReleaseGetType = graphql.NewObject(graphql.ObjectConfig{
		Name: "SpendingReleaseGetType",
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
				Type: graphql.NewList(GetSpendingReleaseType()),
			},
		},
	})
}

func initSpendingReleaseType() {
	spendingReleaseType = graphql.NewObject(graphql.ObjectConfig{
		Name: "SpendingReleaseType",
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
				"account_title": &graphql.Field{
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
				"value": &graphql.Field{
					Type: graphql.String,
				},
				"total_savings": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
				"children": &graphql.Field{
					Type: graphql.NewList(spendingReleaseType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if releaseItem, ok := p.Source.(*dto.SpendingReleaseDTO); ok {
							return releaseItem.Children, nil
						}
						return nil, nil
					},
				},
			}
		}),
	})
}
