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

func (r *Resolver) PublicProcurementPlanItemDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	sortByTitle := params.Args["sort_by_title"]
	sortByPrice := params.Args["sort_by_price"]

	items := []*dto.ProcurementItemResponseItem{}

	if id != nil && id.(int) > 0 {
		item, err := r.Repo.GetProcurementItem(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}

		filter := &dto.GetProcurementArticleListInputMS{}

		if sortByTitle != nil && sortByTitle.(string) != "" {
			value := sortByTitle.(string)
			filter.SortByTitle = &value
		}

		if sortByPrice != nil && sortByPrice.(string) != "" {
			value := sortByPrice.(string)
			filter.SortByPrice = &value
		}

		resItem, _ := buildProcurementItemResponseItem(r.Repo, params.Context, item, nil, filter)
		items = append(items, resItem)
	} else {
		procurements, err := r.Repo.GetProcurementItemList(nil)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		for _, item := range procurements {
			resItem, _ := buildProcurementItemResponseItem(r.Repo, params.Context, item, nil, &dto.GetProcurementArticleListInputMS{})
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

func (r *Resolver) PublicProcurementPlanItemPDFResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	ctx := params.Context

	if params.Args["organization_unit_id"] != nil {
		organizationUnitID := params.Args["organization_unit_id"].(int)
		ctx = context.WithValue(ctx, config.OrganizationUnitIDKey, &organizationUnitID)
	}

	organizationUnitID, _ := ctx.Value(config.OrganizationUnitIDKey).(*int)
	organizationUnitTitle := ""
	if organizationUnitID != nil && *organizationUnitID != 0 {
		organizationUnit, _ := r.Repo.GetOrganizationUnitById(*organizationUnitID)
		organizationUnitTitle = organizationUnit.Title
	}

	item, err := r.Repo.GetProcurementItem(id)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	resItem, _ := buildProcurementItemResponseItem(r.Repo, ctx, item, nil, &dto.GetProcurementArticleListInputMS{})

	if resItem.Status != structs.PostProcurementStatusContracted {
		return errors.HandleAPIError(fmt.Errorf("procurement must be contracted"))
	}

	contract, _ := r.Repo.GetProcurementContract(*resItem.ContractID)
	contractRes, _ := buildProcurementContractResponseItem(r.Repo, contract)

	articles, err := GetProcurementArticles(r.Repo, ctx, resItem.Id)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	// Prepare subtitles
	subtitles := dto.Subtitles{
		PublicProcurement: resItem.Title,
		OrganizationUnit:  organizationUnitTitle,
		Supplier:          contractRes.Supplier.Title,
	}

	var tableData []dto.TableDataRow
	for _, article := range articles {
		articleRes, err := ProcessOrderArticleItem(r.Repo, ctx, article, *organizationUnitID)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		rowData := dto.TableDataRow{
			ProcurementItem:  article.Title,
			KeyFeatures:      article.Description,
			ContractedAmount: fmt.Sprintf("%d", articleRes.Amount),
			AvailableAmount:  fmt.Sprintf("%d", articleRes.Available),
			ConsumedAmount:   fmt.Sprintf("%d", articleRes.Amount-articleRes.Available),
		}
		tableData = append(tableData, rowData)
	}

	responseItem := dto.PdfData{
		Subtitles: subtitles,
		TableData: tableData,
	}

	// Return the path or a URL to the file
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's is your PDF file in base64 encode format!",
		Item:    responseItem,
	}, nil
}

