package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) FlatRatePaymentInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FlatRatePaymentInsertType,
		Description: "Creates new or alter existing flatrate payment",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FlatRatePaymentMutation),
			},
		},
		Resolve: f.Resolvers.FlatRatePaymentInsertResolver,
	}
}

func (f *Field) FlatRatePaymentDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FlatRatePaymentDeleteType,
		Description: "Delete flatrate payment",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.FlatRatePaymentDeleteResolver,
	}
}

func (f *Field) FlatRatePaymentOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FlatRatePaymentOverviewType,
		Description: "Returns a data of flatrate payments items",
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
			"flat_rate_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.FlatRatePaymentOverviewResolver,
	}
}
