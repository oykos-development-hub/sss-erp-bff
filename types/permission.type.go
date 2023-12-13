package types

import (
	"bff/dto"
	"sync"

	"github.com/graphql-go/graphql"
)

var (
	permissionForRoleType          *graphql.Object
	PermissionsForRoleOverviewType *graphql.Object
	once                           sync.Once
)

func GetPermissionForRoleType() *graphql.Object {
	once.Do(func() {
		initPermissionForRoleType()
		initPermissionForRoleOverviewType()
	})
	return permissionForRoleType
}

func initPermissionForRoleType() {
	permissionForRoleType = graphql.NewObject(graphql.ObjectConfig{
		Name: "PermissionForRoleType",
		Fields: (graphql.FieldsThunk)(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"parent_id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"route": &graphql.Field{
					Type: graphql.String,
				},
				"create": &graphql.Field{
					Type: graphql.Boolean,
				},
				"read": &graphql.Field{
					Type: graphql.Boolean,
				},
				"update": &graphql.Field{
					Type: graphql.Boolean,
				},
				"delete": &graphql.Field{
					Type: graphql.Boolean,
				},
				"children": &graphql.Field{
					Type: graphql.NewList(permissionForRoleType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if permissionItem, ok := p.Source.(*dto.PermissionNode); ok {
							return permissionItem.Children, nil
						}
						return nil, nil
					},
				},
			}
		}),
	})
}

func initPermissionForRoleOverviewType() {
	PermissionsForRoleOverviewType = graphql.NewObject(graphql.ObjectConfig{
		Name: "PermissionsForRoleOverview",
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
				Type: permissionForRoleType,
			},
		},
	})
}
