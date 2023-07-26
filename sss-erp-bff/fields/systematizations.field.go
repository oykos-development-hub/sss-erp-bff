package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"
	"github.com/graphql-go/graphql"
)

var SystematizationsOverviewField = &graphql.Field{
	Type:        types.SystematizationsOverviewType,
	Description: "Returns a data of Systematization items",
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
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.SystematizationsOverviewResolver,
}

var SystematizationDetailsField = &graphql.Field{
	Type:        types.SystematizationDetailsType,
	Description: "Returns a data of Systematization item details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.SystematizationResolver,
}

var SystematizationInsertField = &graphql.Field{
	Type:        types.SystematizationDetailsType,
	Description: "Creates new or alter existing Systematization item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.SystematizationInsertMutation),
		},
	},
	Resolve: resolvers.SystematizationInsertResolver,
}

var SystematizationDeleteField = &graphql.Field{
	Type:        types.SystematizationDeleteType,
	Description: "Deletes existing Systematization item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.SystematizationDeleteResolver,
}
