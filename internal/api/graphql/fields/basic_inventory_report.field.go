package fields

import (
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) ReportValueClassInventoryField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ReportValueClassInventoryType,
		Description: "Returns a Report of Value class Basic Inventory items",
		Args: graphql.FieldConfigArgument{
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"class_type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.ReportValueClassInventoryResolver,
	}
}

func (f *Field) ReportInventoryList() *graphql.Field {
	return &graphql.Field{
		Type:        types.ReportInventoryListType,
		Description: "Returns a Report of inventory list",
		Args: graphql.FieldConfigArgument{
			"date": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"source_type": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"office_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"type": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"is_lager": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
		},
		Resolve: f.Resolvers.ReportInventoryListResolver,
	}
}
