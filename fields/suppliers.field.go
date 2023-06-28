package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"
	"github.com/graphql-go/graphql"
)

var SuppliersOverviewField = &graphql.Field{
	Type:        types.SuppliersOverviewType,
	Description: "Returns a data of Suppliers items",
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"search": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.SuppliersOverviewResolver,
}

var SuppliersInsertField = &graphql.Field{
	Type:        types.SuppliersInsertType,
	Description: "Creates new or alter existing Supplier item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.SuppliersInsertMutation),
		},
	},
	Resolve: resolvers.SuppliersInsertResolver,
}

var SuppliersDeleteField = &graphql.Field{
	Type:        types.SuppliersDeleteType,
	Description: "Deletes existing Supplier item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.SuppliersDeleteResolver,
}
