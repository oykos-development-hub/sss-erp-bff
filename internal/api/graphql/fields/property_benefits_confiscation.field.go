package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) PropBenConfInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PropBenConfInsertType,
		Description: "Creates new or alter existing property benefits confiscation",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.PropBenConfMutation),
			},
		},
		Resolve: f.Resolvers.PropBenConfInsertResolver,
	}
}

func (f *Field) PropBenConfDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PropBenConfDeleteType,
		Description: "Delete property benefits confiscation",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PropBenConfDeleteResolver,
	}
}

func (f *Field) PropBenConfOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PropBenConfOverviewType,
		Description: "Returns a data of property benefit confiscation items",
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
			"property_benefits_confiscation_type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.PropBenConfOverviewResolver,
	}
}
