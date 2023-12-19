package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var OrganizationUnitsField = &graphql.Field{
	Type:        types.OrganizationUnitsType,
	Description: "Returns a list of Organization Units",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"search": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"parent_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"settings": &graphql.ArgumentConfig{
			Type:         graphql.Boolean,
			DefaultValue: false,
		},
	},
	Resolve: resolvers.OrganizationUnitsResolver,
}

var OrganizationUnitInsertField = &graphql.Field{
	Type:        types.OrganizationUnitInsertType,
	Description: "Creates new or alter existing Organization Unit",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.OrganizationUnitInsertMutation),
		},
	},
	Resolve: resolvers.OrganizationUnitInsertResolver,
}

var OrganizationUnitDeleteField = &graphql.Field{
	Type:        types.OrganizationUnitDeleteType,
	Description: "Deletes existing Organization Unit",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.OrganizationUnitDeleteResolver,
}
