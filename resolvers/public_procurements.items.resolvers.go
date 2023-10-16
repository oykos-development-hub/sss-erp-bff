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

var PublicProcurementPlanItemDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	items := []*dto.ProcurementItemResponseItem{}
	var authToken = params.Context.Value(config.TokenKey).(string)
	loggedInAccount, err := getLoggedInUser(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	if id != nil && id.(int) > 0 {
		item, err := getProcurementItem(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}

		resItem, _ := buildProcurementItemResponseItem(item, *loggedInAccount)
		items = append(items, resItem)
	} else {
		procurements, err := getProcurementItemList(nil)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, item := range procurements {
			resItem, _ := buildProcurementItemResponseItem(item, *loggedInAccount)
			items = append(items, resItem)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   len(items),
	}, nil
}

var PublicProcurementPlanItemInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	response := dto.ResponseSingle{
		Status: "success",
	}
	var authToken = params.Context.Value(config.TokenKey).(string)
	loggedInAccount, err := getLoggedInUser(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	var data structs.PublicProcurementItem

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateProcurementItem(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resItem, _ := buildProcurementItemResponseItem(res, *loggedInAccount)

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		res, err := createProcurementItem(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		resItem, _ := buildProcurementItemResponseItem(res, *loggedInAccount)

		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil

}

var PublicProcurementPlanItemDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteProcurementItem(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func createProcurementItem(item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ITEMS_ENDPOINT, item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementItem(id int, item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ITEMS_ENDPOINT+"/"+strconv.Itoa(id), item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteProcurementItem(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.ITEMS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getProcurementItem(id int) (*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ITEMS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getProcurementItemList(input *dto.GetProcurementItemListInputMS) ([]*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ITEMS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getProcurementStatus(procurementID int, organizationUnitID *int) (*structs.ProcurementStatus, error) {
	procurementStatus := structs.ProcurementStatusInProgress

	if organizationUnitID != nil {

		articles, err := getProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &procurementID})
		if err != nil {
			return nil, err
		}

		filledArticles, _ := getProcurementOUArticleList(
			&dto.GetProcurementOrganizationUnitArticleListInputDTO{
				OrganizationUnitID: organizationUnitID,
			},
		)
		var matchedArticleCount int
		for _, ouArticle := range filledArticles {
			article, _ := getProcurementArticle(ouArticle.PublicProcurementArticleId)
			if article.PublicProcurementId == procurementID {
				matchedArticleCount++
			}
		}

		if matchedArticleCount >= len(articles) {
			procurementStatus = structs.ProcurementStatusProcessed
		}
	}

	return &procurementStatus, nil
}

func buildProcurementItemResponseItem(item *structs.PublicProcurementItem, loggedInAccount structs.UserAccounts) (*dto.ProcurementItemResponseItem, error) {
	plan, _ := getProcurementPlan(item.PlanId)
	planDropdown := dto.DropdownSimple{Id: plan.Id, Title: plan.Title}

	var articles []*dto.ProcurementArticleResponseItem
	articlesRaw, err := getProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &item.Id})
	if err != nil {
		return nil, err
	}
	for _, article := range articlesRaw {
		articleResItem, err := buildProcurementArticleResponseItem(article)
		if err != nil {
			return nil, err
		}
		articles = append(articles, articleResItem)
	}

	userProfile, _ := getUserProfileByUserAccountID(loggedInAccount.Id)
	organizationUnitID, _ := getOrganizationUnitIdByUserProfile(userProfile.Id)

	planStatus, err := BuildStatus(plan, loggedInAccount)
	if err != nil {
		return nil, err
	}

	var procurementStatus structs.ProcurementStatus

	if planStatus == string(structs.PostProcurementStatusCompleted) {
		procurementStatus = structs.PostProcurementStatusCompleted
	} else if planStatus == string(structs.PreProcurementStatusCompleted) {
		procurementStatus = structs.PostProcurementStatusCompleted
	} else {
		status, err := getProcurementStatus(item.Id, organizationUnitID)
		if err != nil {
			return nil, err
		}
		procurementStatus = *status
	}

	account, err := getAccountItemById(item.BudgetIndentId)
	if err != nil {
		return nil, err
	}

	res := dto.ProcurementItemResponseItem{
		Id:    item.Id,
		Title: item.Title,
		BudgetIndent: dto.DropdownBudgetIndent{
			Id:           account.Id,
			Title:        account.Title,
			SerialNumber: account.SerialNumber,
		},
		Plan:              planDropdown,
		IsOpenProcurement: item.IsOpenProcurement,
		ArticleType:       item.ArticleType,
		Status:            procurementStatus,
		SerialNumber:      item.SerialNumber,
		DateOfAwarding:    item.DateOfAwarding,
		DateOfPublishing:  item.DateOfPublishing,
		FileId:            item.FileId,
		Articles:          articles,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	return &res, nil
}
