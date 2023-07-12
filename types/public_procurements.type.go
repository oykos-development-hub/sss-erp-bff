package types

import "github.com/graphql-go/graphql"

var PublicProcurementPlanDetailsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanDetailsItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"pre_budget_plan": &graphql.Field{
			Type: DropdownItemType,
		},
		"is_pre_budget": &graphql.Field{
			Type: graphql.Boolean,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"year": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_publishing": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_closing": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanItemDetailsItemType),
		},
	},
})

var PublicProcurementPlansOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlansOverview",
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
			Type: graphql.NewList(PublicProcurementPlanDetailsItemType),
		},
	},
})

var PublicProcurementPlanDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanDetails",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanDetailsItemType),
		},
	},
})

var PublicProcurementPlanInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanDetailsItemType),
		},
	},
})

var PublicProcurementPlanDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementPlanItemDetailsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemDetailsItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"budget_indent": &graphql.Field{
			Type: DropdownItemType,
		},
		"plan": &graphql.Field{
			Type: DropdownItemType,
		},
		"is_open_procurement": &graphql.Field{
			Type: graphql.Boolean,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"article_type": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_publishing": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_awarding": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanItemArticleItemType),
		},
	},
})

var PublicProcurementPlanItemDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemDetails",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanItemDetailsItemType),
		},
	},
})

var PublicProcurementPlanItemInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanItemDetailsItemType),
		},
	},
})

var PublicProcurementPlanItemDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementPlanItemLimitItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemLimitItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"public_procurement": &graphql.Field{
			Type: DropdownItemType,
		},
		"limit": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementPlanItemLimitsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemLimits",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanItemLimitItemType),
		},
	},
})

var PublicProcurementPlanItemLimitInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemLimitInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanItemLimitItemType),
		},
	},
})

var PublicProcurementPlanItemArticleItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemArticleItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"budget_indent": &graphql.Field{
			Type: DropdownItemType,
		},
		"public_procurement": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"net_price": &graphql.Field{
			Type: graphql.String,
		},
		"vat_percentage": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementPlanItemArticleInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemArticleInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: PublicProcurementPlanItemArticleItemType,
		},
	},
})

var PublicProcurementPlanItemArticleDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemArticleDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementArticleItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementArticleItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"net_price": &graphql.Field{
			Type: graphql.String,
		},
		"vat_percentage": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementOrganizationUnitArticleItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementOrganizationUnitArticleItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"public_procurement_article": &graphql.Field{
			Type: PublicProcurementArticleItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"is_rejected": &graphql.Field{
			Type: graphql.Boolean,
		},
		"rejected_description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementOrganizationUnitArticlesOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementOrganizationUnitArticlesOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementOrganizationUnitArticleItemType),
		},
	},
})

var PublicProcurementArticleDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementArticleDetails",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"budget_indent": &graphql.Field{
			Type: DropdownItemType,
		},
		"plan": &graphql.Field{
			Type: DropdownItemType,
		},
		"is_open_procurement": &graphql.Field{
			Type: graphql.Boolean,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"article_type": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_publishing": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_awarding": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(PublicProcurementOrganizationUnitArticleItemType),
		},
	},
})

var PublicProcurementOrganizationUnitArticlesDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementOrganizationUnitArticlesDetails",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementArticleDetailsType),
		},
	},
})

var PublicProcurementOrganizationUnitArticleInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementOrganizationUnitArticleInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: PublicProcurementOrganizationUnitArticleItemType,
		},
	},
})

var PublicProcurementContractItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"public_procurement": &graphql.Field{
			Type: DropdownItemType,
		},
		"supplier": &graphql.Field{
			Type: DropdownItemType,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_signing": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_expiry": &graphql.Field{
			Type: graphql.String,
		},
		"net_value": &graphql.Field{
			Type: graphql.String,
		},
		"gross_value": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var PublicProcurementContractsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractsOverview",
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
			Type: graphql.NewList(PublicProcurementContractItemType),
		},
	},
})

var PublicProcurementContractsInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: PublicProcurementContractItemType,
		},
	},
})

var PublicProcurementContractsDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementContractArticleItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractArticleItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"public_procurement_article": &graphql.Field{
			Type: DropdownPublicProcurementParticleItemType,
		},
		"contract": &graphql.Field{
			Type: DropdownItemType,
		},
		"amount": &graphql.Field{
			Type: graphql.String,
		},
		"net_value": &graphql.Field{
			Type: graphql.String,
		},
		"gross_value": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var DropdownPublicProcurementParticleItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DropdownPublicProcurementParticleItemType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"vat_percentage": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementContractArticlesOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractArticlesOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementContractArticleItemType),
		},
	},
})

var PublicProcurementContractArticleInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractArticleInsert",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: PublicProcurementContractArticleItemType,
		},
	},
})
