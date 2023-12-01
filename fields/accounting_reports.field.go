package fields

import (
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var OverallSpendingField = &graphql.Field{
	Type:        types.OverallSpendingType,
	Description: "Returns a data for overall spending report",
	Args: graphql.FieldConfigArgument{
		"year": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"office_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"search": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"exception": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: resolvers.OverallSpendingResolver,
}
