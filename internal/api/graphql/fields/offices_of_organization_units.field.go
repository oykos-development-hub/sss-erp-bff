package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) OfficesOfOrganizationUnitOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OfficesOfOrganizationUnitOverviewType,
		Description: "Returns a data of Offices Of Organization Unit items",
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
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.OfficesOfOrganizationUnitOverviewResolver,
	}
}
func (f *Field) OfficesOfOrganizationUnitInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OfficesOfOrganizationUnitInsertType,
		Description: "Creates new or alter existing Offices Of Organization Unit",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.OfficesOfOrganizationUnitMutation),
			},
		},
		Resolve: f.Resolvers.OfficesOfOrganizationUnitInsertResolver,
	}
}
func (f *Field) OfficesOfOrganizationUnitDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OfficesOfOrganizationUnitDeleteType,
		Description: "Delete existing Offices Of Organization Unit",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.OfficesOfOrganizationUnitDeleteResolver,
	}
}
