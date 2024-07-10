package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) DepositPaymentOrderOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		DepositPaymentOrder, err := r.Repo.GetDepositPaymentOrderByID(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		res, err := buildDepositPaymentOrder(*DepositPaymentOrder, r)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var resItems []dto.DepositPaymentOrderResponse
	for _, item := range items {
		resItem, err := buildDepositPaymentOrder(item, r)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if data.OrganizationUnitID == 0 {

		organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return errors.HandleAPPError(fmt.Errorf("user does not have organization unit assigned"))
		}

		data.OrganizationUnitID = *organizationUnitID

	}

	var item *structs.DepositPaymentOrder

	if data.ID == 0 {
		item, err = r.Repo.CreateDepositPaymentOrder(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdateDepositPaymentOrder(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	singleItem, err := buildDepositPaymentOrder(*item, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) DepositPaymentOrderDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteDepositPaymentOrder(params.Context, itemID)
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

func (r *Resolver) PayDepositOrderResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)
	IDOfStatement := params.Args["id_of_statement"].(string)
	DateOfStatement := params.Args["date_of_statement"].(string)

	dateOfStatement, err := parseDate(DateOfStatement)

	if err != nil {
		fmt.Printf("Paying the order failed because this error - %s.\n", err)
		return dto.ResponseSingle{
			Status: "failed",
		}, nil
	}

	paymentOrder := structs.DepositPaymentOrder{
		ID:              itemID,
		IDOfStatement:   &IDOfStatement,
		DateOfStatement: &dateOfStatement,
	}

	err = r.Repo.PayDepositPaymentOrder(params.Context, paymentOrder)
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

func buildDepositPaymentOrder(item structs.DepositPaymentOrder, r *Resolver) (*dto.DepositPaymentOrderResponse, error) {
	response := dto.DepositPaymentOrderResponse{
		ID:                item.ID,
		CaseNumber:        item.CaseNumber,
		NetAmount:         item.NetAmount,
		BankAccount:       item.BankAccount,
		DateOfPayment:     item.DateOfPayment,
		DateOfStatement:   item.DateOfStatement,
		IDOfStatement:     item.IDOfStatement,
		Status:            item.Status,
		SourceBankAccount: item.SourceBankAccount,
	}

	if item.MunicipalityID != nil && *item.MunicipalityID != 0 {
		supplier, err := r.Repo.GetSupplier(*item.MunicipalityID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get supplier")
		}

		supplierDropdown := dto.DropdownSimple{
			ID:    supplier.ID,
			Title: supplier.Title,
		}

		response.Municipality = supplierDropdown
	}

	if item.TaxAuthorityCodebookID != nil && *item.TaxAuthorityCodebookID != 0 {
		TaxAuthorityCodebook, err := r.Repo.GetTaxAuthorityCodebookByID(*item.TaxAuthorityCodebookID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get tax authority codebook by id")
		}

		OUDropdown := dto.DropdownSimple{
			ID:    TaxAuthorityCodebook.ID,
			Title: TaxAuthorityCodebook.Title,
		}

		response.TaxAuthorityCodebook = OUDropdown
	}

	if item.OrganizationUnitID != 0 {
		value, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
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
			return nil, errors.Wrap(err, "repo get supplier")
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.Supplier = dropdown
	}

	if item.SubjectTypeID != 0 {
		value, err := r.Repo.GetSupplier(item.SubjectTypeID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get supplier")
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.SubjectType = dropdown
	}

	for _, additionalExpense := range item.AdditionalExpenses {
		builtItem, err := buildDepositPaymentAdditionalExpense(r, additionalExpense)

		if err != nil {
			return nil, errors.Wrap(err, "build deposit payment additional expense")
		}

		response.AdditionalExpenses = append(response.AdditionalExpenses, *builtItem)
	}

	for _, additionalExpense := range item.AdditionalExpensesForPaying {
		builtItem, err := buildDepositPaymentAdditionalExpense(r, additionalExpense)

		if err != nil {
			return nil, errors.Wrap(err, "build deposit payment additional expense")
		}

		response.AdditionalExpensesForPaying = append(response.AdditionalExpensesForPaying, *builtItem)
	}

	if item.FileID != 0 {
		file, err := r.Repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}
		fileDropdown := dto.FileDropdownSimple{
			ID:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		}

		response.File = fileDropdown
	}

	return &response, nil
}

func (r *Resolver) DepositPaymentAdditionalExpensesOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.DepositPaymentAdditionalExpensesListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["payment_order_id"].(int); ok && value != 0 {
		input.PaymentOrderID = &value
	}

	if value, ok := params.Args["subject_id"].(int); ok && value != 0 {
		input.SubjectID = &value
	}

	if value, ok := params.Args["status"].(string); ok && value != "" {
		input.Status = &value
	}

	if value, ok := params.Args["source_bank_account"].(string); ok && value != "" {
		input.SourceBankAccount = &value
	}

	if value, ok := params.Args["year"].(int); ok && value != 0 {
		input.Year = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	}

	additionalExpenses, total, err := r.Repo.GetDepositPaymentAdditionalExpenses(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	builtAdditionalExpenses, err := buildDepositPaymentAdditionalExpenseItemList(r, additionalExpenses)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   builtAdditionalExpenses,
		Total:   total,
	}, nil
}

func buildDepositPaymentAdditionalExpenseItemList(r *Resolver, itemList []structs.DepositPaymentAdditionalExpenses) ([]*dto.DepositPaymentAdditionalExpensesResponse, error) {
	var items []*dto.DepositPaymentAdditionalExpensesResponse

	for _, item := range itemList {
		singleItem, err := buildDepositPaymentAdditionalExpense(r, item)

		if err != nil {
			return nil, errors.Wrap(err, "build deposit payment additional expense")
		}

		items = append(items, singleItem)

	}

	return items, nil
}

func buildDepositPaymentAdditionalExpense(r *Resolver, item structs.DepositPaymentAdditionalExpenses) (*dto.DepositPaymentAdditionalExpensesResponse, error) {
	response := dto.DepositPaymentAdditionalExpensesResponse{
		ID:                item.ID,
		Title:             item.Title,
		CaseNumber:        item.CaseNumber,
		Price:             item.Price,
		BankAccount:       item.BankAccount,
		Status:            item.Status,
		SourceBankAccount: item.SourceBankAccount,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	if item.AccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(item.AccountID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get account item by id")
		}

		response.Account = dto.DropdownSimple{
			ID:    account.ID,
			Title: account.Title,
		}
	}

	if item.PaymentOrderID != 0 {
		invoice, err := r.Repo.GetDepositPaymentOrderByID(item.PaymentOrderID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get deposit payment order by id")
		}

		response.PaymentOrder = dto.DropdownSimple{
			ID:    invoice.ID,
			Title: invoice.CaseNumber,
		}
	}

	if item.SubjectID != 0 {
		supplier, err := r.Repo.GetSupplier(item.SubjectID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get supplier")
		}

		response.Subject = dto.DropdownSimple{
			ID:    supplier.ID,
			Title: supplier.Title,
		}
	}

	if item.OrganizationUnitID != 0 {
		orgUnit, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}

		response.OrganizationUnit = dto.DropdownSimple{
			ID:    orgUnit.ID,
			Title: orgUnit.Title,
		}
	}
	return &response, nil
}
