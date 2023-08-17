package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var JobPositionsField = &graphql.Field{
	Type:        types.JobPositionsType,
	Description: "Returns a list of Job Positions",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"search": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JobPositionsResolver,
}

var JobPositionsOrganizationUnitField = &graphql.Field{
	Type:        types.JobPositionsOrganizationUnitType,
	Description: "Returns a list of Job Positions",
	Args: graphql.FieldConfigArgument{
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JobPositionsOrganizationUnitResolver,
}

var JobPositionInsertField = &graphql.Field{
	Type:        types.JobPositionInsertType,
	Description: "Creates new or alter existing Job Position",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.JobPositionInsertMutation),
		},
	},
	Resolve: resolvers.JobPositionInsertResolver,
}

var JobPositionDeleteField = &graphql.Field{
	Type:        types.JobPositionDeleteType,
	Description: "Deletes existing Job Position",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JobPositionDeleteResolver,
}

var JobPositionInOrganizationUnitInsertField = &graphql.Field{
	Type:        types.JobPositionInOrganizationUnitInsertType,
	Description: "Creates new or alter existing Job Position in Organization Unit",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.JobPositionInOrganizationUnitInsertMutation),
		},
	},
	Resolve: resolvers.JobPositionInOrganizationUnitInsertResolver,
}

var JobPositionInOrganizationUnitDeleteField = &graphql.Field{
	Type:        types.JobPositionInOrganizationUnitDeleteType,
	Description: "Deletes existing Job Position in Organization Unit",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JobPositionInOrganizationUnitDeleteResolver,
}
