package types

import (
	"github.com/graphql-go/graphql"
)

var OrganizationUnitItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrganizationUnitItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"parent_id": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_judges": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"order_id": &graphql.Field{
			Type: graphql.Int,
		},
		"abbreviation": &graphql.Field{
			Type: graphql.String,
		},
		"pib": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"bank_accounts": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"folder_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var OrganizationUnitParentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrganizationUnitParent",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"parent_id": &graphql.Field{
			Type: graphql.Int,
		},
		"number_of_judges": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"abbreviation": &graphql.Field{
			Type: graphql.String,
		},
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"pib": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"order_id": &graphql.Field{
			Type: graphql.Int,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"children": &graphql.Field{
			Type: graphql.NewList(OrganizationUnitItemType),
		},
		"bank_accounts": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
		"folder_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var OrganizationUnitsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrganizationUnits",
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
			Type: graphql.NewList(OrganizationUnitParentType),
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var OrganizationUnitInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrganizationUnitInsert",
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
			Type: OrganizationUnitItemType,
		},
	},
})

var OrganizationUnitOrderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrganizationUnitOrder",
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
			Type: graphql.NewList(OrganizationUnitParentType),
		},
	},
})

var OrganizationUnitDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrganizationUnitDelete",
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
