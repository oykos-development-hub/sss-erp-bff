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
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) PublicProcurementPlansOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	items := []*dto.ProcurementPlanResponseItem{}
	var total int

	page := params.Args["page"].(int)
	size := params.Args["size"].(int)
	status := params.Args["status"]
	year := params.Args["year"]
	isPreBudget := params.Args["is_pre_budget"]
	hasContract := params.Args["contract"]
	sortByTitle := params.Args["sort_by_title"]
	sortByYear := params.Args["sort_by_year"]
	sortByDateOfPublishing := params.Args["sort_by_date_of_publishing"]

	input := dto.GetProcurementPlansInput{}

	if year != nil && year.(string) != "" {
		yearValue := year.(string)
		input.Year = &yearValue
	}
	if isPreBudget != nil {
		isPreBudgetValue := isPreBudget.(bool)
		input.IsPreBudget = &isPreBudgetValue
	}
	if sortByYear != nil && sortByYear.(string) != "" {
		value := sortByYear.(string)
		input.SortByYear = &value
	}
	if sortByTitle != nil && sortByTitle.(string) != "" {
		value := sortByTitle.(string)
		input.SortByTitle = &value
	}
	if sortByDateOfPublishing != nil && sortByDateOfPublishing.(string) != "" {
		value := sortByDateOfPublishing.(string)
		input.SortByDateOfPublishing = &value
	}

	plans, err := r.Repo.GetProcurementPlanList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, plan := range plans {
		var contract *bool

		if hasContract != nil {
			pomContract := hasContract.(bool)
			contract = &pomContract
		}

		resItem, err := buildProcurementPlanResponseItem(params.Context, r.Repo, plan, contract, &dto.GetProcurementItemListInputMS{})

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		if resItem == nil {
			//fmt.Printf("user does not have access to this plan id: %d", plan.ID)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   paginatedItems,
		Total:   total,
	}, nil
}

func (r *Resolver) PublicProcurementPlanDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	plan, err := r.Repo.GetProcurementPlan(id)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	sortByTitle := params.Args["sort_by_title"]
	sortBySerialNumber := params.Args["sort_by_serial_number"]
	sortByDateOfPublishing := params.Args["sort_by_date_of_publishing"]
	sortByDateOfAwarding := params.Args["sort_by_date_of_awarding"]

	input := dto.GetProcurementItemListInputMS{PlanID: &plan.ID}

	if sortByTitle != nil && sortByTitle.(string) != "" {
		value := sortByTitle.(string)
		input.SortByTitle = &value
	}

	if sortBySerialNumber != nil && sortBySerialNumber.(string) != "" {
		value := sortBySerialNumber.(string)
		input.SortBySerialNumber = &value
	}

	if sortByDateOfAwarding != nil && sortByDateOfAwarding.(string) != "" {
		value := sortByDateOfAwarding.(string)
		input.SortByDateOfAwarding = &value
	}

	if sortByDateOfPublishing != nil && sortByDateOfPublishing.(string) != "" {
		value := sortByDateOfPublishing.(string)
		input.SortByDateOfPublishing = &value
	}

	resItem, err := buildProcurementPlanResponseItem(params.Context, r.Repo, plan, nil, &input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	if resItem == nil {
		return errors.HandleAPPError(fmt.Errorf("user does not have access to this plan id: %d", plan.ID))
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    resItem,
	}, nil
}

