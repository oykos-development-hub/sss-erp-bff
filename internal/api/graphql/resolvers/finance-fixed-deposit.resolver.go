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

func (r *Resolver) FixedDepositOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		fixedDeposit, err := r.Repo.GetFixedDepositByID(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		res, err := buildFixedDeposit(*fixedDeposit, r)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.FixedDepositResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.FixedDepositFilter{}
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

	if value, ok := params.Args["type"].(string); ok && value != "" {
		input.Type = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	} else {
		input.OrganizationUnitID, _ = params.Context.Value(config.OrganizationUnitIDKey).(*int)
	}

	if value, ok := params.Args["judge_id"].(int); ok && value != 0 {
		input.JudgeID = &value
	}

	items, total, err := r.Repo.GetFixedDepositList(input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var resItems []dto.FixedDepositResponse
	for _, item := range items {
		resItem, err := buildFixedDeposit(item, r)

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

func (r *Resolver) FixedDepositInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FixedDeposit
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

	var item *structs.FixedDeposit

	if data.ID == 0 {
		item, err = r.Repo.CreateFixedDeposit(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdateFixedDeposit(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	singleItem, err := buildFixedDeposit(*item, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) FixedDepositDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFixedDeposit(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) FixedDepositItemInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FixedDepositItem
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

	if data.ID == 0 {
		err = r.Repo.CreateFixedDepositItem(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		err = r.Repo.UpdateFixedDepositItem(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	return response, nil
}

func (r *Resolver) FixedDepositItemDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFixedDepositItem(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) FixedDepositDispatchInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FixedDepositDispatch
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

	if data.ID == 0 {
		err = r.Repo.CreateFixedDepositDispatch(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		err = r.Repo.UpdateFixedDepositDispatch(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	return response, nil
}

func (r *Resolver) FixedDepositDispatchDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFixedDepositDispatch(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) FixedDepositJudgeInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FixedDepositJudge
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

	if data.ID == 0 {
		err = r.Repo.CreateFixedDepositJudge(&data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		err = r.Repo.UpdateFixedDepositJudge(&data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	return response, nil
}

func (r *Resolver) FixedDepositJudgeDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFixedDepositJudge(itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) FixedDepositWillOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		fixedDeposit, err := r.Repo.GetFixedDepositWillByID(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		res, err := buildFixedDepositWill(*fixedDeposit, r)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.FixedDepositWillResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.FixedDepositWillFilter{}
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

	items, total, err := r.Repo.GetFixedDepositWillList(input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var resItems []dto.FixedDepositWillResponse
	for _, item := range items {
		resItem, err := buildFixedDepositWill(item, r)

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

func (r *Resolver) FixedDepositWillInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FixedDepositWill
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

	var item *structs.FixedDepositWill

	if data.ID == 0 {
		data.Status = "Depozit"
		item, err = r.Repo.CreateFixedDepositWill(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdateFixedDepositWill(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	singleItem, err := buildFixedDepositWill(*item, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) FixedDepositWillDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFixedDepositWill(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) FixedDepositWillDispatchInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FixedDepositWillDispatch
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

	if data.ID == 0 {
		err = r.Repo.CreateFixedDepositWillDispatch(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		err = r.Repo.UpdateFixedDepositWillDispatch(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	return response, nil
}

func (r *Resolver) FixedDepositWillDispatchDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFixedDepositWillDispatch(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildFixedDeposit(item structs.FixedDeposit, r *Resolver) (*dto.FixedDepositResponse, error) {
	response := dto.FixedDepositResponse{
		ID:                   item.ID,
		Subject:              item.Subject,
		CaseNumber:           item.CaseNumber,
		DateOfRecipiet:       item.DateOfRecipiet,
		DateOfCase:           item.DateOfCase,
		DateOfFinality:       item.DateOfFinality,
		DateOfEnforceability: item.DateOfEnforceability,
		DateOfEnd:            item.DateOfEnd,
		Status:               item.Status,
		Type:                 item.Type,
		CreatedAt:            item.CreatedAt,
		UpdatedAt:            item.UpdatedAt,
	}

	if item.OrganizationUnitID != 0 {
		orgUnit, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}

		response.OrganizationUnit.ID = orgUnit.ID
		response.OrganizationUnit.Title = orgUnit.Title
	}

	if item.JudgeID != 0 {
		judge, err := r.Repo.GetUserProfileByID(item.JudgeID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}

		response.Judge.ID = judge.ID
		response.Judge.Title = judge.FirstName + " " + judge.LastName
	}

	if item.AccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(item.AccountID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get account item by id")
		}

		response.Account.ID = account.ID
		response.Account.Title = account.Title
	}

	if item.FileID != 0 {
		file, err := r.Repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}

		response.File.ID = file.ID
		response.File.Name = file.Name
		response.File.Type = *file.Type
	}

	for _, itemFixed := range item.Items {
		builtItem, err := buildFixedDepositItem(itemFixed, r)

		if err != nil {
			return nil, errors.Wrap(err, "build fixed deposit item")
		}

		response.Items = append(response.Items, *builtItem)
	}

	for _, dispatch := range item.Dispatches {
		builtDispatch, err := buildFixedDepositDispatches(dispatch, r)

		if err != nil {
			return nil, errors.Wrap(err, "build fixed deposit dispatches")
		}

		response.Dispatches = append(response.Dispatches, *builtDispatch)
	}

	for _, judge := range item.Judges {
		builtJudge, err := buildFixedDepositJudges(judge, r)

		if err != nil {
			return nil, errors.Wrap(err, "build fixed deposit judges")
		}

		response.Judges = append(response.Judges, *builtJudge)
	}

	return &response, nil
}

func buildFixedDepositWill(item structs.FixedDepositWill, r *Resolver) (*dto.FixedDepositWillResponse, error) {
	response := dto.FixedDepositWillResponse{
		ID:              item.ID,
		Subject:         item.Subject,
		FatherName:      item.FatherName,
		DateOfBirth:     item.DateOfBirth,
		JMBG:            item.JMBG,
		CaseNumberSI:    item.CaseNumberSI,
		CaseNumberRS:    item.CaseNumberRS,
		Description:     item.Description,
		DateOfReceiptSI: item.DateOfReceiptSI,
		DateOfReceiptRS: item.DateOfReceiptRS,
		DateOfEnd:       item.DateOfEnd,
		Status:          item.Status,
	}

	if item.OrganizationUnitID != 0 {
		orgUnit, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}

		response.OrganizationUnit.ID = orgUnit.ID
		response.OrganizationUnit.Title = orgUnit.Title
	}

	if item.FileID != 0 {
		file, err := r.Repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}

		response.File.ID = file.ID
		response.File.Name = file.Name
		response.File.Type = *file.Type
	}

	for _, dispatch := range item.Dispatches {
		builtDispatch, err := buildFixedDepositWillDispatches(dispatch, r)

		if err != nil {
			return nil, errors.Wrap(err, "build fixed deposit will dispatches")
		}

		response.Dispatches = append(response.Dispatches, *builtDispatch)
	}

	for _, judge := range item.Judges {
		builtJudge, err := buildFixedDepositJudges(judge, r)

		if err != nil {
			return nil, errors.Wrap(err, "build fixed deposit judges")
		}

		response.Judges = append(response.Judges, *builtJudge)
	}

	return &response, nil
}

func buildFixedDepositItem(item structs.FixedDepositItem, r *Resolver) (*dto.FixedDepositItemResponse, error) {
	response := dto.FixedDepositItemResponse{
		ID:                 item.ID,
		DepositID:          item.DepositID,
		Amount:             item.Amount,
		SerialNumber:       item.SerialNumber,
		DateOfConfiscation: item.DateOfConfiscation,
		CaseNumber:         item.CaseNumber,
		Unit:               item.Unit,
		Currency:           item.Currency,
		CreatedAt:          item.CreatedAt,
		UpdatedAt:          item.UpdatedAt,
	}

	if item.CategoryID != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(item.CategoryID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}

		response.Category.ID = setting.ID
		response.Category.Title = setting.Title
	}

	if item.TypeID != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(item.TypeID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}

		response.Type.ID = setting.ID
		response.Type.Title = setting.Title
	}

	if item.FileID != 0 {
		file, err := r.Repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}

		response.File.ID = file.ID
		response.File.Name = file.Name
		response.File.Type = *file.Type
	}

	if item.JudgeID != 0 {
		judge, err := r.Repo.GetUserProfileByID(item.JudgeID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}

		response.Judge.ID = judge.ID
		response.Judge.Title = judge.FirstName + " " + judge.LastName
	}

	return &response, nil
}

func buildFixedDepositDispatches(item structs.FixedDepositDispatch, r *Resolver) (*dto.FixedDepositDispatchResponse, error) {
	response := dto.FixedDepositDispatchResponse{
		ID:           item.ID,
		DepositID:    item.DepositID,
		Amount:       item.Amount,
		SerialNumber: item.SerialNumber,
		DateOfAction: item.DateOfAction,
		CaseNumber:   item.CaseNumber,
		Subject:      item.Subject,
		Unit:         item.Unit,
		Currency:     item.Currency,
		Action:       item.Action,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	if item.CategoryID != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(item.CategoryID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}

		response.Category.ID = setting.ID
		response.Category.Title = setting.Title
	}

	if item.TypeID != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(item.TypeID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}

		response.Type.ID = setting.ID
		response.Type.Title = setting.Title
	}

	if item.FileID != 0 {
		file, err := r.Repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}

		response.File.ID = file.ID
		response.File.Name = file.Name
		response.File.Type = *file.Type
	}

	if item.JudgeID != 0 {
		judge, err := r.Repo.GetUserProfileByID(item.JudgeID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}

		response.Judge.ID = judge.ID
		response.Judge.Title = judge.FirstName + " " + judge.LastName
	}

	return &response, nil
}

func buildFixedDepositWillDispatches(item structs.FixedDepositWillDispatch, r *Resolver) (*dto.FixedDepositWillDispatchResponse, error) {
	response := dto.FixedDepositWillDispatchResponse{
		ID:             item.ID,
		CaseNumber:     item.CaseNumber,
		WillID:         item.WillID,
		DateOfDispatch: item.DateOfDispatch,
		DispatchType:   item.DispatchType,
		CreatedAt:      item.CreatedAt,
		UpdatedAt:      item.UpdatedAt,
	}

	if item.FileID != 0 {
		file, err := r.Repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}

		response.File.ID = file.ID
		response.File.Name = file.Name
		response.File.Type = *file.Type
	}

	if item.JudgeID != 0 {
		judge, err := r.Repo.GetUserProfileByID(item.JudgeID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}

		response.Judge.ID = judge.ID
		response.Judge.Title = judge.FirstName + " " + judge.LastName
	}

	return &response, nil
}

func buildFixedDepositJudges(item structs.FixedDepositJudge, r *Resolver) (*dto.FixedDepositJudgeResponse, error) {
	response := dto.FixedDepositJudgeResponse{
		ID:          item.ID,
		DepositID:   item.DepositID,
		WillID:      item.WillID,
		DateOfStart: item.DateOfStart,
		DateOfEnd:   item.DateOfEnd,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}

	if item.FileID != 0 {
		file, err := r.Repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}

		response.File.ID = file.ID
		response.File.Name = file.Name
		response.File.Type = *file.Type
	}

	if item.JudgeID != 0 {
		judge, err := r.Repo.GetUserProfileByID(item.JudgeID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}

		response.Judge.ID = judge.ID
		response.Judge.Title = judge.FirstName + " " + judge.LastName
	}

	return &response, nil
}
