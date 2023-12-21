package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) SystematizationsOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SystematizationsOverviewType,
		Description: "Returns a data of Systematization items",
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
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"year": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"active": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.SystematizationsOverviewResolver,
	}
}

func (f *Field) SystematizationDetailsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SystematizationDetailsType,
		Description: "Returns a data of Systematization item details",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.SystematizationResolver,
	}
}

func (f *Field) SystematizationInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SystematizationDetailsType,
		Description: "Creates new or alter existing Systematization item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.SystematizationInsertMutation),
			},
		},
		Resolve: f.Resolvers.SystematizationInsertResolver,
	}
}

func (f *Field) SystematizationDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SystematizationDeleteType,
		Description: "Deletes existing Systematization item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.SystematizationDeleteResolver,
	}
}
