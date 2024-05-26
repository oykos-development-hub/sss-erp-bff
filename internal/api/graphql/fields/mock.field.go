package fields

import (
	"github.com/graphql-go/graphql"
)

func (f *Field) InsertCurrentBudgetMock() *graphql.Field {
	return &graphql.Field{
		Description: "Inserts a list of data for current budgets",
		Resolve:     f.Resolvers.CurrentBudgetMockResolver,
	}
}
