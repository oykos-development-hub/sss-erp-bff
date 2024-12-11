package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfileFamilyResolver(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"].(int)

	familyMembers, err := r.Repo.GetEmployeeFamilyMembers(userProfileID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	familyMembersRes, err := buildFamilyResponseItemList(r.Repo, familyMembers)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   familyMembersRes,
	}, nil
}

func (r *Resolver) UserProfileFamilyInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Family
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		res, err := r.Repo.UpdateEmployeeFamilyMember(itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		resItem, err := buildFamilyResponseItem(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Item = resItem
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateEmployeeFamilyMember(&data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		resItem, err := buildFamilyResponseItem(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Item = resItem
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) UserProfileFamilyDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteEmployeeFamilyMember(itemID.(int))
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildFamilyResponseItemList(r repository.MicroserviceRepositoryInterface, family []structs.Family) (familyResponseItemList []dto.Family, err error) {
	for _, member := range family {
		memberResItem, err := buildFamilyResponseItem(r, &member)
		if err != nil {
			return nil, errors.Wrap(err, "build contract response item")
		}
		familyResponseItemList = append(familyResponseItemList, *memberResItem)
	}
	return
}

func buildFamilyResponseItem(r repository.MicroserviceRepositoryInterface, familyMember *structs.Family) (*dto.Family, error) {
	resItem := dto.Family{
		ID:                   familyMember.ID,
		UserProfileID:        familyMember.UserProfileID,
		FirstName:            familyMember.FirstName,
		MiddleName:           familyMember.MiddleName,
		LastName:             familyMember.LastName,
		BirthLastName:        familyMember.BirthLastName,
		FatherName:           familyMember.FatherName,
		MotherName:           familyMember.MotherName,
		MotherBirthLastName:  familyMember.MotherBirthLastName,
		DateOfBirth:          familyMember.DateOfBirth,
		CountryOfBirth:       familyMember.CountryOfBirth,
		CityOfBirth:          familyMember.CityOfBirth,
		Nationality:          familyMember.Nationality,
		NationalMinority:     familyMember.NationalMinority,
		Citizenship:          familyMember.Citizenship,
		Address:              familyMember.Address,
		OfficialPersonalID:   familyMember.OfficialPersonalID,
		Gender:               familyMember.Gender,
		EmployeeRelationship: familyMember.EmployeeRelationship,
		InsuranceCoverage:    familyMember.InsuranceCoverage,
		HandicappedPerson:    familyMember.HandicappedPerson,
		Files:                make([]dto.FileDropdownSimple, 0, len(familyMember.FileIDs)),
		CreatedAt:            familyMember.CreatedAt,
		UpdatedAt:            familyMember.UpdatedAt,
	}

	for i := range familyMember.FileIDs {
		var file dto.FileDropdownSimple

		res, _ := r.GetFileByID(familyMember.FileIDs[i])

		if res != nil {
			file.ID = res.ID
			file.Name = res.Name
			file.Type = *res.Type
		}

		resItem.Files = append(resItem.Files, file)
	}

	return &resItem, nil
}
