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
		"donor": &graphql.Field{
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
		"target_organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
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
			Type: graphql.Float,
		},
		"residual_price": &graphql.Field{
			Type: graphql.Float,
		},
		"purchase_gross_price": &graphql.Field{
			Type: graphql.Float,
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
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"donor_title": &graphql.Field{
			Type: graphql.String,
		},
		"invoice_number": &graphql.Field{
			Type: graphql.String,
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
		"depreciation_rate": &graphql.Field{
			Type: graphql.String,
		},
		"amortization_value": &graphql.Field{
			Type: graphql.Int,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"inactive": &graphql.Field{
			Type: graphql.String,
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
			Type: graphql.Int,
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
			Type: graphql.Float,
		},
		"purchase_gross_price": &graphql.Field{
			Type: graphql.Float,
		},
		"date_of_purchase": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_assessments": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
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
		"location": &graphql.Field{
			Type: graphql.String,
		},
		"has_assessments": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var ReportValueClassInventoryResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ReportValueClassInventoryResponseType",
	Fields: graphql.Fields{
		"values": &graphql.Field{
			Type: graphql.NewList(ReportValueClassInventoryItemType),
		},
		"purchase_gross_price": &graphql.Field{
			Type: graphql.Float,
		},
		"gross_price": &graphql.Field{
			Type: graphql.Float,
		},
		"price_of_assessment": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var ReportValueClassInventoryItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ReportValueClassInventoryItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"class": &graphql.Field{
			Type: graphql.String,
		},
		"purchase_gross_price": &graphql.Field{
			Type: graphql.Float,
		},
		"gross_price": &graphql.Field{
			Type: graphql.Float,
		},
		"price_of_assessment": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var BasicInventoryOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryOverview",
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
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: BasicInventoryDetailsItemType,
		},
	},
})

var BasicInventoryMessageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryMessageType",
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

var ReportValueClassInventoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ReportValueClassInventoryType",
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
			Type: ReportValueClassInventoryResponseType,
		},
	},
})

var BasicInventoryInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BasicInventoryInsert",
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
			Type: graphql.NewList(BasicInventoryDetailsItemType),
		},
	},
})

var FormPDFType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FormPDFType",
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
			Type: graphql.String,
		},
	},
})
