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

var PublicProcurementContractArticlesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.ProcurementContractArticlesResponseItem{}
	var total int

	contract_id := params.Args["contract_id"]

	ctx := params.Context
	if params.Args["organization_unit_id"] != nil {
		organizationUnitID := params.Args["organization_unit_id"].(int)
		ctx = context.WithValue(ctx, config.OrganizationUnitIDKey, &organizationUnitID)
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

	for _, contractArticle := range contractsRes.Data {
		resItem, err := buildProcurementContractArticlesResponseItem(ctx, contractArticle)
		if err != nil {
			return shared.HandleAPIError(err)
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
	var data structs.PublicProcurementContractArticle
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
		res, err := updateProcurementContractArticle(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementContractArticlesResponseItem(params.Context, res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
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
		response.Item = item
	}

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

	overageList, err := getProcurementContractArticleOverageList(&dto.GetProcurementContractArticleOverageInput{
		ContractArticleID:  &item.Id,
		OrganizationUnitID: organizationUnitID,
	})
	if err != nil {
		return nil, err
	}

	overageTotal := 0
	for _, item := range overageList {
		overageTotal += item.Amount
	}

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
		OverageList:  overageList,
		OverageTotal: overageTotal,
		NetValue:     item.NetValue,
		GrossValue:   item.GrossValue,
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
