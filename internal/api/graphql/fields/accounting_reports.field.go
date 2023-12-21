package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) OverallSpendingField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OverallSpendingType,
		Description: "Returns a data for overall spending report",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.OverallSpendingMutation),
			},
		},
		Resolve: f.Resolvers.OverallSpendingResolver,
	}
}
