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

func (r *Resolver) EnforcedPaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		EnforcedPayment, err := r.Repo.GetEnforcedPaymentByID(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		res, err := buildEnforcedPayment(*EnforcedPayment, r)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.EnforcedPaymentResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.EnforcedPaymentFilter{}
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

	if value, ok := params.Args["registred"].(bool); ok {
		input.Registred = &value
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

	items, total, err := r.Repo.GetEnforcedPaymentList(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var resItems []dto.EnforcedPaymentResponse
	for _, item := range items {
		resItem, err := buildEnforcedPayment(item, r)

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

func (r *Resolver) EnforcedPaymentInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.EnforcedPayment
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

	var item *structs.EnforcedPayment

	if data.ID == 0 {
		item, err = r.Repo.CreateEnforcedPayment(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdateEnforcedPayment(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

	}

	singleItem, err := buildEnforcedPayment(*item, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) ReturnEnforcedPaymentResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)
	returnFileID := params.Args["return_file_id"].(int)
	returnDateString := params.Args["return_date"].(string)

	returnDate, err := parseDate(returnDateString)

	if err != nil {
		fmt.Printf("Returning the enforced payment failed because this error - %s.\n", err)
		return dto.ResponseSingle{
			Status: "failed",
		}, nil
	}

	EnforcedPayment := structs.EnforcedPayment{
		ID:           itemID,
		ReturnFileID: &returnFileID,
		ReturnDate:   &returnDate,
	}

	err = r.Repo.ReturnEnforcedPayment(EnforcedPayment)
	if err != nil {
		fmt.Printf("Returning the enforced payment failed because this error - %s.\n", err)
		return dto.ResponseSingle{
			Status: "failed",
		}, nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You returned this item!",
	}, nil
}

func buildEnforcedPayment(item structs.EnforcedPayment, r *Resolver) (*dto.EnforcedPaymentResponse, error) {
	response := dto.EnforcedPaymentResponse{
		ID:            item.ID,
		BankAccount:   item.BankAccount,
		DateOfPayment: item.DateOfPayment,
		IDOfStatement: item.IDOfStatement,
		ReturnDate:    item.ReturnDate,
		SAPID:         item.SAPID,
		DateOfSAP:     item.DateOfSAP,
		DateOfOrder:   item.DateOfOrder,
		Amount:        item.Amount,
		Status:        item.Status,
		Description:   item.Description,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
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

	if item.ReturnFileID != nil && *item.ReturnFileID != 0 {
		file, err := r.Repo.GetFileByID(*item.ReturnFileID)

		if err != nil {
			return nil, err
		}
		fileDropdown := dto.FileDropdownSimple{
			ID:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		}

		response.ReturnFile = fileDropdown
	}

	for _, orderItem := range item.Items {
		builtItem, err := buildEnforcedPaymentItem(orderItem, r)

		if err != nil {
			return nil, err
		}

		response.Items = append(response.Items, *builtItem)
	}

	return &response, nil
}

func buildEnforcedPaymentItem(item structs.EnforcedPaymentItems, r *Resolver) (*dto.EnforcedPaymentItemResponse, error) {
	response := dto.EnforcedPaymentItemResponse{
		ID:             item.ID,
		PaymentOrderID: item.PaymentOrderID,
		InvoiceID:      item.InvoiceID,
		Title:          item.Title,
		Amount:         item.Amount,
		CreatedAt:      item.CreatedAt,
		UpdatedAt:      item.UpdatedAt,
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
