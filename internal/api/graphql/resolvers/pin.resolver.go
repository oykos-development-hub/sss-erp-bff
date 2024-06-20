package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) PinResolver(p graphql.ResolveParams) (interface{}, error) {
	pin := p.Args["pin"].(string)

	err := r.Repo.ValidatePin(pin, p.Context.Value(config.HTTPHeadersKey).(map[string]string))
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "pin is valid",
	}, nil

}
