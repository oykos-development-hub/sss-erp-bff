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

var JobTenderTypesResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	search := params.Args["search"]
	var items []*structs.JobTenderTypes

	if id != nil && shared.IsInteger(id) && id != 0 {
		tenderType, err := getTenderType(id.(int))
		if err != nil {
			return dto.Response{
				Status:  "error",
				Message: err.Error(),
			}, nil
		}
		items = append(items, tenderType)
	} else {
		input := dto.GetJobTenderTypeInputMS{}
		if search != nil {
			search := search.(string)
			input.Search = &search
		}
		tenderTypes, err := getTenderTypeList(&input)
		if err != nil {
			return dto.Response{
				Status:  "error",
				Message: err.Error(),
			}, nil
		}
		items = tenderTypes
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
	}, nil
}

var JobTenderTypeInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobTenderTypes
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateJobTenderType(itemId, &data)
		if err != nil {
			fmt.Printf("Updating Job Tender Type failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error updating Job Tender Type data"), nil
		}

		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := createJobTenderType(&data)
		if err != nil {
			fmt.Printf("Creating Job Tender Type failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating Job Tender Type data"), nil
		}

		response.Item = res
		response.Message = "You created this item!"
	}

	return response, nil
}

var JobTenderTypeDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteJobTenderType(itemId)
	if err != nil {
		fmt.Printf("Deleting job tender type failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func getTenderTypeList(input *dto.GetJobTenderTypeInputMS) ([]*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_TENDER_TYPES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getTenderType(id int) (*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_TENDER_TYPES_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteJobTenderType(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.JOB_TENDER_TYPES_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func createJobTenderType(jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.JOB_TENDER_TYPES_ENDPOINT, jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateJobTenderType(id int, jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.JOB_TENDER_TYPES_ENDPOINT+"/"+strconv.Itoa(id), jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
