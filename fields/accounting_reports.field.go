package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var OverallSpendingField = &graphql.Field{
	Type:        types.OverallSpendingType,
	Description: "Returns a data for overall spending report",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.OverallSpendingMutation),
		},
	},
	Resolve: resolvers.OverallSpendingResolver,
}
