package types

import (
	"github.com/graphql-go/graphql"
)

var LoginType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Login",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
		"role_id": &graphql.Field{
			Type: graphql.Int,
		},
		"folder_id": &graphql.Field{
			Type: graphql.Int,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"refresh_token": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"first_name": &graphql.Field{
			Type: graphql.String,
		},
		"last_name": &graphql.Field{
			Type: graphql.String,
		},
		"birth_last_name": &graphql.Field{
			Type: graphql.String,
		},
		"gender": &graphql.Field{
			Type: graphql.String,
		},
		"contract": &graphql.Field{
			Type: ContractType,
		},
		"organization_unit": &graphql.Field{
			Type: OrganizationUnitItemType,
		},
		"job_position": &graphql.Field{
			Type: JobPositionItemType,
		},
		"engagement": &graphql.Field{
			Type: EngagementType,
		},
		"permissions": &graphql.Field{
			Type: graphql.NewList(PermissionType),
		},
	},
})
