package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"context"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

var PublicProcurementPlanItemDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	items := []*dto.ProcurementItemResponseItem{}

	if id != nil && id.(int) > 0 {
		item, err := getProcurementItem(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}

		resItem, _ := buildProcurementItemResponseItem(params.Context, item)
		items = append(items, resItem)
	} else {
		procurements, err := getProcurementItemList(nil)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, item := range procurements {
			resItem, _ := buildProcurementItemResponseItem(params.Context, item)
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

	var data structs.PublicProcurementItem

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateProcurementItem(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resItem, _ := buildProcurementItemResponseItem(params.Context, res)

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		res, err := createProcurementItem(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		resItem, _ := buildProcurementItemResponseItem(params.Context, res)

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

func isProcurementProcessed(procurementID int, organizationUnitID *int) bool {
	if organizationUnitID == nil {
		return false
	}
	articles, _ := getProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &procurementID})

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

	return matchedArticleCount >= len(articles)
}

func buildProcurementItemResponseItem(context context.Context, item *structs.PublicProcurementItem) (*dto.ProcurementItemResponseItem, error) {
	organizationUnitID, _ := context.Value(config.OrganizationUnitIDKey).(*int) // assert the type

	plan, _ := getProcurementPlan(item.PlanId)
	planDropdown := dto.DropdownSimple{Id: plan.Id, Title: plan.Title}

	var articles []*dto.ProcurementArticleResponseItem
	articlesRaw, err := getProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &item.Id})
	if err != nil {
		return nil, err
	}
	for _, article := range articlesRaw {
		articleResItem, err := buildProcurementArticleResponseItem(context, article)
		if err != nil {
			return nil, err
		}
		articles = append(articles, articleResItem)
	}

	planStatus, err := BuildStatus(context, plan)
	if err != nil {
		return nil, err
	}

	procurementStatus := getProcurementStatus(*item, *plan, planStatus, organizationUnitID)

	account, err := getAccountItemById(item.BudgetIndentId)
	if err != nil {
		return nil, err
	}

	var contractId *int

	if procurementStatus == structs.PostProcurementStatusContracted {
		contracts, err := getProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &item.Id})
		if err != nil {
			return nil, err
		}
		contractId = &contracts.Data[0].Id
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
		ContractID:        contractId,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	return &res, nil
}

func getProcurementStatus(item structs.PublicProcurementItem, plan structs.PublicProcurementPlan, planStatus dto.PlanStatus, organizationUnitID *int) structs.ProcurementStatus {
	if !plan.IsPreBudget && isContracted(item.Id) {
		return structs.PostProcurementStatusContracted
	} else if planStatus == dto.PlanStatusPostBudgetClosed {
		return structs.PostProcurementStatusCompleted
	} else if planStatus == dto.PlanStatusPreBudgetClosed {
		return structs.PreProcurementStatusCompleted
	} else if isProcurementProcessed(item.Id, organizationUnitID) {
		return structs.ProcurementStatusProcessed
	} else {
		return structs.ProcurementStatusInProgress
	}
}

func isContracted(procurementId int) bool {
	contracts, err := getProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &procurementId})
	if err != nil {
		return false
	}
	return contracts != nil && len(contracts.Data) > 0
}
