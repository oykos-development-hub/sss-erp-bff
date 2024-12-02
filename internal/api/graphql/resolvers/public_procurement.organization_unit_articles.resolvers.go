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
	"strconv"

	goerror "errors"

	"github.com/graphql-go/graphql"
)

func buildProcurementOUArticleResponseItemList(context context.Context, r *Resolver, articles []*structs.PublicProcurementOrganizationUnitArticle) (items []*dto.ProcurementOrganizationUnitArticleResponseItem, err error) {
	for _, article := range articles {
		resItem, err := buildProcurementOUArticleResponseItem(context, r, article)
		if err != nil {
			return nil, errors.Wrap(err, "build procurement ou article response item")
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	items, err := buildProcurementOUArticleResponseItemList(params.Context, r, articles)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	article, _ := r.Repo.GetProcurementArticle(data.PublicProcurementArticleID)
	procurement, _ := r.Repo.GetProcurementItem(article.PublicProcurementID)

	itemID := data.ID

	if itemID != 0 {
		oldRequest, err := r.Repo.GetOrganizationUnitArticleByID(itemID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if oldRequest.Status == string(structs.ArticleStatusAccepted) && oldRequest.Amount != data.Amount {
			return errors.HandleAPPError(goerror.New("you request has been accepted already"))
		}

		var notificationContent string

		oldArticle, err := r.Repo.GetProcurementArticle(oldRequest.PublicProcurementArticleID)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		same := true
		var firstStatus string

		publicProcurement, err := r.Repo.GetProcurementItem(oldArticle.PublicProcurementID)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		publicProcurements, err := r.Repo.GetProcurementItemList(&dto.GetProcurementItemListInputMS{
			PlanID: &publicProcurement.PlanID,
		})

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, item := range publicProcurements {

			articles, err := r.Repo.GetProcurementArticlesList(
				&dto.GetProcurementArticleListInputMS{
					ItemID: &item.ID})

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			for _, articleItem := range articles {

				articlesOU, err := r.Repo.GetOrganizationUnitArticlesList(dto.GetProcurementOrganizationUnitArticleListInputDTO{
					OrganizationUnitID: &oldRequest.OrganizationUnitID,
					ArticleID:          &articleItem.ID,
				})

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}

				for _, itemOU := range articlesOU {
					fmt.Println(strconv.Itoa(data.ID) + " " + firstStatus + " " + strconv.Itoa(itemOU.ID) + " " + itemOU.Status)
					if firstStatus == "" {
						firstStatus = itemOU.Status
					} else if firstStatus != itemOU.Status {
						same = false
					}
				}
			}
		}

		fmt.Println(same)

		if same {
			if oldRequest.Status != string(structs.ArticleStatusRejected) && data.IsRejected {
				notificationContent = "Vaš zahtjev je odbijen. Molimo Vas da pregledate komentar i ponovno pošaljete plan."
			} else if oldRequest.Status != string(structs.ArticleStatusAccepted) && !data.IsRejected {
				notificationContent = "Vaš zahtjev za plan je odobren."
			}

			if notificationContent != "" {
				loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
				employees, _ := GetEmployeesOfOrganizationUnit(r.Repo, data.OrganizationUnitID)
				/*if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}*/

				targetUsers, _ := r.Repo.GetUsersByPermission(config.PublicProcurementPlan, config.OperationRead)
				/*if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}*/

				for _, employee := range employees {
					employeeAccount, err := r.Repo.GetUserAccountByID(employee.UserAccountID)
					if err != nil {
						_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
						return errors.HandleAPPError(err)
					}

					for _, user := range targetUsers {
						if user.ID == employee.UserAccountID && loggedInUser.ID != user.ID {
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
								_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
								return errors.HandleAPPError(err)
							}
						}
					}
				}
			}
		}

		res, err := r.Repo.UpdateProcurementOUArticle(itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildProcurementOUArticleResponseItem(params.Context, r, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := r.Repo.CreateProcurementOUArticle(&data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildProcurementOUArticleResponseItem(params.Context, r, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
		return errors.HandleAPPError(fmt.Errorf("manager has no organization unit assigned"))
	}

	ouArticleList, err := GetOrganizationUnitArticles(r.Repo, planID, *organizationUnitID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	isRejected := false

	for _, ouArticle := range ouArticleList {
		if ouArticle.Status == structs.ArticleStatusRejected {
			isRejected = true
		}
		ouArticle.Status = structs.ArticleStatusRevision
		_, err = r.Repo.UpdateProcurementOUArticle(ouArticle.ID, ouArticle)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

	targetUsers, _ := r.Repo.GetUsersByPermission(config.PublicProcurementPlan, config.OperationFullAccess)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	for _, targetUser := range targetUsers {
		if loggedInUser.ID != targetUser.ID {
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
				Content:    content,
				Module:     "Javne nabavke",
				FromUserID: loggedInUser.ID,
				Path:       fmt.Sprintf("/procurements/plans/%d?tab=requests", planID),
				ToUserID:   targetUser.ID,
				//FromContent: fmt.Sprintf("Menadžer %s", unit.Abbreviation),
				IsRead: false,
				Data:   dataJSON,
			})
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
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

	response, err := buildProcurementOUArticleDetailsResponseItem(params.Context, r, planID, organizationUnitID.(int), procurementID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

func buildProcurementOUArticleDetailsResponseItem(context context.Context, r *Resolver, planID *int, unitID int, procurementID *int) ([]*dto.ProcurementItemWithOrganizationUnitArticleResponseItem, error) {
	var responseItemList []*dto.ProcurementItemWithOrganizationUnitArticleResponseItem

	var items []*structs.PublicProcurementItem
	var plan *structs.PublicProcurementPlan
	var err error

	if procurementID != nil {
		item, err := r.Repo.GetProcurementItem(*procurementID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get procurement itme")
		}
		plan, _ = r.Repo.GetProcurementPlan(item.PlanID)
		items = append(items, item)
	} else {
		plan, err = r.Repo.GetProcurementPlan(*planID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get procurement plan")
		}
		items, err = r.Repo.GetProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &plan.ID})
		if err != nil {
			return nil, errors.Wrap(err, "repo get procurement item list")
		}
	}
	planStatus, err := BuildStatus(context, r, plan)
	if err != nil {
		return nil, errors.Wrap(err, "build status")
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
		organizationUnitArticleList, err := GetOrganizationUnitArticles(r.Repo, plan.ID, unitID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit articles")
		}
		for _, ouArticle := range organizationUnitArticleList {
			resItem, err := buildProcurementOUArticleResponseItem(context, r, ouArticle)
			if err != nil {
				return nil, errors.Wrap(err, "build procurement ou article response item")
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

func buildProcurementOUArticleResponseItem(context context.Context, r *Resolver, item *structs.PublicProcurementOrganizationUnitArticle) (*dto.ProcurementOrganizationUnitArticleResponseItem, error) {
	article, err := r.Repo.GetProcurementArticle(item.PublicProcurementArticleID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get procurement article")
	}
	articleResItem, err := buildProcurementArticleResponseItem(context, r, article, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build procurement article response item")
	}

	organizationUnit, _ := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)
	/*if err != nil {
		return nil, errors.Wrap(err, "repo get organization unit by id")
	}*/

	var organizationUnitDropdown dto.DropdownSimple

	if organizationUnit != nil {
		organizationUnitDropdown = dto.DropdownSimple{
			ID:    organizationUnit.ID,
			Title: organizationUnit.Title,
		}
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
		return nil, errors.Wrap(err, "repo get organization unit by id")
	}
	procurements, err := r.GetProcurementItemList(&dto.GetProcurementItemListInputMS{PlanID: &planID})
	if err != nil {
		return nil, errors.Wrap(err, "repo get procurement item list")
	}

	for _, procurement := range procurements {
		relatedArticles, err := r.GetProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &procurement.ID})
		if err != nil {
			return nil, errors.Wrap(err, "repo get procurement articles list")
		}

		for _, article := range relatedArticles {
			ouArticles, err := r.GetProcurementOUArticleList(
				&dto.GetProcurementOrganizationUnitArticleListInputDTO{
					ArticleID:          &article.ID,
					OrganizationUnitID: &organizationUnit.ID,
				},
			)
			if err != nil {
				return nil, errors.Wrap(err, "repo get procurement ou article list")
			}
			ouArticleList = append(ouArticleList, ouArticles...)
		}
	}

	return ouArticleList, nil
}
