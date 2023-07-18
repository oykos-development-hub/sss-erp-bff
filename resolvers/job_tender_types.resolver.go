package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

var JobTenderTypesResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	search := params.Args["search"]
	JobTenderTypeType := &structs.JobTenderTypes{}
	JobTenderTypeData, JobTenderTypeDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tender_types.json", JobTenderTypeType)

	if JobTenderTypeDataErr != nil {
		fmt.Printf("Fetching Job Tender types failed because of this error - %s.\n", JobTenderTypeDataErr)
	}

	if id != nil && shared.IsInteger(id) && id != 0 {
		JobTenderTypeData = shared.FindByProperty(JobTenderTypeData, "Id", id)
	}

	if search != nil && shared.IsString(search) {
		JobTenderTypeData = shared.FindByProperty(JobTenderTypeData, "Title", search, true)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   JobTenderTypeData,
	}, nil
}

var JobTenderTypeInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.JobTenderTypes
	dataBytes, _ := json.Marshal(params.Args["data"])
	JobTenderTypeType := &structs.JobTenderTypes{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	JobTenderTypeData, JobTenderTypeDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tender_types.json", JobTenderTypeType)

	if JobTenderTypeDataErr != nil {
		fmt.Printf("Fetching Job Tender types failed because of this error - %s.\n", JobTenderTypeDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		JobTenderTypeData = shared.FilterByProperty(JobTenderTypeData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(JobTenderTypeData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_tender_types.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var JobTenderTypeDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	JobTenderTypeType := &structs.JobTenderTypes{}
	JobTenderTypeData, JobTenderTypeDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tender_types.json", JobTenderTypeType)

	if JobTenderTypeDataErr != nil {
		fmt.Printf("Fetching Job Tender types failed because of this error - %s.\n", JobTenderTypeDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		JobTenderTypeData = shared.FilterByProperty(JobTenderTypeData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_tender_types.json"), JobTenderTypeData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
