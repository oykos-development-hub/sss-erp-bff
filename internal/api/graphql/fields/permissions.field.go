package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) PermissionsForRoleField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PermissionsForRoleOverviewType,
		Description: "Returns permissions for role",
		Args: graphql.FieldConfigArgument{
			"role_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PermissionsForRoleResolver,
	}
}
func (f *Field) PermissionsUpdate() *graphql.Field {
	return &graphql.Field{
		Type:        types.PermissionsForRoleOverviewType,
		Description: "Sync Permissions data",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.PermissionsInputMutation),
			},
		},
		Resolve: f.Resolvers.PermissionsUpdateResolver,
	}
}
