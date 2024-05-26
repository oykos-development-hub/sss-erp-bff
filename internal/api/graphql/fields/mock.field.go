package fields

import (
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) InsertCurrentBudgetMock() *graphql.Field {
	return &graphql.Field{
		Type:        types.CurrentBudgetMockType,
		Description: "Inserts a list of data for current budgets",
		Resolve:     f.Resolvers.CurrentBudgetMockResolver,
	}
}
