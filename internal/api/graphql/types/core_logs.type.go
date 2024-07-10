package types

import "github.com/graphql-go/graphql"

var LogsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LogsOverviewType",
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
			Type: graphql.NewList(LogsType),
		},
	},
})

var LogsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LogsType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"operation": &graphql.Field{
			Type: graphql.String,
		},
		"entity": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: DropdownItemType,
		},
		"user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"item_id": &graphql.Field{
			Type: graphql.Int,
		},
		"changed_at": &graphql.Field{
			Type: graphql.String,
		},
		"old_state": &graphql.Field{
			Type: JSON,
		},
		"new_state": &graphql.Field{
			Type: JSON,
		},
	},
})

var ErrorLogsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ErrorLogsOverviewType",
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
			Type: graphql.NewList(ErrorLogsType),
		},
	},
})

var ErrorLogsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ErrorLogsType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"entity": &graphql.Field{
			Type: graphql.String,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"error": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})
