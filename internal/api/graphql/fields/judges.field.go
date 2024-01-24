package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) JudgesOverviewField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.JudgesOverviewResolver,
	}
}
func (f *Field) JudgeNormsInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JudgeNormInsertType,
		Description: "Creates new or alter existing Judge Norm item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.JudgeNormInsertMutation),
			},
		},
		Resolve: f.Resolvers.JudgeNormInsertResolver,
	}
}
func (f *Field) JudgeNormsDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JudgeNormDeleteType,
		Description: "Deletes existing Judge Norm item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JudgeNormDeleteResolver,
	}
}
func (f *Field) JudgeResolutionsOverviewField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.JudgeResolutionsResolver,
	}
}
func (f *Field) OrganizationUintCalculateEmployeeStatsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrganizationUintCalculateEmployeeStatsType,
		Description: "Returns a data of Organization Uint Calculate Employee Stats items",
		Args: graphql.FieldConfigArgument{
			"resolution_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"active": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: true,
			},
		},
		Resolve: f.Resolvers.OrganizationUintCalculateEmployeeStats,
	}
}
func (f *Field) JudgeResolutionsInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JudgeResolutionsInsertType,
		Description: "Creates new or alter existing Judge resolution item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.JudgeResolutionInsertMutation),
			},
		},
		Resolve: f.Resolvers.JudgeResolutionInsertResolver,
	}
}
func (f *Field) JudgeResolutionsDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JudgeResolutionsDeleteType,
		Description: "Deletes existing Judge Norm item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JudgeResolutionDeleteResolver,
	}
}
func (f *Field) CheckJudgeAndPresidentIsAvailableField() *graphql.Field {
	return &graphql.Field{
		Type:        types.CheckJudgeAndPresidentIsAvailableType,
		Description: "Deletes existing Judge Check item",
		Args: graphql.FieldConfigArgument{
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.CheckJudgeAndPresidentIsAvailable,
	}
}
