package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"time"

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

func (r *Resolver) SpendingReleaseOverview(params graphql.ResolveParams) (interface{}, error) {
	loggedInOrganizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok {
		return errors.HandleAPPError(errors.NewBadRequestError("Error getting logged in unit"))
	}

	input := &dto.GetSpendingReleaseListInput{}

	if unitID, ok := params.Args["unit_id"].(int); ok && unitID != 0 {
		input.UnitID = unitID
	} else {
		input.UnitID = *loggedInOrganizationUnitID
	}

	if year, ok := params.Args["year"].(int); ok && year != 0 {
		input.Year = &year
	} else {
		year = time.Now().Year()
	}

	if month, ok := params.Args["month"].(int); ok && month != 0 {
		input.Year = &month
	}

	// spendingReleaseList, err := r.Repo.GetSpendingReleaseList(input)
	// if err != nil {
	// 	var apiErr *errors.APIError
	// 	if goerrors.As(err, &apiErr) {
	// 		if apiErr.StatusCode != 404 {
	// 			return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic"))
	// 		}
	// 	}
	// }

	return dto.Response{
		Status:  "success",
		Message: "Here's the data you asked for!",
		// Items:   spendingReleaseList,
	}, nil
}
