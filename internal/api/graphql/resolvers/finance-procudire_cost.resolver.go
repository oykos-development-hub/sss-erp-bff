package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ProcedureCostInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.ProcedureCost
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

	var item *structs.ProcedureCost

	if data.ID == 0 {
		item, err = r.Repo.CreateProcedureCost(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdateProcedureCost(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}
	procedurecostResItem, err := buildProcedureCostResponseItem(*item, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	response.Item = procedurecostResItem

	return response, nil
}

func (r *Resolver) ProcedureCostOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		procedurecost, err := r.Repo.GetProcedureCost(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		procedurecostResItem, err := buildProcedureCostResponseItem(*procedurecost, r)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.ProcedureCostResponseItem{procedurecostResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetProcedureCostListInputMS{}
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

	if value, ok := params.Args["procedure_cost_type_id"].(int); ok && value != 0 {
		input.FilterByProcedureCostTypeID = &value
	}

	procedurecosts, total, err := r.Repo.GetProcedureCostList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	procedurecostResItem, err := buildProcedureCostResponseItemList(procedurecosts, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   procedurecostResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) ProcedureCostDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteProcedureCost(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildProcedureCostResponseItem(procedurecost structs.ProcedureCost, r *Resolver) (*dto.ProcedureCostResponseItem, error) {
	status := dto.DropdownSimple{
		ID:    int(structs.UnpaidProcedureCostStatus),
		Title: string(dto.FinancialProcedureCostUnpaidStatus),
	}

	switch procedurecost.Status {
	case structs.PaidProcedureCostStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PaidProcedureCostStatus),
			Title: string(dto.FinancialProcedureCostPaidStatus),
		}
	case structs.UnpaidProcedureCostStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.UnpaidProcedureCostStatus),
			Title: string(dto.FinancialProcedureCostUnpaidStatus),
		}
	case structs.PartProcedureCostStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PartProcedureCostStatus),
			Title: string(dto.FinancialProcedureCostPartStatus),
		}
	}

	actType := dto.DropdownSimple{
		ID:    int(structs.DissisionProcedureCostType),
		Title: string(dto.FinancialProcedureCostDissisionType),
	}

	switch procedurecost.ProcedureCostType {
	case structs.DissisionProcedureCostType:
		actType = dto.DropdownSimple{
			ID:    int(structs.DissisionProcedureCostType),
			Title: string(dto.FinancialProcedureCostDissisionType),
		}
	case structs.VerdictProcedureCostType:
		actType = dto.DropdownSimple{
			ID:    int(structs.VerdictProcedureCostType),
			Title: string(dto.FinancialProcedureCostVerdictType),
		}

	}

	response := dto.ProcedureCostResponseItem{
		ID:                      procedurecost.ID,
		ActType:                 actType,
		DecisionNumber:          procedurecost.DecisionNumber,
		DecisionDate:            procedurecost.DecisionDate,
		Subject:                 procedurecost.Subject,
		JMBG:                    procedurecost.JMBG,
		Residence:               procedurecost.Residence,
		Amount:                  procedurecost.Amount,
		PaymentReferenceNumber:  procedurecost.PaymentReferenceNumber,
		DebitReferenceNumber:    procedurecost.DebitReferenceNumber,
		ExecutionDate:           procedurecost.ExecutionDate,
		PaymentDeadlineDate:     procedurecost.PaymentDeadlineDate,
		Description:             procedurecost.Description,
		Status:                  status,
		CourtCosts:              procedurecost.CourtCosts,
		ProcedureCostDetailsDTO: procedurecost.ProcedureCostDetailsDTO,
		CreatedAt:               procedurecost.CreatedAt,
		UpdatedAt:               procedurecost.UpdatedAt,
	}

	if len(procedurecost.File) > 0 {
		for _, fileID := range procedurecost.File {
			file, err := r.Repo.GetFileByID(fileID)

			if err != nil {
				return nil, errors.Wrap(err, "repo get file by id")
			}

			FileDropdown := dto.FileDropdownSimple{
				ID:   file.ID,
				Name: file.Name,
				Type: *file.Type,
			}
			response.File = append(response.File, FileDropdown)
		}
	}

	if procedurecost.AccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(procedurecost.AccountID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get account item by id")
		}

		accountDropdown := dto.DropdownSimple{
			ID:    account.ID,
			Title: account.Title,
		}
		response.Account = accountDropdown
	}

	if procedurecost.CourtAccountID != nil {
		courtAccount, err := r.Repo.GetAccountItemByID(*procedurecost.CourtAccountID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get account item by id")
		}

		courtAccountDropdown := &dto.DropdownSimple{
			ID:    courtAccount.ID,
			Title: courtAccount.Title,
		}
		response.CourtAccount = courtAccountDropdown
	}

	return &response, nil
}

func buildProcedureCostResponseItemList(itemList []structs.ProcedureCost, r *Resolver) ([]*dto.ProcedureCostResponseItem, error) {
	var items []*dto.ProcedureCostResponseItem

	for _, item := range itemList {
		singleItem, err := buildProcedureCostResponseItem(item, r)

		if err != nil {
			return nil, errors.Wrap(err, "build procedure cost response item")
		}

		items = append(items, singleItem)

	}

	return items, nil
}
