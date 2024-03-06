package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) FineInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineInsertType,
		Description: "Creates new or alter existing fine",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.FineMutation),
			},
		},
		Resolve: f.Resolvers.FineInsertResolver,
	}
}

func (f *Field) FineDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Delete fine",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.FineDeleteResolver,
	}
}

func (f *Field) FineOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineOverviewType,
		Description: "Returns a data of fines items",
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
			"act_type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.FineOverviewResolver,
	}
}
