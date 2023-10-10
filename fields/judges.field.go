package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var JudgesOverviewField = &graphql.Field{
	Type:        types.JudgesOverviewType,
	Description: "Returns a data of Judge items",
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 1,
		},
		"size": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 10,
		},
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JudgesOverviewResolver,
}

var JudgeNormsInsertField = &graphql.Field{
	Type:        types.JudgeNormInsertType,
	Description: "Creates new or alter existing Judge Norm item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.JudgeNormInsertMutation),
		},
	},
	Resolve: resolvers.JudgeNormInsertResolver,
}

var JudgeNormsDeleteField = &graphql.Field{
	Type:        types.JudgeNormDeleteType,
	Description: "Deletes existing Judge Norm item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JudgeNormDeleteResolver,
}

var JudgeResolutionsOverviewField = &graphql.Field{
	Type:        types.JudgeResolutionsOverviewType,
	Description: "Returns a data of Judge resolution items",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 1,
		},
		"size": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			DefaultValue: 10,
		},
	},
	Resolve: resolvers.JudgeResolutionsResolver,
}

var OrganizationUintCalculateEmployeeStatsField = &graphql.Field{
	Type:        types.OrganizationUintCalculateEmployeeStatsType,
	Description: "Returns a data of Organization Uint Calculate Employee Stats items",
	Args: graphql.FieldConfigArgument{
		"resolution_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.OrganizationUintCalculateEmployeeStats,
}

var JudgeResolutionsInsertField = &graphql.Field{
	Type:        types.JudgeResolutionsInsertType,
	Description: "Creates new or alter existing Judge resolution item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.JudgeResolutionInsertMutation),
		},
	},
	Resolve: resolvers.JudgeResolutionInsertResolver,
}

var JudgeResolutionsDeleteField = &graphql.Field{
	Type:        types.JudgeResolutionsDeleteType,
	Description: "Deletes existing Judge Norm item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JudgeResolutionDeleteResolver,
}

var CheckJudgeAndPresidentIsAvailableField = &graphql.Field{
	Type:        types.CheckJudgeAndPresidentIsAvailableType,
	Description: "Deletes existing Judge Check item",
	Args: graphql.FieldConfigArgument{
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.CheckJudgeAndPresidentIsAvailable,
}
