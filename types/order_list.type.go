package types

import (
	"github.com/graphql-go/graphql"
)

var OrderListItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"data_order": &graphql.Field{
			Type: graphql.String,
		},
		"total_price": &graphql.Field{
			Type: graphql.Int,
		},
		"public_procurement": &graphql.Field{
			Type: DropdownItemType,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(OrderListProcurementAvailableArticlesType),
		},
		"user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})

var OrderListOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(OrderListItemType),
		},
	},
})

var OrderProcurementAvailableType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderProcurementAvailableOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(OrderListProcurementAvailableArticlesType),
		},
	},
})

var OrderItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"data_order": &graphql.Field{
			Type: graphql.String,
		},
		"total_price": &graphql.Field{
			Type: graphql.Int,
		},
		"public_procurement": &graphql.Field{
			Type: OrderListPublicProcurementArticlesType,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var OrderListPublicProcurementArticlesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListPublicProcurementArticles",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(OrderListProcurementAvailableArticlesType),
		},
	},
})

var OrderListProcurementAvailableArticlesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListProcurementAvailableArticles",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"manufacturer": &graphql.Field{
			Type: graphql.String,
		},
		"unit": &graphql.Field{
			Type: graphql.String,
		},
		"available": &graphql.Field{
			Type: graphql.Int,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"total_price": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var OrderListInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: OrderListItemType,
		},
	},
})

var OrderListReceiveType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListReceive",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var OrderListAssetMovementType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListAssetMovement",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
