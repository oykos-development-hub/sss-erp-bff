package types

import "github.com/graphql-go/graphql"

var RevisionsOverviewItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionsOverviewItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"revision_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"revisor_user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"revision_organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"planned_year": &graphql.Field{
			Type: graphql.String,
		},
		"planned_quarter": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RevisionsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionsOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"revisors": &graphql.Field{
			Type: graphql.NewList(DropdownItemType),
		},
		"items": &graphql.Field{
			Type: graphql.NewList(RevisionsOverviewItemType),
		},
	},
})

var RevisionDetailsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionsDetailsItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"revision_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"revisor_user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"revision_organization_unit": &graphql.Field{
			Type: DropdownItemWithValueType,
		},
		"responsible_user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"implementation_user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"planned_year": &graphql.Field{
			Type: graphql.String,
		},
		"planned_quarter": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"priority": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_revision": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_acceptance": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_rejection": &graphql.Field{
			Type: graphql.String,
		},
		"implementation_suggestion": &graphql.Field{
			Type: graphql.String,
		},
		"implementation_month_span": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_implementation": &graphql.Field{
			Type: graphql.String,
		},
		"state_of_implementation": &graphql.Field{
			Type: graphql.String,
		},
		"implementation_failed_description": &graphql.Field{
			Type: graphql.String,
		},
		"second_implementation_month_span": &graphql.Field{
			Type: graphql.String,
		},
		"second_date_of_revision": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"ref_document": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RevisionDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionDetailsType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: RevisionDetailsItemType,
		},
	},
})

var RevisionDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

//----------------------------------------------------------------------

var RevisionPlanType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionPlan",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RevisionPlansType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionPlansDetails",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(RevisionPlanType),
		},
	},
})

var RevisionPlanOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionPlanOverview",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: RevisionPlanType,
		},
	},
})

//---------------------------------------------------

var RevisionOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionOverviewType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(RevisionType),
		},
		"revisors": &graphql.Field{
			Type: graphql.NewList(DropdownItemType),
		},
	},
})

var RevisionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Revision",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"plan_id": &graphql.Field{
			Type: graphql.Int,
		},
		"internal_revision_subject": &graphql.Field{
			Type: DropdownItemType,
		},
		"external_revision_subject": &graphql.Field{
			Type: DropdownItemType,
		},
		"revisor": &graphql.Field{
			Type: DropdownItemType,
		},
		"revision_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_revision": &graphql.Field{
			Type: graphql.String,
		},
		"revision_priority": &graphql.Field{
			Type: graphql.String,
		},
		"revision_quartal": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RevisionDetailType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionDetailType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: RevisionType,
		},
		"revisors": &graphql.Field{
			Type: graphql.NewList(DropdownItemType),
		},
	},
})

//-----------------------------------------------------------

var RevisionTipsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionTipsOverviewType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(RevisionTipsType),
		},
		"revisors": &graphql.Field{
			Type: graphql.NewList(DropdownItemType),
		},
	},
})

var RevisionTipsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionTips",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"revision_id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"responsible_person": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_accept": &graphql.Field{
			Type: graphql.String,
		},
		"due_date": &graphql.Field{
			Type: graphql.Int,
		},
		"new_due_date": &graphql.Field{
			Type: graphql.Int,
		},
		"date_of_reject": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_execution": &graphql.Field{
			Type: graphql.String,
		},
		"recommendation": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"documents": &graphql.Field{
			Type: graphql.String,
		},
		"reasons_for_non_executing": &graphql.Field{
			Type: graphql.String,
		},
		"file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RevisionTipsDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionTipsDetailsType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"data": &graphql.Field{
			Type: JSON,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"item": &graphql.Field{
			Type: RevisionTipsType,
		},
	},
})
