package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) OrderListOverviewField() *graphql.Field {
	return &graphql.Field{
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
			"year": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"finance_overview": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"sort_by_date_order": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"sort_by_total_price": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.OrderListOverviewResolver,
	}
}
func (f *Field) OrderListInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrderListInsertType,
		Description: "Creates new or alter existing Order List",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.OrderListInsertMutation),
			},
		},
		Resolve: f.Resolvers.OrderListInsertResolver,
	}
}
func (f *Field) OrderListReceiveField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrderListReceiveType,
		Description: "Receive new or alter existing Order List",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.OrderListReceiveMutation),
			},
		},
		Resolve: f.Resolvers.OrderListReceiveResolver,
	}
}
func (f *Field) MovementInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrderListAssetMovementType,
		Description: "Asset Movement new or alter existing Order List",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.OrderListAssetMovementMutation),
			},
		},
		Resolve: f.Resolvers.OrderListAssetMovementResolver,
	}
}
func (f *Field) OrderProcurementAvailableField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrderProcurementAvailableType,
		Description: "Returns a data of Order Procurement Available items List",
		Args: graphql.FieldConfigArgument{
			"public_procurement_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"visibility_type": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.OrderProcurementAvailableResolver,
	}
}
func (f *Field) OrderListDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrderListReceiveType,
		Description: "Delete existing Order",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.OrderListDeleteResolver,
	}
}
func (f *Field) RecipientUsersField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RecipientUsersType,
		Description: "Receive users List",
		Args: graphql.FieldConfigArgument{
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RecipientUsersResolver,
	}
}
func (f *Field) OrderListReceiveDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrderListReceiveDeleteType,
		Description: "Delete Receive existing Order",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.OrderListReceiveDeleteResolver,
	}
}
func (f *Field) MovementDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrderListAssetMovementDeleteType,
		Description: "Delete Asset Movement existing Order",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.MovementDeleteResolver,
	}
}
func (f *Field) StockOverviewFiled() *graphql.Field {
	return &graphql.Field{
		Type:        types.StockOverviewType,
		Description: "Returns a data of stock items",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"date": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"sort_by_amount": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"sort_by_year": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.StockOverviewResolver,
	}
}
func (f *Field) MovementOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.MovementOverviewType,
		Description: "Returns a data of movement items",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"recipient_user_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"office_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"sort_by_date_order": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.MovementOverviewResolver,
	}
}
func (f *Field) MovementDetailsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.MovementDetailsType,
		Description: "Returns a data of movement item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.MovementDetailsResolver,
	}
}
func (f *Field) MovementArticlesField() *graphql.Field {
	return &graphql.Field{
		Type:        types.MovementArticlesType,
		Description: "Returns a data of all avaliable articles",
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.MovementArticlesResolver,
	}
}
