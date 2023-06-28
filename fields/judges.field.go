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
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"organization_unit_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"search": &graphql.ArgumentConfig{
			Type: graphql.String,
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
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"year": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.JudgeResolutionsResolver,
}

var JudgeResolutionDetailsField = &graphql.Field{
	Type:        types.JudgeResolutionItemType,
	Description: "Returns Judge resolution item details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.JudgeResolutionsResolver,
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
