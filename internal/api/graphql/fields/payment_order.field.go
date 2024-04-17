package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) PaymentOrderInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PaymentOrderInsertType,
		Description: "Creates new or alter existing fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.PaymentOrderMutation),
			},
		},
		Resolve: f.Resolvers.PaymentOrderInsertResolver,
	}
}

func (f *Field) PaymentOrderDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Delete fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PaymentOrderDeleteResolver,
	}
}

func (f *Field) PaymentOrderOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PaymentOrderOverviewType,
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
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"supplier_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.PaymentOrderOverviewResolver,
	}
}

func (f *Field) ObligationsOverview() *graphql.Field {
	return &graphql.Field{
		Type:        types.ObligationsOverviewType,
		Description: "Returns a data of fixed deposits",
		Args: graphql.FieldConfigArgument{
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"supplier_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"type": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: f.Resolvers.ObligationsOverviewResolver,
	}
}

func (f *Field) PayOrderField() *graphql.Field {
	return &graphql.Field{
		Type:        types.FineDeleteType,
		Description: "Delete fixed deposit will",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"sap_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"date_of_sap": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: f.Resolvers.PayOrderResolver,
	}
}
