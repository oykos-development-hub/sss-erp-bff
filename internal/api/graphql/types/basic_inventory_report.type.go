package types

import "github.com/graphql-go/graphql"

var ReportValueClassInventoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ReportValueClassInventoryType",
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
			Type: ReportValueClassInventoryResponseType,
		},
	},
})

var ReportInventoryListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ReportInventoryListType",
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
			Type: graphql.NewList(ReportInventoryList),
		},
	},
})

var ReportValueClassInventoryResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ReportValueClassInventoryResponseType",
	Fields: graphql.Fields{
		"values": &graphql.Field{
			Type: graphql.NewList(ReportValueClassInventoryItemType),
		},
		"purchase_gross_price": &graphql.Field{
			Type: graphql.Float,
		},
		"lost_value": &graphql.Field{
			Type: graphql.Float,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var ReportValueClassInventoryItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ReportValueClassInventoryItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"class": &graphql.Field{
			Type: graphql.String,
		},
		"purchase_gross_price": &graphql.Field{
			Type: graphql.Float,
		},
		"lost_value": &graphql.Field{
			Type: graphql.Float,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var ReportInventoryList = graphql.NewObject(graphql.ObjectConfig{
	Name: "ReportInventoryList",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"inventory_number": &graphql.Field{
			Type: graphql.String,
		},
		"office": &graphql.Field{
			Type: graphql.String,
		},
		"procurement_price": &graphql.Field{
			Type: graphql.Float,
		},
		"lost_value": &graphql.Field{
			Type: graphql.Float,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
	},
})
