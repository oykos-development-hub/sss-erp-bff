package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var OfficesOfOrganizationUnitOverviewField = &graphql.Field{
	Type:        types.OfficesOfOrganizationUnitOverviewType,
	Description: "Returns a data of Offices Of Organization Unit items",
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
	},
	Resolve: resolvers.OfficesOfOrganizationUnitOverviewResolver,
}

var OfficesOfOrganizationUnitInsertField = &graphql.Field{
	Type:        types.OfficesOfOrganizationUnitInsertType,
	Description: "Creates new or alter existing Offices Of Organization Unit",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.OfficesOfOrganizationUnitMutation),
		},
	},
	Resolve: resolvers.OfficesOfOrganizationUnitInsertResolver,
}

var OfficesOfOrganizationUnitDeleteField = &graphql.Field{
	Type:        types.OfficesOfOrganizationUnitDeleteType,
	Description: "Delete existing Offices Of Organization Unit",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.OfficesOfOrganizationUnitDeleteResolver,
}
