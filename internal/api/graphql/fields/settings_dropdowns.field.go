package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) SettingsDropdownField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SettingsDropdownType,
		Description: "Returns a list of Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"entity": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"value": &graphql.ArgumentConfig{
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
		},
		Resolve: f.Resolvers.SettingsDropdownResolver,
	}
}

func (f *Field) SettingsDropdownInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SettingsDropdownInsertType,
		Description: "Creates new or alter existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.SettingsDropdownInsertMutation),
			},
		},
		Resolve: f.Resolvers.SettingsDropdownInsertResolver,
	}
}

func (f *Field) SettingsDropdownDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.SettingsDropdownDeleteType,
		Description: "Deletes existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.SettingsDropdownDeleteResolver,
	}
}
