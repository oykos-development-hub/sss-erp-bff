package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) UserProfileEvaluationField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileEvaluationType,
		Description: "Returns a data of User Profile for displaying inside Evaluation tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileEvaluationResolver,
	}
}

func (f *Field) UserProfileEvaluationInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileEvaluationInsertType,
		Description: "Creates new or alter existing User Profile's Evaluation item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileEvaluationInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileEvaluationInsertResolver,
	}
}

func (f *Field) UserProfileEvaluationDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileEvaluationDeleteType,
		Description: "Deletes existing User Profile's Evaluation",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileEvaluationDeleteResolver,
	}
}
