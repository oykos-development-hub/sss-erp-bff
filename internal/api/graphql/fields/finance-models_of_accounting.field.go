package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) ModelsOfAccountingOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ModelsOfAccountingOverviewType,
		Description: "Returns a data of fixed deposit wills",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.ModelsOfAccountingOverviewResolver,
	}
}

func (f *Field) ModelsOfAccountingUpdateField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ModelsOfAccountingInsertType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.ModelsOfAccountingMutation),
			},
		},
		Resolve: f.Resolvers.ModelsOfAccountingUpdateResolver,
	}
}
