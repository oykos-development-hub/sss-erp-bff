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

var PublicProcurementPlansOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var authToken = params.Context.Value("token").(string)
	var total int
	var year string
	if params.Args["year"] == nil {
		year = ""
	} else {
		year = params.Args["year"].(string)
	}
	var status string
	if params.Args["status"] == nil {
		status = ""
	} else {
		status = params.Args["status"].(string)
	}
	var isPreBudget = params.Args["is_pre_budget"]
	page := params.Args["page"]
	size := params.Args["size"]

	var items []interface{}
	var plans = shared.FetchByProperty(
		"public_procurement_plan",
		"",
		"",
	)

	if len(plans) > 0 {
		items = PopulatePlanItemProperties(plans, isPreBudget, year, status, authToken)
	}

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

var PublicProcurementPlanDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	plan, err := getProcurementPlan(id)
	if err != nil {
		return dto.Response{
			Status:  "error",
			Message: err.Error(),
		}, nil
	}
	resItem, _ := buildProcurementPlanResponseItem(plan)

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    *resItem,
	}, nil
}

var PublicProcurementPlanInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementPlan
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating procurement plan data"), nil
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateProcurementPlan(itemId, &data)
		if err != nil {
			fmt.Printf("Updating procurement plan failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error updating procurement plan type data"), nil
		}
		item, err := buildProcurementPlanResponseItem(res)
		if err != nil {
			fmt.Printf("Building procurement plan response failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error building response of procurement plan data"), nil
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := createProcurementPlan(&data)
		if err != nil {
			fmt.Printf("Creating procurement plan failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating procurement plan data"), nil
		}
		item, err := buildProcurementPlanResponseItem(res)
		if err != nil {
			fmt.Printf("Building procurement plan response failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error building response of procurement plan data"), nil
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func buildProcurementPlanResponseItem(item *structs.PublicProcurementPlan) (*dto.ProcurementPlanResponseItem, error) {
	status := "TODO"
	var preBudgetPlan structs.SettingsDropdown
	if item.PreBudgetId != nil {
		plan, err := getProcurementPlan(*item.PreBudgetId)
		if err != nil {
			return nil, err
		}
		preBudgetPlan = structs.SettingsDropdown{
			Id:    plan.Id,
			Title: plan.Title,
		}
	}

	res := dto.ProcurementPlanResponseItem{
		Id:               item.Id,
		IsPreBudget:      item.IsPreBudget,
		Active:           item.Active,
		Year:             item.Year,
		Title:            item.Title,
		Status:           &status,
		SerialNumber:     item.SerialNumber,
		DateOfPublishing: (*string)(item.DateOfPublishing),
		DateOfClosing:    (*string)(item.DateOfClosing),
		PreBudgetId:      item.PreBudgetId,
		FileId:           item.FileId,
		PreBudgetPlan:    preBudgetPlan,
		CreatedAt:        item.CreatedAt,
		UpdatedAt:        item.UpdatedAt,
	}

	return &res, nil
}

var PublicProcurementPlanDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteProcurementPlan(itemId)
	if err != nil {
		fmt.Printf("Deleting procurement plan failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func createProcurementPlan(resolution *structs.PublicProcurementPlan) (*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.PLANS_ENDPOINT, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementPlan(id int, resolution *structs.PublicProcurementPlan) (*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.PLANS_ENDPOINT+"/"+strconv.Itoa(id), resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteProcurementPlan(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.PLANS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getProcurementPlan(id int) (*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.PLANS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getProcurementPlanList(input *dto.GetJobTendersInput) ([]*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.PLANS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
