package types

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/shopspring/decimal"
)

var decimalType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Decimal",
	Description: "The `Decimal` scalar type represents a decimal number with precision. It is serialized as a string.",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case decimal.Decimal:
			val, _ := value.Float64()
			return val
		case decimal.NullDecimal:
			if value.Valid {
				val, _ := value.Decimal.Float64()
				return val
			}
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			dec, err := decimal.NewFromString(value)
			if err == nil {
				return decimal.NullDecimal{Decimal: dec, Valid: true}
			}
		case float64:
			return decimal.NullDecimal{Decimal: decimal.NewFromFloat(value), Valid: true}
		}
		return decimal.NullDecimal{Valid: false}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			dec, err := decimal.NewFromString(valueAST.Value)
			if err == nil {
				return decimal.NullDecimal{Decimal: dec, Valid: true}
			}
		}
		return decimal.NullDecimal{Valid: false}
	},
})
