package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"

	"github.com/graphql-go/graphql"
)

var PinResolver = func(p graphql.ResolveParams) (interface{}, error) {
	pin := p.Args["pin"].(string)

	err := validatePin(pin, p.Context.Value(config.HttpHeadersKey).(map[string]string))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "pin is valid",
	}, nil

}

func validatePin(pin string, headers map[string]string) error {
	reqBody := dto.PinRequestMS{
		Pin: pin,
	}

	_, err := shared.MakeAPIRequest("POST", config.PIN_ENDPOINT, reqBody, nil, headers)
	if err != nil {
		return err
	}

	return nil
}
