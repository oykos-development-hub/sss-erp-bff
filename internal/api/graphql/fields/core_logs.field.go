package fields

import (
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) LogsOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.LogsOverviewType,
		Description: "Returns a data of fixed deposits",
		Args: graphql.FieldConfigArgument{
			"module": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"entity": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"operation": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"item_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"user_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.LogsOverviewResolver,
	}
}
