package types

import "github.com/graphql-go/graphql"

var UserProfileSalaryParamsItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileSalaryParamsItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile": &graphql.Field{
			Type: DropdownItemType,
		},
		"organization_unit": &graphql.Field{
			Type: DropdownItemType,
		},
		"benefited_track": &graphql.Field{
			Type: graphql.Boolean,
		},
		"without_raise": &graphql.Field{
			Type: graphql.Boolean,
		},
		"insurance_basis": &graphql.Field{
			Type: graphql.String,
		},
		"salary_rank": &graphql.Field{
			Type: graphql.String,
		},
		"daily_work_hours": &graphql.Field{
			Type: graphql.String,
		},
		"weekly_work_hours": &graphql.Field{
			Type: graphql.String,
		},
		"education_rank": &graphql.Field{
			Type: graphql.String,
		},
		"education_naming": &graphql.Field{
			Type: graphql.String,
		},
		"resolution": &graphql.Field{
			Type: UserProfileResolutionItemType,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserProfileSalaryParamsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileSalaryParams",
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
			Type: graphql.NewList(UserProfileSalaryParamsItemType),
		},
	},
})

var UserProfileSalaryParamsInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileSalaryParamsInsert",
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
			Type: UserProfileSalaryParamsItemType,
		},
	},
})

var UserProfileSalaryParamsDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SalaryParamsDelete",
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
