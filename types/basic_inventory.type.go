package types

import "github.com/graphql-go/graphql"

var BasicInventoryDetailsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDetailsItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"article_id": &graphql.Field{
			Type: graphql.Int,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"source_type": &graphql.Field{
			Type: graphql.String,
		},
		"class_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"depreciation_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"real_estate": &graphql.Field{
			Type: BasicInventoryRealEstatesItemType,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"inventory_number": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"abbreviation": &graphql.Field{
			Type: graphql.String,
		},
		"internal_ownership": &graphql.Field{
			Type: graphql.Boolean,
		},
		"office": &graphql.Field{
			Type: DropdownItemType,
		},
		"location": &graphql.Field{
			Type: graphql.String,
		},
		"target_user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"unit": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"net_price": &graphql.Field{
			Type: graphql.Int,
		},
		"gross_price": &graphql.Field{
			Type: graphql.Int,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_purchase": &graphql.Field{
			Type: graphql.String,
		},
		"source": &graphql.Field{
			Type: graphql.String,
		},
		"donor_title": &graphql.Field{
			Type: graphql.String,
		},
		"invoice_number": &graphql.Field{
			Type: graphql.Int,
		},
		"price_of_assessment": &graphql.Field{
			Type: graphql.Int,
		},
		"date_of_assessment": &graphql.Field{
			Type: graphql.String,
		},
		"lifetime_of_assessment_in_months": &graphql.Field{
			Type: graphql.Int,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deactivation_description": &graphql.Field{
			Type: graphql.String,
		},
		"assessments": &graphql.Field{
			Type: graphql.NewList(BasicInventoryAssessmentsItemType),
		},
		"movements": &graphql.Field{
			Type: graphql.NewList(BasicInventoryDispatchItemType),
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"invoice_file_id": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var BasicInventoryItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"source_type": &graphql.Field{
			Type: graphql.String,
		},
		"class_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"depreciation_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"real_estate": &graphql.Field{
			Type: BasicInventoryRealEstatesItemType,
		},
		"inventory_number": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"office": &graphql.Field{
			Type: DropdownItemType,
		},
		"target_user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"target_organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"gross_price": &graphql.Field{
			Type: graphql.Int,
		},
		"date_of_purchase": &graphql.Field{
			Type: graphql.String,
		},
		"source": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var BasicInventoryOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryOverview",
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
			Type: graphql.NewList(BasicInventoryItemType),
		},
	},
})

var BasicInventoryDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryDetails",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: graphql.NewList(BasicInventoryDetailsItemType),
		},
	},
})

var BasicInventoryInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: graphql.NewList(BasicInventoryDetailsItemType),
		},
	},
})
