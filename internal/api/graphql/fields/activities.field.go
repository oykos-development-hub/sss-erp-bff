package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) ActivitiesInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ActivitiesInsertType,
		Description: "Creates new or alter existing Activities",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.ActivitiesMutation),
			},
		},
		Resolve: f.Resolvers.ActivityInsertResolver,
	}
}
func (f *Field) ActivitiesDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ActivitiesDeleteType,
		Description: "Deleted Activities",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.ActivityDeleteResolver,
	}
}
func (f *Field) ActivitiesOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ActivitiesOverviewType,
		Description: "Returns a data of Activities items",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"sub_program_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.ActivitiesOverviewResolver,
	}
}
