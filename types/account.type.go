package types

import (
	"bff/dto"

	"github.com/graphql-go/graphql"
)

var AccountType *graphql.Object

func init() {
	initAccountType()
	initAccountOverviewType()
}

func initAccountType() {
	AccountType = graphql.NewObject(graphql.ObjectConfig{
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
					Type: graphql.NewList(AccountType),
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

var AccountOverviewType *graphql.Object

func initAccountOverviewType() {
	AccountOverviewType = graphql.NewObject(graphql.ObjectConfig{
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
			"items": &graphql.Field{
				Type: graphql.NewList(AccountType),
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
