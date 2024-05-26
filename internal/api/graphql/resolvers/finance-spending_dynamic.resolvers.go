package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	goerrors "errors"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) SpendingDynamicInsert(params graphql.ResolveParams) (interface{}, error) {
	var data []structs.SpendingDynamicInsert

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	items, err := r.Repo.CreateSpendingDynamic(params.Context, data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You created this item!",
		Items:   items,
	}, nil
}

func (r *Resolver) SpendingDynamicOverview(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)

	if unitID == 0 {
		loggedInOrganizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok {
			return errors.HandleAPPError(errors.NewBadRequestError("Error getting logged in unit"))
		}

		unitID = *loggedInOrganizationUnitID
	}

	spendingDynamic, err := r.Repo.GetSpendingDynamic(budgetID, unitID)
	if err != nil {
		var apiErr *errors.APIError
		if goerrors.As(err, &apiErr) {
			if apiErr.StatusCode != 404 {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic"))
			}
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the data you asked for!",
		Items:   spendingDynamic,
	}, nil
}

func (r *Resolver) SpendingDynamicHistoryOverview(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)
	version := params.Args["version"].(int)

	if unitID == 0 {
		loggedInOrganizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok {
			return errors.HandleAPPError(errors.NewBadRequestError("Error getting logged in unit"))
		}

		unitID = *loggedInOrganizationUnitID
	}

	input := &dto.GetSpendingDynamicHistoryInput{}
	if version != 0 {
		input.Version = &version
	}

	spendingDynamicHistory, err := r.Repo.GetSpendingDynamicHistory(budgetID, unitID, input)
	if err != nil {
		var apiErr *errors.APIError
		if goerrors.As(err, &apiErr) {
			if apiErr.StatusCode != 404 {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic"))
			}
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the data you asked for!",
		Items:   spendingDynamicHistory,
	}, nil
}
