package types

import (
	"github.com/graphql-go/graphql"
)

var PinType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Pin",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
