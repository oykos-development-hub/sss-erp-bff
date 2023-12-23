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

func buildProcurementOUArticleResponseItemList(context context.Context, r repository.MicroserviceRepositoryInterface, articles []*structs.PublicProcurementOrganizationUnitArticle) (items []*dto.ProcurementOrganizationUnitArticleResponseItem, err error) {
	for _, article := range articles {
		resItem, err := buildProcurementOUArticleResponseItem(context, r, article)
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

	items, err := buildProcurementOUArticleResponseItemList(params.Context, r.Repo, articles)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	filteredItems := make([]*dto.ProcurementOrganizationUnitArticleResponseItem, 0)

	for _, item := range items {
		if procurementID != nil && procurementID.(int) != 0 &&
			procurementID.(int) != item.Article.PublicProcurement.ID {
			continue
		}
		filteredItems = append(filteredItems, item)
	}
	total = len(filteredItems)

	if page != nil && page.(int) > 0 && size != nil && size.(int) > 0 {
		paginatedItems, err := shared.Paginate(filteredItems, page.(int), size.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}
		filteredItems = paginatedItems.([]*dto.ProcurementOrganizationUnitArticleResponseItem)
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

	article, _ := r.Repo.GetProcurementArticle(data.PublicProcurementArticleID)
	procurement, _ := r.Repo.GetProcurementItem(article.PublicProcurementID)

	itemID := data.ID

	if itemID != 0 {
		oldRequest, err := r.Repo.GetOrganizationUnitArticleByID(itemID)
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
			employees, err := GetEmployeesOfOrganizationUnit(r.Repo, data.OrganizationUnitID)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			for _, employee := range employees {
				employeeAccount, err := r.Repo.GetUserAccountByID(employee.UserAccountID)
				if err != nil {
					return errors.HandleAPIError(err)
				}
				if employeeAccount.RoleID == structs.UserRoleManagerOJ {
					plan, _ := r.Repo.GetProcurementPlan(procurement.PlanID)
					data := dto.ProcurementPlanNotification{
						ID:          plan.ID,
						IsPreBudget: plan.IsPreBudget,
						Year:        plan.Year,
					}
					dataJSON, _ := json.Marshal(data)
					_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
						Content:     notificationContent,
						Module:      "Javne nabavke",
						FromUserID:  loggedInUser.ID,
						ToUserID:    employeeAccount.ID,
						FromContent: "Službenik za javne nabavke",
						IsRead:      false,
						Data:        dataJSON,
						Path:        fmt.Sprintf("/procurements/plans/%d", procurement.PlanID),
					})
					if err != nil {
						return errors.HandleAPIError(err)
					}
				}
			}
		}

		res, err := r.Repo.UpdateProcurementOUArticle(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		item, err := buildProcurementOUArticleResponseItem(params.Context, r.Repo, res)
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
		item, err := buildProcurementOUArticleResponseItem(params.Context, r.Repo, res)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) PublicProcurementSendPlanOnRevisionResolver(params graphql.ResolveParams) (interface{}, error) {
	planID := params.Args["plan_id"].(int)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return errors.HandleAPIError(fmt.Errorf("manager has no organization unit assigned"))
	}

	ouArticleList, err := GetOrganizationUnitArticles(r.Repo, planID, *organizationUnitID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	isRejected := false

	for _, ouArticle := range ouArticleList {
		if ouArticle.Status == structs.ArticleStatusRejected {
			isRejected = true
		}
		ouArticle.Status = structs.ArticleStatusRevision
		_, err = r.Repo.UpdateProcurementOUArticle(ouArticle.ID, ouArticle)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	unitID := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	unit, err := r.Repo.GetOrganizationUnitByID(*unitID)
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
		plan, _ := r.Repo.GetProcurementPlan(planID)
		data := dto.ProcurementPlanNotification{
			ID:          plan.ID,
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
			FromUserID:  loggedInUser.ID,
			Path:        fmt.Sprintf("/procurements/plans/%d?tab=requests", planID),
			ToUserID:    targetUser.ID,
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
	organizationUnitID := params.Args["organization_unit_id"]

	var procurementID *int
	if params.Args["procurement_id"] != nil {
		procurementIDParam := params.Args["procurement_id"].(int)
		procurementID = &procurementIDParam
	}

	var planID *int
	if params.Args["plan_id"] != nil {
		planIDParam := params.Args["plan_id"].(int)
		planID = &planIDParam
	}

	if organizationUnitID == nil {
		organizationUnitIDFromContext, _ := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		organizationUnitID = *organizationUnitIDFromContext
	}

	response, err := buildProcurementOUArticleDetailsResponseItem(params.Context, r.Repo, planID, organizationUnitID.(int), procurementID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

func buildProcurementOUArticleDetailsResponseItem(context context.Context, r repository.MicroserviceRepositoryInterface, planID *int, unitID int, procurementID *int) ([]*dto.ProcurementItemWithOrganizationUnitArticleResponseItem, error) {
	var responseItemList []*dto.ProcurementItemWithOrganizationUnitArticleResponseItem

	var items []*structs.PublicProcurementItem
	var plan *structs.PublicProcurementPlan
	var err error

	if procurementID != nil {
		item, err := r.GetProcurementItem(*procurementID)
		if err != nil {
			return nil, err
		}
		plan, _ = r.GetProcurementPlan(item.PlanID)
		items = append(items, item)
	} else {
		plan, err = r.GetProcurementPlan(*planID)
		if err != nil {
			return nil, err
		}
		items, err = r.GetProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &plan.ID})
		if err != nil {
			return nil, err
		}
	}
	planStatus, err := BuildStatus(context, r, plan)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		status := getProcurementStatus(r, *item, *plan, planStatus, &unitID)
		responseItem := dto.ProcurementItemWithOrganizationUnitArticleResponseItem{
			ID: item.ID,
			Plan: dto.DropdownSimple{
				ID:    plan.ID,
				Title: plan.Title,
			},
			IsOpenProcurement: item.IsOpenProcurement,
			Title:             item.Title,
			ArticleType:       item.ArticleType,
			Status:            status,
			SerialNumber:      item.SerialNumber,
			DateOfPublishing:  (*string)(item.DateOfPublishing),
			DateOfAwarding:    (*string)(item.DateOfAwarding),
			FileID:            item.FileID,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
		}
		organizationUnitArticleList, err := GetOrganizationUnitArticles(r, plan.ID, unitID)
		if err != nil {
			return nil, err
		}
		for _, ouArticle := range organizationUnitArticleList {
			resItem, err := buildProcurementOUArticleResponseItem(context, r, ouArticle)
			if err != nil {
				return nil, err
			}
			if resItem.Article.PublicProcurement.ID != item.ID {
				continue
			}
			responseItem.Articles = append(responseItem.Articles, resItem)
		}
		responseItemList = append(responseItemList, &responseItem)
	}

	return responseItemList, nil
}

func buildProcurementOUArticleResponseItem(context context.Context, r repository.MicroserviceRepositoryInterface, item *structs.PublicProcurementOrganizationUnitArticle) (*dto.ProcurementOrganizationUnitArticleResponseItem, error) {
	article, err := r.GetProcurementArticle(item.PublicProcurementArticleID)
	if err != nil {
		return nil, err
	}
	articleResItem, err := buildProcurementArticleResponseItem(context, r, article, nil)
	if err != nil {
		return nil, err
	}

	organizationUnit, err := r.GetOrganizationUnitByID(item.OrganizationUnitID)
	if err != nil {
		return nil, err
	}
	organizationUnitDropdown := dto.DropdownSimple{
		ID:    organizationUnit.ID,
		Title: organizationUnit.Title,
	}

	res := dto.ProcurementOrganizationUnitArticleResponseItem{
		ID:                  item.ID,
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

func GetOrganizationUnitArticles(r repository.MicroserviceRepositoryInterface, planID int, unitID int) ([]*structs.PublicProcurementOrganizationUnitArticle, error) {
	var ouArticleList []*structs.PublicProcurementOrganizationUnitArticle

	organizationUnit, err := r.GetOrganizationUnitByID(unitID)
	if err != nil {
		return nil, err
	}
	procurements, err := r.GetProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &planID})
	if err != nil {
		return nil, err
	}

	for _, procurement := range procurements {
		relatedArticles, err := r.GetProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &procurement.ID})
		if err != nil {
			return nil, err
		}

		for _, article := range relatedArticles {
			ouArticles, err := r.GetProcurementOUArticleList(
				&dto.GetProcurementOrganizationUnitArticleListInputDTO{
					ArticleID:          &article.ID,
					OrganizationUnitID: &organizationUnit.ID,
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
