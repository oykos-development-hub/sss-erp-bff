package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var BasicInventoryDepreciationTypesOverviewField = &graphql.Field{
	Type:        types.BasicInventoryDepreciationTypesOverviewType,
	Description: "Returns a data of Basic Inventory Depreciation Types items",
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
	Resolve: resolvers.BasicInventoryDepreciationTypesOverviewResolver,
}

var BasicInventoryDepreciationTypesInsertField = &graphql.Field{
	Type:        types.BasicInventoryDepreciationTypesInsertType,
	Description: "Creates new or alter existing Basic Inventory Dispatch",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.BasicInventoryDepreciationTypesMutation),
		},
	},
	Resolve: resolvers.BasicInventoryDepreciationTypesInsertResolver,
}

var BasicInventoryDepreciationTypesDeleteField = &graphql.Field{
	Type:        types.BasicInventoryDepreciationTypesDeleteType,
	Description: "Delete existing Basic Inventory Dispatch",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BasicInventoryDepreciationTypesDeleteResolver,
}
