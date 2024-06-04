package types

import "github.com/graphql-go/graphql"

var InternalReallocationOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InternalReallocationOverviewType",
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
			Type: graphql.NewList(InternalReallocationType),
		},
	},
})

var InternalReallocationInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InternalReallocationInsertType",
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
			Type: InternalReallocationType,
		},
	},
})

var InternalReallocationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InternalReallocationType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"organization_unit": &graphql.Field{
			Type: OrganizationUnitParentType,
		},
		"date_of_request": &graphql.Field{
			Type: graphql.String,
		},
		"requested_by": &graphql.Field{
			Type: DropdownItemType,
		},
		"budget": &graphql.Field{
			Type: DropdownItemType,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(InternalReallocationItemType),
		},
	},
})

var InternalReallocationItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InternalReallocationItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"source_account": &graphql.Field{
			Type: DropdownItemType,
		},
		"destination_account": &graphql.Field{
			Type: DropdownItemType,
		},
		"amount": &graphql.Field{
			Type: graphql.String,
		},
	},
})
