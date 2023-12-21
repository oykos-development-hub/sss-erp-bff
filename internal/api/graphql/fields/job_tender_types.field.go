package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) JobTenderTypesField() *graphql.Field {
	return &graphql.Field{
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
		Resolve: f.Resolvers.JobTenderTypesResolver,
	}
}
func (f *Field) JobTenderTypeInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobTenderTypeInsertType,
		Description: "Creates new or alter existing Job Tender type",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.JobTenderTypeInsertMutation),
			},
		},
		Resolve: f.Resolvers.JobTenderTypeInsertResolver,
	}
}
func (f *Field) JobTenderTypeDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobTenderTypeDeleteType,
		Description: "Deletes existing Job Tender type",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JobTenderTypeDeleteResolver,
	}
}
