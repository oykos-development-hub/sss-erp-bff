package types

import "github.com/graphql-go/graphql"

var CurrentBudgetMockType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CurrentBudgetMock",
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
