package resolvers

import (
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) GetObligationsForAccountingResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.ObligationsFilter{}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = value
	}

	items, total, err := r.Repo.GetAllObligationsForAccounting(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	message := "Here's the list you asked for!"

	if len(items) == 0 {
		message = "There aren't items!"
	}

	return dto.Response{
		Status:  "success",
		Message: message,
		Items:   items,
		Total:   total,
	}, nil
}

func (r *Resolver) BuildAccountingOrderForObligationsResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.AccountingOrderForObligationsData
	response := dto.Response{
		Status:  "success",
		Message: "You built items!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		return apierrors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	items, err := r.Repo.BuildAccountingOrderForObligations(data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	responseData := dto.AccountingOrderForObligationsResponse{
		DateOfBooking: items.DateOfBooking,
		CreditAmount:  items.CreditAmount,
		DebitAmount:   items.DebitAmount,
	}

	orgUnit, err := r.Repo.GetOrganizationUnitByID(items.OrganizationUnitID)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	dropdown := dto.DropdownSimple{
		ID:    orgUnit.ID,
		Title: orgUnit.Title,
	}

	responseData.OrganizationUnit = dropdown

	for _, item := range items.Items {
		builtItem, err := buildAccountingOrderItemForObligations(item, r)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		responseData.Items = append(responseData.Items, *builtItem)
	}

	response.Items = responseData

	return response, nil
}

func buildAccountingOrderItemForObligations(item dto.AccountingOrderItemsForObligations, r *Resolver) (*dto.AccountingOrderItemsForObligationsResponse, error) {
	response := dto.AccountingOrderItemsForObligationsResponse{
		Title:        item.Title,
		CreditAmount: item.CreditAmount,
		DebitAmount:  item.DebitAmount,
		Type:         item.Type,
		Salary:       item.Salary,
		Invoice:      item.Invoice,
	}

	if item.AccountID != 0 {
		value, err := r.Repo.GetAccountItemByID(item.AccountID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.Account = dropdown
	}

	return &response, nil
}
