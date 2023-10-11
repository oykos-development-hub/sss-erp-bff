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

var PublicProcurementOrganizationUnitArticlesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.ProcurementOrganizationUnitArticleResponseItem{}
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
	total = len(articles)

	for _, article := range articles {
		resItem, err := buildProcurementOUArticleResponseItem(article)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		if procurementID != nil && procurementID.(int) != 0 &&
			procurementID.(int) != resItem.Article.PublicProcurement.Id {
			total--
			continue
		}
		items = append(items, *resItem)
	}

	if page != nil && page.(int) > 0 && size != nil && size.(int) > 0 {
		paginatedItems, err := shared.Paginate(items, page.(int), size.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		} else {
			items = paginatedItems.([]dto.ProcurementOrganizationUnitArticleResponseItem)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
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
		item, err := buildProcurementOUArticleResponseItem(res)
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
		item, err := buildProcurementOUArticleResponseItem(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

var PublicProcurementSendPlanOnRevisionResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var authToken = params.Context.Value(config.TokenKey).(string)
	loggedInProfile, err := getLoggedInUserProfile(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	plan_id := params.Args["plan_id"].(int)

	organizationUnitId, err := getOrganizationUnitIdByUserProfile(loggedInProfile.Id)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	if organizationUnitId == nil {
		return shared.HandleAPIError(fmt.Errorf("manager with id - %d has no organization unit assigned", loggedInProfile.Id))
	}

	ouArticleList, err := getOrganizationUnitArticles(plan_id, *organizationUnitId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, ouArticle := range ouArticleList {
		ouArticle.Status = structs.StatusRevision
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
	planId := params.Args["plan_id"].(int)
	organizationUnitId := params.Args["organization_unit_id"]
	if organizationUnitId == nil {
		var authToken = params.Context.Value(config.TokenKey).(string)
		userProfile, err := getLoggedInUserProfile(authToken)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		unitId, err := getOrganizationUnitIdByUserProfile(userProfile.Id)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		organizationUnitId = *unitId
	}

	response, err := buildProcurementOUArticleDetailsResponseItem(planId, organizationUnitId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

func buildProcurementOUArticleDetailsResponseItem(planID, unitID int) ([]*dto.ProcurementItemWithOrganizationUnitArticleResponseItem, error) {
	var responseItemList []*dto.ProcurementItemWithOrganizationUnitArticleResponseItem

	plan, err := getProcurementPlan(planID)
	if err != nil {
		return nil, err
	}

	items, err := getProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &planID})
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		status, _ := getProcurementStatus(item.Id, &unitID)
		responseItem := dto.ProcurementItemWithOrganizationUnitArticleResponseItem{
			Id:           item.Id,
			BudgetIndent: dto.DropdownSimple{},
			Plan: dto.DropdownSimple{
				Id:    plan.Id,
				Title: plan.Title,
			},
			IsOpenProcurement: item.IsOpenProcurement,
			Title:             item.Title,
			ArticleType:       item.ArticleType,
			Status:            *status,
			SerialNumber:      item.SerialNumber,
			DateOfPublishing:  (*string)(item.DateOfPublishing),
			DateOfAwarding:    (*string)(item.DateOfAwarding),
			FileID:            item.FileId,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
		}
		organizationUnitArticleList, err := getOrganizationUnitArticles(planID, unitID)
		if err != nil {
			return nil, err
		}
		for _, ouArticle := range organizationUnitArticleList {
			resItem, err := buildProcurementOUArticleResponseItem(ouArticle)
			if err != nil {
				return nil, err
			}
			responseItem.Articles = append(responseItem.Articles, resItem)
		}
		responseItemList = append(responseItemList, &responseItem)
	}

	return responseItemList, nil
}

func buildProcurementOUArticleResponseItem(item *structs.PublicProcurementOrganizationUnitArticle) (*dto.ProcurementOrganizationUnitArticleResponseItem, error) {
	article, err := getProcurementArticle(item.PublicProcurementArticleId)
	if err != nil {
		return nil, err
	}
	articleResItem, err := buildProcurementArticleResponseItem(article)
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
