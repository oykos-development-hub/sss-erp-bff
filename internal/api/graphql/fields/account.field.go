package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) AccountInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.AccountInsertType,
		Description: "Creates new or alter existing Account",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.AccountMutation),
			},
		},
		Resolve: f.Resolvers.AccountInsertResolver,
	}
}
func (f *Field) AccountDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.AccountDeleteType,
		Description: "Deleted Account",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.AccountDeleteResolver,
	}
}
func (f *Field) AccountOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.GetAccountOverviewType(),
		Description: "Returns a data of Account items",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"tree": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"level": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"version": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.AccountOverviewResolver,
	}
}
