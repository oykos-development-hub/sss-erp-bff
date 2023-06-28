package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"
	"github.com/graphql-go/graphql"
)

var JobTenderTypesField = &graphql.Field{
	Type:        types.JobTenderTypesType,
	Description: "Returns a list of Job Tender types",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"search": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.JobTenderTypesResolver,
}

var JobTenderTypeInsertField = &graphql.Field{
	Type:        types.JobTenderTypeInsertType,
	Description: "Creates new or alter existing Job Tender type",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.JobTenderTypeInsertMutation),
		},
	},
	Resolve: resolvers.JobTenderTypeInsertResolver,
}

var JobTenderTypeDeleteField = &graphql.Field{
	Type:        types.JobTenderTypeDeleteType,
	Description: "Deletes existing Job Tender type",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JobTenderTypeDeleteResolver,
}
