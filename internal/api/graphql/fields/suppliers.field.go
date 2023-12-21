package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) SuppliersOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SuppliersOverviewType,
		Description: "Returns a data of Suppliers items",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"entity": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.SuppliersOverviewResolver,
	}
}

func (f *Field) SuppliersInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SuppliersInsertType,
		Description: "Creates new or alter existing Supplier item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.SuppliersInsertMutation),
			},
		},
		Resolve: f.Resolvers.SuppliersInsertResolver,
	}
}

func (f *Field) SuppliersDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SuppliersDeleteType,
		Description: "Deletes existing Supplier item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.SuppliersDeleteResolver,
	}
}
