package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

var PublicProcurementContractArticlesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.ProcurementContractArticlesResponseItem{}
	var total int

	contract_id := params.Args["contract_id"]

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
		resItem, err := buildProcurementContractArticlesResponseItem(contractArticle)
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
		item, err := buildProcurementContractArticlesResponseItem(res)
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
		item, err := buildProcurementContractArticlesResponseItem(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func buildProcurementContractArticlesResponseItem(item *structs.PublicProcurementContractArticle) (*dto.ProcurementContractArticlesResponseItem, error) {
	article, err := getProcurementArticle(item.PublicProcurementArticleId)
	if err != nil {
		return nil, err
	}
	contract, err := getProcurementContract(item.PublicProcurementContractId)
	if err != nil {
		return nil, err
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
		Amount:     item.Amount,
		NetValue:   item.NetValue,
		GrossValue: item.GrossValue,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}

	return &res, nil
}

func createProcurementContractArticle(resolution *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error) {
	res := &dto.GetProcurementContractArticleResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.CONTRACT_ARTICLE_ENDPOINT, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementContractArticle(id int, resolution *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error) {
	res := &dto.GetProcurementContractArticleResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.CONTRACT_ARTICLE_ENDPOINT+"/"+strconv.Itoa(id), resolution, res)
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
