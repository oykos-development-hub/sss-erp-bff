package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var BasicInventoryOverviewField = &graphql.Field{
	Type:        types.BasicInventoryOverviewType,
	Description: "Returns a data of Basic Inventory items",
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
		"type": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"class_type_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"office_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"search": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"source_type": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"depreciation_type_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BasicInventoryOverviewResolver,
}

var BasicInventoryDetailsField = &graphql.Field{
	Type:        types.BasicInventoryDetailsType,
	Description: "Returns a data of Basic Inventory item Details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.BasicInventoryDetailsResolver,
}

var BasicInventoryInsertField = &graphql.Field{
	Type:        types.BasicInventoryInsertType,
	Description: "Creates new or alter existing Basic Inventory",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewList(mutations.BasicInventoryInsertMutation),
		},
	},
	Resolve: resolvers.BasicInventoryInsertResolver,
}
