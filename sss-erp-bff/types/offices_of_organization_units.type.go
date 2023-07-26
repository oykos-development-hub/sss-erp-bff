package types

import "github.com/graphql-go/graphql"

var OfficesOfOrganizationUnitItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OfficesOfOrganizationUnitItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"abbreviation": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var OfficesOfOrganizationUnitOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OfficesOfOrganizationUnitOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(OfficesOfOrganizationUnitItemType),
		},
	},
})

var OfficesOfOrganizationUnitInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OfficesOfOrganizationUnitInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: OfficesOfOrganizationUnitItemType,
		},
	},
})

var OfficesOfOrganizationUnitDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OfficesOfOrganizationUnitDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
