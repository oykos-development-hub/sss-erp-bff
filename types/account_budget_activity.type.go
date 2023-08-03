package types

import (
	"bff/structs"

	"github.com/graphql-go/graphql"
)

var AccountBudgetActivityType *graphql.Object

func init() {
	initAccountBudgetActivityType()
	initAccountBudgetActivityOverviewType()
}

func initAccountBudgetActivityType() {
	AccountBudgetActivityType = graphql.NewObject(graphql.ObjectConfig{
		Name: "AccountBudgetActivityType",
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
				"value_current_year": &graphql.Field{
					Type: graphql.Int,
				},
				"value_next_year": &graphql.Field{
					Type: graphql.Int,
				},
				"value_after_next_year": &graphql.Field{
					Type: graphql.Int,
				},
				"children": &graphql.Field{
					Type: graphql.NewList(AccountBudgetActivityType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if accountItem, ok := p.Source.(*structs.AccountItemNode); ok {
							return accountItem.Children, nil
						}
						return nil, nil
					},
				},
			}
		}),
	})
}

var AccountBudgetActivityOverviewType *graphql.Object

func initAccountBudgetActivityOverviewType() {
	AccountBudgetActivityOverviewType = graphql.NewObject(graphql.ObjectConfig{
		Name: "AccountBudgetActivityOverviewType",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"items": &graphql.Field{
				Type: graphql.NewList(AccountBudgetActivityType),
			},
		},
	})
}

var AccountBudgetActivityInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccountBudgetActivityInsert",
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
