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
		"group_of_articles_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"account_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"articles": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ArticlesInsertMutation),
		},
		"order_file": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"passed_to_finance": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"used_in_finance": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"is_pro_forma_invoice": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"pro_forma_invoice_date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"pro_forma_invoice_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
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
			Type: graphql.String,
		},
		"invoice_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"receive_file": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"articles": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ArticlesInsertMutation),
		},
	},
})

var OrderListAssetMovementMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OrderListAssetMovementMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"office_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"recipient_user_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_order": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"articles": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(ArticlesMovement),
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
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"net_price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"vat_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"order_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var ArticlesMovement = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ArticlesMovement",
	Fields: graphql.InputObjectConfigFieldMap{
		"quantity": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
