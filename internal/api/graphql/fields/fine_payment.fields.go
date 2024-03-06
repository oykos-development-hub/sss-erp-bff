package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) FinePaymentInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FinePaymentInsertType,
		Description: "Creates new or alter existing fine payment",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FinePaymentMutation),
			},
		},
		Resolve: f.Resolvers.FinePaymentInsertResolver,
	}
}

func (f *Field) FinePaymentDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FinePaymentDeleteType,
		Description: "Delete fine payment",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.FinePaymentDeleteResolver,
	}
}

func (f *Field) FinePaymentOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FinePaymentOverviewType,
		Description: "Returns a data of fine payments items",
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
			"fine_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.FinePaymentOverviewResolver,
	}
}
