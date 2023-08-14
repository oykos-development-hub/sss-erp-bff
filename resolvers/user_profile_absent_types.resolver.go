package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

var AbsentTypeResolver = func(params graphql.ResolveParams) (interface{}, error) {
	absentTypesAll, err := getAbsentTypes()
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   absentTypesAll.Data,
		Total:   absentTypesAll.Total,
	}, nil
}

var AbsentTypeInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var data structs.AbsentType

	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		item, err := updateAbsentType(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := createAbsentType(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

var AbsentTypeDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := deleteAbsentType(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func deleteAbsentType(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.ABSENT_TYPE+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func updateAbsentType(id int, absent *structs.AbsentType) (*structs.AbsentType, error) {
	res := &dto.GetAbsentTypeResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ABSENT_TYPE+"/"+strconv.Itoa(id), absent, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createAbsentType(absent *structs.AbsentType) (*structs.AbsentType, error) {
	res := &dto.GetAbsentTypeResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ABSENT_TYPE, absent, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getAbsentTypes() (*dto.GetAbsentTypeListResponseMS, error) {
	res := &dto.GetAbsentTypeListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ABSENT_TYPE, nil, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getAbsentTypeById(absentTypeId int) (*structs.AbsentType, error) {
	res := &dto.GetAbsentTypeResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ABSENT_TYPE+"/"+strconv.Itoa(absentTypeId), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
