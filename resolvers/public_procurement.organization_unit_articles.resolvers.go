package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

func buildProcurementOUArticleResponseItemList(context context.Context, articles []*structs.PublicProcurementOrganizationUnitArticle) (items []*dto.ProcurementOrganizationUnitArticleResponseItem, err error) {
	for _, article := range articles {
		resItem, err := buildProcurementOUArticleResponseItem(context, article)
		if err != nil {
			return nil, err
		}
		items = append(items, resItem)
	}
	return
}

var PublicProcurementOrganizationUnitArticlesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var total int

	page := params.Args["page"]
	size := params.Args["size"]
	procurementID := params.Args["procurement_id"]

	input := dto.GetProcurementOrganizationUnitArticleListInputDTO{}

	if organizationUnitID, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitID != 0 {
		input.OrganizationUnitID = &organizationUnitID
	}

	articles, err := getProcurementOUArticleList(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	items, err := buildProcurementOUArticleResponseItemList(params.Context, articles)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	filteredItems := make([]*dto.ProcurementOrganizationUnitArticleResponseItem, 0)

	for _, item := range items {
		if procurementID != nil && procurementID.(int) != 0 &&
			procurementID.(int) != item.Article.PublicProcurement.Id {
			continue
		}
		filteredItems = append(filteredItems, item)
	}
	total = len(filteredItems)

	if page != nil && page.(int) > 0 && size != nil && size.(int) > 0 {
		paginatedItems, err := shared.Paginate(filteredItems, page.(int), size.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		} else {
			filteredItems = paginatedItems.([]*dto.ProcurementOrganizationUnitArticleResponseItem)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   filteredItems,
		Total:   total,
	}, nil
}

var PublicProcurementOrganizationUnitArticleInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementOrganizationUnitArticle
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateProcurementOUArticle(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementOUArticleResponseItem(params.Context, res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := createProcurementOUArticle(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementOUArticleResponseItem(params.Context, res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

var PublicProcurementSendPlanOnRevisionResolver = func(params graphql.ResolveParams) (interface{}, error) {
	plan_id := params.Args["plan_id"].(int)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return shared.HandleAPIError(fmt.Errorf("manager has no organization unit assigned"))
	}

	ouArticleList, err := getOrganizationUnitArticles(plan_id, *organizationUnitID)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, ouArticle := range ouArticleList {
		ouArticle.Status = structs.ArticleStatusRevision
		_, err = updateProcurementOUArticle(ouArticle.Id, ouArticle)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}, nil
}

var PublicProcurementOrganizationUnitArticlesDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	organizationUnitId := params.Args["organization_unit_id"]

	var procurementID *int
	if params.Args["procurement_id"] != nil {
		procurementIDParam := params.Args["procurement_id"].(int)
		procurementID = &procurementIDParam
	}

	var planId *int
	if params.Args["plan_id"] != nil {
		planIDParam := params.Args["plan_id"].(int)
		planId = &planIDParam
	}

	if organizationUnitId == nil {
		organizationUnitID, _ := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		organizationUnitId = *organizationUnitID
	}

	response, err := buildProcurementOUArticleDetailsResponseItem(params.Context, planId, organizationUnitId.(int), procurementID)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

func buildProcurementOUArticleDetailsResponseItem(context context.Context, planID *int, unitID int, procurementID *int) ([]*dto.ProcurementItemWithOrganizationUnitArticleResponseItem, error) {
	var responseItemList []*dto.ProcurementItemWithOrganizationUnitArticleResponseItem

	var items []*structs.PublicProcurementItem
	var plan *structs.PublicProcurementPlan
	var err error

	if procurementID != nil {
		item, err := getProcurementItem(*procurementID)
		if err != nil {
			return nil, err
		}
		plan, _ = getProcurementPlan(item.PlanId)
		items = append(items, item)
	} else {
		plan, err = getProcurementPlan(*planID)
		if err != nil {
			return nil, err
		}
		items, err = getProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &plan.Id})
		if err != nil {
			return nil, err
		}
	}
	planStatus, err := BuildStatus(context, plan)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		status := getProcurementStatus(*item, *plan, planStatus, &unitID)
		responseItem := dto.ProcurementItemWithOrganizationUnitArticleResponseItem{
			Id: item.Id,
			Plan: dto.DropdownSimple{
				Id:    plan.Id,
				Title: plan.Title,
			},
			IsOpenProcurement: item.IsOpenProcurement,
			Title:             item.Title,
			ArticleType:       item.ArticleType,
			Status:            status,
			SerialNumber:      item.SerialNumber,
			DateOfPublishing:  (*string)(item.DateOfPublishing),
			DateOfAwarding:    (*string)(item.DateOfAwarding),
			FileID:            item.FileId,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
		}
		organizationUnitArticleList, err := getOrganizationUnitArticles(plan.Id, unitID)
		if err != nil {
			return nil, err
		}
		for _, ouArticle := range organizationUnitArticleList {
			resItem, err := buildProcurementOUArticleResponseItem(context, ouArticle)
			if err != nil {
				return nil, err
			}
			if procurementID != nil && resItem.Article.PublicProcurement.Id != *procurementID {
				continue
			}
			responseItem.Articles = append(responseItem.Articles, resItem)
		}
		responseItemList = append(responseItemList, &responseItem)
	}

	return responseItemList, nil
}

func buildProcurementOUArticleResponseItem(context context.Context, item *structs.PublicProcurementOrganizationUnitArticle) (*dto.ProcurementOrganizationUnitArticleResponseItem, error) {
	article, err := getProcurementArticle(item.PublicProcurementArticleId)
	if err != nil {
		return nil, err
	}
	articleResItem, err := buildProcurementArticleResponseItem(context, article, nil)
	if err != nil {
		return nil, err
	}

	organizationUnit, err := getOrganizationUnitById(item.OrganizationUnitId)
	if err != nil {
		return nil, err
	}
	organizationUnitDropdown := dto.DropdownSimple{
		Id:    organizationUnit.Id,
		Title: organizationUnit.Title,
	}

	res := dto.ProcurementOrganizationUnitArticleResponseItem{
		Id:                  item.Id,
		Article:             *articleResItem,
		OrganizationUnit:    organizationUnitDropdown,
		Amount:              item.Amount,
		Status:              item.Status,
		IsRejected:          item.IsRejected,
		RejectedDescription: item.RejectedDescription,
		CreatedAt:           item.CreatedAt,
		UpdatedAt:           item.UpdatedAt,
	}

	return &res, nil
}

func getOrganizationUnitArticles(planId int, unitId int) ([]*structs.PublicProcurementOrganizationUnitArticle, error) {
	var ouArticleList []*structs.PublicProcurementOrganizationUnitArticle

	organizationUnit, err := getOrganizationUnitById(unitId)
	if err != nil {
		return nil, err
	}
	procurements, err := getProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &planId})
	if err != nil {
		return nil, err
	}

	for _, procurement := range procurements {
		relatedArticles, err := getProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &procurement.Id})
		if err != nil {
			return nil, err
		}

		for _, article := range relatedArticles {
			ouArticles, err := getProcurementOUArticleList(
				&dto.GetProcurementOrganizationUnitArticleListInputDTO{
					ArticleID:          &article.Id,
					OrganizationUnitID: &organizationUnit.Id,
				},
			)
			if err != nil {
				return nil, err
			}
			ouArticleList = append(ouArticleList, ouArticles...)
		}
	}

	return ouArticleList, nil
}

func createProcurementOUArticle(article *structs.PublicProcurementOrganizationUnitArticle) (*structs.PublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ORGANIZATION_UNIT_ARTICLE_ENDPOINT, article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementOUArticle(id int, article *structs.PublicProcurementOrganizationUnitArticle) (*structs.PublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ORGANIZATION_UNIT_ARTICLE_ENDPOINT+"/"+strconv.Itoa(id), article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getProcurementOUArticleList(input *dto.GetProcurementOrganizationUnitArticleListInputDTO) ([]*structs.PublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ORGANIZATION_UNIT_ARTICLE_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
