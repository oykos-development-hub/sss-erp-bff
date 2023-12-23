package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func buildSalaryResponseItemList(r repository.MicroserviceRepositoryInterface, items []*structs.SalaryParams) (resItemList []*dto.SalaryParams, err error) {
	for _, item := range items {
		resItem, err := buildSalaryResponseItem(r, item)
		if err != nil {
			return nil, err
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

func buildSalaryResponseItem(r repository.MicroserviceRepositoryInterface, item *structs.SalaryParams) (resItem *dto.SalaryParams, err error) {
	organizationUnit, err := r.GetOrganizationUnitByID(item.OrganizationUnitID)
	if err != nil {
		return nil, err
	}

	userProfile, err := r.GetUserProfileByID(item.UserProfileID)
	if err != nil {
		return nil, err
	}

	salaryParams := &dto.SalaryParams{
		ID: item.ID,
		UserProfile: dto.DropdownSimple{
			ID:    userProfile.ID,
			Title: userProfile.GetFullName(),
		},
		OrganizationUnit: dto.DropdownSimple{
			ID:    organizationUnit.ID,
			Title: organizationUnit.Title,
		},
		BenefitedTrack:      item.BenefitedTrack,
		WithoutRaise:        item.WithoutRaise,
		InsuranceBasis:      item.InsuranceBasis,
		SalaryRank:          item.SalaryRank,
		DailyWorkHours:      item.DailyWorkHours,
		WeeklyWorkHours:     item.WeeklyWorkHours,
		EducationRank:       item.EducationRank,
		EducationNaming:     item.EducationNaming,
		ObligationReduction: item.ObligationReduction,
		CreatedAt:           item.CreatedAt,
		UpdatedAt:           item.UpdatedAt,
	}

	if item.UserResolutionID != nil {
		resolution, err := r.GetEmployeeResolution(*item.UserResolutionID)
		if err != nil {
			return nil, err
		}
		resolutionResItem, err := buildResolutionResItem(r, resolution)
		if err != nil {
			return nil, err
		}
		salaryParams.Resolution = *resolutionResItem
	}

	return salaryParams, nil
}

func (r *Resolver) UserProfileSalaryParamsResolver(params graphql.ResolveParams) (interface{}, error) {
	profileID := params.Args["user_profile_id"].(int)

	salaries, err := r.Repo.GetEmployeeSalaryParams(profileID)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	salaryResponseItemList, err := buildSalaryResponseItemList(r.Repo, salaries)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   salaryResponseItemList,
	}, nil
}

func (r *Resolver) UserProfileSalaryParamsInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error

	var data structs.SalaryParams
	response := dto.ResponseSingle{
		Status: "success",
	}

	var item *structs.SalaryParams

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.ErrorResponse("Error updating settings data"), nil
	}

	itemID := data.ID
	if itemID != 0 {
		item, err = r.Repo.UpdateEmployeeSalaryParams(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
	} else {
		item, err = r.Repo.CreateEmployeeSalaryParams(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
	}

	salaryResItem, err := buildSalaryResponseItem(r.Repo, item)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	response.Item = salaryResItem

	return response, nil
}

func (r *Resolver) UserProfileSalaryParamsDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]
	err := r.Repo.DeleteSalaryParams(itemID.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
