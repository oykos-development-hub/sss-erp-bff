package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var ProgramInsertField = &graphql.Field{
	Type:        types.ProgramInsertType,
	Description: "Creates new or alter existing Program",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.ProgramMutation),
		},
	},
	Resolve: resolvers.ProgramInsertResolver,
}

var ProgramDeleteField = &graphql.Field{
	Type:        types.ProgramDeleteType,
	Description: "Deleted Program",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.ProgramDeleteResolver,
}

var ProgramOverviewField = &graphql.Field{
	Type:        types.ProgramsOverviewType,
	Description: "Returns a data of Program items",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"program": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: resolvers.ProgramOverviewResolver,
}
