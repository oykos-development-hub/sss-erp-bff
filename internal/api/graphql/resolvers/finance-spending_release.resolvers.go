package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) SpendingReleaseInsert(params graphql.ResolveParams) (interface{}, error) {
	var data structs.SpendingReleaseInsert

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	item, err := r.Repo.CreateSpendingRelease(params.Context, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
		Item:    item,
	}, nil
}
