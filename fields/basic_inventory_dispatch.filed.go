package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var BasicInventoryDispatchOverviewField = &graphql.Field{
	Type:        types.BasicInventoryDispatchOverviewType,
	Description: "Returns a data of Basic Inventory Dispatch with is activated false items",
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
		"source_organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"accepted": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"inventory_type": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.BasicInventoryDispatchOverviewResolver,
}

var BasicInventoryDispatchInsertField = &graphql.Field{
	Type:        types.BasicInventoryDispatchInsertType,
	Description: "Creates new or alter existing Basic Inventory Dispatch",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.BasicInventoryDispatchMutation),
		},
	},
	Resolve: resolvers.BasicInventoryDispatchInsertResolver,
}

var BasicInventoryDispatchDeleteField = &graphql.Field{
	Type:        types.BasicInventoryDispatchDeleteType,
	Description: "Delete existing Basic Inventory Dispatch",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.BasicInventoryDispatchDeleteResolver,
}

var BasicInventoryDispatchAcceptField = &graphql.Field{
	Type:        types.BasicInventoryDispatchDeleteType,
	Description: "Accept existing Basic Inventory Dispatch",
	Args: graphql.FieldConfigArgument{
		"dispatch_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.BasicInventoryDispatchAcceptResolver,
}
