package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"
	"github.com/graphql-go/graphql"
)

var UserProfileEvaluationField = &graphql.Field{
	Type:        types.UserProfileEvaluationType,
	Description: "Returns a data of User Profile for displaying inside Evaluation tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"user_account_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileEvaluationResolver,
}

var UserProfileEvaluationInsertField = &graphql.Field{
	Type:        types.UserProfileEvaluationInsertType,
	Description: "Creates new or alter existing User Profile's Evaluation item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileEvaluationInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileEvaluationInsertResolver,
}

var UserProfileEvaluationDeleteField = &graphql.Field{
	Type:        types.UserProfileEvaluationDeleteType,
	Description: "Deletes existing User Profile's Evaluation",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileEvaluationDeleteResolver,
}
