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

func (r *Resolver) DepositPaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		DepositPayment, err := r.Repo.GetDepositPaymentByID(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		res, err := buildDepositPayment(*DepositPayment, r)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var resItems []dto.DepositPaymentResponse
	for _, item := range items {
		resItem, err := buildDepositPayment(item, r)

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

func (r *Resolver) GetInitialStateOverviewResolver(params graphql.ResolveParams) (interface{}, error) {

	input := dto.DepositInitialStateFilter{}
	if value, ok := params.Args["date"].(string); ok && value != "" {
		date, err := parseDate(value)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		input.Date = date
	}

	if value, ok := params.Args["bank_account"].(string); ok && value != "" {
		input.BankAccount = &value
	}

	if value, ok := params.Args["transitional_bank_account"].(bool); ok && value {
		input.TransitionalBankAccount = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	} else {
		input.OrganizationUnitID, _ = params.Context.Value(config.OrganizationUnitIDKey).(*int)
	}

	var resItems []structs.DepositPayment
	var items []structs.DepositPayment
	var err error
	if input.TransitionalBankAccount != nil && input.OrganizationUnitID != nil {
		items, err = r.Repo.GetInitialState(input)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		resItems = append(resItems, items...)
	} else if input.BankAccount != nil {
		items, err = r.Repo.GetInitialState(input)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		resItems = append(resItems, items...)
	} else if input.OrganizationUnitID != nil && input.BankAccount == nil {
		orgUnit, err := r.Repo.GetOrganizationUnitByID(*input.OrganizationUnitID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		for _, bankAccount := range orgUnit.BankAccounts {
			input.BankAccount = &bankAccount
			items, err = r.Repo.GetInitialState(input)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			resItems = append(resItems, items...)
		}

	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   resItems,
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

	var item *structs.DepositPayment

	if data.ID == 0 {
		item, err = r.Repo.CreateDepositPayment(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdateDepositPayment(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	singleItem, err := buildDepositPayment(*item, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) DepositPaymentCaseNumberResolver(params graphql.ResolveParams) (interface{}, error) {
	caseNumber := params.Args["case_number"].(string)
	bankAccount := params.Args["bank_account"].(string)

	res, err := r.Repo.GetDepositPaymentCaseNumber(caseNumber, bankAccount)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    res,
	}, nil
}

func (r *Resolver) DepositCaseNumberResolver(params graphql.ResolveParams) (interface{}, error) {
	organizationUnitID := params.Args["organization_unit_id"].(int)
	bankAccount := params.Args["bank_account"].(string)

	res, err := r.Repo.GetCaseNumber(organizationUnitID, bankAccount)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   res,
	}, nil
}

func (r *Resolver) DepositPaymentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteDepositPayment(params.Context, itemID)
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
		CurrentBankAccount:        item.CurrentBankAccount,
		Status:                    item.Status,
		CreatedAt:                 item.CreatedAt,
		UpdatedAt:                 item.UpdatedAt,
	}

	if item.OrganizationUnitID != 0 {
		value, _ := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}*/

		if value != nil {

			dropdown := dto.DropdownSimple{
				ID:    value.ID,
				Title: value.Title,
			}

			response.OrganizationUnit = dropdown
		}
	}

	if item.AccountID != 0 {
		value, _ := r.Repo.GetAccountItemByID(item.AccountID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get account item by id")
		}*/

		if value != nil {

			dropdown := dto.DropdownSimple{
				ID:    value.ID,
				Title: value.Title,
			}

			response.Account = dropdown
		}
	}

	if item.FileID != 0 {
		file, _ := r.Repo.GetFileByID(item.FileID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}*/

		if file != nil {
			fileDropdown := dto.FileDropdownSimple{
				ID:   file.ID,
				Name: file.Name,
				Type: *file.Type,
			}

			response.File = fileDropdown
		}
	}

	return &response, nil
}
