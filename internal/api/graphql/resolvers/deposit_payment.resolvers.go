package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) DepositPaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		DepositPayment, err := r.Repo.GetDepositPaymentByID(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		res, err := buildDepositPayment(*DepositPayment, r)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.DepositPaymentResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.DepositPaymentFilter{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["status"].(string); ok && value != "" {
		input.Status = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	} else {
		input.OrganizationUnitID, _ = params.Context.Value(config.OrganizationUnitIDKey).(*int)
	}

	items, total, err := r.Repo.GetDepositPaymentList(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var resItems []dto.DepositPaymentResponse
	for _, item := range items {
		resItem, err := buildDepositPayment(item, r)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		resItems = append(resItems, *resItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   resItems,
		Total:   total,
	}, nil
}

func (r *Resolver) DepositPaymentInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.DepositPayment
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		return apierrors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if data.OrganizationUnitID == 0 {

		organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}

		data.OrganizationUnitID = *organizationUnitID

	}

	var item *structs.DepositPayment

	if data.ID == 0 {
		item, err = r.Repo.CreateDepositPayment(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdateDepositPayment(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

	}

	singleItem, err := buildDepositPayment(*item, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) DepositPaymentCaseNumberResolver(params graphql.ResolveParams) (interface{}, error) {
	caseNumber := params.Args["case_number"].(string)

	res, err := r.Repo.GetDepositPaymentCaseNumber(caseNumber)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    res,
	}, nil
}

func (r *Resolver) DepositPaymentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteDepositPayment(itemID)
	if err != nil {
		fmt.Printf("Deleting fixed deposit failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildDepositPayment(item structs.DepositPayment, r *Resolver) (*dto.DepositPaymentResponse, error) {
	response := dto.DepositPaymentResponse{
		ID:                        item.ID,
		Payer:                     item.Payer,
		CaseNumber:                item.CaseNumber,
		PartyName:                 item.PartyName,
		NumberOfBankStatement:     item.NumberOfBankStatement,
		DateOfBankStatement:       item.DateOfBankStatement,
		Amount:                    item.Amount,
		MainBankAccount:           item.MainBankAccount,
		DateOfTransferMainAccount: item.DateOfTransferMainAccount,
		CreatedAt:                 item.CreatedAt,
		UpdatedAt:                 item.UpdatedAt,
	}

	if item.OrganizationUnitID != 0 {
		value, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.OrganizationUnit = dropdown
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
