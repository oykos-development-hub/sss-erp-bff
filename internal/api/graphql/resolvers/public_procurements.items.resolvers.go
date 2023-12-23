package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
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

		resItem, _ := buildProcurementItemResponseItem(params.Context, r.Repo, item, nil, filter)
		items = append(items, resItem)
	} else {
		procurements, err := r.Repo.GetProcurementItemList(nil)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		for _, item := range procurements {
			resItem, _ := buildProcurementItemResponseItem(params.Context, r.Repo, item, nil, &dto.GetProcurementArticleListInputMS{})
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
		organizationUnit, _ := r.Repo.GetOrganizationUnitByID(*organizationUnitID)
		organizationUnitTitle = organizationUnit.Title
	}

	item, err := r.Repo.GetProcurementItem(id)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	resItem, _ := buildProcurementItemResponseItem(ctx, r.Repo, item, nil, &dto.GetProcurementArticleListInputMS{})

	if resItem.Status != structs.PostProcurementStatusContracted {
		return errors.HandleAPIError(fmt.Errorf("procurement must be contracted"))
	}

	contract, _ := r.Repo.GetProcurementContract(*resItem.ContractID)
	contractRes, _ := buildProcurementContractResponseItem(r.Repo, contract)

	articles, err := GetProcurementArticles(ctx, r.Repo, resItem.ID)
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
		articleRes, err := ProcessOrderArticleItem(r.Repo, article, *organizationUnitID)

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

	itemID := data.ID

	if itemID != 0 {
		res, err := r.Repo.UpdateProcurementItem(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resItem, _ := buildProcurementItemResponseItem(params.Context, r.Repo, res, nil, &dto.GetProcurementArticleListInputMS{})

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		res, err := r.Repo.CreateProcurementItem(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resItem, _ := buildProcurementItemResponseItem(params.Context, r.Repo, res, nil, &dto.GetProcurementArticleListInputMS{})

		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil

}

func (r *Resolver) PublicProcurementPlanItemDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteProcurementItem(itemID)
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
		article, _ := r.GetProcurementArticle(ouArticle.PublicProcurementArticleID)
		if article.PublicProcurementID == procurementID {
			matchedArticleCount++
		}
	}

	return matchedArticleCount >= len(articles)
}

func buildProcurementItemResponseItem(context context.Context, r repository.MicroserviceRepositoryInterface, item *structs.PublicProcurementItem, organizationUnitID *int, filter *dto.GetProcurementArticleListInputMS) (*dto.ProcurementItemResponseItem, error) {
	if organizationUnitID == nil {
		organizationUnitID, _ = context.Value(config.OrganizationUnitIDKey).(*int)
	}

	plan, _ := r.GetProcurementPlan(item.PlanID)
	planDropdown := dto.DropdownSimple{ID: plan.ID, Title: plan.Title}
	var totalGross float32
	var totalNet float32

	var articles []*dto.ProcurementArticleResponseItem
	filter.ItemID = &item.ID
	articlesRaw, err := r.GetProcurementArticlesList(filter)
	if err != nil {
		return nil, err
	}
	for _, article := range articlesRaw {
		articleResItem, err := buildProcurementArticleResponseItem(context, r, article, organizationUnitID)
		if err != nil {
			return nil, err
		}
		totalGross += articleResItem.GrossPrice * float32(articleResItem.TotalAmount)
		totalNet += articleResItem.NetPrice * float32(articleResItem.TotalAmount)
		articles = append(articles, articleResItem)
	}

	planStatus, err := BuildStatus(context, r, plan)
	if err != nil {
		return nil, err
	}

	procurementStatus := getProcurementStatus(r, *item, *plan, planStatus, organizationUnitID)

	account, err := r.GetAccountItemByID(item.BudgetIndentID)
	if err != nil {
		return nil, err
	}

	var contractID *int

	if procurementStatus == structs.PostProcurementStatusContracted {
		contracts, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &item.ID})
		if err != nil {
			return nil, err
		}
		contractID = &contracts.Data[0].ID
	}

	typeOfProcedure := "Otvoreni postupak"
	if !item.IsOpenProcurement {
		typeOfProcedure = "Jednostavna nabavka"
	}

	res := dto.ProcurementItemResponseItem{
		ID:    item.ID,
		Title: item.Title,
		BudgetIndent: dto.DropdownBudgetIndent{
			ID:           account.ID,
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
		FileID:            item.FileID,
		Articles:          articles,
		ContractID:        contractID,
		TotalGross:        totalGross,
		TotalNet:          totalNet,
		TypeOfProcedure:   typeOfProcedure,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	return &res, nil
}

func getProcurementStatus(r repository.MicroserviceRepositoryInterface, item structs.PublicProcurementItem, plan structs.PublicProcurementPlan, planStatus dto.PlanStatus, organizationUnitID *int) structs.ProcurementStatus {
	if !plan.IsPreBudget && isContracted(r, item.ID) {
		return structs.PostProcurementStatusContracted
	} else if planStatus == dto.PlanStatusPostBudgetClosed {
		return structs.PostProcurementStatusCompleted
	} else if planStatus == dto.PlanStatusPreBudgetClosed {
		return structs.PreProcurementStatusCompleted
	} else if isProcurementProcessed(r, item.ID, organizationUnitID) {
		return structs.ProcurementStatusProcessed
	}
	return structs.ProcurementStatusInProgress
}

func isContracted(r repository.MicroserviceRepositoryInterface, procurementID int) bool {
	contracts, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &procurementID})
	if err != nil {
		return false
	}
	return contracts != nil && len(contracts.Data) > 0
}
