package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) ProgramInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProgramInsertType,
		Description: "Creates new or alter existing Program",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.ProgramMutation),
			},
		},
		Resolve: f.Resolvers.ProgramInsertResolver,
	}
}

func (f *Field) ProgramDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProgramDeleteType,
		Description: "Deleted Program",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.ProgramDeleteResolver,
	}
}

func (f *Field) ProgramOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProgramsOverviewType,
		Description: "Returns a data of Program items",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"is_program": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.ProgramOverviewResolver,
	}
}
