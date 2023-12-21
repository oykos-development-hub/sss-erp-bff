package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) UserProfileResolutionField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileResolutionType,
		Description: "Returns a data of User Profile for displaying inside Resolution tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileResolutionResolver,
	}
}

func (f *Field) UserProfileResolutionInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileResolutionInsertType,
		Description: "Creates new or alter existing User Profile's Resolution item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileResolutionInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileResolutionInsertResolver,
	}
}

func (f *Field) UserProfileResolutionDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileResolutionDeleteType,
		Description: "Deletes existing User Profile's Resolution",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileResolutionDeleteResolver,
	}
}
