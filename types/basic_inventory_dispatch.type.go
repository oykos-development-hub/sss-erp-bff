package types

import "github.com/graphql-go/graphql"

var BasicInventoryDispatchItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDispatchItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"dispatch_id": &graphql.Field{
			Type: graphql.Int,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"source_user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"target_user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"source_organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"target_organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"office": &graphql.Field{
			Type: DropdownItemType,
		},
		"is_accepted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"dispatch_description": &graphql.Field{
			Type: graphql.String,
		},
		"inventory_type": &graphql.Field{
			Type: graphql.String,
		},
		"inventory": &graphql.Field{
			Type: graphql.NewList(BasicInventoryItemType),
		},
		"date": &graphql.Field{
			Type: graphql.String,
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"deactivation_description": &graphql.Field{
			Type: graphql.String,
		},
		"deactivation_file_id": &graphql.Field{
			Type: FileDropdownItemType,
		},
		"date_of_deactivation": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var BasicInventoryDispatchItemInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDispatchItemInsertType",
	Fields: graphql.Fields{
		"source_user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"target_organization_unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"dispatch_description": &graphql.Field{
			Type: graphql.String,
		},
		"inventory_id": &graphql.Field{
			Type: graphql.NewList(graphql.Int),
		},
	},
})

var BasicInventoryDispatchOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDispatchOverview",
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
			Type: graphql.NewList(BasicInventoryDispatchItemType),
		},
	},
})

var BasicInventoryDispatchInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDispatchInsert",
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
			Type: BasicInventoryDispatchItemType,
		},
	},
})

var BasicInventoryDispatchDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDispatchDelete",
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
