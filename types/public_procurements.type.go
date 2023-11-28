package types

import (
	"bff/dto"

	"github.com/graphql-go/graphql"
)

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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if plan, ok := p.Source.(*dto.ProcurementPlanResponseItem); ok {
					return string(plan.Status), nil
				}
				return nil, nil
			},
		},
		"data": &graphql.Field{
			Type: JSON,
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
		"requests": &graphql.Field{
			Type: graphql.Int,
		},
		"rejected_description": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanItemDetailsItemType),
		},
		"total_net": &graphql.Field{
			Type: graphql.Int,
		},
		"total_gross": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var PublicProcurementPlansOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlansOverview",
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
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: PublicProcurementPlanDetailsItemType,
		},
	},
})

var PublicProcurementPlanInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanInsert",
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
			Type: PublicProcurementPlanDetailsItemType,
		},
	},
})

var PublicProcurementPlanDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanDelete",
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

var PublicProcurementPlanItemDetailsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemDetailsItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"budget_indent": &graphql.Field{
			Type: DropdownBudgetIndentItemType,
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
		"data": &graphql.Field{
			Type: JSON,
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
		"contract_id": &graphql.Field{
			Type: graphql.Int,
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(PublicProcurementPlanItemArticleItemType),
		},
	},
})

var ContractPdfReportItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ContractPdfReportItem",
	Fields: graphql.Fields{
		"subtitles": &graphql.Field{
			Type: subtitlesType,
		},
		"table_data": &graphql.Field{
			Type: graphql.NewList(tableDataRowType),
		},
	},
})

var PlanPdfReportItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PlanPdfReportItem",
	Fields: graphql.Fields{
		"plan_id": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.String,
		},
		"published_date": &graphql.Field{
			Type: graphql.String,
		},
		"total_gross": &graphql.Field{
			Type: graphql.String,
		},
		"total_vat": &graphql.Field{
			Type: graphql.String,
		},
		"table_data": &graphql.Field{
			Type: graphql.NewList(PlanTableDataRowType),
		},
	},
})

var subtitlesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subtitles",
	Fields: graphql.Fields{
		"public_procurement": &graphql.Field{
			Type: graphql.String,
		},
		"organization_unit": &graphql.Field{
			Type: graphql.String,
		},
		"supplier": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var tableDataRowType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TableDataRow",
	Fields: graphql.Fields{
		"procurement_item": &graphql.Field{
			Type: graphql.String,
		},
		"key_features": &graphql.Field{
			Type: graphql.String,
		},
		"contracted_amount": &graphql.Field{
			Type: graphql.String,
		},
		"available_amount": &graphql.Field{
			Type: graphql.String,
		},
		"consumed_amount": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PlanTableDataRowType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PlanTableDataRow",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"article_type": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"total_gross": &graphql.Field{
			Type: graphql.String,
		},
		"total_vat": &graphql.Field{
			Type: graphql.String,
		},
		"type_of_procedure": &graphql.Field{
			Type: graphql.String,
		},
		"budget_indent": &graphql.Field{
			Type: graphql.String,
		},
		"funding_source": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementPlanItemDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemDetails",
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
			Type: graphql.NewList(PublicProcurementPlanItemDetailsItemType),
		},
	},
})

var PublicProcurementPlanItemPDFType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemPDF",
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
			Type: ContractPdfReportItemType,
		},
	},
})

var PublicProcurementPlanPDFType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanPDF",
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
			Type: PlanPdfReportItemType,
		},
	},
})

var PublicProcurementPlanItemInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemInsert",
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
			Type: PublicProcurementPlanItemDetailsItemType,
		},
	},
})

var PublicProcurementPlanItemDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemDelete",
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
			Type: graphql.Int,
		},
	},
})

var PublicProcurementPlanItemLimitsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemLimits",
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
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: PublicProcurementPlanItemLimitItemType,
		},
	},
})

var PublicProcurementPlanItemArticleItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemArticleItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
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
		"manufacturer": &graphql.Field{
			Type: graphql.String,
		},
		"net_price": &graphql.Field{
			Type: graphql.Float,
		},
		"vat_percentage": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"total_amount": &graphql.Field{
			Type: graphql.Int,
		},
		"visibility_type": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var PublicProcurementPlanItemArticleInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemArticleInsert",
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
			Type: graphql.NewList(PublicProcurementPlanItemArticleItemType),
		},
	},
})

var PublicProcurementPlanItemArticleDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementPlanItemArticleDelete",
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
		"data": &graphql.Field{
			Type: JSON,
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
		"data": &graphql.Field{
			Type: JSON,
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
		"data": &graphql.Field{
			Type: JSON,
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
		"data": &graphql.Field{
			Type: JSON,
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
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: PublicProcurementOrganizationUnitArticleItemType,
		},
	},
})

var PublicProcurementSendPlanOnRevisionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementSendPlanOnRevisionType",
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
			Type: graphql.Float,
		},
		"gross_value": &graphql.Field{
			Type: graphql.Float,
		},
		"vat_value": &graphql.Field{
			Type: graphql.Float,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"days_until_expiry": &graphql.Field{
			Type: graphql.Int,
		},
		"file": &graphql.Field{
			Type: graphql.NewList(FileDropdownItemType),
		},
	},
})

var PublicProcurementContractsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractsOverview",
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
		"data": &graphql.Field{
			Type: JSON,
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
		"data": &graphql.Field{
			Type: JSON,
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
			Type: DropdownPublicProcurementArticleItemType,
		},
		"contract": &graphql.Field{
			Type: DropdownItemType,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"used_articles": &graphql.Field{
			Type: graphql.Int,
		},
		"overages": &graphql.Field{
			Type: graphql.NewList(PublicProcurementContractArticleOverageItemType),
		},
		"overage_total": &graphql.Field{
			Type: graphql.Int,
		},
		"net_value": &graphql.Field{
			Type: graphql.String,
		},
		"gross_value": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var DropdownPublicProcurementArticleItemType = graphql.NewObject(graphql.ObjectConfig{
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
		"data": &graphql.Field{
			Type: JSON,
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
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: graphql.NewList(PublicProcurementContractArticleItemType),
		},
	},
})

var PublicProcurementContractArticleOverageItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractArticleOverageItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"article_id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.Field{
			Type: graphql.Int,
		},
		"amount": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var PublicProcurementContractArticleOverageInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractArticleOverageInsert",
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
			Type: PublicProcurementContractArticleOverageItemType,
		},
	},
})

var PublicProcurementContractArticleOverageDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PublicProcurementContractArticleOverageDelete",
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
