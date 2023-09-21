package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var RevisionsOverviewField = &graphql.Field{
	Type:        types.RevisionsOverviewType,
	Description: "Returns a data of Revision items",
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
		"internal": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"revisor_user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.RevisionsOverviewResolver,
}

var RevisionDetailsField = &graphql.Field{
	Type:        types.RevisionDetailsType,
	Description: "Returns a data of Revision item details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.RevisionDetailsResolver,
}

var RevisionInsertField = &graphql.Field{
	Type:        types.RevisionDetailsType,
	Description: "Creates new or alter existing Revision item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.RevisionInsertMutation),
		},
	},
	Resolve: resolvers.RevisionInsertResolver,
}

var RevisionDeleteField = &graphql.Field{
	Type:        types.RevisionDeleteType,
	Description: "Deletes existing Revision item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.RevisionDeleteResolver,
}
