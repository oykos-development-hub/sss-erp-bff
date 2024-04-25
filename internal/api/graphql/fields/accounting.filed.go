package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) GetObligationsForAccounting() *graphql.Field {
	return &graphql.Field{
		Type:        types.ObligationsForAccountingOverviewType,
		Description: "Returns a data of fixed deposits",
		Args: graphql.FieldConfigArgument{
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.GetObligationsForAccountingResolver,
	}
}

func (f *Field) BuildAccountingOrderForObligationsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.AccountingOrderForObligationsOverviewType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.AccountingOrderForObligationsMutation),
			},
		},
		Resolve: f.Resolvers.BuildAccountingOrderForObligationsResolver,
	}
}
