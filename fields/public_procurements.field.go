package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var PublicProcurementPlansOverviewField = &graphql.Field{
	Type:        types.PublicProcurementPlansOverviewType,
	Description: "Returns a data of Public Procurement Plan items",
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"sort_by_year": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"sort_by_title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"sort_by_date_of_publishing": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"year": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"is_pre_budget": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"contract": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: resolvers.PublicProcurementPlansOverviewResolver,
}

var PublicProcurementPlanDetailsField = &graphql.Field{
	Type:        types.PublicProcurementPlanDetailsType,
	Description: "Returns Public Procurement Plan item details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"sort_by_title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"sort_by_serial_number": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"sort_by_date_of_publishing": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"sort_by_date_of_awarding": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.PublicProcurementPlanDetailsResolver,
}

var PublicProcurementPlanInsertField = &graphql.Field{
	Type:        types.PublicProcurementPlanInsertType,
	Description: "Creates new or alter existing Public Procurement Plan item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.PublicProcurementPlanInsertMutation),
		},
	},
	Resolve: resolvers.PublicProcurementPlanInsertResolver,
}

var PublicProcurementPlanDeleteField = &graphql.Field{
	Type:        types.PublicProcurementPlanDeleteType,
	Description: "Deletes existing Public Procurement Plan item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.PublicProcurementPlanDeleteResolver,
}

var PublicProcurementPlanItemDetailsField = &graphql.Field{
	Type:        types.PublicProcurementPlanItemDetailsType,
	Description: "Returns Public Procurement item details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"sort_by_title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"sort_by_price": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.PublicProcurementPlanItemDetailsResolver,
}

var PublicProcurementPlanItemPDFField = &graphql.Field{
	Type:        types.PublicProcurementPlanItemPDFType,
	Description: "Returns the PDF URL of Public Procurement",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.PublicProcurementPlanItemPDFResolver,
}

var PublicProcurementPlanPDFField = &graphql.Field{
	Type:        types.PublicProcurementPlanPDFType,
	Description: "Returns the PDF URL of Public Procurement Plan",
	Args: graphql.FieldConfigArgument{
		"plan_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.PublicProcurementPlanPDFResolver,
}

var PublicProcurementPlanItemInsertField = &graphql.Field{
	Type:        types.PublicProcurementPlanItemInsertType,
	Description: "Creates new or alter existing Public Procurement item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.PublicProcurementPlanItemInsertMutation),
		},
	},
	Resolve: resolvers.PublicProcurementPlanItemInsertResolver,
}

var PublicProcurementPlanItemDeleteField = &graphql.Field{
	Type:        types.PublicProcurementPlanItemDeleteType,
	Description: "Deletes existing Public Procurement item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.PublicProcurementPlanItemDeleteResolver,
}

var PublicProcurementPlanItemLimitsField = &graphql.Field{
	Type:        types.PublicProcurementPlanItemLimitsType,
	Description: "Returns all Limits for a specific Public Procurement item",
	Args: graphql.FieldConfigArgument{
		"procurement_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.PublicProcurementPlanItemLimitsResolver,
}

var PublicProcurementPlanItemLimitInsertField = &graphql.Field{
	Type:        types.PublicProcurementPlanItemLimitInsertType,
	Description: "Creates new or alter existing Limits for a specific Public Procurement item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.PublicProcurementPlanItemLimitInsertMutation),
		},
	},
	Resolve: resolvers.PublicProcurementPlanItemLimitInsertResolver,
}

var PublicProcurementPlanItemArticleInsertField = &graphql.Field{
	Type:        types.PublicProcurementPlanItemArticleInsertType,
	Description: "Creates new or alter existing Article for a specific Public Procurement item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewList(mutations.PublicProcurementPlanItemArticleInsertMutation),
		},
	},
	Resolve: resolvers.PublicProcurementPlanItemArticleInsertResolver,
}

var PublicProcurementPlanItemArticleDeleteField = &graphql.Field{
	Type:        types.PublicProcurementPlanItemArticleDeleteType,
	Description: "Deletes existing Article for a specific Public Procurement item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.PublicProcurementPlanItemArticleDeleteResolver,
}

