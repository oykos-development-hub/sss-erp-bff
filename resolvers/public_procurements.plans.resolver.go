package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

var PublicProcurementPlansOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	items := []*dto.ProcurementPlanResponseItem{}
	var total int

	page := params.Args["page"].(int)
	size := params.Args["size"].(int)
	status := params.Args["status"]
	year := params.Args["year"]
	isPreBudget := params.Args["is_pre_budget"]
	hasContract := params.Args["contract"]

	input := dto.GetProcurementPlansInput{}

	if year != nil && year.(string) != "" {
		yearValue := year.(string)
		input.Year = &yearValue
	}
	if isPreBudget != nil {
		isPreBudgetValue := isPreBudget.(bool)
		input.IsPreBudget = &isPreBudgetValue
	}

	plans, err := getProcurementPlanList(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, plan := range plans {
		var contract *bool = nil

		if hasContract != nil {
			pomContract := hasContract.(bool)
			contract = &pomContract
		}

		resItem, err := buildProcurementPlanResponseItem(params.Context, plan, contract)

		if err != nil {
			return shared.HandleAPIError(err)
		}
		if resItem == nil {
			fmt.Printf("user does not have access to this plan id: %d", plan.Id)
			continue
		}
		if status != nil && dto.PlanStatus(status.(string)) != resItem.Status {
			continue
		}
		items = append(items, resItem)
	}
	total = len(items)

	paginatedItems, err := shared.Paginate(items, page, size)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   paginatedItems,
		Total:   total,
	}, nil
}

var PublicProcurementPlanDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	plan, err := getProcurementPlan(id)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	resItem, err := buildProcurementPlanResponseItem(params.Context, plan, nil)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	if resItem == nil {
		return shared.HandleAPIError(fmt.Errorf("user does not have access to this plan id: %d", plan.Id))
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    resItem,
	}, nil
}

var PublicProcurementPlanInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementPlan
	response := dto.ResponseSingle{
		Status: "success",
	}

	// @TODO: adjust logic of activating plans
	data.Active = true

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateProcurementPlan(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementPlanResponseItem(params.Context, res, nil)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := createProcurementPlan(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementPlanResponseItem(params.Context, res, nil)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func buildProcurementPlanResponseItem(context context.Context, plan *structs.PublicProcurementPlan, hasContract *bool) (*dto.ProcurementPlanResponseItem, error) {
	items := []*dto.ProcurementItemResponseItem{}
	rawItems, err := getProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &plan.Id})
	if err != nil {
		return nil, err
	}

	contYes := true
	contNo := false

	for _, item := range rawItems {

		if hasContract != nil && *hasContract == contYes {
			filter := dto.GetProcurementContractsInput{
				ProcurementID: &item.Id,
			}

			contract, err := getProcurementContractsList(&filter)
			if err != nil {
				return nil, err
			}

			if len(contract.Data) == 0 {
				continue
			}

		} else if hasContract != nil && *hasContract == contNo {
			filter := dto.GetProcurementContractsInput{
				ProcurementID: &item.Id,
			}

			contract, err := getProcurementContractsList(&filter)
			if err != nil {
				return nil, err
			}

			if len(contract.Data) != 0 {
				continue
			}
		}

		resItem, err := buildProcurementItemResponseItem(context, item, nil)
		if err != nil {
			return nil, err
		}
		items = append(items, resItem)
	}

	status, err := BuildStatus(context, plan)
	if err != nil {
		return nil, err
	}

	if status == dto.PlanStatusNotAccessible {
		return nil, nil
	}

	res := dto.ProcurementPlanResponseItem{
		Id:               plan.Id,
		IsPreBudget:      plan.IsPreBudget,
		Active:           plan.Active,
		Year:             plan.Year,
		Title:            plan.Title,
		Status:           status,
		SerialNumber:     plan.SerialNumber,
		DateOfPublishing: plan.DateOfPublishing,
		DateOfClosing:    plan.DateOfClosing,
		PreBudgetId:      plan.PreBudgetId,
		FileId:           plan.FileId,
		CreatedAt:        plan.CreatedAt,
		UpdatedAt:        plan.UpdatedAt,
		Items:            items,
	}

	if plan.PreBudgetId != nil {
		plan, err := getProcurementPlan(*plan.PreBudgetId)
		if err != nil {
			return nil, err
		}
		res.PreBudgetPlan = &dto.DropdownSimple{
			Id:    plan.Id,
			Title: plan.Title,
		}
	}

	uniqueOrganizationUnits := make(map[int]bool)

	organizationUnitID, _ := context.Value(config.OrganizationUnitIDKey).(*int) // assert the type

	for _, procurement := range res.Items {
		if len(procurement.Articles) > 0 {
			firstArticle := procurement.Articles[0]

			oUArticles, err := getProcurementOUArticleList(&dto.GetProcurementOrganizationUnitArticleListInputDTO{ArticleID: &firstArticle.Id})
			if err != nil {
				return nil, err
			}

			for _, ouArticle := range oUArticles {
				if _, recorded := uniqueOrganizationUnits[ouArticle.OrganizationUnitId]; !recorded {
					uniqueOrganizationUnits[ouArticle.OrganizationUnitId] = true
					updateRejectedDescriptionIfNeeded(organizationUnitID, ouArticle, &res)
				}
			}
		}
	}
	if status == dto.PlanStatusAdminPublished {
		res.Requests = len(uniqueOrganizationUnits)
	}

	return &res, nil
}

