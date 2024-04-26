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

func (f *Field) AccountingEntryOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.AccountingEntryOverviewType,
		Description: "Returns a data of fixed deposit wills",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.AccountingEntryOverviewResolver,
	}
}

func (f *Field) AccountingEntryInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.AccountingEntryInsertType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.AccountingEntryMutation),
			},
		},
		Resolve: f.Resolvers.AccountingEntryInsertResolver,
	}
}

func (f *Field) AccountingEntryDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.AccountingEntryDeleteResolver,
	}
}