var PublicProcurementOrganizationUnitArticlesOverviewField = &graphql.Field{
	Type:        types.PublicProcurementOrganizationUnitArticlesOverviewType,
	Description: "Returns a data of Public Procurement articles' amount set by Organization Units",
	Args: graphql.FieldConfigArgument{
		"procurement_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.PublicProcurementOrganizationUnitArticlesOverviewResolver,
}

var PublicProcurementOrganizationUnitArticlesDetailsField = &graphql.Field{
	Type:        types.PublicProcurementOrganizationUnitArticlesDetailsType,
	Description: "Returns a details of the request made by Organization Unit for all Public Procurements' Articles inside of one Plan",
	Args: graphql.FieldConfigArgument{
		"plan_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"procurement_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.PublicProcurementOrganizationUnitArticlesDetailsResolver,
}

var PublicProcurementOrganizationUnitArticleInsertField = &graphql.Field{
	Type:        types.PublicProcurementOrganizationUnitArticleInsertType,
	Description: "Creates new or alter existing Public Procurement articles' amount set by Organization Units",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.PublicProcurementOrganizationUnitArticleInsertMutation),
		},
	},
	Resolve: resolvers.PublicProcurementOrganizationUnitArticleInsertResolver,
}

var PublicProcurementSendPlanOnRevision = &graphql.Field{
	Type:        types.PublicProcurementSendPlanOnRevisionType,
	Description: "Send plan on revision",
	Args: graphql.FieldConfigArgument{
		"plan_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.PublicProcurementSendPlanOnRevisionResolver,
}

var PublicProcurementContractsOverviewField = &graphql.Field{
	Type:        types.PublicProcurementContractsOverviewType,
	Description: "Returns a data of Public Procurement Contract items",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"procurement_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"year": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"supplier_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"sort_by_date_of_expiry": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"sort_by_date_of_signing": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"sort_by_gross_value": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"sort_by_serial_number": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.PublicProcurementContractsOverviewResolver,
}

var PublicProcurementContractsInsertField = &graphql.Field{
	Type:        types.PublicProcurementContractsInsertType,
	Description: "Creates new or alter existing Public Procurement Contract item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.PublicProcurementContractInsertMutation),
		},
	},
	Resolve: resolvers.PublicProcurementContractInsertResolver,
}

var PublicProcurementContractsDeleteField = &graphql.Field{
	Type:        types.PublicProcurementContractsDeleteType,
	Description: "Deletes existing Public Procurement Contract item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.PublicProcurementContractDeleteResolver,
}

var PublicProcurementContractArticlesOverviewField = &graphql.Field{
	Type:        types.PublicProcurementContractArticlesOverviewType,
	Description: "Returns a data of Public Procurement Contract articles",
	Args: graphql.FieldConfigArgument{
		"contract_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"visibility_type": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.PublicProcurementContractArticlesOverviewResolver,
}

var PublicProcurementContractOrganizationUnitArticlesOverviewField = &graphql.Field{
	Type:        types.PublicProcurementContractArticlesOverviewType,
	Description: "Returns a data of Public Procurement Contract articles",
	Args: graphql.FieldConfigArgument{
		"contract_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"visibility_type": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.PublicProcurementContractArticlesOrganizationUnitResponseItem,
}

var PublicProcurementContractArticleInsertField = &graphql.Field{
	Type:        types.PublicProcurementContractArticleInsertType,
	Description: "Creates new or alter existing Public Procurement Contract article",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewList(mutations.PublicProcurementContractArticleInsertMutation),
		},
	},
	Resolve: resolvers.PublicProcurementContractArticleInsertResolver,
}

var PublicProcurementContractArticleOverageInsertField = &graphql.Field{
	Type:        types.PublicProcurementContractArticleOverageInsertType,
	Description: "Creates new or alter existing Public Procurement Contract Article Overage",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.PublicProcurementContractArticleOverageInsertMutation),
		},
	},
	Resolve: resolvers.PublicProcurementContractArticleOverageInsertResolver,
}

var PublicProcurementContractArticleOverageDeleteField = &graphql.Field{
	Type:        types.PublicProcurementContractArticleOverageDeleteType,
	Description: "Deletes existing Public Procurement Contract Article Overage Item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.PublicProcurementContractArticleOverageDeleteResolver,
}