func (r *Resolver) PublicProcurementPlanItemInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	response := dto.ResponseSingle{
		Status: "success",
	}

	var data structs.PublicProcurementItem

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := r.Repo.UpdateProcurementItem(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resItem, _ := buildProcurementItemResponseItem(r.Repo, params.Context, res, nil, &dto.GetProcurementArticleListInputMS{})

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		res, err := r.Repo.CreateProcurementItem(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resItem, _ := buildProcurementItemResponseItem(r.Repo, params.Context, res, nil, &dto.GetProcurementArticleListInputMS{})

		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil

}

func (r *Resolver) PublicProcurementPlanItemDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := r.Repo.DeleteProcurementItem(itemId)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func isProcurementProcessed(r repository.MicroserviceRepositoryInterface, procurementID int, organizationUnitID *int) bool {
	if organizationUnitID == nil {
		return false
	}
	articles, _ := r.GetProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &procurementID})

	if len(articles) == 0 {
		return false
	}

	filledArticles, _ := r.GetProcurementOUArticleList(
		&dto.GetProcurementOrganizationUnitArticleListInputDTO{
			OrganizationUnitID: organizationUnitID,
		},
	)
	var matchedArticleCount int
	for _, ouArticle := range filledArticles {
		article, _ := r.GetProcurementArticle(ouArticle.PublicProcurementArticleId)
		if article.PublicProcurementId == procurementID {
			matchedArticleCount++
		}
	}

	return matchedArticleCount >= len(articles)
}

func buildProcurementItemResponseItem(r repository.MicroserviceRepositoryInterface, context context.Context, item *structs.PublicProcurementItem, organizationUnitID *int, filter *dto.GetProcurementArticleListInputMS) (*dto.ProcurementItemResponseItem, error) {
	if organizationUnitID == nil {
		organizationUnitID, _ = context.Value(config.OrganizationUnitIDKey).(*int)
	}

	plan, _ := r.GetProcurementPlan(item.PlanId)
	planDropdown := dto.DropdownSimple{Id: plan.Id, Title: plan.Title}
	var totalGross float32
	var totalNet float32

	var articles []*dto.ProcurementArticleResponseItem
	filter.ItemID = &item.Id
	articlesRaw, err := r.GetProcurementArticlesList(filter)
	if err != nil {
		return nil, err
	}
	for _, article := range articlesRaw {
		articleResItem, err := buildProcurementArticleResponseItem(r, context, article, organizationUnitID)
		if err != nil {
			return nil, err
		}
		totalGross += articleResItem.GrossPrice * float32(articleResItem.TotalAmount)
		totalNet += articleResItem.NetPrice * float32(articleResItem.TotalAmount)
		articles = append(articles, articleResItem)
	}

	planStatus, err := BuildStatus(r, context, plan)
	if err != nil {
		return nil, err
	}

	procurementStatus := getProcurementStatus(r, *item, *plan, planStatus, organizationUnitID)

	account, err := r.GetAccountItemById(item.BudgetIndentId)
	if err != nil {
		return nil, err
	}

	var contractId *int

	if procurementStatus == structs.PostProcurementStatusContracted {
		contracts, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &item.Id})
		if err != nil {
			return nil, err
		}
		contractId = &contracts.Data[0].Id
	}

	typeOfProcedure := "Otvoreni postupak"
	if !item.IsOpenProcurement {
		typeOfProcedure = "Jednostavna nabavka"
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
		TotalGross:        totalGross,
		TotalNet:          totalNet,
		TypeOfProcedure:   typeOfProcedure,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	return &res, nil
}

func getProcurementStatus(r repository.MicroserviceRepositoryInterface, item structs.PublicProcurementItem, plan structs.PublicProcurementPlan, planStatus dto.PlanStatus, organizationUnitID *int) structs.ProcurementStatus {
	if !plan.IsPreBudget && isContracted(r, item.Id) {
		return structs.PostProcurementStatusContracted
	} else if planStatus == dto.PlanStatusPostBudgetClosed {
		return structs.PostProcurementStatusCompleted
	} else if planStatus == dto.PlanStatusPreBudgetClosed {
		return structs.PreProcurementStatusCompleted
	} else if isProcurementProcessed(r, item.Id, organizationUnitID) {
		return structs.ProcurementStatusProcessed
	} else {
		return structs.ProcurementStatusInProgress
	}
}

func isContracted(r repository.MicroserviceRepositoryInterface, procurementId int) bool {
	contracts, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &procurementId})
	if err != nil {
		return false
	}
	return contracts != nil && len(contracts.Data) > 0
}
