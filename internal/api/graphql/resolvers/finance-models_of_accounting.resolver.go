package resolvers

import (
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ModelsOfAccountingOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		ModelsOfAccounting, err := r.Repo.GetModelsOfAccountingByID(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		res, err := buildModelsOfAccounting(*ModelsOfAccounting, r)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.ModelsOfAccountingResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.ModelsOfAccountingFilter{}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	items, total, err := r.Repo.GetModelsOfAccountingList(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var resItems []dto.ModelsOfAccountingResponse
	for _, item := range items {
		resItem, err := buildModelsOfAccounting(item, r)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		resItems = append(resItems, *resItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   resItems,
		Total:   total,
	}, nil
}

func (r *Resolver) ModelsOfAccountingUpdateResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.ModelsOfAccounting
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		return apierrors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var item *structs.ModelsOfAccounting

	item, err = r.Repo.UpdateModelsOfAccounting(params.Context, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	singleItem, err := buildModelsOfAccounting(*item, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func buildModelsOfAccounting(item structs.ModelsOfAccounting, r *Resolver) (*dto.ModelsOfAccountingResponse, error) {
	response := dto.ModelsOfAccountingResponse{
		ID:        item.ID,
		Title:     item.Title,
		Type:      item.Type,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	for _, orderItem := range item.Items {

		item := dto.ModelOfAccountingItemResponse{
			ID:    orderItem.ID,
			Title: orderItem.Title,
		}

		if orderItem.DebitAccountID != 0 {
			account, err := r.Repo.GetAccountItemByID(orderItem.DebitAccountID)

			if err != nil {
				return nil, err
			}

			builtAccount := dto.AccountItemResponseItem{
				ID:           account.ID,
				SerialNumber: account.SerialNumber,
				Title:        account.Title,
			}

			item.DebitAccount = builtAccount
		}

		if orderItem.CreditAccountID != 0 {
			account, err := r.Repo.GetAccountItemByID(orderItem.CreditAccountID)

			if err != nil {
				return nil, err
			}

			builtAccount := dto.AccountItemResponseItem{
				ID:           account.ID,
				SerialNumber: account.SerialNumber,
				Title:        account.Title,
			}

			item.CreditAccount = builtAccount
		}

		response.Items = append(response.Items, item)
	}

	return &response, nil
}
