package mutations

import "github.com/graphql-go/graphql"

var RevisionInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RevisionInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revision_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"revisor_user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"revisor_user_profile": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"internal_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"external_organization_unit_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"responsible_user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"responsible_user_profile": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"implementation_user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"implementation_user_profile": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"planned_year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"planned_quarter": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"priority": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_revision": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_acceptance": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_rejection": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"implementation_suggestion": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"implementation_month_span": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_implementation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"state_of_implementation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"implementation_failed_description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"second_implementation_month_span": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"second_date_of_revision": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"ref_document": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

//---------------------------------------------------------
var RevisionPlanInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RevisionPlanInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"year": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var RevisionsInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RevisionsInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"plan_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"revision_type_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"serial_number": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"internal_revision_subject_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"external_revision_subject_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_revision": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revision_priority": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revision_quartal": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revisor_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

var RevisionTipsInsertMutation = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RevisionTipsInsertMutation",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"responsible_person": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"revision_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_accept": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"due_date": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"new_due_date": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date_of_reject": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"date_of_execution": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"recommendation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"documents": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"reasons_for_non_executing": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"user_profile_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"file_id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})
