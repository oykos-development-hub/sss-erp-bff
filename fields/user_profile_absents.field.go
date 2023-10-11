package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var UserProfileAbsentField = &graphql.Field{
	Type:        types.UserProfileAbsentType,
	Description: "Returns a data of User Profile for displaying inside Absent tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileAbsentResolver,
}

var UserProfileVacationField = &graphql.Field{
	Type:        types.UserProfileVacationType,
	Description: "Returns a data of User Profile's Vacations for displaying inside Vacations tab",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileVacationResolver,
}

var UserProfileVacationInsertField = &graphql.Field{
	Type:        types.UserProfileVacationInsertType,
	Description: "Creates new or alter existing User Profile's Absent item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileVacationInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileVacationResolutionInsertResolver,
}

var UserProfileAbsentInsertField = &graphql.Field{
	Type:        types.UserProfileAbsentInsertType,
	Description: "Creates new or alter existing User Profile's Absent item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.UserProfileAbsentInsertMutation),
		},
	},
	Resolve: resolvers.UserProfileAbsentInsertResolver,
}

var UserProfileAbsentDeleteField = &graphql.Field{
	Type:        types.UserProfileAbsentDeleteType,
	Description: "Deletes existing User Profile's Absent",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.UserProfileAbsentDeleteResolver,
}

var TerminateEmployment = &graphql.Field{
	Type:        types.UserProfileAbsentDeleteType,
	Description: "Terminate user employment",
	Args: graphql.FieldConfigArgument{
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.TerminateEmployment,
}
