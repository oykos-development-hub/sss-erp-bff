package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) EnforcedPaymentInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.EnforcedPaymentInsertType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.EnforcedPaymentMutation),
			},
		},
		Resolve: f.Resolvers.EnforcedPaymentInsertResolver,
	}
}

func (f *Field) EnforcedPaymentOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.EnforcedPaymentOverviewType,
		Description: "Returns a data of fixed deposit wills",
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
			"year": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"supplier_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.EnforcedPaymentOverviewResolver,
	}
}

func (f *Field) ReturnEnforcedPaymentField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Delete fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"return_file_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"return_date": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: f.Resolvers.ReturnEnforcedPaymentResolver,
	}
}
