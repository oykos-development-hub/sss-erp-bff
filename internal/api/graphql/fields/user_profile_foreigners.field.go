package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) UserProfileForeignerField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileForeignerType,
		Description: "Returns a data of User Profile for displaying inside Foreigner tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileForeignerResolver,
	}
}

func (f *Field) UserProfileForeignerInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileForeignerInsertType,
		Description: "Creates new or alter existing User Profile's Foreigner item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileForeignerInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileForeignerInsertResolver,
	}
}

func (f *Field) UserProfileForeignerDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileForeignerDeleteType,
		Description: "Deletes existing User Profile's Foreigner",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileForeignerDeleteResolver,
	}
}
