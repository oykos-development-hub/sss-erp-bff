package types

import "github.com/graphql-go/graphql"

var ExternalReallocationOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExternalReallocationOverviewType",
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
			Type: graphql.NewList(ExternalReallocationType),
		},
	},
})

var ExternalReallocationInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExternalReallocationInsertType",
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
			Type: ExternalReallocationType,
		},
	},
})

var ExternalReallocationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExternalReallocationType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"source_organization_unit": &graphql.Field{
			Type: OrganizationUnitParentType,
		},
		"destination_organization_unit": &graphql.Field{
			Type: OrganizationUnitParentType,
		},
		"date_of_request": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_action_dest_org_unit": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_action_sss": &graphql.Field{
			Type: graphql.String,
		},
		"requested_by": &graphql.Field{
			Type: DropdownItemType,
		},
		"accepted_by": &graphql.Field{
			Type: DropdownItemType,
		},
		"budget": &graphql.Field{
			Type: DropdownItemType,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"destination_org_unit_file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"sss_file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(ExternalReallocationItemType),
		},
	},
})

var ExternalReallocationItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExternalReallocationItemType",
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
