package types

import "github.com/graphql-go/graphql"

var ActivitiesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActivitiesType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"subroutine": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
	},
})

var ActivitiesOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActivitiesOverview",
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
			Type: graphql.NewList(ActivitiesType),
		},
	},
})

var ActivitiesDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActivitiesDelete",
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

var ActivitiesInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ActivitiesInsert",
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
			Type: ActivitiesType,
		},
	},
})
