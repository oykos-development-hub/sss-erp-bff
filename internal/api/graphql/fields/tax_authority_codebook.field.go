package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) TaxAuthorityCodebooksOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.TaxAuthorityCodebooksOverviewType,
		Description: "Returns a list of Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.TaxAuthorityCodebooksOverviewResolver,
	}
}

func (f *Field) TaxAuthorityCodebooksInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.TaxAuthorityCodebooksInsertType,
		Description: "Creates new or alter existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.TaxAuthorityCodebookInsertMutation),
			},
		},
		Resolve: f.Resolvers.TaxAuthorityCodebooksInsertResolver,
	}
}

func (f *Field) TaxAuthorityCodebooksDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.TaxAuthorityCodebooksDeleteType,
		Description: "Deletes existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.TaxAuthorityCodebooksDeleteResolver,
	}
}
