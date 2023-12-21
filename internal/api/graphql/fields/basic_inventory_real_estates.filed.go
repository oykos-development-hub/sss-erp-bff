package fields

import (
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) BasicInventoryRealEstatesOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.BasicInventoryRealEstatesOverviewType,
		Description: "Returns a data of Basic Inventory Real Estates items",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.BasicInventoryRealEstatesOverviewResolver,
	}
}
