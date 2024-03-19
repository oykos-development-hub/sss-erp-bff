package types

import "github.com/graphql-go/graphql"

var InvoiceOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceOverview",
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
			Type: graphql.NewList(InvoiceType),
		},
	},
})

var InvoiceInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceInsert",
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
			Type: InvoiceType,
		},
	},
})

var InvoiceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"type_of_subject": &graphql.Field{
			Type: DropdownItemType,
		},
		"type_of_contract": &graphql.Field{
			Type: DropdownItemType,
		},
		"source_of_funding": &graphql.Field{
			Type: DropdownItemType,
		},
		"supplier_title": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_start": &graphql.Field{
			Type: graphql.String,
		},
		"invoice_number": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"order_id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"date_of_invoice": &graphql.Field{
			Type: graphql.String,
		},
		"receipt_date": &graphql.Field{
			Type: graphql.String,
		},
		"sss_invoice_receipt_date": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_payment": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"bank_account": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.Int,
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(InvoiceArticlesType),
		},
	},
})

var InvoiceDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceDelete",
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

var InvoiceArticlesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceArticles",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"net_price": &graphql.Field{
			Type: graphql.Float,
		},
		"vat_price": &graphql.Field{
			Type: graphql.Float,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"account": &graphql.Field{
			Type: DropdownItemType,
		},
		"cost_account": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})
