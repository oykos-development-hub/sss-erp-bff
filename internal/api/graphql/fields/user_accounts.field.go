package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) UserAccountField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserAccountsOverviewType,
		Description: "Returns a data of User Accounts for displaying on Settings screen",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"is_active": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: nil,
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.UserAccountsOverviewResolver,
	}
}

func (f *Field) UserAccountInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserAccountInsertType,
		Description: "Inserts a data of User Account for displaying on Settings screen",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserAccountInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserAccountBasicInsertResolver,
	}
}

func (f *Field) UserAccountDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserAccountDeleteType,
		Description: "Deletes existing User Account's data",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserAccountDeleteResolver,
	}
}
