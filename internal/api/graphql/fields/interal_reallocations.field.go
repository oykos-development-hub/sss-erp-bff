package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) InternalReallocationOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.InternalReallocationOverviewType,
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
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"year": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"requested_by": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.InternalReallocationOverviewResolver,
	}
}

func (f *Field) InternalReallocationInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.InternalReallocationInsertType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.InternalReallocationMutation),
			},
		},
		Resolve: f.Resolvers.InternalReallocationInsertResolver,
	}
}

func (f *Field) InternalReallocationDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.InternalReallocationDeleteResolver,
	}
}
