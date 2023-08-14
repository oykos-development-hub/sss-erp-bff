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

var UserProfileSalaryParamsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	profileId := params.Args["user_profile_id"].(int)

	res, err := getEmployeeSalaryParams(profileId)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	for _, salary := range res {
		organizationUnit, err := getOrganizationUnitById(salary.OrganizationUnitID)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		salary.OrganizationUnit = structs.SettingsDropdown{Id: organizationUnit.Id, Title: organizationUnit.Title}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   res,
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

	organizationUnit, err := getOrganizationUnitById(item.OrganizationUnitID)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	item.OrganizationUnit = structs.SettingsDropdown{Id: organizationUnit.Id, Title: organizationUnit.Title}

	response.Item = item

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
