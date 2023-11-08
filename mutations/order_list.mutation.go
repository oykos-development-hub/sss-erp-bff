package mutations

import "github.com/graphql-go/graphql"

var OrderListInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OrderListInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_order": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"total_price": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"public_procurement_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"articles": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ArticlesInsertMutation),
		},
		"file": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
})

var OrderListReceiveMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OrderListReceiveMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"order_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"date_system": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"invoice_date": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"invoice_number": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var OrderListAssetMovementMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OrderListAssetMovementMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"order_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"office_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"recipient_user_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var ArticlesInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ArticlesInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
