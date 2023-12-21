package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) RoleDetailsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RoleDetailsType,
		Description: "Returns details for role",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.RoleDetailsResolver,
	}
}
func (f *Field) RoleOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RoleOverviewType,
		Description: "Returns roles",
		Args:        graphql.FieldConfigArgument{},
		Resolve:     f.Resolvers.RoleOverviewResolver,
	}
}
func (f *Field) RoleInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RoleInsertType,
		Description: "Inserts a data of Role",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.RolesInsertMutation),
			},
		},
		Resolve: f.Resolvers.RolesInsertResolver,
	}
}
