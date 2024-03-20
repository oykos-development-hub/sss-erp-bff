package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) FlatRateInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FlatRateInsertType,
		Description: "Creates new or alter existing flatrate",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FlatRateMutation),
			},
		},
		Resolve: f.Resolvers.FlatRateInsertResolver,
	}
}

func (f *Field) FlatRateDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FlatRateDeleteType,
		Description: "Delete flatrate",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.FlatRateDeleteResolver,
	}
}

func (f *Field) FlatRateOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FlatRateOverviewType,
		Description: "Returns a data of flatrates items",
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
			"subject": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"flat_rate_type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.FlatRateOverviewResolver,
	}
}
