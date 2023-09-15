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
		"active": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
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
			Type: graphql.Int,
		},
		"plan_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"is_open_procurement": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"article_type": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
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
			Type: graphql.Int,
		},
		"public_procurement_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"limit": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var PublicProcurementPlanItemArticleInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PublicProcurementPlanItemArticleInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"budget_indent_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"public_procurement_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"net_price": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"vat_percentage": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
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
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"is_rejected": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"rejected_description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
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
			Type: graphql.Int,
		},
		"supplier_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_signing": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_expiry": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"net_value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"gross_value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
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
			Type: graphql.Int,
		},
		"public_procurement_contract_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"net_value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"gross_value": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
