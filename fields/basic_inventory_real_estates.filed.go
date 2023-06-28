package fields

import (
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var BasicInventoryRealEstatesOverviewField = &graphql.Field{
	Type:        types.BasicInventoryRealEstatesOverviewType,
	Description: "Returns a data of Basic Inventory Real Estates items",
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BasicInventoryRealEstatesOverviewResolver,
}
