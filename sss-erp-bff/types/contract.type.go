package types

import "github.com/graphql-go/graphql"

var ContractTypeType = graphql.NewObject(graphql.ObjectConfig{
    Name: "ContractType",
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

var ContractType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Contract",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.Int,
        },
        "title": &graphql.Field{
            Type: graphql.String,
        },
        "type": &graphql.Field{
            Type: ContractTypeType,
        },
        "abbreviation": &graphql.Field{
            Type: graphql.String,
        },
        "description": &graphql.Field{
            Type: graphql.String,
        },
        "active": &graphql.Field{
            Type: graphql.Boolean,
        },
        "serial_number": &graphql.Field{
            Type: graphql.String,
        },
        "net_salary": &graphql.Field{
            Type: graphql.String,
        },
        "gross_salary": &graphql.Field{
            Type: graphql.String,
        },
        "bank_account": &graphql.Field{
            Type: graphql.String,
        },
        "bank_name": &graphql.Field{
            Type: graphql.String,
        },
        "date_of_signature": &graphql.Field{
            Type: graphql.String,
        },
        "date_of_eligibility": &graphql.Field{
            Type: graphql.String,
        },
        "date_of_start": &graphql.Field{
            Type: graphql.String,
        },
        "date_of_end": &graphql.Field{
            Type: graphql.String,
        },
        "color": &graphql.Field{
            Type: graphql.String,
        },
        "icon": &graphql.Field{
            Type: graphql.String,
        },
        "file_id": &graphql.Field{
            Type: graphql.Int,
        },
    },
})