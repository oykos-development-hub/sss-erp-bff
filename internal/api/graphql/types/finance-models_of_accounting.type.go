package types

import "github.com/graphql-go/graphql"

var ModelsOfAccountingOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ModelsOfAccountingOverviewType",
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
			Type: graphql.NewList(ModelsOfAccountingType),
		},
	},
})

var ModelsOfAccountingInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ModelsOfAccountingInsertType",
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
			Type: ModelsOfAccountingType,
		},
	},
})

var ModelsOfAccountingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ModelsOfAccountingType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(ModelOfAccountingItemType),
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var ModelOfAccountingItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ModelOfAccountingItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"debit_account": &graphql.Field{
			Type: AccountItemType,
		},
		"credit_account": &graphql.Field{
			Type: AccountItemType,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
