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
	active, activeOk := params.Args["active"].(bool)

	var (
		items []dto.TaxAuthorityCodebookResponse
		total int
	)

	if id != nil && id != 0 {
		setting, err := r.Repo.GetTaxAuthorityCodebookByID(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}

		responseItem, err := buildTaxAuthorityCodeBook(*setting, r)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		items = []dto.TaxAuthorityCodebookResponse{*responseItem}
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

		for _, item := range res.Data {
			resItem, err := buildTaxAuthorityCodeBook(item, r)

			if err != nil {
				return errors.HandleAPIError(err)
			}

			items = append(items, *resItem)
		}

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

		if data.Code == "" && data.Title == "" {
			err := r.Repo.DeactivateTaxAuthorityCodebook(params.Context, itemID, data.Active)

			if err != nil {
				return errors.HandleAPIError(err)
			}
		} else {
			itemRes, err := r.Repo.UpdateTaxAuthorityCodebook(params.Context, itemID, &data)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			response.Message = "You updated this item!"
			responseItem, err := buildTaxAuthorityCodeBook(*itemRes, r)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			response.Item = responseItem
		}
	} else {
		itemRes, err := r.Repo.CreateTaxAuthorityCodebook(params.Context, &data)
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
		ID:             item.ID,
		Title:          item.Title,
		Code:           item.Code,
		Active:         item.Active,
		AmountLess700:  item.AmountLess700,
		AmountLess1000: item.AmountLess1000,
		AmountMore1000: item.AmountMore1000,
		IncludeSubtax:  item.IncludeSubtax,
	}

	if item.TaxPercentage != 0 {
		response.TaxPercentage = &item.TaxPercentage
	}
	if item.ReleasePercentage != 0 {
		response.ReleasePercentage = &item.ReleasePercentage
	}
	if item.ReleaseAmount != 0 {
		response.ReleaseAmount = &item.ReleaseAmount
	}
	if item.PioPercentage != 0 {
		response.PioPercentage = &item.PioPercentage
	}
	if item.PioPercentageEmployerPercentage != 0 {
		response.PioPercentageEmployerPercentage = &item.PioPercentageEmployerPercentage
	}
	if item.PioPercentageEmployeePercentage != 0 {
		response.PioPercentageEmployeePercentage = &item.PioPercentageEmployeePercentage
	}
	if item.UnemploymentPercentage != 0 {
		response.UnemploymentPercentage = &item.UnemploymentPercentage
	}
	if item.UnemploymentEmployerPercentage != 0 {
		response.UnemploymentEmployerPercentage = &item.UnemploymentEmployerPercentage
	}
	if item.UnemploymentEmployeePercentage != 0 {
		response.UnemploymentEmployeePercentage = &item.UnemploymentEmployeePercentage
	}
	if item.LaborFund != 0 {
		response.LaborFund = &item.LaborFund
	}
	if item.PreviousIncomePercentageLessThan700 != 0 {
		response.PreviousIncomePercentageLessThan700 = &item.PreviousIncomePercentageLessThan700
	}
	if item.PreviousIncomePercentageLessThan1000 != 0 {
		response.PreviousIncomePercentageLessThan1000 = &item.PreviousIncomePercentageLessThan1000
	}
	if item.PreviousIncomePercentageMoreThan1000 != 0 {
		response.PreviousIncomePercentageMoreThan1000 = &item.PreviousIncomePercentageMoreThan1000
	}
	if item.Coefficient != 0 {
		response.Coefficient = &item.Coefficient
	}
	if item.CoefficientLess700 != 0 {
		response.CoefficientLess700 = &item.CoefficientLess700
	}
	if item.CoefficientLess1000 != 0 {
		response.CoefficientLess1000 = &item.CoefficientLess1000
	}
	if item.CoefficientMore1000 != 0 {
		response.CoefficientMore1000 = &item.CoefficientMore1000
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