func updateRejectedDescriptionIfNeeded(organizationUnitID *int, ouArticle *structs.PublicProcurementOrganizationUnitArticle, res *dto.ProcurementPlanResponseItem) {
	if organizationUnitID != nil && *organizationUnitID == ouArticle.OrganizationUnitId && ouArticle.IsRejected {
		res.RejectedDescription = &ouArticle.RejectedDescription
	}
}

func BuildStatus(context context.Context, plan *structs.PublicProcurementPlan) (dto.PlanStatus, error) {
	loggedInAccount, _ := context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	organizationUnitID, _ := context.Value(config.OrganizationUnitIDKey).(*int)

	var isPublished = plan.DateOfPublishing != nil && *plan.DateOfPublishing != ""
	var isClosed = plan.DateOfClosing != nil && *plan.DateOfClosing != ""
	var isPreBudget = plan.IsPreBudget
	var isConverted = false
	var isSentOnRevision = false
	var isRejected = false
	var isAccepted = false

	conversionTargetPlans, err := getProcurementPlanList(&dto.GetProcurementPlansInput{TargetBudgetID: &plan.Id})
	if err != nil {
		return "", err
	}

	if len(conversionTargetPlans) > 0 {
		isConverted = true
	}

	isAdmin := loggedInAccount.RoleId == 1 || loggedInAccount.RoleId == 3

	if !isAdmin {
		if organizationUnitID == nil {
			return "", errors.New("manager has no organization unit assigned")
		}

		ouArticleList, err := getOrganizationUnitArticles(plan.Id, *organizationUnitID)
		if err != nil {
			return "", err
		}

		isSentOnRevision, isRejected, isAccepted = checkArticlesStatusFlags(ouArticleList)
	}

	if isAdmin {
		if isPublished {
			if isClosed {
				if isPreBudget {
					if isConverted {
						// Admin converted a closed pre-budget Plan into a new post-budget Plan
						return dto.PlanStatusPreBudgetConverted, nil
					}
					// Admin closed a pre-budget Plan that can't be edited any further
					return dto.PlanStatusPreBudgetClosed, nil
				}
				// Admin closed a post-budget Plan that can't be edited any further
				return dto.PlanStatusPostBudgetClosed, nil
			}
			// Admin published a Plan that can be seen by Users in Organization units
			return dto.PlanStatusAdminPublished, nil
		}
		// Draft version of the Plan before it has been published
		return dto.PlanStatusAdminInProggress, nil
	} else {
		if isPublished {
			if isClosed {
				if isPreBudget {
					if isConverted {
						// Admin converted a closed pre-budget Plan into a new post-budget Plan
						return dto.PlanStatusPreBudgetConverted, nil
					}
					// Admin closed a pre-budget Plan that can't be edited any further
					return dto.PlanStatusPreBudgetClosed, nil
				}
				// Admin closed a post-budget Plan that can't be edited any further
				return dto.PlanStatusPostBudgetClosed, nil
			}
			if isSentOnRevision {
				if isRejected {
					return dto.PlanStatusUserRejected, nil
				}
				if isAccepted {
					return dto.PlanStatusUserAccepted, nil
				}

				return dto.PlanStatusUserRequested, nil
			}
			// Users in Organization units can see Plan and request Articles after it has been published
			return dto.PlanStatusUserPublished, nil
		} else {
			// Not accessible for Users in Organization units before Plan has been published
			return dto.PlanStatusNotAccessible, nil
		}
	}
}

func checkArticlesStatusFlags(articles []*structs.PublicProcurementOrganizationUnitArticle) (isSentOnRevision, isRejected, isAccepted bool) {
	var revisionCount, acceptedCount int

	for _, article := range articles {
		switch article.Status {
		case structs.ArticleStatusRejected:
			return false, true, false
		case structs.ArticleStatusRevision:
			revisionCount++
		case structs.ArticleStatusAccepted:
			acceptedCount++
		}
	}

	isSentOnRevision = revisionCount > 0 && revisionCount == len(articles)
	isAccepted = acceptedCount > 0 && acceptedCount == len(articles)

	return isSentOnRevision, false, isAccepted
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

func getProcurementPlanList(input *dto.GetProcurementPlansInput) ([]*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.PLANS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
