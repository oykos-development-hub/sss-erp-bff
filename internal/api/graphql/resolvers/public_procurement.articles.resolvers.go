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
	"strconv"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) PublicProcurementPlanItemArticleInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data []structs.PublicProcurementArticle
	response := dto.Response{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	var items []*dto.ProcurementArticleResponseItem
	for _, item := range data {
		itemId := item.Id

		if shared.IsInteger(itemId) && itemId != 0 {
			res, err := r.Repo.UpdateProcurementArticle(itemId, &item)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			item, err := buildProcurementArticleResponseItem(r.Repo, params.Context, res, nil)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			items = append(items, item)
		} else {
			res, err := r.Repo.CreateProcurementArticle(&item)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			item, err := buildProcurementArticleResponseItem(r.Repo, params.Context, res, nil)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			items = append(items, item)
			response.Message = "You created this item!"
		}
	}
	response.Items = items
	return response, nil
}

func (r *Resolver) PublicProcurementPlanItemArticleDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := r.Repo.DeleteProcurementArticle(itemId)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildProcurementArticleResponseItem(r repository.MicroserviceRepositoryInterface, context context.Context, item *structs.PublicProcurementArticle, organizationUnitID *int) (*dto.ProcurementArticleResponseItem, error) {
	if organizationUnitID == nil {
		organizationUnitID, _ = context.Value(config.OrganizationUnitIDKey).(*int)
	}
	procurement, err := r.GetProcurementItem(item.PublicProcurementId)
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
		VisibilityType:    int(item.VisibilityType),
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	ouArticles, _ := r.GetProcurementOUArticleList(&dto.GetProcurementOrganizationUnitArticleListInputDTO{ArticleID: &item.Id})

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

	// if it's simple procurement, then amount is entered directly to article by official
	if !procurement.IsOpenProcurement {
		res.Amount = item.Amount
		res.TotalAmount = item.Amount
	} else {
		res.TotalAmount = totalAmount
	}

	vatPercentage, _ := strconv.ParseFloat(item.VatPercentage, 32)
	res.GrossPrice = item.NetPrice + item.NetPrice*float32(vatPercentage)/100

	return &res, nil
}
