package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

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
		return errors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	var item *structs.ProcedureCost

	if data.ID == 0 {
		item, err = r.Repo.CreateProcedureCost(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdateProcedureCost(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}
	procedurecostResItem, err := buildProcedureCostResponseItem(*item, r)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	response.Item = procedurecostResItem

	return response, nil
}

func (r *Resolver) ProcedureCostOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		procedurecost, err := r.Repo.GetProcedureCost(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		procedurecostResItem, err := buildProcedureCostResponseItem(*procedurecost, r)
		if err != nil {
			return errors.HandleAPIError(err)
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
		return errors.HandleAPIError(err)
	}

	procedurecostResItem, err := buildProcedureCostResponseItemList(procedurecosts, r)
	if err != nil {
		return errors.HandleAPIError(err)
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

	err := r.Repo.DeleteProcedureCost(itemID)
	if err != nil {
		fmt.Printf("Deleting procedure cost item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildProcedureCostResponseItem(procedurecost structs.ProcedureCost, r *Resolver) (*dto.ProcedureCostResponseItem, error) {
	status := dto.FinancialProcedureCostUnpaidStatus
	switch procedurecost.Status {
	case structs.PaidProcedureCostStatus:
		status = dto.FinancialProcedureCostPaidStatus
	case structs.UnpaidProcedureCostStatus:
		status = dto.FinancialProcedureCostUnpaidStatus
	case structs.PartProcedureCostStatus:
		status = dto.FinancialProcedureCostPartStatus
	}

	actType := dto.FinancialProcedureCostDissisionType
	switch procedurecost.ProcedureCostType {
	case structs.VerdictProcedureCostType:
		actType = dto.FinancialProcedureCostVerdictType
	case structs.DissisionProcedureCostType:
		actType = dto.FinancialProcedureCostDissisionType
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
				return nil, err
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
			return nil, err
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
			return nil, err
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
			return nil, err
		}

		items = append(items, singleItem)

	}

	return items, nil
}
