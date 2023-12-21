package types

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

var JSON = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "JSON",
	Description: "The `JSON` scalar type represents JSON values as specified by [ECMA-404](http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf).",
	Serialize: func(value interface{}) interface{} {
		return value
	},
	ParseValue: func(value interface{}) interface{} {
		return value
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			var result interface{}
			err := json.Unmarshal([]byte(valueAST.Value), &result)
			if err != nil {
				return nil
			}
			return result
		}
		return nil
	},
})
