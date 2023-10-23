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

var PublicProcurementPlanItemArticleInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementArticle
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
		res, err := updateProcurementArticle(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementArticleResponseItem(params.Context, res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := createProcurementArticle(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementArticleResponseItem(params.Context, res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

var PublicProcurementPlanItemArticleDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteProcurementArticle(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildProcurementArticleResponseItem(context context.Context, item *structs.PublicProcurementArticle) (*dto.ProcurementArticleResponseItem, error) {
	organizationUnitID, _ := context.Value(config.OrganizationUnitIDKey).(*int)
	procurement, err := getProcurementItem(item.PublicProcurementId)
	if err != nil {
		return nil, err
	}
	procurementDropdown := dto.DropdownSimple{Id: procurement.Id, Title: procurement.Title}

	res := dto.ProcurementArticleResponseItem{
		Id:                item.Id,
		PublicProcurement: procurementDropdown,
		Title:             item.Title,
		Description:       item.Description,
		NetPrice:          item.NetPrice,
		VATPercentage:     &item.VatPercentage,
		Manufacturer:      item.Manufacturer,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	ouArticles, _ := getProcurementOUArticleList(&dto.GetProcurementOrganizationUnitArticleListInputDTO{ArticleID: &item.Id})

	totalAmount := 0

	for _, ouArticle := range ouArticles {
		if ouArticle.Status == structs.ArticleStatusInProgress {
			continue
		}
		totalAmount += ouArticle.Amount
		if organizationUnitID != nil && ouArticle.OrganizationUnitId == *organizationUnitID {
			res.Amount = ouArticle.Amount
		}
	}

	res.TotalAmount = totalAmount

	return &res, nil
}

func createProcurementArticle(article *structs.PublicProcurementArticle) (*structs.PublicProcurementArticle, error) {
	res := &dto.GetProcurementArticleResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ARTICLES_ENDPOINT, article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementArticle(id int, article *structs.PublicProcurementArticle) (*structs.PublicProcurementArticle, error) {
	res := &dto.GetProcurementArticleResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ARTICLES_ENDPOINT+"/"+strconv.Itoa(id), article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteProcurementArticle(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.ARTICLES_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getProcurementArticlesList(input *dto.GetProcurementArticleListInputMS) ([]*structs.PublicProcurementArticle, error) {
	res := &dto.GetProcurementArticleListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ARTICLES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getProcurementArticle(id int) (*structs.PublicProcurementArticle, error) {
	res := &dto.GetProcurementArticleResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ARTICLES_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
