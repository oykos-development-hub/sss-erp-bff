package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) ExternalReallocationOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ExternalReallocationOverviewType,
		Description: "Returns a data of fixed deposit wills",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"source_organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"destination_organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"year": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"requested_by": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"budget_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.ExternalReallocationOverviewResolver,
	}
}

func (f *Field) ExternalReallocationInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ExternalReallocationInsertType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.ExternalReallocationMutation),
			},
		},
		Resolve: f.Resolvers.ExternalReallocationInsertResolver,
	}
}

func (f *Field) ExternalReallocationDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.ExternalReallocationDeleteResolver,
	}
}

func (f *Field) ExternalReallocationOUAcceptField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ExternalReallocationInsertType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.ExternalReallocationMutation),
			},
		},
		Resolve: f.Resolvers.ExternalReallocationOUAcceptResolver,
	}
}

func (f *Field) ExternalReallocationOURejectField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.ExternalReallocationOURejectResolver,
	}
}
