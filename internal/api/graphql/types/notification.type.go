package types

import (
	"github.com/graphql-go/graphql"
)

var NotificationReadType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationRead",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
	},
})

var NotificationsGetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationsGet",
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
		"items": &graphql.Field{
			Type: graphql.NewList(NotificationItemType),
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var NotificationItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationItem",
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
		"from_content": &graphql.Field{
			Type: graphql.String,
		},
		"from_user_id": &graphql.Field{
			Type: graphql.String,
		},
		"to_user_id": &graphql.Field{
			Type: graphql.Int,
		},
		"path": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
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

var NotificationDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "NotificationDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
	},
})
