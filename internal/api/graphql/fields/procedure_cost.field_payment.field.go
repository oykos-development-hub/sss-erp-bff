package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) ProcedureCostPaymentInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProcedureCostPaymentInsertType,
		Description: "Creates new or alter existing procedure cost payment",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.ProcedureCostPaymentMutation),
			},
		},
		Resolve: f.Resolvers.ProcedureCostPaymentInsertResolver,
	}
}

func (f *Field) ProcedureCostPaymentDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProcedureCostPaymentDeleteType,
		Description: "Delete procedure cost payment",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.ProcedureCostPaymentDeleteResolver,
	}
}

func (f *Field) ProcedureCostPaymentOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProcedureCostPaymentOverviewType,
		Description: "Returns a data of procedure cost payments items",
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
			"procedure_cost_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.ProcedureCostPaymentOverviewResolver,
	}
}
