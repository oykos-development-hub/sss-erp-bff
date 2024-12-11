package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) PublicProcurementOrganizationUnitArticleInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementOrganizationUnitArticleInsertType,
		Description: "Creates new or alter existing Public Procurement articles' amount set by Organization Units",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.PublicProcurementOrganizationUnitArticleInsertMutation),
			},
		},
		Resolve: f.Resolvers.PublicProcurementOrganizationUnitArticleInsertResolver,
	}
}

func (f *Field) PublicProcurementSendPlanOnRevision() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementSendPlanOnRevisionType,
		Description: "Send plan on revision",
		Args: graphql.FieldConfigArgument{
			"plan_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PublicProcurementSendPlanOnRevisionResolver,
	}
}

func (f *Field) PublicProcurementPlanInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementPlanInsertType,
		Description: "Creates new or alter existing Public Procurement Plan item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.PublicProcurementPlanInsertMutation),
			},
		},
		Resolve: f.Resolvers.PublicProcurementPlanInsertResolver,
	}
}

func (f *Field) PublicProcurementPlansOverviewField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.PublicProcurementPlansOverviewResolver,
	}
}
func (f *Field) PublicProcurementPlanDetailsField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.PublicProcurementPlanDetailsResolver,
	}
}
func (f *Field) PublicProcurementPlanDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementPlanDeleteType,
		Description: "Deletes existing Public Procurement Plan item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PublicProcurementPlanDeleteResolver,
	}
}
func (f *Field) PublicProcurementPlanItemDetailsField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.PublicProcurementPlanItemDetailsResolver,
	}
}
func (f *Field) PublicProcurementPlanItemPDFField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.PublicProcurementPlanItemPDFResolver,
	}
}
func (f *Field) PublicProcurementPlanPDFField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.PublicProcurementPlanPDFResolver,
	}
}
func (f *Field) PublicProcurementPlanItemInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementPlanItemInsertType,
		Description: "Creates new or alter existing Public Procurement item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.PublicProcurementPlanItemInsertMutation),
			},
		},
		Resolve: f.Resolvers.PublicProcurementPlanItemInsertResolver,
	}
}
func (f *Field) PublicProcurementPlanItemDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementPlanItemDeleteType,
		Description: "Deletes existing Public Procurement item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PublicProcurementPlanItemDeleteResolver,
	}
}
func (f *Field) PublicProcurementPlanItemLimitsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementPlanItemLimitsType,
		Description: "Returns all Limits for a specific Public Procurement item",
		Args: graphql.FieldConfigArgument{
			"procurement_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PublicProcurementPlanItemLimitsResolver,
	}
}
func (f *Field) PublicProcurementPlanItemLimitInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementPlanItemLimitInsertType,
		Description: "Creates new or alter existing Limits for a specific Public Procurement item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.PublicProcurementPlanItemLimitInsertMutation),
			},
		},
		Resolve: f.Resolvers.PublicProcurementPlanItemLimitInsertResolver,
	}
}
func (f *Field) PublicProcurementPlanItemArticleInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementPlanItemArticleInsertType,
		Description: "Creates new or alter existing Article for a specific Public Procurement item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.PublicProcurementPlanItemArticleInsertMutation),
			},
		},
		Resolve: f.Resolvers.PublicProcurementPlanItemArticleInsertResolver,
	}
}
func (f *Field) PublicProcurementPlanItemArticleDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementPlanItemArticleDeleteType,
		Description: "Deletes existing Article for a specific Public Procurement item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PublicProcurementPlanItemArticleDeleteResolver,
	}
}
func (f *Field) PublicProcurementOrganizationUnitArticlesOverviewField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.PublicProcurementOrganizationUnitArticlesOverviewResolver,
	}
}
func (f *Field) PublicProcurementOrganizationUnitArticlesDetailsField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.PublicProcurementOrganizationUnitArticlesDetailsResolver,
	}
}
func (f *Field) PublicProcurementContractsOverviewField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.PublicProcurementContractsOverviewResolver,
	}
}
func (f *Field) PublicProcurementContractsInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementContractsInsertType,
		Description: "Creates new or alter existing Public Procurement Contract item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.PublicProcurementContractInsertMutation),
			},
		},
		Resolve: f.Resolvers.PublicProcurementContractInsertResolver,
	}
}
func (f *Field) PublicProcurementContractsDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementContractsDeleteType,
		Description: "Deletes existing Public Procurement Contract item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PublicProcurementContractDeleteResolver,
	}
}
func (f *Field) PublicProcurementContractArticlesOverviewField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.PublicProcurementContractArticlesOverviewResolver,
	}
}
func (f *Field) PublicProcurementContractOrganizationUnitArticlesOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementContractArticlesOverviewType,
		Description: "Returns a data of Public Procurement Contract articles",
		Args: graphql.FieldConfigArgument{
			"contract_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PublicProcurementContractArticlesOrganizationUnitResponseItem,
	}
}
func (f *Field) PublicProcurementContractArticleInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementContractArticleInsertType,
		Description: "Creates new or alter existing Public Procurement Contract article",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.PublicProcurementContractArticleInsertMutation),
			},
		},
		Resolve: f.Resolvers.PublicProcurementContractArticleInsertResolver,
	}
}
func (f *Field) PublicProcurementContractArticleOverageInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementContractArticleOverageInsertType,
		Description: "Creates new or alter existing Public Procurement Contract Article Overage",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.PublicProcurementContractArticleOverageInsertMutation),
			},
		},
		Resolve: f.Resolvers.PublicProcurementContractArticleOverageInsertResolver,
	}
}
func (f *Field) PublicProcurementContractArticleOverageDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PublicProcurementContractArticleOverageDeleteType,
		Description: "Deletes existing Public Procurement Contract Article Overage Item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PublicProcurementContractArticleOverageDeleteResolver,
	}
}
