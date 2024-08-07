package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var items []*dto.ProcurementArticleResponseItem
	for _, item := range data {
		itemID := item.ID

		if itemID != 0 {
			res, err := r.Repo.UpdateProcurementArticle(itemID, &item)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			item, err := buildProcurementArticleResponseItem(params.Context, r, res, nil)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			items = append(items, item)
		} else {
			res, err := r.Repo.CreateProcurementArticle(&item)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			item, err := buildProcurementArticleResponseItem(params.Context, r, res, nil)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			items = append(items, item)
			response.Message = "You created this item!"
		}
	}
	response.Items = items
	return response, nil
}

func (r *Resolver) PublicProcurementPlanItemArticleDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteProcurementArticle(itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildProcurementArticleResponseItem(context context.Context, r *Resolver, item *structs.PublicProcurementArticle, organizationUnitID *int) (*dto.ProcurementArticleResponseItem, error) {
	if organizationUnitID == nil {
		organizationUnitID, _ = context.Value(config.OrganizationUnitIDKey).(*int)
	}
	procurement, err := r.Repo.GetProcurementItem(item.PublicProcurementID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get procurement item")
	}
	procurementDropdown := dto.DropdownSimple{ID: procurement.ID, Title: procurement.Title}

	res := dto.ProcurementArticleResponseItem{
		ID:                item.ID,
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

	ouArticles, _ := r.Repo.GetProcurementOUArticleList(&dto.GetProcurementOrganizationUnitArticleListInputDTO{ArticleID: &item.ID})

	totalAmount := 0

	for _, ouArticle := range ouArticles {
		if ouArticle.Status == structs.ArticleStatusInProgress {
			continue
		}
		totalAmount += ouArticle.Amount
		if organizationUnitID != nil && ouArticle.OrganizationUnitID == *organizationUnitID {
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
