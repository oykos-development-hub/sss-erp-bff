package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) FlatRateInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FlatRate
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

	var item *structs.FlatRate

	if data.ID == 0 {
		item, err = r.Repo.CreateFlatRate(params.Context, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdateFlatRate(params.Context, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

	}

	singleItem, err := buildFlatRateResponseItem(*item, r)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) FlatRateOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		flatrate, err := r.Repo.GetFlatRate(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		flatrateResItem, err := buildFlatRateResponseItem(*flatrate, r)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.FlatRateResponseItem{flatrateResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetFlatRateListInputMS{}
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

	if value, ok := params.Args["flat_rate_type_id"].(int); ok && value != 0 {
		input.FilterByTypeID = &value
	}

	flatrates, total, err := r.Repo.GetFlatRateList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	flatrateResItem, err := buildFlatRateResponseItemList(flatrates, r)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   flatrateResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) FlatRateDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFlatRate(params.Context, itemID)
	if err != nil {
		fmt.Printf("Deleting flatrate item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildFlatRateResponseItem(flatrate structs.FlatRate, r *Resolver) (*dto.FlatRateResponseItem, error) {
	status := dto.DropdownSimple{
		ID:    int(structs.UnpaidFlatRateStatus),
		Title: string(dto.FinancialFlatRateUnpaidStatus),
	}
	switch flatrate.Status {
	case structs.UnpaidFlatRateStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.UnpaidFlatRateStatus),
			Title: string(dto.FinancialFlatRateUnpaidStatus),
		}
	case structs.PaidFlatRateStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PaidFlatRateStatus),
			Title: string(dto.FinancialFlatRatePaidStatus),
		}
	case structs.PartFlatRateStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PartFlatRateStatus),
			Title: string(dto.FinancialFlatRatePartStatus),
		}
	}

	flatRateType := dto.DropdownSimple{
		ID:    int(structs.DissisionType),
		Title: string(dto.FinancialFlatRateDissisionType),
	}
	switch flatrate.FlatRateType {
	case structs.DissisionType:
		flatRateType = dto.DropdownSimple{
			ID:    int(structs.DissisionType),
			Title: string(dto.FinancialFlatRateDissisionType),
		}
	case structs.VerdictType:
		flatRateType = dto.DropdownSimple{
			ID:    int(structs.VerdictType),
			Title: string(dto.FinancialFlatRateVerdictType),
		}
	}

	response := dto.FlatRateResponseItem{
		ID:                     flatrate.ID,
		FlatRateType:           flatRateType,
		DecisionNumber:         flatrate.DecisionNumber,
		DecisionDate:           flatrate.DecisionDate,
		Subject:                flatrate.Subject,
		JMBG:                   flatrate.JMBG,
		Residence:              flatrate.Residence,
		Amount:                 flatrate.Amount,
		PaymentReferenceNumber: flatrate.PaymentReferenceNumber,
		DebitReferenceNumber:   flatrate.DebitReferenceNumber,
		ExecutionDate:          flatrate.ExecutionDate,
		PaymentDeadlineDate:    flatrate.PaymentDeadlineDate,
		Description:            flatrate.Description,
		Status:                 status,
		CourtCosts:             flatrate.CourtCosts,
		FlatRateDetailsDTO:     flatrate.FlatRateDetailsDTO,
		CreatedAt:              flatrate.CreatedAt,
		UpdatedAt:              flatrate.UpdatedAt,
	}

	if len(flatrate.File) > 0 {
		for _, fileID := range flatrate.File {
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

	if flatrate.AccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(flatrate.AccountID)

		if err != nil {
			return nil, err
		}

		accountDropdown := dto.DropdownSimple{
			ID:    account.ID,
			Title: account.Title,
		}
		response.Account = accountDropdown
	}

	if flatrate.CourtAccountID != nil {
		courtAccount, err := r.Repo.GetAccountItemByID(*flatrate.CourtAccountID)

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

func buildFlatRateResponseItemList(itemList []structs.FlatRate, r *Resolver) ([]*dto.FlatRateResponseItem, error) {
	var items []*dto.FlatRateResponseItem

	for _, item := range itemList {
		singleItem, err := buildFlatRateResponseItem(item, r)

		if err != nil {
			return nil, err
		}

		items = append(items, singleItem)

	}

	return items, nil
}
