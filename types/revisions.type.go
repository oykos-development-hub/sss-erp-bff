package types

import "github.com/graphql-go/graphql"

var RevisionsOverviewItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionsOverviewItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
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
			Type: graphql.Boolean,
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
		"revision_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"revisor_user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"revision_organization_unit": &graphql.Field{
			Type: DropdownItemType,
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
	},
})

var RevisionDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionDetailsType",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"items": &graphql.Field{
			Type: graphql.NewList(RevisionDetailsItemType),
		},
	},
})

var RevisionDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RevisionDelete",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
