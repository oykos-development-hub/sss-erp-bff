package fields

import (
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) PinField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PinType,
		Description: "Validates user pin",
		Args: graphql.FieldConfigArgument{
			"pin": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: f.Resolvers.PinResolver,
	}
}
