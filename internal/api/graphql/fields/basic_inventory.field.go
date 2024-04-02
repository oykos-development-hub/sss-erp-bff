package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) BasicInventoryOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BasicInventoryOverviewType,
		Description: "Returns a data of Basic Inventory items",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"type": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"class_type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"office_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"source_type": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"type_of_immovable_property": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"depreciation_type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"expire": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"is_external_donation": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
		},
		Resolve: f.Resolvers.BasicInventoryOverviewResolver,
	}
}
func (f *Field) BasicInventoryDeactivateField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BasicInventoryMessageType,
		Description: "Returns a data of Basic Inventory item Details",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"inactive": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"deactivation_description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"file_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.BasicInventoryDeactivateResolver,
	}
}

func (f *Field) BasicInventoryDetailsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BasicInventoryDetailsType,
		Description: "Returns a data of Basic Inventory item Details",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BasicInventoryDetailsResolver,
	}
}

func (f *Field) BasicInventoryInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BasicInventoryInsertType,
		Description: "Creates new or alter existing Basic Inventory",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.BasicInventoryInsertMutation),
			},
		},
		Resolve: f.Resolvers.BasicInventoryInsertResolver,
	}
}

func (f *Field) InvoicesForInventoryOverview() *graphql.Field {
	return &graphql.Field{
		Type:        types.InvoiceArticleType,
		Description: "Returns a data of Basic Inventory item Details",
		Args: graphql.FieldConfigArgument{
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"supplier_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.InvoicesForInventoryOverview,
	}
}
