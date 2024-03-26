package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) TaxAuthorityCodebooksOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	search, searchOk := params.Args["search"].(string)

	var (
		items []structs.TaxAuthorityCodebook
		total int
	)

	if id != nil && id != 0 {
		setting, err := r.Repo.GetTaxAuthorityCodebookByID(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}
		items = []structs.TaxAuthorityCodebook{*setting}
		total = 1
	} else {
		input := dto.TaxAuthorityCodebookFilter{}
		if searchOk && search != "" {
			input.Search = &search
		}

		res, err := r.Repo.GetTaxAuthorityCodebooks(input)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		items = res.Data
		total = res.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   total,
	}, nil
}

func (r *Resolver) TaxAuthorityCodebooksInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.TaxAuthorityCodebook
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	response := dto.ResponseSingle{
		Status: "success",
	}

	if itemID != 0 {
		itemRes, err := r.Repo.UpdateTaxAuthorityCodebook(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = itemRes

	} else {
		itemRes, err := r.Repo.CreateTaxAuthorityCodebook(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = itemRes
	}

	return response, nil

}

func (r *Resolver) TaxAuthorityCodebooksDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	/*itemID := params.Args["id"].(int)

	err := r.Repo.DeleteDropdownSettings(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil*/
	return dto.ResponseSingle{
		Status:  "failed",
		Message: "You can not delete this item!",
	}, nil
}
