package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func buildProcurementOUArticleResponseItemList(r repository.MicroserviceRepositoryInterface, context context.Context, articles []*structs.PublicProcurementOrganizationUnitArticle) (items []*dto.ProcurementOrganizationUnitArticleResponseItem, err error) {
	for _, article := range articles {
		resItem, err := buildProcurementOUArticleResponseItem(r, context, article)
		if err != nil {
			return nil, err
		}
		items = append(items, resItem)
	}
	return
}

func (r *Resolver) PublicProcurementOrganizationUnitArticlesOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var total int

	page := params.Args["page"]
	size := params.Args["size"]
	procurementID := params.Args["procurement_id"]

	input := dto.GetProcurementOrganizationUnitArticleListInputDTO{}

	if organizationUnitID, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitID != 0 {
		input.OrganizationUnitID = &organizationUnitID
	}

	articles, err := r.Repo.GetProcurementOUArticleList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	items, err := buildProcurementOUArticleResponseItemList(r.Repo, params.Context, articles)
	if err != nil {
		return errors.HandleAPIError(err)
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
			return errors.HandleAPIError(err)
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

func (r *Resolver) PublicProcurementOrganizationUnitArticleInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementOrganizationUnitArticle
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	article, _ := r.Repo.GetProcurementArticle(data.PublicProcurementArticleId)
	procurement, _ := r.Repo.GetProcurementItem(article.PublicProcurementId)

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		oldRequest, err := r.Repo.GetOrganizationUnitArticleByID(itemId)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		var notificationContent string

		if oldRequest.Status != string(structs.ArticleStatusRejected) && data.IsRejected {
			notificationContent = "Vaš zahtjev je odbijen. Molimo Vas da pregledate komentar i ponovno pošaljete plan."
		} else if oldRequest.Status != string(structs.ArticleStatusAccepted) && !data.IsRejected {
			notificationContent = "Vaš zahtjev za plan je odobren."
		}

		if notificationContent != "" {
			loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
			employees, err := GetEmployeesOfOrganizationUnit(r.Repo, data.OrganizationUnitId)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			for _, employee := range employees {
				employeeAccount, err := r.Repo.GetUserAccountById(employee.UserAccountId)
				if err != nil {
					return errors.HandleAPIError(err)
				}
				if employeeAccount.RoleId == structs.UserRoleManagerOJ {
					plan, _ := r.Repo.GetProcurementPlan(procurement.PlanId)
					data := dto.ProcurementPlanNotification{
						ID:          plan.Id,
						IsPreBudget: plan.IsPreBudget,
						Year:        plan.Year,
					}
					dataJSON, _ := json.Marshal(data)
					_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
						Content:     notificationContent,
						Module:      "Javne nabavke",
						FromUserID:  loggedInUser.Id,
						ToUserID:    employeeAccount.Id,
						FromContent: "Službenik za javne nabavke",
						IsRead:      false,
						Data:        dataJSON,
						Path:        fmt.Sprintf("/procurements/plans/%d", procurement.PlanId),
					})
					if err != nil {
						return errors.HandleAPIError(err)
					}
				}
			}
		}

		res, err := r.Repo.UpdateProcurementOUArticle(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		item, err := buildProcurementOUArticleResponseItem(r.Repo, params.Context, res)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := r.Repo.CreateProcurementOUArticle(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		item, err := buildProcurementOUArticleResponseItem(r.Repo, params.Context, res)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) PublicProcurementSendPlanOnRevisionResolver(params graphql.ResolveParams) (interface{}, error) {
	plan_id := params.Args["plan_id"].(int)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return errors.HandleAPIError(fmt.Errorf("manager has no organization unit assigned"))
	}

	ouArticleList, err := GetOrganizationUnitArticles(r.Repo, plan_id, *organizationUnitID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	isRejected := false

	for _, ouArticle := range ouArticleList {
		if ouArticle.Status == structs.ArticleStatusRejected {
			isRejected = true
		}
		ouArticle.Status = structs.ArticleStatusRevision
		_, err = r.Repo.UpdateProcurementOUArticle(ouArticle.Id, ouArticle)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	unitID := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	unit, err := r.Repo.GetOrganizationUnitById(*unitID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	oficialOfProcurementsRole := structs.UserRoleOfficialForPublicProcurements
	targetUsers, err := r.Repo.GetUserAccounts(&dto.GetUserAccountListInput{
		RoleID: &oficialOfProcurementsRole,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, targetUser := range targetUsers.Data {
		plan, _ := r.Repo.GetProcurementPlan(plan_id)
		data := dto.ProcurementPlanNotification{
			ID:          plan.Id,
			IsPreBudget: plan.IsPreBudget,
			Year:        plan.Year,
		}
		dataJSON, _ := json.Marshal(data)

		var content string
		if isRejected {
			content = "Zahtjev sa novim izmjenama je ažuriran i proslijeđen."
		} else {
			content = "Zahtjev je proslijeđen."
		}
		_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
			Content:     content,
			Module:      "Javne nabavke",
			FromUserID:  loggedInUser.Id,
			Path:        fmt.Sprintf("/procurements/plans/%d?tab=requests", plan_id),
			ToUserID:    targetUser.Id,
			FromContent: fmt.Sprintf("Menadžer %s", unit.Abbreviation),
			IsRead:      false,
			Data:        dataJSON,
		})
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}, nil
}

func (r *Resolver) PublicProcurementOrganizationUnitArticlesDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
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

	response, err := buildProcurementOUArticleDetailsResponseItem(r.Repo, params.Context, planId, organizationUnitId.(int), procurementID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

func buildProcurementOUArticleDetailsResponseItem(r repository.MicroserviceRepositoryInterface, context context.Context, planID *int, unitID int, procurementID *int) ([]*dto.ProcurementItemWithOrganizationUnitArticleResponseItem, error) {
	var responseItemList []*dto.ProcurementItemWithOrganizationUnitArticleResponseItem

	var items []*structs.PublicProcurementItem
	var plan *structs.PublicProcurementPlan
	var err error

	if procurementID != nil {
		item, err := r.GetProcurementItem(*procurementID)
		if err != nil {
			return nil, err
		}
		plan, _ = r.GetProcurementPlan(item.PlanId)
		items = append(items, item)
	} else {
		plan, err = r.GetProcurementPlan(*planID)
		if err != nil {
			return nil, err
		}
		items, err = r.GetProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &plan.Id})
		if err != nil {
			return nil, err
		}
	}
	planStatus, err := BuildStatus(r, context, plan)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		status := getProcurementStatus(r, *item, *plan, planStatus, &unitID)
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
		organizationUnitArticleList, err := GetOrganizationUnitArticles(r, plan.Id, unitID)
		if err != nil {
			return nil, err
		}
		for _, ouArticle := range organizationUnitArticleList {
			resItem, err := buildProcurementOUArticleResponseItem(r, context, ouArticle)
			if err != nil {
				return nil, err
			}
			if resItem.Article.PublicProcurement.Id != item.Id {
				continue
			}
			responseItem.Articles = append(responseItem.Articles, resItem)
		}
		responseItemList = append(responseItemList, &responseItem)
	}

	return responseItemList, nil
}

func buildProcurementOUArticleResponseItem(r repository.MicroserviceRepositoryInterface, context context.Context, item *structs.PublicProcurementOrganizationUnitArticle) (*dto.ProcurementOrganizationUnitArticleResponseItem, error) {
	article, err := r.GetProcurementArticle(item.PublicProcurementArticleId)
	if err != nil {
		return nil, err
	}
	articleResItem, err := buildProcurementArticleResponseItem(r, context, article, nil)
	if err != nil {
		return nil, err
	}

	organizationUnit, err := r.GetOrganizationUnitById(item.OrganizationUnitId)
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

func GetOrganizationUnitArticles(r repository.MicroserviceRepositoryInterface, planId int, unitId int) ([]*structs.PublicProcurementOrganizationUnitArticle, error) {
	var ouArticleList []*structs.PublicProcurementOrganizationUnitArticle

	organizationUnit, err := r.GetOrganizationUnitById(unitId)
	if err != nil {
		return nil, err
	}
	procurements, err := r.GetProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &planId})
	if err != nil {
		return nil, err
	}

	for _, procurement := range procurements {
		relatedArticles, err := r.GetProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &procurement.Id})
		if err != nil {
			return nil, err
		}

		for _, article := range relatedArticles {
			ouArticles, err := r.GetProcurementOUArticleList(
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
