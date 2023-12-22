package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) BasicInventoryAssessmentsInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BasicInventoryAssessmentsInsertType,
		Description: "Creates new or alter existing Basic Inventory Assessment",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.BasicInventoryAssessmentsMutation),
			},
		},
		Resolve: f.Resolvers.BasicInventoryAssessmentsInsertResolver,
	}
}

func (f *Field) BasicEXCLInventoryAssessmentsInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BasicEXCLInventoryAssessmentsInsertType,
		Description: "Creates new or alter existing EXCL Basic Inventory Assessment",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.BasicInventoryAssessmentsMutation),
			},
		},
		Resolve: f.Resolvers.BasicEXCLInventoryAssessmentsInsertResolver,
	}
}

func (f *Field) BasicInventoryAssessmentsDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BasicInventoryAssessmentsDeleteType,
		Description: "Deletes existing Basic Inventory Assessment",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.BasicInventoryAssessmentDeleteResolver,
	}
}
