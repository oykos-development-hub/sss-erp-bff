package types

import "github.com/graphql-go/graphql"

var UserProfileEvaluationItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileEvaluationItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_profile_id": &graphql.Field{
			Type: graphql.Int,
		},
		"evaluation_type": &graphql.Field{
			Type: DropdownItemType,
		},
		"date_of_evaluation": &graphql.Field{
			Type: graphql.String,
		},
		"score": &graphql.Field{
			Type: graphql.String,
		},
		"evaluator": &graphql.Field{
			Type: graphql.String,
		},
		"is_relevant": &graphql.Field{
			Type: graphql.Boolean,
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
		"file": &graphql.Field{
			Type: FileDropdownItemType,
		},
	},
})

var UserProfileEvaluationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileEvaluation",
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
			Type: graphql.NewList(UserProfileEvaluationItemType),
		},
	},
})

var UserProfileEvaluationInsertType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserProfileEvaluationInsert",
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
			Type: UserProfileEvaluationItemType,
		},
	},
})

var UserProfileEvaluationDeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EvaluationDelete",
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
