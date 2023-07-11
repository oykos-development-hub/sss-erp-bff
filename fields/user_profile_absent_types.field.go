package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var AbsentTypeField = &graphql.Field{
	Type:        types.AbsentTypeItemType,
	Description: "Returns a data of Absent Type",
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.AbsentTypeResolver,
}

var AbsentTypeInsertField = &graphql.Field{
	Type:        types.AbsentTypeInsertType,
	Description: "Creates new or alter existing Absent Type item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.AbsentTypeInsertMutation),
		},
	},
	Resolve: resolvers.AbsentTypeInsertResolver,
}

var AbsentTypeDeleteField = &graphql.Field{
	Type:        types.UserProfileAbsentDeleteType,
	Description: "Deletes existing Absent Type",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.AbsentTypeDeleteResolver,
}
