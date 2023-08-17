package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var BasicInventoryAssessmentsInsertField = &graphql.Field{
	Type:        types.BasicInventoryAssessmentsInsertType,
	Description: "Creates new or alter existing Basic Inventory Assessment",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.BasicInventoryAssessmentsMutation),
		},
	},
	Resolve: resolvers.BasicInventoryAssessmentsInsertResolver,
}

var BasicInventoryAssessmentsDeleteField = &graphql.Field{
	Type:        types.BasicInventoryAssessmentsDeleteType,
	Description: "Deletes existing Basic Inventory Assessment",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.BasicInventoryAssessmentDeleteResolver,
}
