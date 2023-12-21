package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) JobTendersOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobTenderDetailsType,
		Description: "Returns a data of Job Tenders items",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"active": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: nil,
			},
			"type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JobTenderResolver,
	}
}
func (f *Field) JobTenderInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobTenderInsertType,
		Description: "Creates new or alter existing Job Tender",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.JobTenderInsertMutation),
			},
		},
		Resolve: f.Resolvers.JobTenderInsertResolver,
	}
}
func (f *Field) JobTenderDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobTenderDeleteType,
		Description: "Deletes existing Job Tender",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JobTenderDeleteResolver,
	}
}
func (f *Field) JobTenderApplicationsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobTenderApplicationsOverviewType,
		Description: "Returns a data of Job Tender application items",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"job_tender_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JobTenderApplicationsResolver,
	}
}
func (f *Field) JobTenderApplicationsInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobTenderApplicationInsertType,
		Description: "Creates new or alter existing Job Tender",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.JobTenderApplicationInsertMutation),
			},
		},
		Resolve: f.Resolvers.JobTenderApplicationInsertResolver,
	}
}
func (f *Field) JobTenderApplicationsDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.JobTenderApplicationDeleteType,
		Description: "Deletes existing Job Tender",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.JobTenderApplicationDeleteResolver,
	}
}
