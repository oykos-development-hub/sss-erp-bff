package types

import "github.com/graphql-go/graphql"

var UserProfileForeignerItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileForeignerItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"work_permit_number": &graphql.Field{
			Type: graphql.String,
		},
		"work_permit_issuer": &graphql.Field{
			Type: graphql.String,
		},
		"work_permit_date_of_start": &graphql.Field{
			Type: graphql.String,
		},
		"work_permit_date_of_end": &graphql.Field{
			Type: graphql.String,
		},
		"work_permit_indefinite_length": &graphql.Field{
			Type: graphql.Boolean,
		},
		"residence_permit_number": &graphql.Field{
			Type: graphql.String,
		},
		"residence_permit_date_of_end": &graphql.Field{
			Type: graphql.String,
		},
		"residence_permit_indefinite_length": &graphql.Field{
			Type: graphql.Boolean,
		},
		"country_of_origin": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
		"work_permit_file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"residence_permit_file_id": &graphql.Field{
			Type: graphql.Int,
		},
		"files": &graphql.Field{
			Type: graphql.NewList(FileDropdownItemType),
		},
	},
})

var UserProfileForeignerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileForeigner",
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
			Type: graphql.NewList(UserProfileForeignerItemType),
		},
	},
})

var UserProfileForeignerInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileForeignerInsert",
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
			Type: UserProfileForeignerItemType,
		},
	},
})

var UserProfileForeignerDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ForeignerDelete",
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
