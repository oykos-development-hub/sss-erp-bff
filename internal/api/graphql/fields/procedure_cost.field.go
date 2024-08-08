package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) ProcedureCostInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProcedureCostInsertType,
		Description: "Creates new or alter existing procedure cost",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.ProcedureCostMutation),
			},
		},
		Resolve: f.Resolvers.ProcedureCostInsertResolver,
	}
}

func (f *Field) ProcedureCostDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProcedureCostDeleteType,
		Description: "Delete procedure cost",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.ProcedureCostDeleteResolver,
	}
}

func (f *Field) ProcedureCostOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProcedureCostOverviewType,
		Description: "Returns a data of procedure costs items",
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
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"subject": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"procedure_cost_type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.ProcedureCostOverviewResolver,
	}
}
