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

func (r *Resolver) GetObligationsForAccountingResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.ObligationsFilter{}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = value
	}

	items, total, err := r.Repo.GetAllObligationsForAccounting(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for i := 0; i < len(items); i++ {
		if items[i].SupplierID != nil && *items[i].SupplierID != 0 {
			supplier, err := r.Repo.GetSupplier(*items[i].SupplierID)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			items[i].Supplier.ID = supplier.ID
			items[i].Supplier.Title = supplier.Title
		}
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

func (r *Resolver) GetPaymentOrdersForAccountingResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.ObligationsFilter{}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = value
	}

	items, total, err := r.Repo.GetAllPaymentOrdersForAccounting(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for i := 0; i < len(items); i++ {
		if items[i].SupplierID != nil && *items[i].SupplierID != 0 {
			supplier, err := r.Repo.GetSupplier(*items[i].SupplierID)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			items[i].Supplier.ID = supplier.ID
			items[i].Supplier.Title = supplier.Title
		}
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

	orgUnit, err := r.Repo.GetOrganizationUnitByID(data.OrganizationUnitID)

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

func (r *Resolver) AccountingEntryOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		AccountingEntry, err := r.Repo.GetAccountingEntryByID(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		res, err := buildAccountingEntry(*AccountingEntry, r)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.AccountingEntryResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.AccountingEntryFilter{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	} else {
		input.OrganizationUnitID, _ = params.Context.Value(config.OrganizationUnitIDKey).(*int)
	}

	items, total, err := r.Repo.GetAccountingEntryList(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var resItems []dto.AccountingEntryResponse
	for _, item := range items {
		resItem, err := buildAccountingEntry(item, r)

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

func (r *Resolver) AccountingEntryInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.AccountingEntry
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

	var item *structs.AccountingEntry

	item, err = r.Repo.CreateAccountingEntry(&data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	singleItem, err := buildAccountingEntry(*item, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) AccountingEntryDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteAccountingEntry(itemID)
	if err != nil {
		fmt.Printf("Deleting accounting entry failed because of this error - %s.\n", err)
		return dto.ResponseSingle{
			Status: "failed",
		}, nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildAccountingOrderItemForObligations(item dto.AccountingOrderItemsForObligations, r *Resolver) (*dto.AccountingOrderItemsForObligationsResponse, error) {
	response := dto.AccountingOrderItemsForObligationsResponse{
		CreditAmount: item.CreditAmount,
		DebitAmount:  item.DebitAmount,
		Type:         item.Type,
		Salary:       item.Salary,
		Invoice:      item.Invoice,
		PaymentOrder: item.PaymentOrder,
		Title:        item.Title,
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

	if item.SupplierID != 0 {
		value, err := r.Repo.GetSupplier(item.SupplierID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.Title = dropdown.Title
	}

	return &response, nil
}

func buildAccountingEntry(item structs.AccountingEntry, r *Resolver) (*dto.AccountingEntryResponse, error) {
	response := dto.AccountingEntryResponse{
		ID:            item.ID,
		Title:         item.Title,
		IDOfEntry:     item.IDOfEntry,
		DateOfBooking: item.DateOfBooking,
		CreditAmount:  item.CreditAmount,
		DebitAmount:   item.DebitAmount,
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

	for _, orderItem := range item.Items {
		builtItem, err := buildAccountingEntryItem(orderItem, r)

		if err != nil {
			return nil, err
		}

		response.Items = append(response.Items, *builtItem)
	}

	return &response, nil
}

func buildAccountingEntryItem(item structs.AccountingEntryItems, r *Resolver) (*dto.AccountingEntryItemResponse, error) {
	response := dto.AccountingEntryItemResponse{
		ID:           item.ID,
		Title:        item.Title,
		EntryID:      item.EntryID,
		CreditAmount: item.CreditAmount,
		DebitAmount:  item.DebitAmount,
		Type:         item.Type,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
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

	if item.SalaryID != nil && *item.SalaryID != 0 {
		value, err := r.Repo.GetSalaryByID(*item.SalaryID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Month,
		}

		response.Salary = dropdown
	}

	if item.InvoiceID != nil && *item.InvoiceID != 0 {
		value, err := r.Repo.GetInvoice(*item.InvoiceID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.InvoiceNumber,
		}

		response.Invoice = dropdown
	}

	return &response, nil
}
