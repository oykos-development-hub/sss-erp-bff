package types

import "github.com/graphql-go/graphql"

var EngagementType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Engagement",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.Int,
        },
        "title": &graphql.Field{
            Type: graphql.String,
        },
        "abbreviation": &graphql.Field{
            Type: graphql.String,
        },
        "color": &graphql.Field{
            Type: graphql.String,
        },
        "icon": &graphql.Field{
            Type: graphql.String,
        },
    },
})