package mutations

import "github.com/graphql-go/graphql"

var PublicProcurementPlanInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementPlanInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"pre_budget_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"is_pre_budget": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"year": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_publishing": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_closing": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var PublicProcurementPlanItemInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementPlanItemInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_indent_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"plan_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"is_open_procurement": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"article_type": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_publishing": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_awarding": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var PublicProcurementPlanItemLimitInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementPlanItemLimitInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"public_procurement_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"limit": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

var PublicProcurementPlanItemArticleInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementPlanItemArticleInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"public_procurement_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"net_price": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"vat_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"manufacturer": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"visibility_type": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var PublicProcurementOrganizationUnitArticleInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementOrganizationUnitArticleInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"public_procurement_article_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"status": &graphql.InputObjectFieldConfig{
			Type:         graphql.String,
			DefaultValue: "in_progress",
		},
		"is_rejected": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"rejected_description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var PublicProcurementPlanSendOnRevisionMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementPlanSendOnRevisionMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"plan_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var PublicProcurementContractInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementContractInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"public_procurement_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"date_of_signing": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"date_of_expiry": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"net_value": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"gross_value": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"vat_value": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"file": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
})

var PublicProcurementContractArticleInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementContractArticleInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"public_procurement_article_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"public_procurement_contract_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"net_value": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"gross_value": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})

var PublicProcurementContractArticleOverageInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementContractArticleOverageInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"article_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})
