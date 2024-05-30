package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"
	"math"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) GetObligationsForAccountingResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.ObligationsFilter{}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = value
	}

	if value, ok := params.Args["type"].(string); ok && value != "" {
		input.Type = &value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["date_of_start"].(string); ok && value != "" {
		dateOfStart, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfStart = &dateOfStart
	}

	if value, ok := params.Args["date_of_end"].(string); ok && value != "" {
		dateOfEnd, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfEnd = &dateOfEnd
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

		//id ne znaci nista, znaci samo kolegama sa fe, jer imaju felericnu tabelu
		items[i].ID = i + 1

		items[i].Price = math.Round(items[i].Price*100) / 100
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

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["date_of_start"].(string); ok && value != "" {
		dateOfStart, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfStart = &dateOfStart
	}

	if value, ok := params.Args["date_of_end"].(string); ok && value != "" {
		dateOfEnd, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfEnd = &dateOfEnd
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

			//zbog fronta dodato
			items[i].ID = i + 1

			items[i].Price = math.Round(items[i].Price*100) / 100
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

func (r *Resolver) GetEnforcedPaymentsForAccountingResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.ObligationsFilter{}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["date_of_start"].(string); ok && value != "" {
		dateOfStart, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfStart = &dateOfStart
	}

	if value, ok := params.Args["date_of_end"].(string); ok && value != "" {
		dateOfEnd, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfEnd = &dateOfEnd
	}

	items, total, err := r.Repo.GetAllEnforcedPaymentsForAccounting(input)
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

			//zbog fronta dodato
			items[i].ID = i + 1

			items[i].Price = math.Round(items[i].Price*100) / 100
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

func (r *Resolver) GetReturnedEnforcedPaymentsForAccountingResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.ObligationsFilter{}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["date_of_start"].(string); ok && value != "" {
		dateOfStart, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfStart = &dateOfStart
	}

	if value, ok := params.Args["date_of_end"].(string); ok && value != "" {
		dateOfEnd, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfEnd = &dateOfEnd
	}

	items, total, err := r.Repo.GetAllReturnedEnforcedPaymentsForAccounting(input)
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

			//zbog fronta dodato
			items[i].ID = i + 1

			items[i].Price = math.Round(items[i].Price*100) / 100
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

	for id, item := range items.Items {
		builtItem, err := buildAccountingOrderItemForObligations(item, r)

		//za front i felericne tabele
		builtItem.ID = id + 1

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

	if value, ok := params.Args["type"].(string); ok && value != "" {
		input.Type = &value
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

func (r *Resolver) AnalyticalCardOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.AnalyticalCardFilter{}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = value
	}

	if value, ok := params.Args["supplier_id"].(int); ok && value != 0 {
		input.SupplierID = &value
	}

	if value, ok := params.Args["date_of_start"].(string); ok && value != "" {
		dateOfStart, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfStart = &dateOfStart
	}

	if value, ok := params.Args["date_of_end"].(string); ok && value != "" {
		dateOfEnd, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfEnd = &dateOfEnd
	}

	if value, ok := params.Args["date_of_start_booking"].(string); ok && value != "" {
		dateOfStart, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfStartBooking = &dateOfStart
	}

	if value, ok := params.Args["date_of_end_booking"].(string); ok && value != "" {
		dateOfEnd, err := parseDate(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		input.DateOfEndBooking = &dateOfEnd
	}

	if value, ok := params.Args["account_id"].(int); ok && value != 0 {
		account, err := r.Repo.GetAccountItemByID(value)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{
			SerialNumber: &account.SerialNumber,
		})

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		for _, account := range accounts.Data {
			input.AccountID = append(input.AccountID, account.ID)
		}
	}

	items, err := r.Repo.GetAnalyticalCard(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for i := 0; i < len(items); i++ {
		if input.DateOfStart != nil && input.DateOfEnd != nil {
			items[i].DateOfStart = input.DateOfStart
			items[i].DateOfEnd = input.DateOfEnd
		} else {
			items[i].DateOfStart = input.DateOfStartBooking
			items[i].DateOfEnd = input.DateOfEndBooking
		}

		items[i].OrganizationUnitID = input.OrganizationUnitID
	}

	response, err := buildAnalyticalCardResponse(items, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	orgUnitID, err := r.Repo.GetOrganizationUnitByID(input.OrganizationUnitID)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for i := 0; i < len(response); i++ {
		response[i].OrganizationUnit.ID = orgUnitID.ID
		response[i].OrganizationUnit.Title = orgUnitID.Title
	}

	return dto.Response{
		Status: "success",
		Items:  response,
	}, nil
}

func buildAccountingOrderItemForObligations(item dto.AccountingOrderItemsForObligations, r *Resolver) (*dto.AccountingOrderItemsForObligationsResponse, error) {

	response := dto.AccountingOrderItemsForObligationsResponse{
		CreditAmount:          item.CreditAmount,
		DebitAmount:           item.DebitAmount,
		Type:                  item.Type,
		Salary:                item.Salary,
		Invoice:               item.Invoice,
		PaymentOrder:          item.PaymentOrder,
		EnforcedPayment:       item.EnforcedPayment,
		ReturnEnforcedPayment: item.ReturnEnforcedPayment,
		Date:                  item.Date,
		Title:                 item.Title,
		SupplierID:            item.SupplierID,
	}

	if item.AccountID != 0 {
		value, err := r.Repo.GetAccountItemByID(item.AccountID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title + " - " + value.SerialNumber,
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

		if response.Title == structs.SupplierTitle {
			response.Title = dropdown.Title
		}
	}

	return &response, nil
}

func buildTypeForAccountingOrder(itemType string) string {
	switch itemType {
	case config.TypeInvoice:
		return config.TitleInvoice
	case config.TypeDecision:
		return config.TitleDecision
	case config.TypeContract:
		return config.TitleContract
	case config.TypeSalary:
		return config.TitleSalary
	case config.TypeObligations:
		return config.TitleObligations
	case config.TypePaymentOrder:
		return config.TitlePaymentOrder
	case config.TypeEnforcedPayment:
		return config.TitleEnforcedPayment
	case config.TypeReturnEnforcedPayment:
		return config.TitleReturnEnforcedPayment

	}
	return ""
}

func buildAccountingEntry(item structs.AccountingEntry, r *Resolver) (*dto.AccountingEntryResponse, error) {

	responseType := buildTypeForAccountingOrder(item.Type)

	year := item.DateOfBooking.Year()
	yearLastTwoDigits := year % 100

	// Format the string
	formatedIDOfEntry := fmt.Sprintf("%02d-%03d", yearLastTwoDigits, item.IDOfEntry)

	response := dto.AccountingEntryResponse{
		ID:                item.ID,
		Title:             item.Title,
		Type:              responseType,
		IDOfEntry:         item.IDOfEntry,
		FormatedIDOfEntry: formatedIDOfEntry,
		DateOfBooking:     item.DateOfBooking,
		CreditAmount:      item.CreditAmount,
		DebitAmount:       item.DebitAmount,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	if item.OrganizationUnitID != 0 {
		value, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.OrganizationUnitsOverviewResponse{
			ID:      value.ID,
			Title:   value.Title,
			Address: value.Address,
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
		EntryNumber:  item.EntryNumber,
		EntryDate:    item.EntryDate,
		CreditAmount: item.CreditAmount,
		DebitAmount:  item.DebitAmount,
		Type:         item.Type,
		Date:         item.Date,
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
			Title: value.SerialNumber + " - " + value.Title,
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

	if item.PaymentOrderID != nil && *item.PaymentOrderID != 0 {
		value, err := r.Repo.GetPaymentOrderByID(*item.PaymentOrderID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: *value.SAPID,
		}

		response.PaymentOrder = dropdown
	}

	if item.EnforcedPaymentID != nil && *item.EnforcedPaymentID != 0 {
		value, err := r.Repo.GetEnforcedPaymentByID(*item.EnforcedPaymentID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: *value.SAPID,
		}

		response.EnforcedPayment = dropdown
	}

	if item.ReturnEnforcedPaymentID != nil && *item.ReturnEnforcedPaymentID != 0 {
		value, err := r.Repo.GetEnforcedPaymentByID(*item.ReturnEnforcedPaymentID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: *value.SAPID,
		}

		response.ReturnEnforcedPayment = dropdown
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

	return &response, nil
}

func buildAnalyticalCardResponse(items []structs.AnalyticalCard, r *Resolver) ([]dto.AnalyticalCardDTO, error) {
	var response []dto.AnalyticalCardDTO

	for _, item := range items {
		if item.SupplierID != 0 {
			supplier, err := r.Repo.GetSupplier(item.SupplierID)

			if err != nil {
				return nil, err
			}

			responseItem := dto.AnalyticalCardDTO{
				InitialState:            item.InitialState,
				SumCreditAmount:         item.SumCreditAmount,
				SumDebitAmount:          item.SumDebitAmount,
				SumCreditAmountInPeriod: item.SumCreditAmountInPeriod,
				SumDebitAmountInPeriod:  item.SumDebitAmountInPeriod,
				DateOfStart:             item.DateOfStart,
				DateOfEnd:               item.DateOfEnd,
				Supplier: dto.DropdownSimple{
					ID:    supplier.ID,
					Title: supplier.Title,
				}}

			for _, entryItem := range item.Items {

				dateOfBooking, err := parseDate(entryItem.DateOfBooking)

				if err != nil {
					return nil, err
				}

				year := dateOfBooking.Year()
				yearLastTwoDigits := year % 100

				// Format the string
				formatedIDOfEntry := fmt.Sprintf("%02d-%03d", yearLastTwoDigits, entryItem.IDOfEntry)

				itemType := buildTypeForAccountingOrder(entryItem.Type)

				if entryItem.Date == config.DefaultDateString {
					entryItem.Date = ""
				}

				responseItem.Items = append(responseItem.Items, dto.AnalyticalCardItemsDTO{
					ID:                entryItem.ID,
					Title:             entryItem.Title,
					Type:              itemType,
					CreditAmount:      entryItem.CreditAmount,
					DebitAmount:       entryItem.DebitAmount,
					Balance:           entryItem.Balance,
					DateOfBooking:     entryItem.DateOfBooking,
					Date:              entryItem.Date,
					DocumentNumber:    entryItem.DocumentNumber,
					FormatedIDOfEntry: formatedIDOfEntry,
				})
			}

			response = append(response, responseItem)
		}
	}

	return response, nil
}
