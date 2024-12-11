package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) UserProfileContractsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileContractsType,
		Description: "Returns a data of User Profile's contracts",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.UserProfileContractsResolver,
	}
}

func (f *Field) ActiveContractInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileContractInsertType,
		Description: "Inserts or updates contract of User Profile",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.ActiveContractInsertMutation),
			},
		},
		Resolve: f.Resolvers.ActiveContractMutateResolver,
	}
}

func (f *Field) TerminateEmployment() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileContractDeleteType,
		Description: "Terminate user employment",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"file_ids": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.Int),
			},
		},
		Resolve: f.Resolvers.TerminateEmployment,
	}
}
