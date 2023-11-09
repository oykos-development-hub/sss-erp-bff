package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var OrderListOverviewField = &graphql.Field{
	Type:        types.OrderListOverviewType,
	Description: "Returns a data of Order List items",
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
		"supplier_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"search": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"public_procurement_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"active_plan": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: resolvers.OrderListOverviewResolver,
}

var OrderListInsertField = &graphql.Field{
	Type:        types.OrderListInsertType,
	Description: "Creates new or alter existing Order List",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.OrderListInsertMutation),
		},
	},
	Resolve: resolvers.OrderListInsertResolver,
}

var OrderListReceiveField = &graphql.Field{
	Type:        types.OrderListReceiveType,
	Description: "Receive new or alter existing Order List",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.OrderListReceiveMutation),
		},
	},
	Resolve: resolvers.OrderListReceiveResolver,
}

var OrderListAssetMovementField = &graphql.Field{
	Type:        types.OrderListAssetMovementType,
	Description: "Asset Movement new or alter existing Order List",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.OrderListAssetMovementMutation),
		},
	},
	Resolve: resolvers.OrderListAssetMovementResolver,
}

var OrderProcurementAvailableField = &graphql.Field{
	Type:        types.OrderProcurementAvailableType,
	Description: "Returns a data of Order Procurement Available items List",
	Args: graphql.FieldConfigArgument{
		"public_procurement_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.OrderProcurementAvailableResolver,
}

var OrderListDeleteField = &graphql.Field{
	Type:        types.OrderListReceiveType,
	Description: "Delete existing Order",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.OrderListDeleteResolver,
}

var RecipientUsersField = &graphql.Field{
	Type:        types.RecipientUsersType,
	Description: "Receive users List",
	Args: graphql.FieldConfigArgument{
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.RecipientUsersResolver,
}

var OrderListReceiveDeleteField = &graphql.Field{
	Type:        types.OrderListReceiveDeleteType,
	Description: "Delete Receive existing Order",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.OrderListReceiveDeleteResolver,
}

var OrderListAssetMovementDeleteField = &graphql.Field{
	Type:        types.OrderListAssetMovementDeleteType,
	Description: "Delete Asset Movement existing Order",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.OrderListAssetMovementDeleteResolver,
}
