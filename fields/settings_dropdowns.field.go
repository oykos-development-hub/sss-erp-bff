package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"
	"github.com/graphql-go/graphql"
)

var SettingsDropdownField = &graphql.Field{
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
	},
	Resolve: resolvers.SettingsDropdownResolver,
}

var SettingsDropdownInsertField = &graphql.Field{
	Type:        types.SettingsDropdownInsertType,
	Description: "Creates new or alter existing Settings Dropdown options",
	Args: graphql.FieldConfigArgument{
		"entity": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.SettingsDropdownInsertMutation),
		},
	},
	Resolve: resolvers.SettingsDropdownInsertResolver,
}

var SettingsDropdownDeleteField = &graphql.Field{
	Type:        types.SettingsDropdownDeleteType,
	Description: "Deletes existing Settings Dropdown options",
	Args: graphql.FieldConfigArgument{
		"entity": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.SettingsDropdownDeleteResolver,
}
