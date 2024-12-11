package fields

import (
	"bff/internal/api/graphql/mutations"
	"bff/internal/api/graphql/types"

	"github.com/graphql-go/graphql"
)

func (f *Field) RevisionsOverviewField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionsOverviewType,
		Description: "Returns a data of Revision items",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"organization_unit_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"internal": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"revisor_user_profile_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionsOverviewResolver,
	}
}
func (f *Field) RevisionDetailsField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionDetailsType,
		Description: "Returns a data of Revision item details",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.RevisionDetailsResolver,
	}
}
func (f *Field) RevisionInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionDetailsType,
		Description: "Creates new or alter existing Revision item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.RevisionInsertMutation),
			},
		},
		Resolve: f.Resolvers.RevisionInsertResolver,
	}
}
func (f *Field) RevisionDeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionDeleteType,
		Description: "Deletes existing Revision item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionDeleteResolver,
	}
}

//----------------------------------------------------------------------
//nova polja

func (f *Field) RevisionPlansOverview() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionPlansType,
		Description: "Returns a data of Revision plans",
		Args: graphql.FieldConfigArgument{
			"year": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionPlansOverviewResolver,
	}
}
func (f *Field) RevisionPlansDetails() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionPlanOverviewType,
		Description: "Returns a data of Revision plan details",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionPlansDetailsResolver,
	}
}
func (f *Field) RevisionPlansInsertField() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionPlanOverviewType,
		Description: "Creates new or alter existing Revision plan item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.RevisionPlanInsertMutation),
			},
		},
		Resolve: f.Resolvers.RevisionPlanInsertResolver,
	}
}
func (f *Field) RevisionPlansDelete() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionDeleteType,
		Description: "Deletes existing Revision item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionPlanDeleteResolver,
	}
}

//---------------------------------------

func (f *Field) RevisionOverview() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionOverviewType,
		Description: "Returns a data of Revision item details",
		Args: graphql.FieldConfigArgument{
			"plan_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"internal_revision_subject_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"revision_type_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"revisor_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionOverviewResolver,
	}
}
func (f *Field) RevisionDetails() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionDetailsType,
		Description: "Returns a data of Revision item details",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.RevisionDetailResolver,
	}
}
func (f *Field) RevisionInsert() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionDetailType,
		Description: "Creates new or alter existing Revision item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.RevisionsInsertMutation),
			},
		},
		Resolve: f.Resolvers.RevisionsInsertResolver,
	}
}
func (f *Field) RevisionDelete() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionDeleteType,
		Description: "Deletes existing Revision item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionsDeleteResolver,
	}
}

//---------------------------------------------

func (f *Field) RevisionTipsOverview() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionTipsOverviewType,
		Description: "Returns a data of Revision tips items",
		Args: graphql.FieldConfigArgument{
			"revision_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionTipsOverviewResolver,
	}
}
func (f *Field) RevisionTipsDetails() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionTipsDetailsType,
		Description: "Returns a data of Revision item details",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: f.Resolvers.RevisionTipsDetailsResolver,
	}
}
func (f *Field) RevisionTipsInsert() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionTipsDetailsType,
		Description: "Creates new or alter existing Revision item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.RevisionTipsInsertMutation),
			},
		},
		Resolve: f.Resolvers.RevisionTipsInsertResolver,
	}
}
func (f *Field) RevisionTipsDelete() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionDeleteType,
		Description: "Deletes existing Revision item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionTipsDeleteResolver,
	}
}

func (f *Field) RevisionTipImplementationOverview() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionTipImplementationOverviewType,
		Description: "Returns a data of Revision tip implementation",
		Args: graphql.FieldConfigArgument{
			"tip_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionTipImplementationOverviewResolver,
	}
}

func (f *Field) RevisionTipImplementationInsert() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionTipsImplementationDetailsType,
		Description: "Creates new or alter existing revision tip implementation item",
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(mutations.RevisionTipImplementationInsertMutation),
			},
		},
		Resolve: f.Resolvers.RevisionTipImplementationInsertResolver,
	}
}

func (f *Field) RevisionTipImplementationDelete() *graphql.Field {
	return &graphql.Field{
		Type:        types.RevisionDeleteType,
		Description: "Deletes existing revision tip implementation item",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: f.Resolvers.RevisionTipImplementationDeleteResolver,
	}
}
