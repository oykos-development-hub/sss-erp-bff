package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) AbsentTypeField() *graphql.Field {
	return &graphql.Field{
		Type:        types.AbsentTypeItemType,
		Description: "Returns a data of Absent Type",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"parent_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.AbsentTypeResolver,
	}
}

func (f *Field) AbsentTypeInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.AbsentTypeInsertType,
		Description: "Creates new or alter existing Absent Type item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.AbsentTypeInsertMutation),
			},
		},
		Resolve: f.Resolvers.AbsentTypeInsertResolver,
	}
}

func (f *Field) AbsentTypeDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserProfileAbsentDeleteType,
		Description: "Deletes existing Absent Type",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.AbsentTypeDeleteResolver,
	}
}
