package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/graphql-go/graphql"
)

func PopulateJobTenderItemProperties(jobTenders []interface{}, id int, organizationUnitId int, typeParam string, isActive ...interface{}) []interface{} {
	var items []interface{}

	for _, item := range jobTenders {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}
		// Filtering by Type
		if len(typeParam) > 0 && typeParam != mergedItem["type"] {
			continue
		}
		// Filtering by status
		if len(isActive) > 0 && isActive[0] != nil && isActive[0] != mergedItem["active"] {
			continue
		}
		// Filtering by Organization Unit
		var relatedJobPositionInOrganizationUnit = shared.FetchByProperty("job_position_in_organization_unit", "Id", mergedItem["position_in_organization_unit_id"])

		if relatedJobPositionInOrganizationUnit == nil || len(relatedJobPositionInOrganizationUnit) < 1 {
			continue
		}

		relatedJobPositionInOrganizationUnitValue := reflect.ValueOf(relatedJobPositionInOrganizationUnit[0])

		if relatedJobPositionInOrganizationUnitValue.Kind() == reflect.Ptr {
			relatedJobPositionInOrganizationUnitValue = relatedJobPositionInOrganizationUnitValue.Elem()
		}

		var relatedOrganizationUnit = shared.FetchByProperty("organization_unit", "Id", relatedJobPositionInOrganizationUnitValue.FieldByName("ParentOrganizationUnitId").Interface())

		if relatedOrganizationUnit == nil || len(relatedOrganizationUnit) < 1 {
			continue
		}

		relatedOrganizationUnitValue := reflect.ValueOf(relatedOrganizationUnit[0])

		if relatedOrganizationUnitValue.Kind() == reflect.Ptr {
			relatedOrganizationUnitValue = relatedOrganizationUnitValue.Elem()
		}

		if shared.IsInteger(organizationUnitId) && organizationUnitId != 0 && relatedOrganizationUnitValue.FieldByName("Id").Interface() != organizationUnitId {
			continue
		}

		mergedItem["organization_unit"] = map[string]interface{}{
			"title": relatedOrganizationUnitValue.FieldByName("Title").Interface(),
			"id":    relatedOrganizationUnitValue.FieldByName("Id").Interface(),
		}

		var relatedJobPosition = shared.FetchByProperty("job_position", "Id", relatedJobPositionInOrganizationUnitValue.FieldByName("JobPositionId").Interface())

		if len(relatedJobPosition) > 0 {
			relatedJobPositionValue := reflect.ValueOf(relatedJobPosition[0])

			if relatedJobPositionValue.Kind() == reflect.Ptr {
				relatedJobPositionValue = relatedJobPositionValue.Elem()
			}

			mergedItem["job_position"] = map[string]interface{}{
				"title": relatedJobPositionValue.FieldByName("Title").Interface(),
				"id":    relatedJobPositionValue.FieldByName("Id").Interface(),
			}
		}

		items = append(items, mergedItem)
	}

	return items
}

func PopulateJobTenderApplicationProperties(jobTenderApplications []interface{}, id int, jobTenderId int) []interface{} {
	var items []interface{}

	for _, item := range jobTenderApplications {

		var mergedItem = shared.WriteStructToInterface(item)
		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}
		// Filtering by Job Tender ID
		if shared.IsInteger(jobTenderId) && jobTenderId != 0 && jobTenderId != mergedItem["job_tender_id"] {
			continue
		}
		// Populating Job Tender data
		if shared.IsInteger(mergedItem["job_tender_id"]) && mergedItem["job_tender_id"].(int) > 0 {
			var relatedJobTender = shared.FetchByProperty(
				"job_tender",
				"Id",
				mergedItem["job_tender_id"],
			)

			if relatedJobTender == nil || len(relatedJobTender) < 1 {
				continue
			}

			var jobTender = shared.WriteStructToInterface(relatedJobTender[0])

			mergedItem["job_tender"] = map[string]interface{}{
				"title": jobTender["serial_number"].(string),
				"id":    jobTender["id"],
			}
		}
		// Populating User Profile data
		if shared.IsInteger(mergedItem["user_profile_id"]) && mergedItem["user_profile_id"].(int) > 0 {
			var relatedUserProfile = shared.FetchByProperty(
				"user_profile",
				"Id",
				mergedItem["user_profile_id"],
			)

			if relatedUserProfile == nil || len(relatedUserProfile) < 1 {
				continue
			}

			var userProfile = shared.WriteStructToInterface(relatedUserProfile[0])

			mergedItem["user_profile"] = map[string]interface{}{
				"title": userProfile["first_name"].(string) + " " + userProfile["last_name"].(string),
				"id":    userProfile["id"],
			}
			mergedItem["first_name"] = userProfile["first_name"].(string)
			mergedItem["last_name"] = userProfile["last_name"].(string)
			mergedItem["official_personal_id"] = userProfile["official_personal_id"].(string)
			mergedItem["date_of_birth"] = userProfile["date_of_birth"].(string)
			mergedItem["nationality"] = userProfile["nationality"].(string)
		}

		items = append(items, mergedItem)
	}

	return items
}

var JobTendersOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}
	var organizationUnitId int
	if params.Args["organization_unit_id"] == nil {
		organizationUnitId = 0
	} else {
		organizationUnitId = params.Args["organization_unit_id"].(int)
	}

	page := params.Args["page"]
	size := params.Args["size"]
	typeParam := params.Args["type"]

	JobTendersType := &structs.JobTenders{}
	JobTendersData, JobTendersDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tenders.json", JobTendersType)

	if JobTendersDataErr != nil {
		fmt.Printf("Fetching Job Tenders failed because of this error - %s.\n", JobTendersDataErr)
	}

	// Populate data for each Job Tender with Organization Unit and Job Position
	items = PopulateJobTenderItemProperties(JobTendersData, id, organizationUnitId, typeParam.(string), params.Args["active"])

	total = len(items)

	// Filtering by Pagination params
	if shared.IsInteger(page) && page != 0 && shared.IsInteger(size) && size != 0 {
		items = shared.Pagination(items, page.(int), size.(int))
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"total":   total,
		"items":   items,
	}, nil
}

var JobTenderResolver = func(params graphql.ResolveParams) (interface{}, error) {
	JobTenderType := &structs.JobTenders{}
	JobTenderData, JobTenderDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tenders.json", JobTenderType)

	var id int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	if JobTenderDataErr != nil {
		fmt.Printf("Fetching Job Tenders failed because of this error - %s.\n", JobTenderDataErr)
	}

	// Populate data for each Job Tender with Organization Unit and Job Position
	var items = PopulateJobTenderItemProperties(JobTenderData, id, 0, "")

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   items,
	}, nil
}

var JobTenderInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.JobTenders
	dataBytes, _ := json.Marshal(params.Args["data"])
	JobTenderType := &structs.JobTenders{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	jobTenderData, jobTenderDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tenders.json", JobTenderType)

	if jobTenderDataErr != nil {
		fmt.Printf("Fetching Job Tenders failed because of this error - %s.\n", jobTenderDataErr)
	}

	sliceData := []interface{}{data}

	// Populate data for each Job Tender with Organization Unit and Job Position
	var populatedData = PopulateJobTenderItemProperties(sliceData, itemId, 0, "")

	if shared.IsInteger(itemId) && itemId != 0 {
		jobTenderData = shared.FilterByProperty(jobTenderData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(jobTenderData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_tenders.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var JobTenderDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	JobTenderType := &structs.JobTenders{}
	jobTenderData, jobTenderDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tenders.json", JobTenderType)

	if jobTenderDataErr != nil {
		fmt.Printf("Fetching Job Tender failed because of this error - %s.\n", jobTenderDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		jobTenderData = shared.FilterByProperty(jobTenderData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_tenders.json"), jobTenderData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

var JobTenderApplicationsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}
	var jobTenderId int
	if params.Args["job_tender_id"] == nil {
		jobTenderId = 0
	} else {
		jobTenderId = params.Args["job_tender_id"].(int)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	JobTenderApplicationsType := &structs.JobTenderApplications{}
	JobTenderApplicationsData, JobTenderApplicationsDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tender_applications.json", JobTenderApplicationsType)

	if JobTenderApplicationsDataErr != nil {
		fmt.Printf("Fetching Job Tenders failed because of this error - %s.\n", JobTenderApplicationsDataErr)
	}

	// Populate data for each Job Tender with Organization Unit and Job Position
	items = PopulateJobTenderApplicationProperties(JobTenderApplicationsData, id, jobTenderId)

	total = len(items)

	// Filtering by Pagination params
	if shared.IsInteger(page) && page != 0 && shared.IsInteger(size) && size != 0 {
		items = shared.Pagination(items, page.(int), size.(int))
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"total":   total,
		"items":   items,
	}, nil
}

var JobTenderApplicationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.JobTenderApplications
	dataBytes, _ := json.Marshal(params.Args["data"])
	JobTenderApplicationType := &structs.JobTenderApplications{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	JobTenderApplicationData, JobTenderApplicationDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tender_applications.json", JobTenderApplicationType)

	if JobTenderApplicationDataErr != nil {
		fmt.Printf("Fetching Job Tender Applications failed because of this error - %s.\n", JobTenderApplicationDataErr)
	}

	sliceData := []interface{}{data}

	// Populate data for each Job Tender with Organization Unit and Job Position
	var populatedData = PopulateJobTenderApplicationProperties(sliceData, itemId, 0)

	if shared.IsInteger(itemId) && itemId != 0 {
		JobTenderApplicationData = shared.FilterByProperty(JobTenderApplicationData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(JobTenderApplicationData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_tender_applications.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var JobTenderApplicationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	JobTenderApplicationType := &structs.JobTenderApplications{}
	JobTenderApplicationData, JobTenderApplicationDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_tender_applications.json", JobTenderApplicationType)

	if JobTenderApplicationDataErr != nil {
		fmt.Printf("Fetching Job Tender Applications failed because of this error - %s.\n", JobTenderApplicationDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		JobTenderApplicationData = shared.FilterByProperty(JobTenderApplicationData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_tender_applications.json"), JobTenderApplicationData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
