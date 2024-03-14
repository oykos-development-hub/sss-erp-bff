package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) FeePaymentInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FeePaymentInsertType,
		Description: "Creates new or alter existing fee payment",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FeePaymentMutation),
			},
		},
		Resolve: f.Resolvers.FeePaymentInsertResolver,
	}
}

func (f *Field) FeePaymentDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FeePaymentDeleteType,
		Description: "Delete fee payment",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.FeePaymentDeleteResolver,
	}
}

func (f *Field) FeePaymentOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FeePaymentOverviewType,
		Description: "Returns a data of fee payments items",
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
			"fee_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.FeePaymentOverviewResolver,
	}
}
