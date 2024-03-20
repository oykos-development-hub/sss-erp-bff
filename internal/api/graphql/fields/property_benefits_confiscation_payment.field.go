package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) PropBenConfPaymentInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PropBenConfPaymentInsertType,
		Description: "Creates new or alter existing property benefit confiscation payment",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.PropBenConfPaymentMutation),
			},
		},
		Resolve: f.Resolvers.PropBenConfPaymentInsertResolver,
	}
}

func (f *Field) PropBenConfPaymentDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PropBenConfPaymentDeleteType,
		Description: "Delete  property benefit confiscation",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.PropBenConfPaymentDeleteResolver,
	}
}

func (f *Field) PropBenConfPaymentOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PropBenConfPaymentOverviewType,
		Description: "Returns a data of property benefit confiscation items",
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
			"property_benefits_confiscation_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.PropBenConfPaymentOverviewResolver,
	}
}
