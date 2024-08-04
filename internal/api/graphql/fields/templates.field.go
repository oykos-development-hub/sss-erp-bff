package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) TemplateField() *graphql.Field {
	return &graphql.Field{
		Type:        types.TemplateType,
		Description: "Returns a list of Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"template_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.TemplateResolver,
	}
}

func (f *Field) TemplateInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.TemplateInsertType,
		Description: "Creates new or alter existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.TemplateInsertMutation),
			},
		},
		Resolve: f.Resolvers.TemplateInsertResolver,
	}
}

func (f *Field) TemplateItemUpdateField() *graphql.Field {
	return &graphql.Field{
		Type:        types.TemplateInsertType,
		Description: "Creates new or alter existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.TemplateInsertMutation),
			},
		},
		Resolve: f.Resolvers.TemplateItemUpdateResolver,
	}
}

func (f *Field) TemplateUpdateField() *graphql.Field {
	return &graphql.Field{
		Type:        types.TemplateInsertType,
		Description: "Creates new or alter existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.TemplateInsertMutation),
			},
		},
		Resolve: f.Resolvers.TemplateUpdateResolver,
	}
}

func (f *Field) TemplateDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.TemplateDeleteType,
		Description: "Deletes existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.TemplateDeleteResolver,
	}
}

func (f *Field) CustomerSupportUpdateFiled() *graphql.Field {
	return &graphql.Field{
		Type:        types.CustomerSupportInsertType,
		Description: "Creates new or alter existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"user_documentation_file_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.CustomerSupportUpdateResolver,
	}
}

func (f *Field) CustomerSupportOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.CustomerSupportInsertType,
		Description: "Creates new or alter existing Settings Dropdown options",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.CustomerSupportOverviewResolver,
	}
}

func (f *Field) ListOfParametersOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.ListOfParametersOverviewType,
		Description: "Returns a list of parameters",
		Resolve:     f.Resolvers.ListOfParametersOverviewResolver,
	}
}
