package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) FineInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Fine
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

	var item *structs.Fine

	if data.OrganizationUnitID == 0 {
		organizationUnitID, _ := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if organizationUnitID != nil {
			data.OrganizationUnitID = *organizationUnitID
		}
	}

	if data.ID == 0 {
		item, err = r.Repo.CreateFine(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdateFine(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	singleItem, err := buildFineResponseItem(*item, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) FineOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		fine, err := r.Repo.GetFine(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		fineResItem, err := buildFineResponseItem(*fine, r)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.FineResponseItem{fineResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetFineListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["subject"].(string); ok && value != "" {
		input.Subject = &value
	}

	if value, ok := params.Args["act_type_id"].(int); ok && value != 0 {
		input.FilterByActTypeID = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	}

	fines, total, err := r.Repo.GetFineList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	fineResItem, err := buildFineResponseItemList(fines, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   fineResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) FineDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFine(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildFineResponseItem(fine structs.Fine, r *Resolver) (*dto.FineResponseItem, error) {
	status := dto.DropdownSimple{
		ID:    int(structs.UnpaidFineStatus),
		Title: string(dto.FinancialFineUnpaidFineStatus),
	}

	switch fine.Status {
	case structs.PaidFineStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PaidFineStatus),
			Title: string(dto.FinancialFinePaidFineStatus),
		}
	case structs.UnpaidFineStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.UnpaidFineStatus),
			Title: string(dto.FinancialFineUnpaidFineStatus),
		}
	case structs.PartFineStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PartFineStatus),
			Title: string(dto.FinancialFinePartFineStatus),
		}
	}

	actType := dto.DropdownSimple{
		ID:    int(structs.DissisionActType),
		Title: string(dto.FinancialFineDissisionActType),
	}
	switch fine.ActType {
	case structs.VerdictActType:
		actType = dto.DropdownSimple{
			ID:    int(structs.VerdictActType),
			Title: string(dto.FinancialFineVerdictActType),
		}
	case structs.DissisionActType:
		actType = dto.DropdownSimple{
			ID:    int(structs.DissisionActType),
			Title: string(dto.FinancialFineDissisionActType),
		}
	}

	response := dto.FineResponseItem{
		ID:                     fine.ID,
		ActType:                actType,
		DecisionNumber:         fine.DecisionNumber,
		DecisionDate:           fine.DecisionDate,
		Subject:                fine.Subject,
		JMBG:                   fine.JMBG,
		Residence:              fine.Residence,
		Amount:                 fine.Amount,
		PaymentReferenceNumber: fine.PaymentReferenceNumber,
		DebitReferenceNumber:   fine.DebitReferenceNumber,
		ExecutionDate:          fine.ExecutionDate,
		PaymentDeadlineDate:    fine.PaymentDeadlineDate,
		Description:            fine.Description,
		Status:                 status,
		CourtCosts:             fine.CourtCosts,
		FineFeeDetailsDTO:      fine.FineFeeDetailsDTO,
		CreatedAt:              fine.CreatedAt,
		UpdatedAt:              fine.UpdatedAt,
	}

	if len(fine.File) > 0 {
		for _, fileID := range fine.File {
			file, _ := r.Repo.GetFileByID(fileID)

			/*if err != nil {
				return nil, errors.Wrap(err, "repo get file by id")
			}*/

			if file != nil {

				FileDropdown := dto.FileDropdownSimple{
					ID:   file.ID,
					Name: file.Name,
					Type: *file.Type,
				}
				response.File = append(response.File, FileDropdown)
			}
		}
	}

	if fine.AccountID != 0 {
		account, _ := r.Repo.GetAccountItemByID(fine.AccountID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get account item by id")
		}*/

		if account != nil {

			accountDropdown := dto.DropdownSimple{
				ID:    account.ID,
				Title: account.Title,
			}
			response.Account = accountDropdown
		}
	}

	if fine.CourtAccountID != nil {
		courtAccount, _ := r.Repo.GetAccountItemByID(*fine.CourtAccountID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get account item by id")
		}*/

		if courtAccount != nil {

			courtAccountDropdown := &dto.DropdownSimple{
				ID:    courtAccount.ID,
				Title: courtAccount.Title,
			}
			response.CourtAccount = courtAccountDropdown
		}
	}

	if fine.OrganizationUnitID != 0 {
		organizationUnit, _ := r.Repo.GetOrganizationUnitByID(fine.OrganizationUnitID)
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}*/

		if organizationUnit != nil {

			orgUnitDropdown := dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			}

			response.OrganizationUnit = orgUnitDropdown
		}

	}
	return &response, nil
}

func buildFineResponseItemList(itemList []structs.Fine, r *Resolver) ([]*dto.FineResponseItem, error) {
	var items []*dto.FineResponseItem

	for _, item := range itemList {
		singleItem, err := buildFineResponseItem(item, r)

		if err != nil {
			return nil, errors.Wrap(err, "build fine response item")
		}

		items = append(items, singleItem)

	}

	return items, nil
}
