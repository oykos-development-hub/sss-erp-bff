package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) SalaryInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SalaryInsertType,
		Description: "Creates new or alter existing fixed deposit",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.SalaryMutation),
			},
		},
		Resolve: f.Resolvers.SalaryInsertResolver,
	}
}

func (f *Field) SalaryDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Delete fixed deposit",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.SalaryDeleteResolver,
	}
}

func (f *Field) SalaryOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SalaryOverviewType,
		Description: "Returns a data of fixed deposits",
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
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"activity_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"year": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.SalaryOverviewResolver,
	}
}
