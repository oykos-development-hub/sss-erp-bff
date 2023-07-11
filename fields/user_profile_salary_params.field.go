package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var UserProfileSalaryParamsField = &graphql.Field{
	Type:        types.UserProfileSalaryParamsType,
	Description: "Returns a data of User Profile for displaying inside SalaryParams tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.UserProfileSalaryParamsResolver,
}

var UserProfileSalaryParamsInsertField = &graphql.Field{
	Type:        types.UserProfileSalaryParamsInsertType,
	Description: "Creates new or alter existing User Profile's SalaryParams item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileSalaryParamsInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileSalaryParamsInsertResolver,
}

var UserProfileSalaryParamsDeleteField = &graphql.Field{
	Type:        types.UserProfileSalaryParamsDeleteType,
	Description: "Deletes existing User Profile's SalaryParams",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.UserProfileSalaryParamsDeleteResolver,
}
