package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/pdfutils"
	"bff/shared"
	"bff/structs"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

var PublicProcurementPlanItemDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	items := []*dto.ProcurementItemResponseItem{}

	if id != nil && id.(int) > 0 {
		item, err := getProcurementItem(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}

		resItem, _ := buildProcurementItemResponseItem(params.Context, item, nil)
		items = append(items, resItem)
	} else {
		procurements, err := getProcurementItemList(nil)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, item := range procurements {
			resItem, _ := buildProcurementItemResponseItem(params.Context, item, nil)
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

var PublicProcurementPlanItemPDFResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	ctx := params.Context

	if params.Args["organization_unit_id"] != nil {
		organizationUnitID := params.Args["organization_unit_id"].(int)
		ctx = context.WithValue(ctx, config.OrganizationUnitIDKey, &organizationUnitID)
	}

	organizationUnitID, _ := ctx.Value(config.OrganizationUnitIDKey).(*int)
	organizationUnitTitle := ""
	if organizationUnitID != nil && *organizationUnitID != 0 {
		organizationUnit, _ := getOrganizationUnitById(*organizationUnitID)
		organizationUnitTitle = organizationUnit.Title
	}

	item, err := getProcurementItem(id)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	resItem, _ := buildProcurementItemResponseItem(ctx, item, nil)

	if resItem.Status != structs.PostProcurementStatusContracted {
		return shared.HandleAPIError(errors.New("procurement must be contracted"))
	}

	contract, _ := getProcurementContract(*resItem.ContractID)
	contractRes, _ := buildProcurementContractResponseItem(contract)
	contractArticles, _ := getProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{ContractID: &contract.Id})

	c := creator.New()
	c.SetPageMargins(50, 50, 50, 50)
	c.NewPage()

	// Define fonts
	fontRegular, err := model.NewCompositePdfFontFromTTFFile(config.BASE_APP_DIR + "fonts/RobotoSlab-VariableFont_wght.ttf")
	if err != nil {
		return shared.HandleAPIError(err)
	}

	fontBold, err := model.NewCompositePdfFontFromTTFFile(config.BASE_APP_DIR + "fonts/RobotoSlab-Bold.ttf")
	if err != nil {
		return shared.HandleAPIError(err)
	}

	// Add Title
	title, err := pdfutils.CreateTitle(c, "Izvještaj o potrošnji i dostupnim količinama", fontBold)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	err = c.Draw(title)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	subtitles := []string{
		"JAVNA NABAVKA: " + resItem.Title,
		"ORGANIZACIONA JEDINICA: " + organizationUnitTitle,
		"DOBAVLJAČ: " + contractRes.Supplier.Title,
	}

	for _, text := range subtitles {
		subtitle, err := pdfutils.CreateSubtitle(c, text, fontRegular)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		err = c.Draw(subtitle)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	// Set the headers for the table.
	headers := []string{"OPIS PREDMETA NABAVKE", "BITNE KARAKTERISTIKE", "UGOVORENA KOLIČINA", "DOSTUPNA KOLIČINA", "POTROŠENA KOLIČINA"}
	tableData := [][]string{}

	for _, article := range contractArticles.Data {
		articleRes, _ := buildProcurementContractArticlesResponseItem(ctx, article)
		articleOrderItem, _ := processContractArticle(ctx, article)

		rowData := []string{
			articleRes.Article.Title,
			articleRes.Article.Description,
			fmt.Sprintf("%d", articleRes.Amount),
			fmt.Sprintf("%d", articleOrderItem.Available),
			fmt.Sprintf("%d", articleRes.Amount-articleOrderItem.Available),
		}
		tableData = append(tableData, rowData)
	}

	table, err := pdfutils.CreateTable(c, headers, tableData, fontRegular)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	err = c.Draw(table)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	// Add a hook that is called before drawing each page
	c.DrawFooter(func(block *creator.Block, args creator.FooterFunctionArgs) {
		p := c.NewParagraph(time.Now().Format("2006-01-02"))
		p.SetFont(fontRegular)
		p.SetFontSize(10)
		p.SetPos(50, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		_ = block.Draw(p)
	})

	var buf bytes.Buffer
	err = c.Write(&buf)
	if err != nil {
		return nil, err
	}

	encodedStr := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Return the path or a URL to the file
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's is your PDF file in base64 encode format!",
		Item:    encodedStr,
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
		resItem, _ := buildProcurementItemResponseItem(params.Context, res, nil)

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		res, err := createProcurementItem(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		resItem, _ := buildProcurementItemResponseItem(params.Context, res, nil)

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

func buildProcurementItemResponseItem(context context.Context, item *structs.PublicProcurementItem, organizationUnitID *int) (*dto.ProcurementItemResponseItem, error) {
	if organizationUnitID == nil {
		organizationUnitID, _ = context.Value(config.OrganizationUnitIDKey).(*int) // assert the type
	}

	plan, _ := getProcurementPlan(item.PlanId)
	planDropdown := dto.DropdownSimple{Id: plan.Id, Title: plan.Title}

	var articles []*dto.ProcurementArticleResponseItem
	articlesRaw, err := getProcurementArticlesList(&dto.GetProcurementArticleListInputMS{ItemID: &item.Id})
	if err != nil {
		return nil, err
	}
	for _, article := range articlesRaw {
		articleResItem, err := buildProcurementArticleResponseItem(context, article, organizationUnitID)
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
