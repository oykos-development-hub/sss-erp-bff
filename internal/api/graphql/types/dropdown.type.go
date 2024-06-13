package types

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var DropdownItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DropdownItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var FileDropdownItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FileDropdownItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var DropdownBudgetIndentItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DropdownBudgetIndentItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"serial_number": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var DropdownItemWithValueType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DropdownWithValueItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"value": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var DropdownOUItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DropdownOUItem",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
	},
})
var JSONType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "JSON",
	Description: "The `JSON` scalar type represents JSON values.",
	Serialize: func(value interface{}) interface{} {
		return value
	},
	ParseValue: func(value interface{}) interface{} {
		return value
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch value := valueAST.(type) {
		case *ast.StringValue:
			var result interface{}
			err := json.Unmarshal([]byte(value.Value), &result)
			if err != nil {
				return nil
			}
			return result
		default:
			return nil
		}
	},
})
