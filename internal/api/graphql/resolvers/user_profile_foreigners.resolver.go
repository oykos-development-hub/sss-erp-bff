package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfileForeignerResolver(params graphql.ResolveParams) (interface{}, error) {
	profileID := params.Args["user_profile_id"].(int)

	items, err := r.Repo.GetEmployeeForeigners(profileID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	resItems, err := buildForeignerResponseItemList(r.Repo, items)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   resItems,
	}, nil
}

func (r *Resolver) UserProfileForeignerInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error

	var data structs.Foreigners
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	itemID := data.ID
	if itemID != 0 {
		item, err := r.Repo.UpdateEmployeeForeigner(itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		resItem, err := buildForeignerResponseItem(r.Repo, item)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Item = resItem
		response.Message = "You updated this item!"
	} else {
		item, err := r.Repo.CreateEmployeeForeigner(&data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		resItem, err := buildForeignerResponseItem(r.Repo, item)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Item = resItem
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) UserProfileForeignerDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteForeigner(itemID.(int))
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildForeignerResponseItemList(r repository.MicroserviceRepositoryInterface, foreigners []structs.Foreigners) (foreignerResponseItemList []dto.Foreigner, err error) {
	for _, foreigner := range foreigners {
		memberResItem, err := buildForeignerResponseItem(r, &foreigner)
		if err != nil {
			return nil, errors.Wrap(err, "build contract response item")
		}
		foreignerResponseItemList = append(foreignerResponseItemList, *memberResItem)
	}

	return
}

func buildForeignerResponseItem(r repository.MicroserviceRepositoryInterface, foreigner *structs.Foreigners) (*dto.Foreigner, error) {
	resItem := dto.Foreigner{
		ID:                              foreigner.ID,
		UserProfileID:                   foreigner.UserProfileID,
		WorkPermitNumber:                foreigner.WorkPermitNumber,
		WorkPermitIssuer:                foreigner.WorkPermitIssuer,
		WorkPermitDateOfStart:           foreigner.WorkPermitDateOfStart,
		WorkPermitDateOfEnd:             foreigner.WorkPermitDateOfEnd,
		WorkPermitIndefiniteLength:      foreigner.WorkPermitIndefiniteLength,
		ResidencePermitDateOfStart:      foreigner.ResidencePermitDateOfStart,
		ResidencePermitDateOfEnd:        foreigner.ResidencePermitDateOfEnd,
		ResidencePermitIndefiniteLength: foreigner.ResidencePermitIndefiniteLength,
		ResidencePermitNumber:           foreigner.ResidencePermitNumber,
		ResidencePermitIssuer:           foreigner.ResidencePermitIssuer,
		CountryOfOrigin:                 foreigner.CountryOfOrigin,
		CreatedAt:                       foreigner.CreatedAt,
		UpdatedAt:                       foreigner.UpdatedAt,
		WorkPermitFileID:                foreigner.WorkPermitFileID,
		ResidencePermitFileID:           foreigner.ResidencePermitFileID,
		Files:                           make([]dto.FileDropdownSimple, 0, len(foreigner.FileIDs)),
	}

	for i := range foreigner.FileIDs {
		var file dto.FileDropdownSimple

		res, _ := r.GetFileByID(foreigner.FileIDs[i])

		if res != nil {
			file.ID = res.ID
			file.Name = res.Name
			file.Type = *res.Type
		}

		resItem.Files = append(resItem.Files, file)
	}

	return &resItem, nil
}
