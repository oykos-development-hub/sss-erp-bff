package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) OrganizationUnitsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrganizationUnitsType,
		Description: "Returns a list of Organization Units",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"parent_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"settings": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
			"disable_filters": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
			"has_president": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
		},
		Resolve: f.Resolvers.OrganizationUnitsResolver,
	}
}
func (f *Field) OrganizationUnitInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrganizationUnitInsertType,
		Description: "Creates new or alter existing Organization Unit",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.OrganizationUnitInsertMutation),
			},
		},
		Resolve: f.Resolvers.OrganizationUnitInsertResolver,
	}
}

func (f *Field) OrganizationUnitOrderField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrganizationUnitOrderType,
		Description: "Updates array of existing Organization Unit",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewList(mutations.OrganizationUnitInsertMutation),
			},
		},
		Resolve: f.Resolvers.OrganizationUnitOrderResolver,
	}
}

func (f *Field) OrganizationUnitDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.OrganizationUnitDeleteType,
		Description: "Deletes existing Organization Unit",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.OrganizationUnitDeleteResolver,
	}
}
