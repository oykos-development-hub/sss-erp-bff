package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) UserProfileAbsentField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileAbsentType,
		Description: "Returns a data of User Profile for displaying inside Absent tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileAbsentResolver,
	}
}

func (f *Field) UserProfileVacationField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileVacationType,
		Description: "Returns a data of User Profile's Vacations for displaying inside Vacations tab",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileVacationResolver,
	}
}

func (f *Field) UserProfileVacationInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileVacationInsertType,
		Description: "Creates new or alter existing User Profile's Absent item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileVacationInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileVacationResolutionInsertResolver,
	}
}

func (f *Field) UserProfileAbsentInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileAbsentInsertType,
		Description: "Creates new or alter existing User Profile's Absent item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.UserProfileAbsentInsertMutation),
			},
		},
		Resolve: f.Resolvers.UserProfileAbsentInsertResolver,
	}
}

func (f *Field) UserProfileAbsentDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileAbsentDeleteType,
		Description: "Deletes existing User Profile's Absent",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.UserProfileAbsentDeleteResolver,
	}
}

func (f *Field) TerminateEmployment() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileAbsentDeleteType,
		Description: "Terminate user employment",
		Args: graphql.FieldConfigArgument{
			"user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.TerminateEmployment,
	}
}
