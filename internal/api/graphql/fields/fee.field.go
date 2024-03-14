package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) FeeInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FeeInsertType,
		Description: "Creates new or alter existing fee",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FeeMutation),
			},
		},
		Resolve: f.Resolvers.FeeInsertResolver,
	}
}

func (f *Field) FeeDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FeeDeleteType,
		Description: "Delete fee",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.FeeDeleteResolver,
	}
}

func (f *Field) FeeOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FeeOverviewType,
		Description: "Returns a data of fees items",
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
			"fee_type": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"fee_subcategory": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.FeeOverviewResolver,
	}
}
