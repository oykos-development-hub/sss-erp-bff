package types

import "github.com/graphql-go/graphql"

var JobTenderItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"type": &graphql.Field{
			Type: DropdownItemType,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"date_of_start": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_end": &graphql.Field{
			Type: graphql.String,
		},
		"number_of_vacant_seats": &graphql.Field{
			Type: graphql.Int,
		},
		"file": &graphql.Field{
			Type: FileDropdownItemType,
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

var JobTenderDetailsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenders",
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
		"items": &graphql.Field{
			Type: graphql.NewList(JobTenderItemType),
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var JobTenderInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderInsert",
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
			Type: JobTenderItemType,
		},
	},
})

var JobTenderDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderDelete",
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

var JobTenderApplicationItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderApplicationItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"job_tender": &graphql.Field{
			Type: DropdownItemType,
		},
		"tender_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"first_name": &graphql.Field{
			Type: graphql.String,
		},
		"last_name": &graphql.Field{
			Type: graphql.String,
		},
		"official_personal_id": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_birth": &graphql.Field{
			Type: graphql.String,
		},
		"citizenship": &graphql.Field{
			Type: graphql.String,
		},
		"evaluation": &graphql.Field{
			Type: graphql.String,
		},
		"date_of_application": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
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

var JobTenderApplicationsOverviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderApplicationsOverview",
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
			Type: graphql.NewList(JobTenderApplicationItemType),
		},
	},
})

var JobTenderApplicationInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderApplicationInsert",
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
			Type: JobTenderApplicationItemType,
		},
	},
})

var JobTenderApplicationDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "JobTenderApplicationDelete",
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
