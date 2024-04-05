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
	active, activeOk := params.Args["search"].(bool)

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

		if activeOk {
			input.Active = &active
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
		responseItem, err := buildTaxAuthorityCodeBook(*itemRes, r)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = responseItem

	} else {
		itemRes, err := r.Repo.CreateTaxAuthorityCodebook(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		responseItem, err := buildTaxAuthorityCodeBook(*itemRes, r)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = responseItem
	}

	return response, nil

}

func buildTaxAuthorityCodeBook(item structs.TaxAuthorityCodebook, r *Resolver) (*dto.TaxAuthorityCodebookResponse, error) {
	response := &dto.TaxAuthorityCodebookResponse{
		ID:                                   item.ID,
		Title:                                item.Title,
		Code:                                 item.Code,
		Active:                               item.Active,
		TaxPercentage:                        item.TaxPercentage,
		ReleasePercentage:                    item.ReleasePercentage,
		PioPercentage:                        item.PioPercentage,
		PioPercentageEmployerPercentage:      item.PioPercentageEmployerPercentage,
		PioPercentageEmployeePercentage:      item.PioPercentageEmployeePercentage,
		UnemploymentPercentage:               item.UnemploymentPercentage,
		UnemploymentEmployerPercentage:       item.UnemploymentEmployerPercentage,
		UnemploymentEmployeePercentage:       item.UnemploymentEmployeePercentage,
		LaborFund:                            item.LaborFund,
		PreviousIncomePercentageLessThan700:  item.PreviousIncomePercentageLessThan700,
		PreviousIncomePercentageLessThan1000: item.PreviousIncomePercentageLessThan1000,
		PreviousIncomePercentageMoreThan1000: item.PreviousIncomePercentageMoreThan1000,
		Coefficient:                          item.Coefficient,
	}

	if item.TaxSupplierID != 0 {
		supplier, err := r.Repo.GetSupplier(item.TaxSupplierID)

		if err != nil {
			return nil, err
		}

		response.TaxSupplier.ID = supplier.ID
		response.TaxSupplier.Title = supplier.Title
	}

	if item.PioSupplierID != 0 {
		supplier, err := r.Repo.GetSupplier(item.PioSupplierID)

		if err != nil {
			return nil, err
		}

		response.PioSupplier.ID = supplier.ID
		response.PioSupplier.Title = supplier.Title
	}

	if item.LaborFundSupplierID != 0 {
		supplier, err := r.Repo.GetSupplier(item.LaborFundSupplierID)

		if err != nil {
			return nil, err
		}

		response.LaborFundSupplier.ID = supplier.ID
		response.LaborFundSupplier.Title = supplier.Title
	}

	if item.PioEmployeeSupplierID != 0 {
		supplier, err := r.Repo.GetSupplier(item.PioEmployeeSupplierID)

		if err != nil {
			return nil, err
		}

		response.PioEmployeeSupplier.ID = supplier.ID
		response.PioEmployeeSupplier.Title = supplier.Title
	}

	if item.UnemploymentSupplierID != 0 {
		supplier, err := r.Repo.GetSupplier(item.UnemploymentSupplierID)

		if err != nil {
			return nil, err
		}

		response.UnemploymentSupplier.ID = supplier.ID
		response.UnemploymentSupplier.Title = supplier.Title
	}

	return response, nil
}
