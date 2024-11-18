package dto

import (
	"fmt"
	"strings"
)

type Response struct {
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	Total       int         `json:"total"`
	Price       float64     `json:"price"`
	Items       interface{} `json:"items"`
	Summary     interface{} `json:"summary"`
	Version     int         `json:"version"`
	Validator   interface{} `json:"validator"`
	SortedItems interface{} `json:"sorted_items"`
}

func ErrorResponse(err error) Response {
	return Response{
		Status:  "error",
		Message: err.Error(),
	}
}

type ResponseSingle struct {
	Status         string      `json:"status"`
	Message        string      `json:"message"`
	EncryptedEmail string      `json:"encrypted_email"`
	Item           interface{} `json:"item,omitempty"`
}

// FormatToEuro takes a float64 and returns it in European currency format with the Euro symbol.
func FormatToEuro(number float64) string {
	// Convert to European euro format with 2 decimal places
	euroString := fmt.Sprintf("%.2f", number)
	euroString = strings.ReplaceAll(euroString, ".", ",") // Replace dot with comma for European style

	// Split the string to add the thousands separator
	parts := strings.Split(euroString, ",")
	parts[0] = addThousandsSeparator(parts[0], '.')

	// Combine the parts and add the Euro symbol at the end
	return strings.Join(parts, ",") + " €"
}

// addThousandsSeparator inserts a thousands separator into a number string.
func addThousandsSeparator(numberString string, separator rune) string {
	var result strings.Builder
	n := len(numberString)

	for i, char := range numberString {
		if i > 0 && (n-i)%3 == 0 {
			result.WriteRune(separator)
		}
		result.WriteRune(char)
	}

	return result.String()
}
