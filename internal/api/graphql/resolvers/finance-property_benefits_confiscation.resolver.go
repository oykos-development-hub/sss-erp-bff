package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) PropBenConfInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PropBenConf
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

	var item *structs.PropBenConf

	if data.ID == 0 {
		item, err = r.Repo.CreatePropBenConf(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdatePropBenConf(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}
	propBenConfResItem, err := buildPropBenConfResponseItem(*item, r)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	response.Item = propBenConfResItem

	return response, nil
}

func (r *Resolver) PropBenConfOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		PropBenConf, err := r.Repo.GetPropBenConf(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		PropBenConfResItem, err := buildPropBenConfResponseItem(*PropBenConf, r)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.PropBenConfResponseItem{PropBenConfResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetPropBenConfListInputMS{}
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

	if value, ok := params.Args["property_benefits_confiscation_type_id"].(int); ok && value != 0 {
		input.FilterByPropBenConfTypeID = &value
	}

	PropBenConfs, total, err := r.Repo.GetPropBenConfList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	PropBenConfResItem, err := buildPropBenConfResponseItemList(PropBenConfs, r)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   PropBenConfResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) PropBenConfDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeletePropBenConf(itemID)
	if err != nil {
		fmt.Printf("Deleting property benefit confiscation item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildPropBenConfResponseItem(PropBenConf structs.PropBenConf, r *Resolver) (*dto.PropBenConfResponseItem, error) {
	status := dto.FinancialPropBenConfUnpaidStatus
	switch PropBenConf.Status {
	case structs.PaidPropBenConfStatus:
		status = dto.FinancialPropBenConfPaidStatus
	case structs.UnpaidPropBenConfStatus:
		status = dto.FinancialPropBenConfUnpaidStatus
	case structs.PartPropBenConfStatus:
		status = dto.FinancialPropBenConfPartStatus
	}

	propBenConfType := dto.FinancialPropBenConfDissisionType
	switch PropBenConf.PropBenConfType {
	case structs.VerdictPropBenConfType:
		propBenConfType = dto.FinancialPropBenConfVerdictType
	case structs.DissisionPropBenConfType:
		propBenConfType = dto.FinancialPropBenConfDissisionType
	}

	response := dto.PropBenConfResponseItem{
		ID:                     PropBenConf.ID,
		PropBenConfType:        propBenConfType,
		DecisionNumber:         PropBenConf.DecisionNumber,
		DecisionDate:           PropBenConf.DecisionDate,
		Subject:                PropBenConf.Subject,
		JMBG:                   PropBenConf.JMBG,
		Residence:              PropBenConf.Residence,
		Amount:                 PropBenConf.Amount,
		PaymentReferenceNumber: PropBenConf.PaymentReferenceNumber,
		DebitReferenceNumber:   PropBenConf.DebitReferenceNumber,
		ExecutionDate:          PropBenConf.ExecutionDate,
		PaymentDeadlineDate:    PropBenConf.PaymentDeadlineDate,
		Description:            PropBenConf.Description,
		Status:                 status,
		CourtCosts:             PropBenConf.CourtCosts,
		PropBenConfDetailsDTO:  PropBenConf.PropBenConfDetailsDTO,
		CreatedAt:              PropBenConf.CreatedAt,
		UpdatedAt:              PropBenConf.UpdatedAt,
	}

	if len(PropBenConf.File) > 0 {
		for _, fileID := range PropBenConf.File {
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

	if PropBenConf.AccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(PropBenConf.AccountID)

		if err != nil {
			return nil, err
		}

		accountDropdown := dto.DropdownSimple{
			ID:    account.ID,
			Title: account.Title,
		}
		response.Account = accountDropdown
	}

	if PropBenConf.CourtAccountID != nil {
		courtAccount, err := r.Repo.GetAccountItemByID(*PropBenConf.CourtAccountID)

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

func buildPropBenConfResponseItemList(itemList []structs.PropBenConf, r *Resolver) ([]*dto.PropBenConfResponseItem, error) {
	var items []*dto.PropBenConfResponseItem

	for _, item := range itemList {
		singleItem, err := buildPropBenConfResponseItem(item, r)

		if err != nil {
			return nil, err
		}

		items = append(items, singleItem)
	}

	return items, nil
}
