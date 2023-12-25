package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) JobPositionsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobPositionsType,
		Description: "Returns a list of Job Positions",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JobPositionsResolver,
	}
}
func (f *Field) JobPositionsOrganizationUnitField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobPositionsOrganizationUnitType,
		Description: "Returns a list of Job Positions",
		Args: graphql.FieldConfigArgument{
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JobPositionsOrganizationUnitResolver,
	}
}
func (f *Field) JobPositionInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobPositionInsertType,
		Description: "Creates new or alter existing Job Position",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.JobPositionInsertMutation),
			},
		},
		Resolve: f.Resolvers.JobPositionInsertResolver,
	}
}
func (f *Field) JobPositionDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobPositionDeleteType,
		Description: "Deletes existing Job Position",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.JobPositionDeleteResolver,
	}
}
func (f *Field) JobPositionAvailableInOrganizationUnitField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobPositionInOrganizationUnitType,
		Description: "Creates new or alter existing Job Position in Organization Unit",
		Args: graphql.FieldConfigArgument{
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"office_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JobPositionInOrganizationUnitResolver,
	}
}
func (f *Field) JobPositionInOrganizationUnitInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobPositionInOrganizationUnitInsertType,
		Description: "Creates new or alter existing Job Position in Organization Unit",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.JobPositionInOrganizationUnitInsertMutation),
			},
		},
		Resolve: f.Resolvers.JobPositionInOrganizationUnitInsertResolver,
	}
}
func (f *Field) JobPositionInOrganizationUnitDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobPositionInOrganizationUnitDeleteType,
		Description: "Deletes existing Job Position in Organization Unit",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.JobPositionInOrganizationUnitDeleteResolver,
	}
}
