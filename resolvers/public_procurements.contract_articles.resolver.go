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

var PublicProcurementContractArticlesOrganizationUnitResponseItem = func(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.ProcurementContractArticlesResponseItem{}
	var total int

	contractId := params.Args["contract_id"].(int)

	articles, err := getProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{ContractID: &contractId})

	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, article := range articles.Data {
		articlesInOrgUnit, err := getProcurementOUArticleList(&dto.GetProcurementOrganizationUnitArticleListInputDTO{
			ArticleID: &article.PublicProcurementArticleId})

		if err != nil {
			return shared.HandleAPIError(err)
		}
		amount := 0
		for _, articleInOrgUnit := range articlesInOrgUnit {
			amount += articleInOrgUnit.Amount
		}

		inventors, err := getAllInventoryItem(dto.InventoryItemFilter{
			ArticleId: &article.PublicProcurementArticleId,
		})

		amount = amount - inventors.Total
		if err != nil {
			return shared.HandleAPIError(err)
		}

		articleData, err := getProcurementArticle(article.PublicProcurementArticleId)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		articleOverages, err := getProcurementContractArticleOverageList(&dto.GetProcurementContractArticleOverageInput{
			ContractArticleID: &articleData.Id,
		})

		if err != nil {
			return shared.HandleAPIError(err)
		}

		var total int

		for _, overage := range articleOverages {
			total += overage.Amount
		}

		amount += total
		vatPercentageFloat, err := strconv.ParseFloat(articleData.VatPercentage, 32)

		var grossValue float32

		if err != nil {
			grossValue = 0
		} else {
			grossValue = article.NetValue + article.NetValue*float32(vatPercentageFloat)/100
		}

		if amount > 0 {
			items = append(items, dto.ProcurementContractArticlesResponseItem{
				Id: article.PublicProcurementArticleId,
				Article: dto.DropdownProcurementArticle{
					Id:            articleData.Id,
					Title:         articleData.Title,
					VatPercentage: articleData.VatPercentage,
					Description:   articleData.Description,
				},
				Amount:       amount,
				NetValue:     articleData.NetPrice,
				GrossValue:   grossValue,
				OverageList:  articleOverages,
				OverageTotal: total,
			})
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   total,
	}, nil
}

var PublicProcurementContractArticlesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.ProcurementContractArticlesResponseItem{}
	var total int

	contract_id := params.Args["contract_id"]
	visibilityType := params.Args["visibility_type"]

	ctx := params.Context
	var orgUnitID *int
	if params.Args["organization_unit_id"] != nil {
		organizationUnitID := params.Args["organization_unit_id"].(int)
		ctx = context.WithValue(ctx, config.OrganizationUnitIDKey, &organizationUnitID)
		orgUnitID = &organizationUnitID
	}

	input := dto.GetProcurementContractArticlesInput{}

	if shared.IsInteger(contract_id) && contract_id.(int) > 0 {
		contractID := contract_id.(int)
		input.ContractID = &contractID
	}

	contractsRes, err := getProcurementContractArticlesList(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	total = contractsRes.Total

	contract, err := getProcurementContract(*input.ContractID)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	procurement, err := getProcurementItem(contract.PublicProcurementId)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, contractArticle := range contractsRes.Data {
		article, _ := getProcurementArticle(contractArticle.PublicProcurementArticleId)
		if visibilityType != nil && visibilityType.(int) != int(article.VisibilityType) {
			continue
		}
		resItem, err := buildProcurementContractArticlesResponseItem(ctx, contractArticle)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		filter := dto.GetProcurementOrganizationUnitArticleListInputDTO{
			ArticleID: &contractArticle.PublicProcurementArticleId,
		}

		if orgUnitID != nil && *orgUnitID > 0 {
			filter.OrganizationUnitID = orgUnitID
		}

		orgUnitArticles, err := getOrganizationUnitArticlesList(filter)

		if err != nil {
			return nil, err
		}

		amount := 0

		for _, orgArticle := range orgUnitArticles {
			amount += orgArticle.Amount
		}

		resItem.Amount = amount

		if !procurement.IsOpenProcurement {
			resItem.Amount = article.Amount
		}

		items = append(items, *resItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   total,
	}, nil
}

var PublicProcurementContractArticleInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data []structs.PublicProcurementContractArticle
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	var items []*dto.ProcurementContractArticlesResponseItem
	for _, data := range data {
		itemId := data.Id

		if shared.IsInteger(itemId) && itemId != 0 {
			res, err := updateProcurementContractArticle(itemId, &data)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			item, err := buildProcurementContractArticlesResponseItem(params.Context, res)
			if err != nil {
				return shared.HandleAPIError(err)
			}

			response.Message = "You updated this item!"
			items = append(items, item)
		} else {
			res, err := createProcurementContractArticle(&data)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			item, err := buildProcurementContractArticlesResponseItem(params.Context, res)
			if err != nil {
				return shared.HandleAPIError(err)
			}

			response.Message = "You created this item!"
			items = append(items, item)
		}
	}
	response.Item = items
	return response, nil
}

var PublicProcurementContractArticleOverageInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementContractArticleOverage
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
		res, err := updateProcurementContractArticleOverage(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = res
	} else {
		res, err := createProcurementContractArticleOverage(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = res
	}

	return response, nil
}

var PublicProcurementContractArticleOverageDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteProcurementContractArticleOverage(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildProcurementContractArticlesResponseItem(context context.Context, item *structs.PublicProcurementContractArticle) (*dto.ProcurementContractArticlesResponseItem, error) {
	organizationUnitID, _ := context.Value(config.OrganizationUnitIDKey).(*int)

	article, err := getProcurementArticle(item.PublicProcurementArticleId)
	if err != nil {
		return nil, err
	}
	articleResItem, err := buildProcurementArticleResponseItem(context, article, organizationUnitID)
	if err != nil {
		return nil, err
	}
	contract, err := getProcurementContract(item.PublicProcurementContractId)
	if err != nil {
		return nil, err
	}

	overageInput := dto.GetProcurementContractArticleOverageInput{ContractArticleID: &item.Id}
	if organizationUnitID != nil && *organizationUnitID != 0 {
		overageInput.OrganizationUnitID = organizationUnitID
	}
	overageList, err := getProcurementContractArticleOverageList(&overageInput)
	if err != nil {
		return nil, err
	}

	overageTotal := 0
	for _, item := range overageList {
		overageTotal += item.Amount
	}

	vatPercentage, err := strconv.ParseFloat(article.VatPercentage, 64)

	if err != nil {
		return nil, err
	}

	grossPrice := item.NetValue + item.NetValue*float32(vatPercentage)/100

	res := dto.ProcurementContractArticlesResponseItem{
		Id: item.Id,
		Article: dto.DropdownProcurementArticle{
			Id:            article.Id,
			Title:         article.Title,
			VatPercentage: article.VatPercentage,
			Description:   article.Description,
		},
		Contract: dto.DropdownSimple{
			Id:    contract.Id,
			Title: contract.SerialNumber,
		},
		UsedArticles: item.UsedArticles,
		OverageList:  overageList,
		OverageTotal: overageTotal,
		NetValue:     item.NetValue,
		GrossValue:   grossPrice,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	if organizationUnitID != nil && *organizationUnitID == 0 {
		res.Amount = articleResItem.TotalAmount
	} else {
		res.Amount = articleResItem.Amount
	}

	return &res, nil
}

func createProcurementContractArticle(article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error) {
	res := &dto.GetProcurementContractArticleResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.CONTRACT_ARTICLE_ENDPOINT, article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementContractArticle(id int, article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error) {
	res := &dto.GetProcurementContractArticleResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.CONTRACT_ARTICLE_ENDPOINT+"/"+strconv.Itoa(id), article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getProcurementContractArticlesList(input *dto.GetProcurementContractArticlesInput) (*dto.GetProcurementContractArticlesListResponseMS, error) {
	res := &dto.GetProcurementContractArticlesListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.CONTRACT_ARTICLE_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func createProcurementContractArticleOverage(articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error) {
	res := &dto.GetProcurementContractArticleOverageResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.CONTRACT_ARTICLE_OVERAGE_ENDPOINT, articleOverage, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementContractArticleOverage(id int, articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error) {
	res := &dto.GetProcurementContractArticleOverageResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.CONTRACT_ARTICLE_OVERAGE_ENDPOINT+"/"+strconv.Itoa(id), articleOverage, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getProcurementContractArticleOverageList(input *dto.GetProcurementContractArticleOverageInput) ([]*structs.PublicProcurementContractArticleOverage, error) {
	res := &dto.GetProcurementContractArticleOverageListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.CONTRACT_ARTICLE_OVERAGE_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func deleteProcurementContractArticleOverage(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.CONTRACT_ARTICLE_OVERAGE_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
