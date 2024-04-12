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

func (r *Resolver) DepositPaymentOrderOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		DepositPaymentOrder, err := r.Repo.GetDepositPaymentOrderByID(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		res, err := buildDepositPaymentOrder(*DepositPaymentOrder, r)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.DepositPaymentOrderResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.DepositPaymentOrderFilter{}
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

	if value, ok := params.Args["supplier_id"].(int); ok && value != 0 {
		input.SupplierID = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	} else {
		input.OrganizationUnitID, _ = params.Context.Value(config.OrganizationUnitIDKey).(*int)
	}

	items, total, err := r.Repo.GetDepositPaymentOrderList(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var resItems []dto.DepositPaymentOrderResponse
	for _, item := range items {
		resItem, err := buildDepositPaymentOrder(item, r)

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

func (r *Resolver) DepositPaymentOrderInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.DepositPaymentOrder
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

	var item *structs.DepositPaymentOrder

	if data.ID == 0 {
		item, err = r.Repo.CreateDepositPaymentOrder(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdateDepositPaymentOrder(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

	}

	singleItem, err := buildDepositPaymentOrder(*item, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) DepositPaymentOrderDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteDepositPaymentOrder(itemID)
	if err != nil {
		fmt.Printf("Deleting fixed deposit failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildDepositPaymentOrder(item structs.DepositPaymentOrder, r *Resolver) (*dto.DepositPaymentOrderResponse, error) {
	response := dto.DepositPaymentOrderResponse{
		ID:              item.ID,
		CaseNumber:      item.CaseNumber,
		NetAmount:       item.NetAmount,
		BankAccount:     item.BankAccount,
		DateOfPayment:   item.DateOfPayment,
		DateOfStatement: item.DateOfStatement,
		IDOfStatement:   item.IDOfStatement,
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

	if item.SupplierID != 0 {
		value, err := r.Repo.GetAccountItemByID(item.SupplierID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.Supplier = dropdown
	}

	return &response, nil
}
