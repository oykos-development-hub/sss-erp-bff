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
		"date_order": &graphql.Field{
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
		"data": &graphql.Field{
			Type: JSON,
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(OrderListProcurementAvailableArticlesType),
		},
		"office": &graphql.Field{
			Type: DropdownItemType,
		},
		"recipient_user": &graphql.Field{
			Type: DropdownItemType,
		},
		"invoice_date": &graphql.Field{
			Type: graphql.String,
		},
		"invoice_number": &graphql.Field{
			Type: graphql.String,
		},
		"date_system": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var OrderListOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
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
		"data": &graphql.Field{
			Type: JSON,
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
			Type: graphql.Float,
		},
	},
})

var OrderListInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
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
		"data": &graphql.Field{
			Type: JSON,
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
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RecipientUsersType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RecipientUsersOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(DropdownItemType),
		},
	},
})

var OrderListReceiveDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListReceiveDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var OrderListAssetMovementDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrderListAssetMovementDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
