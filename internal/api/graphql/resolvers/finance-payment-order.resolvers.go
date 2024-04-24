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

func (r *Resolver) PaymentOrderOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		PaymentOrder, err := r.Repo.GetPaymentOrderByID(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		res, err := buildPaymentOrder(*PaymentOrder, r)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.PaymentOrderResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.PaymentOrderFilter{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["year"].(int); ok && value != 0 {
		input.Year = &value
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

	items, total, err := r.Repo.GetPaymentOrderList(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var resItems []dto.PaymentOrderResponse
	for _, item := range items {
		resItem, err := buildPaymentOrder(item, r)

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

func (r *Resolver) PaymentOrderInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PaymentOrder
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

	var item *structs.PaymentOrder

	if data.ID == 0 {
		item, err = r.Repo.CreatePaymentOrder(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdatePaymentOrder(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

	}

	singleItem, err := buildPaymentOrder(*item, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) PaymentOrderDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeletePaymentOrder(itemID)
	if err != nil {
		fmt.Printf("Deleting fixed deposit failed because of this error - %s.\n", err)
		return dto.ResponseSingle{
			Status: "failed",
		}, nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) ObligationsOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.ObligationsFilter{}

	if value, ok := params.Args["supplier_id"].(int); ok && value != 0 {
		input.SupplierID = value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = value
	}

	if value, ok := params.Args["type"].(string); ok && value != "" {
		input.Type = &value
	}

	items, total, err := r.Repo.GetAllObligations(input)
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

func (r *Resolver) PayOrderResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)
	SAPID := params.Args["sap_id"].(string)
	DateOfSAP := params.Args["date_of_sap"].(string)

	dateOfSAP, err := parseDate(DateOfSAP)

	if err != nil {
		fmt.Printf("Paying the order failed because this error - %s.\n", err)
		return dto.ResponseSingle{
			Status: "failed",
		}, nil
	}

	paymentOrder := structs.PaymentOrder{
		ID:        itemID,
		SAPID:     &SAPID,
		DateOfSAP: &dateOfSAP,
	}

	err = r.Repo.PayPaymentOrder(paymentOrder)
	if err != nil {
		fmt.Printf("Paying the order failed because this error - %s.\n", err)
		return dto.ResponseSingle{
			Status: "failed",
		}, nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You paid this item!",
	}, nil
}

func buildPaymentOrder(item structs.PaymentOrder, r *Resolver) (*dto.PaymentOrderResponse, error) {
	response := dto.PaymentOrderResponse{
		ID:              item.ID,
		BankAccount:     item.BankAccount,
		DateOfPayment:   item.DateOfPayment,
		IDOfStatement:   item.IDOfStatement,
		SAPID:           item.SAPID,
		DateOfSAP:       item.DateOfSAP,
		DateOfOrder:     item.DateOfOrder,
		Amount:          item.Amount,
		Status:          item.Status,
		SourceOfFunding: item.SourceOfFunding,
		Description:     item.Description,
		CreatedAt:       item.CreatedAt,
		UpdatedAt:       item.UpdatedAt,
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
		value, err := r.Repo.GetSupplier(item.SupplierID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.Supplier = dropdown
	}

	if item.FileID != nil && *item.FileID != 0 {
		file, err := r.Repo.GetFileByID(*item.FileID)

		if err != nil {
			return nil, err
		}
		fileDropdown := dto.FileDropdownSimple{
			ID:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		}

		response.File = fileDropdown
	}

	for _, orderItem := range item.Items {
		builtItem, err := buildPaymentOrderItem(orderItem, r)

		if err != nil {
			return nil, err
		}

		response.Items = append(response.Items, *builtItem)
	}

	return &response, nil
}

func buildPaymentOrderItem(item structs.PaymentOrderItems, r *Resolver) (*dto.PaymentOrderItemResponse, error) {
	response := dto.PaymentOrderItemResponse{
		ID:                        item.ID,
		PaymentOrderID:            item.PaymentOrderID,
		InvoiceID:                 item.InvoiceID,
		AdditionalExpenseID:       item.AdditionalExpenseID,
		SalaryAdditionalExpenseID: item.SalaryAdditionalExpenseID,
		Type:                      item.Type,
		Title:                     item.Title,
		Amount:                    item.Amount,
		CreatedAt:                 item.CreatedAt,
		UpdatedAt:                 item.UpdatedAt,
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
