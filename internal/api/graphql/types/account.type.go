package types

import (
	"bff/internal/api/dto"
	"sync"

	"github.com/graphql-go/graphql"
)

var (
	accountType                  *graphql.Object
	accountOverviewType          *graphql.Object
	onceInitalizeAccount         sync.Once
	onceInitalizeAccountOverview sync.Once
)

func GetAccountType() *graphql.Object {
	onceInitalizeAccount.Do(func() {
		initAccountType()
	})
	return accountType
}

func GetAccountOverviewType() *graphql.Object {
	onceInitalizeAccountOverview.Do(func() {
		initAccountOverviewType()
	})
	return accountOverviewType
}

func initAccountType() {
	accountType = graphql.NewObject(graphql.ObjectConfig{
		Name: "AccountType",
		Fields: (graphql.FieldsThunk)(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"parent_id": &graphql.Field{
					Type: graphql.Int,
				},
				"serial_number": &graphql.Field{
					Type: graphql.String,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"children": &graphql.Field{
					Type: graphql.NewList(accountType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if accountItem, ok := p.Source.(*dto.AccountItemResponseItem); ok {
							return accountItem.Children, nil
						}
						return nil, nil
					},
				},
			}
		}),
	})
}

func initAccountOverviewType() {
	accountOverviewType = graphql.NewObject(graphql.ObjectConfig{
		Name: "AccountOverviewType",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"total": &graphql.Field{
				Type: graphql.Int,
			},
			"version": &graphql.Field{
				Type: graphql.Int,
			},
			"items": &graphql.Field{
				Type: graphql.NewList(GetAccountType()),
			},
		},
	})
}

var AccountItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountItem",
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
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var AccountInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountInsert",
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
			Type: AccountItemType,
		},
	},
})

var AccountDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountDelete",
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
