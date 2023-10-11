package fields

import (
	"bff/mutations"
	"bff/resolvers"
	"bff/types"

	"github.com/graphql-go/graphql"
)

var RevisionsOverviewField = &graphql.Field{
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
	Resolve: resolvers.RevisionsOverviewResolver,
}

var RevisionDetailsField = &graphql.Field{
	Type:        types.RevisionDetailsType,
	Description: "Returns a data of Revision item details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.RevisionDetailsResolver,
}

var RevisionInsertField = &graphql.Field{
	Type:        types.RevisionDetailsType,
	Description: "Creates new or alter existing Revision item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.RevisionInsertMutation),
		},
	},
	Resolve: resolvers.RevisionInsertResolver,
}

var RevisionDeleteField = &graphql.Field{
	Type:        types.RevisionDeleteType,
	Description: "Deletes existing Revision item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.RevisionDeleteResolver,
}

//----------------------------------------------------------------------
//nova polja

var RevisionPlansOverview = &graphql.Field{
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
	Resolve: resolvers.RevisionPlansOverviewResolver,
}

var RevisionPlansDetails = &graphql.Field{
	Type:        types.RevisionPlanOverviewType,
	Description: "Returns a data of Revision plan details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.RevisionPlansDetailsResolver,
}

var RevisionPlansInsertField = &graphql.Field{
	Type:        types.RevisionPlanOverviewType,
	Description: "Creates new or alter existing Revision plan item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.RevisionPlanInsertMutation),
		},
	},
	Resolve: resolvers.RevisionPlanInsertResolver,
}

var RevisionPlansDelete = &graphql.Field{
	Type:        types.RevisionDeleteType,
	Description: "Deletes existing Revision item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.RevisionPlanDeleteResolver,
}

//---------------------------------------

var RevisionOverview = &graphql.Field{
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
	Resolve: resolvers.RevisionOverviewResolver,
}

var RevisionDetails = &graphql.Field{
	Type:        types.RevisionDetailsType,
	Description: "Returns a data of Revision item details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.RevisionDetailResolver,
}

var RevisionInsert = &graphql.Field{
	Type:        types.RevisionDetailType,
	Description: "Creates new or alter existing Revision item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.RevisionsInsertMutation),
		},
	},
	Resolve: resolvers.RevisionsInsertResolver,
}

var RevisionDelete = &graphql.Field{
	Type:        types.RevisionDeleteType,
	Description: "Deletes existing Revision item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.RevisionsDeleteResolver,
}

//---------------------------------------------

var RevisionTipsOverview = &graphql.Field{
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
	Resolve: resolvers.RevisionTipsOverviewResolver,
}

var RevisionTipsDetails = &graphql.Field{
	Type:        types.RevisionTipsDetailsType,
	Description: "Returns a data of Revision item details",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.RevisionTipsDetailsResolver,
}

var RevisionTipsInsert = &graphql.Field{
	Type:        types.RevisionTipsDetailsType,
	Description: "Creates new or alter existing Revision item",
	Args: graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(mutations.RevisionTipsInsertMutation),
		},
	},
	Resolve: resolvers.RevisionTipsInsertResolver,
}

var RevisionTipsDelete = &graphql.Field{
	Type:        types.RevisionDeleteType,
	Description: "Deletes existing Revision item",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.RevisionTipsDeleteResolver,
}
