package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var item *structs.PropBenConf

	if data.OrganizationUnitID == 0 {
		organizationUnitID, _ := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if organizationUnitID != nil {
			data.OrganizationUnitID = *organizationUnitID
		}
	}

	if data.ID == 0 {
		item, err = r.Repo.CreatePropBenConf(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdatePropBenConf(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}
	propBenConfResItem, err := buildPropBenConfResponseItem(*item, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	response.Item = propBenConfResItem

	return response, nil
}

func (r *Resolver) PropBenConfOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		PropBenConf, err := r.Repo.GetPropBenConf(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		PropBenConfResItem, err := buildPropBenConfResponseItem(*PropBenConf, r)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	}

	PropBenConfs, total, err := r.Repo.GetPropBenConfList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	PropBenConfResItem, err := buildPropBenConfResponseItemList(PropBenConfs, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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

	err := r.Repo.DeletePropBenConf(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildPropBenConfResponseItem(PropBenConf structs.PropBenConf, r *Resolver) (*dto.PropBenConfResponseItem, error) {
	status := dto.DropdownSimple{
		ID:    int(structs.UnpaidPropBenConfStatus),
		Title: string(dto.FinancialPropBenConfUnpaidStatus),
	}
	switch PropBenConf.Status {
	case structs.PaidPropBenConfStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PaidPropBenConfStatus),
			Title: string(dto.FinancialPropBenConfPaidStatus),
		}
	case structs.UnpaidPropBenConfStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.UnpaidPropBenConfStatus),
			Title: string(dto.FinancialPropBenConfUnpaidStatus),
		}
	case structs.PartPropBenConfStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PartPropBenConfStatus),
			Title: string(dto.FinancialPropBenConfPartStatus),
		}
	}

	propBenConfType := dto.DropdownSimple{
		ID:    int(structs.DissisionPropBenConfType),
		Title: string(dto.FinancialPropBenConfDissisionType),
	}
	switch PropBenConf.PropBenConfType {
	case structs.DissisionPropBenConfType:
		propBenConfType = dto.DropdownSimple{
			ID:    int(structs.DissisionPropBenConfType),
			Title: string(dto.FinancialPropBenConfDissisionType),
		}
	case structs.VerdictPropBenConfType:
		propBenConfType = dto.DropdownSimple{
			ID:    int(structs.VerdictPropBenConfType),
			Title: string(dto.FinancialPropBenConfVerdictType),
		}

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

	if PropBenConf.AccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(PropBenConf.AccountID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get account item by id")
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
			return nil, errors.Wrap(err, "repo get account item by id")
		}

		courtAccountDropdown := &dto.DropdownSimple{
			ID:    courtAccount.ID,
			Title: courtAccount.Title,
		}
		response.CourtAccount = courtAccountDropdown
	}

	if PropBenConf.OrganizationUnitID != 0 {
		organizationUnit, err := r.Repo.GetOrganizationUnitByID(PropBenConf.OrganizationUnitID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}

		orgUnitDropdown := dto.DropdownSimple{
			ID:    organizationUnit.ID,
			Title: organizationUnit.Title,
		}

		response.OrganizationUnit = orgUnitDropdown
	}

	return &response, nil
}

func buildPropBenConfResponseItemList(itemList []structs.PropBenConf, r *Resolver) ([]*dto.PropBenConfResponseItem, error) {
	var items []*dto.PropBenConfResponseItem

	for _, item := range itemList {
		singleItem, err := buildPropBenConfResponseItem(item, r)

		if err != nil {
			return nil, errors.Wrap(err, "build prop ben conf response item")
		}

		items = append(items, singleItem)
	}

	return items, nil
}
