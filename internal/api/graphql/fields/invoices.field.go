package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) InvoiceInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.InvoiceInsertType,
		Description: "Creates new or alter existing invoice",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.InvoiceMutation),
			},
		},
		Resolve: f.Resolvers.InvoiceInsertResolver,
	}
}

func (f *Field) InvoiceDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.InvoiceDeleteType,
		Description: "Delete invoice",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.InvoiceDeleteResolver,
	}
}

func (f *Field) InvoiceOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.InvoiceOverviewType,
		Description: "Returns a data of invoice items",
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
			"supplier_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"year": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"type": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.InvoiceOverviewResolver,
	}
}

func (f *Field) AdditionalExpensesOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.AdditionalExpensesType,
		Description: "Returns a data of additional expenses",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"subject_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"year": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.AdditionalExpensesOverviewResolver,
	}
}

func (f *Field) CalculateAdditionalExpenses() *graphql.Field {
	return &graphql.Field{
		Type:        types.AdditionalExpensesOverviewType,
		Description: "Returns a data of additional expenses",
		Args: graphql.FieldConfigArgument{
			"tax_authority_codebook_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"previous_income": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: f.Resolvers.CalculateAdditionalExpensesResolver,
	}
}
