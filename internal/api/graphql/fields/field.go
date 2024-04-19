package fields

import "bff/internal/api/graphql/resolvers"

type Field struct {
	Resolvers *resolvers.Resolver
}

func NewFields(resolvers *resolvers.Resolver) *Field {
	return &Field{
		Resolvers: resolvers,
	}
}
