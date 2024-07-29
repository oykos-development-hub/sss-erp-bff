package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"context"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) PublicProcurementContractArticlesOrganizationUnitResponseItem(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.ProcurementContractArticlesResponseItem{}
	var total int

	contractID := params.Args["contract_id"].(int)

	articles, err := r.Repo.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{ContractID: &contractID})

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, article := range articles.Data {
		articlesInOrgUnit, err := r.Repo.GetProcurementOUArticleList(&dto.GetProcurementOrganizationUnitArticleListInputDTO{
			ArticleID: &article.PublicProcurementArticleID})

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		amount := 0
		for _, articleInOrgUnit := range articlesInOrgUnit {
			amount += articleInOrgUnit.Amount
		}

		inventors, err := r.Repo.GetAllInventoryItem(dto.InventoryItemFilter{
			ArticleID: &article.PublicProcurementArticleID,
		})

		amount = amount - inventors.Total
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		articleData, err := r.Repo.GetProcurementArticle(article.PublicProcurementArticleID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		articleOverages, err := r.Repo.GetProcurementContractArticleOverageList(&dto.GetProcurementContractArticleOverageInput{
			ContractArticleID: &articleData.ID,
		})

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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

		if amount > 0 && articleData.VisibilityType == 1 {
			items = append(items, dto.ProcurementContractArticlesResponseItem{
				ID: article.PublicProcurementArticleID,
				Article: dto.DropdownProcurementArticle{
					ID:            articleData.ID,
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

func (r *Resolver) PublicProcurementContractArticlesOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.ProcurementContractArticlesResponseItem{}
	var total int

	contractID := params.Args["contract_id"]
	visibilityType := params.Args["visibility_type"]

	ctx := params.Context
	var orgUnitID *int
	if params.Args["organization_unit_id"] != nil {
		organizationUnitID := params.Args["organization_unit_id"].(int)
		ctx = context.WithValue(ctx, config.OrganizationUnitIDKey, &organizationUnitID)
		orgUnitID = &organizationUnitID
	}

	input := dto.GetProcurementContractArticlesInput{}

	if contractID != nil && contractID.(int) > 0 {
		contractID := contractID.(int)
		input.ContractID = &contractID
	}

	contractsRes, err := r.Repo.GetProcurementContractArticlesList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	total = contractsRes.Total

	contract, err := r.Repo.GetProcurementContract(*input.ContractID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	procurement, err := r.Repo.GetProcurementItem(contract.PublicProcurementID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, contractArticle := range contractsRes.Data {
		article, _ := r.Repo.GetProcurementArticle(contractArticle.PublicProcurementArticleID)
		if visibilityType != nil && visibilityType.(int) != int(article.VisibilityType) {
			continue
		}
		resItem, err := buildProcurementContractArticlesResponseItem(ctx, r.Repo, contractArticle)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		filter := dto.GetProcurementOrganizationUnitArticleListInputDTO{
			ArticleID: &contractArticle.PublicProcurementArticleID,
		}

		if orgUnitID != nil && *orgUnitID > 0 {
			filter.OrganizationUnitID = orgUnitID
		}

		orgUnitArticles, err := r.Repo.GetOrganizationUnitArticlesList(filter)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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

func (r *Resolver) PublicProcurementContractArticleInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data []structs.PublicProcurementContractArticle
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var items []*dto.ProcurementContractArticlesResponseItem
	for _, data := range data {
		itemID := data.ID

		if itemID != 0 {
			res, err := r.Repo.UpdateProcurementContractArticle(itemID, &data)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			item, err := buildProcurementContractArticlesResponseItem(params.Context, r.Repo, res)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			response.Message = "You updated this item!"
			items = append(items, item)
		} else {
			res, err := r.Repo.CreateProcurementContractArticle(&data)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			item, err := buildProcurementContractArticlesResponseItem(params.Context, r.Repo, res)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			response.Message = "You created this item!"
			items = append(items, item)
		}
	}
	response.Item = items
	return response, nil
}

func (r *Resolver) PublicProcurementContractArticleOverageInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementContractArticleOverage
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		res, err := r.Repo.UpdateProcurementContractArticleOverage(itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Message = "You updated this item!"
		response.Item = res
	} else {
		res, err := r.Repo.CreateProcurementContractArticleOverage(&data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Message = "You created this item!"
		response.Item = res
	}

	return response, nil
}

func (r *Resolver) PublicProcurementContractArticleOverageDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteProcurementContractArticleOverage(itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildProcurementContractArticlesResponseItem(context context.Context, r repository.MicroserviceRepositoryInterface, item *structs.PublicProcurementContractArticle) (*dto.ProcurementContractArticlesResponseItem, error) {
	organizationUnitID, _ := context.Value(config.OrganizationUnitIDKey).(*int)

	article, err := r.GetProcurementArticle(item.PublicProcurementArticleID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get procurement article")
	}
	articleResItem, err := buildProcurementArticleResponseItem(context, r, article, organizationUnitID)
	if err != nil {
		return nil, errors.Wrap(err, "build procurement article response item")
	}
	contract, err := r.GetProcurementContract(item.PublicProcurementContractID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get procurement contract")
	}

	overageInput := dto.GetProcurementContractArticleOverageInput{ContractArticleID: &item.ID}
	if organizationUnitID != nil && *organizationUnitID != 0 {
		overageInput.OrganizationUnitID = organizationUnitID
	}
	overageList, err := r.GetProcurementContractArticleOverageList(&overageInput)
	if err != nil {
		return nil, errors.Wrap(err, "repo get procurement contract article overage list")
	}

	overageTotal := 0
	for _, item := range overageList {
		overageTotal += item.Amount
	}

	vatPercentage, err := strconv.ParseFloat(article.VatPercentage, 64)

	if err != nil {
		return nil, errors.Wrap(err, "strconv parse float")
	}

	grossPrice := item.NetValue + item.NetValue*float32(vatPercentage)/100

	res := dto.ProcurementContractArticlesResponseItem{
		ID: item.ID,
		Article: dto.DropdownProcurementArticle{
			ID:            article.ID,
			Title:         article.Title,
			VatPercentage: article.VatPercentage,
			Description:   article.Description,
		},
		Contract: dto.DropdownSimple{
			ID:    contract.ID,
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
