package types

import "github.com/graphql-go/graphql"

var NotificationsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationsItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"module": &graphql.Field{
			Type: graphql.String,
		},
		"to_user_id": &graphql.Field{
			Type: graphql.Int,
		},
		"from_user_id": &graphql.Field{
			Type: graphql.Int,
		},
		"from_content": &graphql.Field{
			Type: graphql.String,
		},
		"is_read": &graphql.Field{
			Type: graphql.Boolean,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var NotificationsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationsOverview",
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
			Type: graphql.NewList(NotificationsItemType),
		},
	},
})

var NotificationsInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationsInsert",
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
			Type: NotificationsItemType,
		},
	},
})

var NotificationsDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationsDelete",
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
