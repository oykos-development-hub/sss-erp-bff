package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

var STATUSES = map[string]string{
	"not_accessible":       "Nedostupan",
	"admin_in_progress":    "U toku",
	"admin_published":      "Poslat",
	"user_published":       "Obradi",
	"user_requested":       "Na čekanju",
	"user_accepted":        "Odobren",
	"user_rejected":        "Odbijen",
	"pre_budget_closed":    "Zaključen",
	"pre_budget_converted": "Konvertovan",
	"post_budget_closed":   "Objavljen",
}

var PublicProcurementPlansOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var authToken = params.Context.Value(config.TokenKey).(string)

	loggedInAccount, err := getLoggedInUser(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	items := []dto.ProcurementPlanResponseItem{}
	var total int

	page := params.Args["page"].(int)
	size := params.Args["size"].(int)
	status := params.Args["status"]
	year := params.Args["year"]
	isPreBudget := params.Args["is_pre_budget"]

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
	total = len(plans)

	for _, plan := range plans {
		resItem, err := buildProcurementPlanResponseItem(plan, loggedInAccount)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		if resItem == nil {
			return shared.HandleAPIError(fmt.Errorf("user with id: %d do not have access to this plan id: %d", loggedInAccount.Id, plan.Id))
		}
		if status != nil && status.(string) != "" && status.(string) != *resItem.Status {
			total--
			continue
		}
		items = append(items, *resItem)
	}

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
	var authToken = params.Context.Value(config.TokenKey).(string)
	loggedInAccount, err := getLoggedInUser(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	id := params.Args["id"].(int)

	plan, err := getProcurementPlan(id)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	resItem, err := buildProcurementPlanResponseItem(plan, loggedInAccount)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	if resItem == nil {
		return shared.HandleAPIError(fmt.Errorf("user with id: %d do not have access to this plan id: %d", loggedInAccount.Id, plan.Id))
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    *resItem,
	}, nil
}

var PublicProcurementPlanInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var authToken = params.Context.Value(config.TokenKey).(string)
	loggedInAccount, err := getLoggedInUser(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	var data structs.PublicProcurementPlan
	response := dto.ResponseSingle{
		Status: "success",
	}

	// @TODO: adjust logic of activating plans
	data.Active = true

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateProcurementPlan(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementPlanResponseItem(res, loggedInAccount)
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
		item, err := buildProcurementPlanResponseItem(res, loggedInAccount)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func buildProcurementPlanResponseItem(plan *structs.PublicProcurementPlan, loggedInAccount *structs.UserAccounts) (*dto.ProcurementPlanResponseItem, error) {
	var preBudgetPlan structs.SettingsDropdown
	if plan.PreBudgetId != nil {
		plan, err := getProcurementPlan(*plan.PreBudgetId)
		if err != nil {
			return nil, err
		}
		preBudgetPlan = structs.SettingsDropdown{
			Id:    plan.Id,
			Title: plan.Title,
		}
	}

	items := []*dto.ProcurementItemResponseItem{}
	rawItems, err := getProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &plan.Id})
	if err != nil {
		return nil, err
	}

	for _, item := range rawItems {
		resItem, err := buildProcurementItemResponseItem(item, *loggedInAccount)
		if err != nil {
			return nil, err
		}
		items = append(items, resItem)
	}

	status, err := BuildStatus(plan, *loggedInAccount)
	if err != nil {
		return nil, err
	}

	if status == STATUSES["not_accessible"] {
		return nil, nil
	}

	res := dto.ProcurementPlanResponseItem{
		Id:               plan.Id,
		IsPreBudget:      plan.IsPreBudget,
		Active:           plan.Active,
		Year:             plan.Year,
		Title:            plan.Title,
		Status:           &status,
		SerialNumber:     plan.SerialNumber,
		DateOfPublishing: plan.DateOfPublishing,
		DateOfClosing:    plan.DateOfClosing,
		PreBudgetId:      plan.PreBudgetId,
		FileId:           plan.FileId,
		PreBudgetPlan:    preBudgetPlan,
		CreatedAt:        plan.CreatedAt,
		UpdatedAt:        plan.UpdatedAt,
		Items:            items,
	}

	return &res, nil
}

func BuildStatus(plan *structs.PublicProcurementPlan, userAccount structs.UserAccounts) (string, error) {
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

	isAdmin := userAccount.RoleId == 1 || userAccount.RoleId == 3

	if !isAdmin {
		loggedInProfile, err := getUserProfileByUserAccountID(userAccount.Id)
		if err != nil {
			return "", err
		}
		organizationUnitId, err := getOrganizationUnitIdByUserProfile(loggedInProfile.Id)
		if err != nil {
			return "", err
		}
		if organizationUnitId == nil {
			return "", errors.New("manager has no organization unit assigned")
		}

		ouArticleList, err := getOrganizationUnitArticles(plan.Id, *organizationUnitId)
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
						return STATUSES["pre_budget_converted"], nil
					}
					// Admin closed a pre-budget Plan that can't be edited any further
					return STATUSES["pre_budget_closed"], nil
				}
				// Admin closed a post-budget Plan that can't be edited any further
				return STATUSES["post_budget_closed"], nil
			}
			// Admin published a Plan that can be seen by Users in Organization units
			return STATUSES["admin_published"], nil
		}
		// Draft version of the Plan before it has been published
		return STATUSES["admin_in_progress"], nil
	} else {
		if isPublished {
			if isClosed {
				if isPreBudget {
					if isConverted {
						// Admin converted a closed pre-budget Plan into a new post-budget Plan
						return STATUSES["pre_budget_converted"], nil
					}
					// Admin closed a pre-budget Plan that can't be edited any further
					return STATUSES["pre_budget_closed"], nil
				}
				// Admin closed a post-budget Plan that can't be edited any further
				return STATUSES["post_budget_closed"], nil
			}
			if isSentOnRevision {
				if isRejected {
					return STATUSES["user_rejected"], nil
				}
				if isAccepted {
					return STATUSES["user_accepted"], nil
				}

				return STATUSES["user_requested"], nil
			}
			// Users in Organization units can see Plan and request Articles after it has been published
			return STATUSES["user_published"], nil
		} else {
			// Not accessible for Users in Organization units before Plan has been published
			return STATUSES["not_accessible"], nil
		}
	}
}

func checkArticlesStatusFlags(articles []*structs.PublicProcurementOrganizationUnitArticle) (isSentOnRevision, isRejected, isAccepted bool) {
	var revisionCount, acceptedCount int

	for _, article := range articles {
		switch article.Status {
		case structs.StatusRejected:
			return false, true, false
		case structs.StatusRevision:
			revisionCount++
		case structs.StatusAccepted:
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
