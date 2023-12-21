package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) UserProfileSalaryParamsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileSalaryParamsType,
		Description: "Returns a data of User Profile for displaying inside SalaryParams tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileSalaryParamsResolver,
	}
}

func (f *Field) UserProfileSalaryParamsInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileSalaryParamsInsertType,
		Description: "Creates new or alter existing User Profile's SalaryParams item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileSalaryParamsInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileSalaryParamsInsertResolver,
	}
}

func (f *Field) UserProfileSalaryParamsDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileSalaryParamsDeleteType,
		Description: "Deletes existing User Profile's SalaryParams",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileSalaryParamsDeleteResolver,
	}
}
