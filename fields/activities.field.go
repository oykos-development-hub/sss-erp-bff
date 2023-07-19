package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var ActivitiesInsertField = &graphql.Field{
	Type:        types.ActivitiesInsertType,
	Description: "Creates new or alter existing Activities",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.ActivitiesMutation),
		},
	},
	Resolve: resolvers.ActivityInsertResolver,
}

var ActivitiesDeleteField = &graphql.Field{
	Type:        types.ActivitiesDeleteType,
	Description: "Deleted Activities",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.ActivityDeleteResolver,
}

var ActivitiesOverviewField = &graphql.Field{
	Type:        types.ActivitiesOverviewType,
	Description: "Returns a data of Activities items",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.ActivitiesOverviewResolver,
}
