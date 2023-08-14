package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

func buildSalaryResponseItemList(items []*structs.SalaryParams) (resItemList []*dto.SalaryParams, err error) {
	for _, item := range items {
		resItem, err := buildSalaryResponseItem(item)
		if err != nil {
			return nil, err
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

func buildSalaryResponseItem(item *structs.SalaryParams) (resItem *dto.SalaryParams, err error) {
	var resolutionResItem *dto.Resolution
	organizationUnit, err := getOrganizationUnitById(item.OrganizationUnitID)
	if err != nil {
		return nil, err
	}

	if item.UserResolutionId != nil {
		resolution, err := getEmployeeResolution(*item.UserResolutionId)
		if err != nil {
			return nil, err
		}
		resolutionResItem, err = buildResolutionResItem(resolution)
		if err != nil {
			return nil, err
		}
	}

	userProfile, err := getUserProfileById(item.UserProfileId)
	if err != nil {
		return nil, err
	}

	return &dto.SalaryParams{
		Id: item.Id,
		UserProfile: dto.DropdownSimple{
			Id:    userProfile.Id,
			Title: userProfile.GetFullName(),
		},
		OrganizationUnit: dto.DropdownSimple{
			Id:    organizationUnit.Id,
			Title: organizationUnit.Title,
		},
		Resolution:      *resolutionResItem,
		BenefitedTrack:  item.BenefitedTrack,
		WithoutRaise:    item.WithoutRaise,
		InsuranceBasis:  item.InsuranceBasis,
		SalaryRank:      item.SalaryRank,
		DailyWorkHours:  item.DailyWorkHours,
		WeeklyWorkHours: item.WeeklyWorkHours,
		EducationRank:   item.EducationRank,
		EducationNaming: item.EducationNaming,
		CreatedAt:       item.CreatedAt,
		UpdatedAt:       item.UpdatedAt,
	}, nil
}

var UserProfileSalaryParamsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	profileId := params.Args["user_profile_id"].(int)

	salaries, err := getEmployeeSalaryParams(profileId)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	salaryResponseItemList, err := buildSalaryResponseItemList(salaries)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   salaryResponseItemList,
	}, nil
}

var UserProfileSalaryParamsInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
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
		return shared.ErrorResponse("Error updating settings data"), nil
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		item, err = updateEmployeeSalaryParams(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
	} else {
		item, err = createEmployeeSalaryParams(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You created this item!"
	}

	salaryResItem, err := buildSalaryResponseItem(item)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	response.Item = salaryResItem

	return response, nil
}

var UserProfileSalaryParamsDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]
	err := deleteSalaryParams(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func getEmployeeSalaryParams(userProfileID int) ([]*structs.SalaryParams, error) {
	res := &dto.GetEmployeeSalaryParamsListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(userProfileID)+"/salaries", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func createEmployeeSalaryParams(salaries *structs.SalaryParams) (*structs.SalaryParams, error) {
	res := dto.GetEmployeeSalaryParamsResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.SALARIES, salaries, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func deleteSalaryParams(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.SALARIES+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func updateEmployeeSalaryParams(id int, salaries *structs.SalaryParams) (*structs.SalaryParams, error) {
	res := dto.GetEmployeeSalaryParamsResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.SALARIES+"/"+strconv.Itoa(id), salaries, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