func (r *Resolver) PublicProcurementPlanInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementPlan
	response := dto.ResponseSingle{
		Status: "success",
	}

	// @TODO: adjust logic of activating plans
	data.Active = true

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		oldPlan, err := r.Repo.GetProcurementPlan(itemID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if oldPlan.DateOfPublishing == nil && data.DateOfPublishing != nil {
			loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
			targetUsers, err := r.Repo.GetUsersByPermission(config.PublicProcurementPlan)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			plan, _ := r.Repo.GetProcurementPlan(data.ID)
			planData := dto.ProcurementPlanNotification{
				ID:          plan.ID,
				IsPreBudget: plan.IsPreBudget,
				Year:        plan.Year,
			}
			planDataJSON, _ := json.Marshal(planData)

			for _, targetUser := range targetUsers {
				if targetUser.ID != loggedInUser.ID {
					_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
						Content:     "Proslijeđen je novi plan javnih nabavki na pregled i popunjavanje.",
						Module:      "Javne nabavke",
						FromUserID:  loggedInUser.ID,
						ToUserID:    targetUser.ID,
						FromContent: "Službenik za javne nabavke",
						Path:        fmt.Sprintf("/procurements/plans/%d", data.ID),
						Data:        planDataJSON,
						IsRead:      false,
					})
					if err != nil {
						_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
						return errors.HandleAPPError(err)
					}
				}
			}
		}

		res, err := r.Repo.UpdateProcurementPlan(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildProcurementPlanResponseItem(params.Context, r.Repo, res, nil, &dto.GetProcurementItemListInputMS{})
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := r.Repo.CreateProcurementPlan(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildProcurementPlanResponseItem(params.Context, r.Repo, res, nil, &dto.GetProcurementItemListInputMS{})
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func buildProcurementPlanResponseItem(context context.Context, r repository.MicroserviceRepositoryInterface, plan *structs.PublicProcurementPlan, hasContract *bool, filter *dto.GetProcurementItemListInputMS) (*dto.ProcurementPlanResponseItem, error) {
	items := []*dto.ProcurementItemResponseItem{}
	filter.PlanID = &plan.ID
	rawItems, err := r.GetProcurementItemList(filter)
	if err != nil {
		return nil, errors.Wrap(err, "repo get procurement item list")
	}

	contYes := true
	contNo := false

	var totalNet float32
	var totalGross float32

	for _, item := range rawItems {

		if hasContract != nil && *hasContract == contYes {
			filter := dto.GetProcurementContractsInput{
				ProcurementID: &item.ID,
			}

			contract, err := r.GetProcurementContractsList(&filter)
			if err != nil {
				return nil, errors.Wrap(err, "repo get procurement contracts list")
			}

			if len(contract.Data) == 0 {
				continue
			}

		} else if hasContract != nil && *hasContract == contNo {
			filter := dto.GetProcurementContractsInput{
				ProcurementID: &item.ID,
			}

			contract, err := r.GetProcurementContractsList(&filter)
			if err != nil {
				return nil, errors.Wrap(err, "repo get procurement contracts list")
			}

			if len(contract.Data) != 0 {
				continue
			}
		}

		resItem, err := buildProcurementItemResponseItem(context, r, item, nil, &dto.GetProcurementArticleListInputMS{})
		if err != nil {
			return nil, errors.Wrap(err, "build procurement item response item")
		}
		totalNet += resItem.TotalNet
		totalGross += resItem.TotalGross

		items = append(items, resItem)
	}

	status, err := BuildStatus(context, r, plan)
	if err != nil {
		return nil, errors.Wrap(err, "build status")
	}

	if status == dto.PlanStatusNotAccessible {
		return nil, nil
	}

	res := dto.ProcurementPlanResponseItem{
		ID:               plan.ID,
		IsPreBudget:      plan.IsPreBudget,
		Active:           plan.Active,
		Year:             plan.Year,
		Title:            plan.Title,
		Status:           status,
		SerialNumber:     plan.SerialNumber,
		DateOfPublishing: plan.DateOfPublishing,
		DateOfClosing:    plan.DateOfClosing,
		PreBudgetID:      plan.PreBudgetID,
		FileID:           plan.FileID,
		TotalNet:         totalNet,
		TotalGross:       totalGross,
		CreatedAt:        plan.CreatedAt,
		UpdatedAt:        plan.UpdatedAt,
		Items:            items,
	}

	if plan.PreBudgetID != nil {
		plan, err := r.GetProcurementPlan(*plan.PreBudgetID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get procurement plan")
		}
		res.PreBudgetPlan = &dto.DropdownSimple{
			ID:    plan.ID,
			Title: plan.Title,
		}
	}

	uniqueOrganizationUnits := make(map[int]bool)
	approvedRequestsCount := 0

	organizationUnitID, _ := context.Value(config.OrganizationUnitIDKey).(*int) // assert the type

	for _, procurement := range res.Items {
		if len(procurement.Articles) > 0 {
			firstArticle := procurement.Articles[0]

			oUArticles, err := r.GetProcurementOUArticleList(&dto.GetProcurementOrganizationUnitArticleListInputDTO{ArticleID: &firstArticle.ID})
			if err != nil {
				return nil, errors.Wrap(err, "repo get procurement ou article list")
			}

			for _, ouArticle := range oUArticles {
				if _, recorded := uniqueOrganizationUnits[ouArticle.OrganizationUnitID]; !recorded {
					if ouArticle.Status == structs.ArticleStatusAccepted {
						approvedRequestsCount++
					}
					uniqueOrganizationUnits[ouArticle.OrganizationUnitID] = true
					updateRejectedDescriptionIfNeeded(organizationUnitID, ouArticle, &res)
				}
			}
		}
	}
	if status == dto.PlanStatusAdminPublished {
		res.Requests = len(uniqueOrganizationUnits)
	}

	res.ApprovedRequests = approvedRequestsCount

	return &res, nil
}

func updateRejectedDescriptionIfNeeded(organizationUnitID *int, ouArticle *structs.PublicProcurementOrganizationUnitArticle, res *dto.ProcurementPlanResponseItem) {
	if organizationUnitID != nil && *organizationUnitID == ouArticle.OrganizationUnitID && ouArticle.IsRejected {
		res.RejectedDescription = &ouArticle.RejectedDescription
	}
}

func BuildStatus(context context.Context, r repository.MicroserviceRepositoryInterface, plan *structs.PublicProcurementPlan) (dto.PlanStatus, error) {
	loggedInAccount, _ := context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	organizationUnitID, _ := context.Value(config.OrganizationUnitIDKey).(*int)

	var isPublished = plan.DateOfPublishing != nil && *plan.DateOfPublishing != ""
	var isClosed = plan.DateOfClosing != nil && *plan.DateOfClosing != ""
	var isPreBudget = plan.IsPreBudget
	var isConverted = false
	var isSentOnRevision = false
	var isRejected = false
	var isAccepted = false

	conversionTargetPlans, err := r.GetProcurementPlanList(&dto.GetProcurementPlansInput{TargetBudgetID: &plan.ID})
	if err != nil {
		return "", errors.Wrap(err, "repo get procurement plan list")
	}

	if len(conversionTargetPlans) > 0 {
		isConverted = true
	}

	isAdmin := loggedInAccount.RoleID != nil && (*loggedInAccount.RoleID == 1 || *loggedInAccount.RoleID == 3)

	if !isAdmin {
		if organizationUnitID == nil {
			return "", fmt.Errorf("manager has no organization unit assigned")
		}

		ouArticleList, err := GetOrganizationUnitArticles(r, plan.ID, *organizationUnitID)
		if err != nil {
			return "", errors.Wrap(err, "repo get organization unit articles")
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
	}
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
	}
	// Not accessible for Users in Organization units before Plan has been published
	return dto.PlanStatusNotAccessible, nil
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

func (r *Resolver) PublicProcurementPlanDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteProcurementPlan(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) PublicProcurementPlanPDFResolver(params graphql.ResolveParams) (interface{}, error) {
	planID := params.Args["plan_id"].(int)
	plan, err := r.Repo.GetProcurementPlan(planID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	planResItem, _ := buildProcurementPlanResponseItem(params.Context, r.Repo, plan, nil, &dto.GetProcurementItemListInputMS{})

	dateCurrentLayout := "2006-01-02T15:04:05Z"
	dateOutputLayout := "02.01.2006. 15:04"

	t, err := time.Parse(dateCurrentLayout, *planResItem.DateOfPublishing)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var planPDFResponse dto.PlanPDFResponse

	planPDFResponse.TotalGross = dto.FormatToEuro(planResItem.TotalGross)
	planPDFResponse.TotalVAT = dto.FormatToEuro(planResItem.TotalGross - planResItem.TotalNet)
	planPDFResponse.PublishedDate = t.Format(dateOutputLayout)
	planPDFResponse.Year = planResItem.Year
	planPDFResponse.PlanID = strconv.Itoa(planResItem.ID)

	var tableData []dto.PlanPDFTableDataRow

	for index, procurement := range planResItem.Items {
		rowData := dto.PlanPDFTableDataRow{
			ID:              fmt.Sprintf("%d", index+1),
			ArticleType:     procurement.ArticleType,
			Title:           procurement.Title,
			TotalGross:      dto.FormatToEuro(procurement.TotalGross),
			TotalVAT:        dto.FormatToEuro(procurement.TotalGross - procurement.TotalNet),
			TypeOfProcedure: procurement.TypeOfProcedure,
			BudgetIndent:    procurement.BudgetIndent.SerialNumber,
			FundingSource:   "Budžet CG",
		}
		tableData = append(tableData, rowData)
	}

	planPDFResponse.TableData = tableData

	// Return the path or a URL to the file
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's is your PDF file in base64 encode format!",
		Item:    planPDFResponse,
	}, nil
}
