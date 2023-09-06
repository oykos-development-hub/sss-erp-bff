package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var JobTendersOverviewField = &graphql.Field{
	Type:        types.JobTenderDetailsType,
	Description: "Returns a data of Job Tenders items",
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"active": &graphql.ArgumentConfig{
			Type:         graphql.Boolean,
			DefaultValue: nil,
		},
		"type_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JobTenderResolver,
}

var JobTenderInsertField = &graphql.Field{
	Type:        types.JobTenderInsertType,
	Description: "Creates new or alter existing Job Tender",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.JobTenderInsertMutation),
		},
	},
	Resolve: resolvers.JobTenderInsertResolver,
}

var JobTenderDeleteField = &graphql.Field{
	Type:        types.JobTenderDeleteType,
	Description: "Deletes existing Job Tender",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JobTenderDeleteResolver,
}

var JobTenderApplicationsField = &graphql.Field{
	Type:        types.JobTenderApplicationsOverviewType,
	Description: "Returns a data of Job Tender application items",
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
		"job_tender_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JobTenderApplicationsResolver,
}

var JobTenderApplicationsInsertField = &graphql.Field{
	Type:        types.JobTenderApplicationInsertType,
	Description: "Creates new or alter existing Job Tender",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.JobTenderApplicationInsertMutation),
		},
	},
	Resolve: resolvers.JobTenderApplicationInsertResolver,
}

var JobTenderApplicationsDeleteField = &graphql.Field{
	Type:        types.JobTenderApplicationDeleteType,
	Description: "Deletes existing Job Tender",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JobTenderApplicationDeleteResolver,
}
