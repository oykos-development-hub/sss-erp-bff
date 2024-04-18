package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) InvoiceOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		invoice, err := r.Repo.GetInvoice(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		invoiceResItem, err := buildInvoiceResponseItem(params.Context, r, *invoice)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.InvoiceResponseItem{invoiceResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetInvoiceListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["type"].(string); ok && value != "" {
		input.Type = &value
	}

	if value, ok := params.Args["status"].(string); ok && value != "" {
		input.Status = &value
	}

	if value, ok := params.Args["supplier_id"].(int); ok && value != 0 {
		input.SupplierID = &value
	}

	if value, ok := params.Args["year"].(int); ok && value != 0 {
		input.Year = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	}

	if value, ok := params.Args["passed_to_inventory"].(bool); ok && value {
		input.PassedToInventory = &value
	}

	invoices, total, err := r.Repo.GetInvoiceList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	invoiceResItem, err := buildInvoiceResponseItemList(params.Context, r, invoices)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   invoiceResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) InvoiceInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Invoice
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		return errors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	var item *structs.Invoice

	if data.OrganizationUnitID == 0 {
		organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return errors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}

		data.OrganizationUnitID = *organizationUnitID
	}

	if data.ID == 0 {
		if data.PassedToAccounting {

			orderList := &structs.OrderListItem{
				SupplierID:            &data.SupplierID,
				Status:                "created",
				IsProFormaInvoice:     true,
				ProFormaInvoiceNumber: data.ProFormaInvoiceNumber,
				OrganizationUnitID:    data.OrganizationUnitID,
				OrderFile:             &data.FileID,
				ReceiveFile:           []int{data.ProFormaInvoiceFileID},
				InvoiceNumber:         &data.InvoiceNumber,
			}

			if data.ProFormaInvoiceDate != nil {
				proFormaInvoiceDate := data.ProFormaInvoiceDate.Format("2006-01-02T15:04:05Z")
				orderList.ProFormaInvoiceDate = &proFormaInvoiceDate
				orderList.DateOrder = proFormaInvoiceDate
			}

			if data.DateOfInvoice != nil {
				invoiceDate := data.DateOfInvoice.Format("2006-01-02T15:04:05Z")
				orderList.InvoiceDate = &invoiceDate
				orderList.DateOrder = invoiceDate
			}

			order, err := r.Repo.CreateOrderListItem(orderList)

			if err != nil {
				return errors.HandleAPIError(err)
			}

			for _, article := range data.Articles {
				year := strconv.Itoa(time.Now().Year())
				newArticle := structs.OrderProcurementArticleItem{
					OrderID:       order.ID,
					Year:          year,
					Title:         article.Title,
					Description:   article.Description,
					Amount:        article.Amount,
					NetPrice:      float32(article.NetPrice),
					VatPercentage: article.VatPercentage,
				}

				_, err := r.Repo.CreateOrderProcurementArticle(&newArticle)

				if err != nil {
					return errors.HandleAPIError(err)
				}
			}

			data.OrderID = order.ID

		}

		item, err = r.Repo.CreateInvoice(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		if item.OrderID != 0 {
			order, err := r.Repo.GetOrderListByID(item.OrderID)

			if err != nil {
				return errors.HandleAPIError(err)
			}

			_, err = r.Repo.UpdateOrderListItem(item.OrderID, &structs.OrderListItem{
				ID:                  order.ID,
				DateOrder:           order.DateOrder,
				TotalPrice:          order.TotalPrice,
				PublicProcurementID: order.PublicProcurementID,
				GroupOfArticlesID:   order.GroupOfArticlesID,
				SupplierID:          order.SupplierID,
				Status:              order.Status,
				PassedToFinance:     true,
				UsedInFinance:       true,
				DateSystem:          order.DateSystem,
				InvoiceDate:         order.InvoiceDate,
				InvoiceNumber:       order.InvoiceNumber,
				OrganizationUnitID:  order.OrganizationUnitID,
				OfficeID:            order.OfficeID,
				RecipientUserID:     order.RecipientUserID,
				Description:         order.Description,
				IsUsed:              order.IsUsed,
				OrderFile:           order.OrderFile,
				ReceiveFile:         order.ReceiveFile,
				MovementFile:        order.MovementFile,
			})

			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	} else {

		invoice, err := r.Repo.GetInvoice(data.ID)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		item, err = r.Repo.UpdateInvoice(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		if invoice.DateOfInvoice == nil && data.DateOfInvoice != nil && invoice.InvoiceNumber == "" && data.InvoiceNumber != "" && data.OrderID != 0 {
			order, err := r.Repo.GetOrderListByID(data.OrderID)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			invoiceDate := data.DateOfInvoice.Format("2006-01-02T15:04:05Z")
			order.InvoiceDate = &invoiceDate
			order.InvoiceNumber = &data.InvoiceNumber
			order.ReceiveFile = []int{data.ProFormaInvoiceFileID}

			_, err = r.Repo.UpdateOrderListItem(order.ID, order)

			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	}

	articles, err := r.Repo.GetInvoiceArticleList(item.ID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	articlesForDelete := make(map[int]bool)

	for _, article := range articles {
		articlesForDelete[article.ID] = true
	}

	for _, article := range data.Articles {
		article.InvoiceID = item.ID
		if article.ID != 0 {
			_, err := r.Repo.UpdateInvoiceArticle(&article)

			if err != nil {
				return errors.HandleAPIError(err)
			}
			articlesForDelete[article.ID] = false
		} else {
			_, err := r.Repo.CreateInvoiceArticle(&article)

			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	}

	for id, delete := range articlesForDelete {
		if delete {
			err := r.Repo.DeleteInvoiceArticle(id)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	}

	responseItem, err := buildInvoiceResponseItem(params.Context, r, *item)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	response.Item = *responseItem

	return response, nil
}

func (r *Resolver) InvoiceDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	item, err := r.Repo.GetInvoice(itemID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	err = r.Repo.DeleteInvoice(itemID)
	if err != nil {
		fmt.Printf("Deleting invoice item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	if item.OrderID != 0 && item.InvoiceNumber != "" {
		order, err := r.Repo.GetOrderListByID(item.OrderID)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		_, err = r.Repo.UpdateOrderListItem(item.OrderID, &structs.OrderListItem{
			ID:                  order.ID,
			DateOrder:           order.DateOrder,
			TotalPrice:          order.TotalPrice,
			PublicProcurementID: order.PublicProcurementID,
			GroupOfArticlesID:   order.GroupOfArticlesID,
			SupplierID:          order.SupplierID,
			Status:              order.Status,
			PassedToFinance:     true,
			UsedInFinance:       false,
			DateSystem:          order.DateSystem,
			InvoiceDate:         order.InvoiceDate,
			InvoiceNumber:       order.InvoiceNumber,
			OrganizationUnitID:  order.OrganizationUnitID,
			OfficeID:            order.OfficeID,
			RecipientUserID:     order.RecipientUserID,
			Description:         order.Description,
			IsUsed:              order.IsUsed,
			OrderFile:           order.OrderFile,
			ReceiveFile:         order.ReceiveFile,
			MovementFile:        order.MovementFile,
		})

		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else if item.OrderID != 0 {
		err = r.Repo.DeleteOrderList(item.OrderID)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) AdditionalExpensesOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	input := dto.AdditionalExpensesListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["invoice_id"].(int); ok && value != 0 {
		input.InvoiceID = &value
	}

	if value, ok := params.Args["subject_id"].(int); ok && value != 0 {
		input.SubjectID = &value
	}

	if value, ok := params.Args["status"].(string); ok && value != "" {
		input.Status = &value
	}

	if value, ok := params.Args["year"].(int); ok && value != 0 {
		input.Year = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	}

	additionalExpenses, total, err := r.Repo.GetAdditionalExpenses(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	builedAdditionalExpenses, err := buildAdditionalExpenseItemList(params.Context, r, additionalExpenses)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   builedAdditionalExpenses,
		Total:   total,
	}, nil
}

func calculateCoefficient(item structs.TaxAuthorityCodebook, subTax float64) float64 {
	var coefficient float64

	coefficient = 1
	release := item.ReleasePercentage

	if release != 0 {
		if item.TaxPercentage != 0 {
			coefficient -= (1 - release/100) * (item.TaxPercentage / 100)
			if subTax != 0 {
				coefficient -= (1 - release/100) * (item.TaxPercentage / 100) * (subTax / 100)
			}
		}
		if item.PioPercentage != 0 {
			coefficient -= (1 - release/100) * (item.PioPercentage / 100)
		}
	} else {
		if item.TaxPercentage != 0 {
			coefficient -= 1 - (item.TaxPercentage / 100)
			if subTax != 0 {
				coefficient -= 1 - (item.TaxPercentage/100)*(subTax/100)
			}
		}
		if item.PioPercentage != 0 {
			coefficient -= 1 - (item.PioPercentage / 100)
		}
	}

	return coefficient
}

func calculateCoefficientLess700(item structs.TaxAuthorityCodebook, previousIncomeGross float64, r *Resolver, organizationUnit *structs.OrganizationUnits, municipality structs.Suppliers) (float64, float64, error) {

	additionalExpenses, err := calculateAdditionalExpenses(item, float64(700), previousIncomeGross, r, organizationUnit, municipality)

	if err != nil {
		return float64(0), float64(0), err
	}

	maxNetAmount := additionalExpenses[len(additionalExpenses)-1].Price

	coefficient := maxNetAmount / 700

	return float64(coefficient), float64(maxNetAmount), nil
}

func calculateCoefficientLess1000(item structs.TaxAuthorityCodebook, previousIncomeGross float64, r *Resolver, organizationUnit *structs.OrganizationUnits, municipality structs.Suppliers) (float64, float64, error) {

	additionalExpenses, err := calculateAdditionalExpenses(item, float64(1000), previousIncomeGross, r, organizationUnit, municipality)

	if err != nil {
		return float64(0), float64(0), err
	}

	maxNetAmount := additionalExpenses[len(additionalExpenses)-1].Price

	_, amount, err := calculateCoefficientLess700(item, previousIncomeGross, r, organizationUnit, municipality)

	if err != nil {
		return float64(0), float64(0), err
	}

	coefficient := (maxNetAmount - float32(amount)) / 300

	return float64(coefficient), float64(maxNetAmount), nil
}

func calculateCoefficientMore1000(item structs.TaxAuthorityCodebook, previousIncomeGross float64, r *Resolver, organizationUnit *structs.OrganizationUnits, municipality structs.Suppliers) (float64, float64, error) {

	coefficient := item.PreviousIncomePercentageMoreThan1000 + item.PioPercentage + item.PioPercentageEmployerPercentage +
		item.UnemploymentEmployeePercentage + item.UnemploymentPercentage

	return float64(1 - coefficient/100), float64(0), nil
}

func (r *Resolver) CalculateAdditionalExpensesResolver(params graphql.ResolveParams) (interface{}, error) {
	taxAuthorityCodebookID := params.Args["tax_authority_codebook_id"].(int)

	taxAuthorityCodebook, err := r.Repo.GetTaxAuthorityCodebookByID(taxAuthorityCodebookID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	municipalityID := params.Args["municipality_id"].(int)

	municipality, err := r.Repo.GetSupplier(municipalityID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return errors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	organizationUnit, err := r.Repo.GetOrganizationUnitByID(*organizationUnitID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	previousIncomeGross, previousIncomeGrossOK := params.Args["previous_income_gross"].(float64)
	previousIncomeNet, previousIncomeNetOK := params.Args["previous_income_net"].(float64)

	taxAuthorityCodebook.CoefficientLess700, taxAuthorityCodebook.AmountLess700, err = calculateCoefficientLess700(*taxAuthorityCodebook, 0, r, organizationUnit, *municipality)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	taxAuthorityCodebook.CoefficientLess1000, taxAuthorityCodebook.AmountLess1000, err = calculateCoefficientLess1000(*taxAuthorityCodebook, 0, r, organizationUnit, *municipality)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	taxAuthorityCodebook.CoefficientMore1000, taxAuthorityCodebook.AmountMore1000, err = calculateCoefficientMore1000(*taxAuthorityCodebook, 0, r, organizationUnit, *municipality)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	if !previousIncomeGrossOK {

		//konvertuje neto u bruto
		if previousIncomeNetOK && taxAuthorityCodebook.TaxPercentage != 0 {
			taxAuthorityCodebook.Coefficient = calculateCoefficient(*taxAuthorityCodebook, float64(municipality.TaxPercentage))
			previousIncomeGross = previousIncomeNet / taxAuthorityCodebook.Coefficient
			helper := math.Round(previousIncomeGross*100) / 100
			previousIncomeGross = float64(helper)

		} else if previousIncomeNetOK && taxAuthorityCodebook.TaxPercentage == 0 {

			if previousIncomeNet < taxAuthorityCodebook.AmountLess700 {
				previousIncomeGross = previousIncomeNet / taxAuthorityCodebook.CoefficientLess700
				helper := math.Round(previousIncomeGross*100) / 100
				previousIncomeGross = float64(helper)
			} else if previousIncomeNet > taxAuthorityCodebook.AmountLess1000 {
				previousIncomeGross = (taxAuthorityCodebook.AmountLess700/taxAuthorityCodebook.CoefficientLess700 +
					(taxAuthorityCodebook.AmountLess1000-taxAuthorityCodebook.AmountLess700)/taxAuthorityCodebook.CoefficientLess1000 +
					(previousIncomeNet-taxAuthorityCodebook.AmountLess1000)/taxAuthorityCodebook.CoefficientMore1000)
				helper := math.Round(previousIncomeGross*100) / 100
				previousIncomeGross = float64(helper)
			} else {
				previousIncomeGross = (taxAuthorityCodebook.AmountLess700/taxAuthorityCodebook.CoefficientLess700 +
					(previousIncomeNet-taxAuthorityCodebook.AmountLess700)/taxAuthorityCodebook.CoefficientLess1000)
				helper := math.Round(previousIncomeGross*100) / 100
				previousIncomeGross = float64(helper)
			}
		}

	}

	grossPrice, grossPriceOK := params.Args["gross_price"].(float64)

	if !grossPriceOK {
		netPrice, netPriceOK := params.Args["net_price"].(float64)

		//konvertuje neto u bruto
		if !netPriceOK {
			err := errors.APIError{Message: "you must provide price"}
			return errors.HandleAPIError(err)
		}

		if taxAuthorityCodebook.TaxPercentage != 0 {
			taxAuthorityCodebook.Coefficient = calculateCoefficient(*taxAuthorityCodebook, float64(municipality.TaxPercentage))
			grossPrice = netPrice / taxAuthorityCodebook.Coefficient
			helper := math.Round(grossPrice*100) / 100
			grossPrice = float64(helper)
		} else if previousIncomeNetOK {
			if !previousIncomeGrossOK {
				sumNetPrice := previousIncomeNet + netPrice

				//konvertuje neto u bruto
				if sumNetPrice < taxAuthorityCodebook.AmountLess700 {
					sumNetPrice = sumNetPrice / taxAuthorityCodebook.CoefficientLess700
					helper := math.Round(sumNetPrice*100) / 100
					sumNetPrice = float64(helper)
				} else if sumNetPrice > taxAuthorityCodebook.AmountLess1000 {
					sumNetPrice = (taxAuthorityCodebook.AmountLess700/taxAuthorityCodebook.CoefficientLess700 +
						(taxAuthorityCodebook.AmountLess1000-taxAuthorityCodebook.AmountLess700)/taxAuthorityCodebook.CoefficientLess1000 +
						(sumNetPrice-taxAuthorityCodebook.AmountLess1000)/taxAuthorityCodebook.CoefficientMore1000)
					helper := math.Round(sumNetPrice*100) / 100
					sumNetPrice = float64(helper)
				} else {
					sumNetPrice = (taxAuthorityCodebook.AmountLess700/taxAuthorityCodebook.CoefficientLess700 +
						(sumNetPrice-taxAuthorityCodebook.AmountLess700)/taxAuthorityCodebook.CoefficientLess1000)
					helper := math.Round(sumNetPrice*100) / 100
					sumNetPrice = float64(helper)
				}
				grossPrice = sumNetPrice - previousIncomeGross
			}
		} else {
			if netPrice < taxAuthorityCodebook.AmountLess700 {
				grossPrice = netPrice / taxAuthorityCodebook.CoefficientLess700
				helper := math.Round(grossPrice*100) / 100
				grossPrice = float64(helper)
			} else if netPrice > taxAuthorityCodebook.AmountLess1000 {
				grossPrice = (taxAuthorityCodebook.AmountLess700/taxAuthorityCodebook.CoefficientLess700 +
					(taxAuthorityCodebook.AmountLess1000-taxAuthorityCodebook.AmountLess700)/taxAuthorityCodebook.CoefficientLess1000 +
					(netPrice-taxAuthorityCodebook.AmountLess1000)/taxAuthorityCodebook.CoefficientMore1000)
				helper := math.Round(grossPrice*100) / 100
				grossPrice = float64(helper)
			} else {
				grossPrice = (taxAuthorityCodebook.AmountLess700/taxAuthorityCodebook.CoefficientLess700 +
					(netPrice-taxAuthorityCodebook.AmountLess700)/taxAuthorityCodebook.CoefficientLess1000)
				helper := math.Round(grossPrice*100) / 100
				grossPrice = float64(helper)
			}
		}

	}

	additionalExpenses, err := calculateAdditionalExpenses(*taxAuthorityCodebook, grossPrice, previousIncomeGross, r, organizationUnit, *municipality)

	/*if !taxAuthorityCodebook.IncludeSubtax {
		additionalExpenses[len(additionalExpenses)-1].Price -= additionalExpenses[1].Price
	}*/

	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   additionalExpenses,
		//	Total:   total,
	}, nil
}

func calculateAdditionalExpenses(taxAuthorityCodebook structs.TaxAuthorityCodebook, grossPrice float64, previousIncomeGross float64, r *Resolver, organizationUnit *structs.OrganizationUnits, municipality structs.Suppliers) ([]dto.AdditionalExpensesResponse, error) {
	var additionalExpenses []dto.AdditionalExpensesResponse

	nonReleasedGrossPrice := grossPrice

	//oslobodjenje
	if taxAuthorityCodebook.ReleasePercentage != 0 {
		grossPrice = grossPrice - grossPrice*taxAuthorityCodebook.ReleasePercentage/100
		helper := math.Round(grossPrice*100) / 100
		grossPrice = float64(helper)
	}

	var taxPrice float64
	//porez
	if taxAuthorityCodebook.TaxPercentage != 0 {

		taxPrice = grossPrice * taxAuthorityCodebook.TaxPercentage / 100
		helper := math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		taxSupplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.TaxSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Porez",
			Price:  float32(taxPrice),
			Status: string(structs.AdditionalExpenseStatusCreated),
			OrganizationUnit: dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			},
			Subject: dto.DropdownSimple{
				ID:    taxSupplier.ID,
				Title: taxSupplier.Title,
			},
		}

		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	if taxAuthorityCodebook.PreviousIncomePercentageLessThan700 != 0 || taxAuthorityCodebook.PreviousIncomePercentageLessThan1000 != 0 || taxAuthorityCodebook.PreviousIncomePercentageMoreThan1000 != 0 {

		taxSupplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.TaxSupplierID)

		if err != nil {
			return nil, err
		}

		remainGross := grossPrice
		helpGross := grossPrice + previousIncomeGross

		firstGross := helpGross - 1000
		if previousIncomeGross > 1000 {
			firstGross -= previousIncomeGross - 1000
		}

		if firstGross > 0 {
			taxPrice = firstGross * taxAuthorityCodebook.PreviousIncomePercentageMoreThan1000 / 100
			remainGross -= firstGross
			helper := math.Round(taxPrice*100) / 100
			taxPrice = float64(helper)
		}

		secondGross := remainGross - 700

		if secondGross < 0 {
			secondGross = 0
		} else {
			remainGross -= secondGross
		}

		taxPrice += secondGross * taxAuthorityCodebook.PreviousIncomePercentageLessThan1000 / 100

		helper := math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		taxPrice += remainGross * taxAuthorityCodebook.PreviousIncomePercentageLessThan700 / 100

		helper = math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Porez",
			Price:  float32(taxPrice),
			Status: string(structs.AdditionalExpenseStatusCreated),
			OrganizationUnit: dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			},
			Subject: dto.DropdownSimple{
				ID:    taxSupplier.ID,
				Title: taxSupplier.Title,
			},
		}

		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	subTaxPrice := float64(taxPrice) * float64(municipality.TaxPercentage/100)
	helper := math.Round(subTaxPrice*100) / 100
	subTaxPrice = float64(helper)

	additionalExpenseTax := dto.AdditionalExpensesResponse{
		Title:  "Prirez",
		Price:  float32(subTaxPrice),
		Status: string(structs.AdditionalExpenseStatusCreated),
		OrganizationUnit: dto.DropdownSimple{
			ID:    organizationUnit.ID,
			Title: organizationUnit.Title,
		},
		Subject: dto.DropdownSimple{
			ID:    municipality.ID,
			Title: municipality.Title,
		},
	}

	additionalExpenses = append(additionalExpenses, additionalExpenseTax)

	//fond rada
	if taxAuthorityCodebook.LaborFund != 0 {
		taxPrice := grossPrice * taxAuthorityCodebook.LaborFund / 100

		helper := math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.LaborFundSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Fond rada",
			Price:  float32(taxPrice),
			Status: string(structs.AdditionalExpenseStatusCreated),
			OrganizationUnit: dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			},
			Subject: dto.DropdownSimple{
				ID:    supplier.ID,
				Title: supplier.Title,
			},
		}

		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	//pio
	if taxAuthorityCodebook.PioPercentage != 0 {
		taxPrice := grossPrice * taxAuthorityCodebook.PioPercentage / 100

		helper := math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.PioSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "PIO",
			Price:  float32(taxPrice),
			Status: string(structs.AdditionalExpenseStatusCreated),
			OrganizationUnit: dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			},
			Subject: dto.DropdownSimple{
				ID:    supplier.ID,
				Title: supplier.Title,
			},
		}

		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	//pio na teret zaposlenog
	if taxAuthorityCodebook.PioPercentageEmployeePercentage != 0 {
		taxPrice := grossPrice * taxAuthorityCodebook.PioPercentageEmployerPercentage / 100

		helper := math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.PioSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "PIO na teret zaposlenog",
			Price:  float32(taxPrice),
			Status: string(structs.AdditionalExpenseStatusCreated),
			OrganizationUnit: dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			},
			Subject: dto.DropdownSimple{
				ID:    supplier.ID,
				Title: supplier.Title,
			},
		}

		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	//pio na teret poslodavca
	if taxAuthorityCodebook.PioPercentageEmployerPercentage != 0 {
		taxPrice := grossPrice * taxAuthorityCodebook.PioPercentageEmployeePercentage / 100

		helper := math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.PioSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "PIO na teret poslodavca",
			Price:  float32(taxPrice),
			Status: string(structs.AdditionalExpenseStatusCreated),
			OrganizationUnit: dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			},
			Subject: dto.DropdownSimple{
				ID:    supplier.ID,
				Title: supplier.Title,
			},
		}

		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	//za nezaposlenost
	if taxAuthorityCodebook.UnemploymentPercentage != 0 {
		taxPrice := grossPrice * taxAuthorityCodebook.UnemploymentPercentage / 100

		helper := math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.UnemploymentSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Nezaposlenost",
			Price:  float32(taxPrice),
			Status: string(structs.AdditionalExpenseStatusCreated),
			OrganizationUnit: dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			},
			Subject: dto.DropdownSimple{
				ID:    supplier.ID,
				Title: supplier.Title,
			},
		}

		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	//za nezaposlenost na teret poslodavca
	if taxAuthorityCodebook.UnemploymentEmployerPercentage != 0 {
		taxPrice := grossPrice * taxAuthorityCodebook.UnemploymentEmployerPercentage / 100

		helper := math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.UnemploymentSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Nezaposlenost na teret poslodavca",
			Price:  float32(taxPrice),
			Status: string(structs.AdditionalExpenseStatusCreated),
			OrganizationUnit: dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			},
			Subject: dto.DropdownSimple{
				ID:    supplier.ID,
				Title: supplier.Title,
			},
		}

		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	//za nezaposlenost na teret zaposlenog
	if taxAuthorityCodebook.UnemploymentEmployeePercentage != 0 {
		taxPrice := grossPrice * taxAuthorityCodebook.UnemploymentEmployeePercentage / 100

		helper := math.Round(taxPrice*100) / 100
		taxPrice = float64(helper)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.UnemploymentSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Nezaposlenost na teret zaposlenog",
			Price:  float32(taxPrice),
			Status: string(structs.AdditionalExpenseStatusCreated),
			OrganizationUnit: dto.DropdownSimple{
				ID:    organizationUnit.ID,
				Title: organizationUnit.Title,
			},
			Subject: dto.DropdownSimple{
				ID:    supplier.ID,
				Title: supplier.Title,
			},
		}

		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	for _, item := range additionalExpenses {
		//ako prirez ne ide na teret poslodavca, ili je nesto na teret poslodavca, ne pravi razliku izmedju bruto i neto
		if (item.Title == "Prirez" && taxAuthorityCodebook.IncludeSubtax) ||
			(item.Title == "Nezaposlenost na teret poslodavca") || (item.Title == "PIO na teret poslodavca") || (item.Title == "Fond rada") {
			nonReleasedGrossPrice -= 0
		} else {
			nonReleasedGrossPrice -= float64(item.Price)
		}
	}

	additionalExpenseTax = dto.AdditionalExpensesResponse{
		Title:  "Neto",
		Price:  float32(nonReleasedGrossPrice),
		Status: string(structs.AdditionalExpenseStatusCreated),
		OrganizationUnit: dto.DropdownSimple{
			ID:    organizationUnit.ID,
			Title: organizationUnit.Title,
		},
		/*	ime subjekta kojem se uplacuje
			Subject: dto.DropdownSimple{
				ID:    supplier.ID,
				Title: supplier.Title,
			},*/
	}

	additionalExpenses = append(additionalExpenses, additionalExpenseTax)

	return additionalExpenses, nil
}

func buildInvoiceResponseItemList(ctx context.Context, r *Resolver, itemList []structs.Invoice) ([]*dto.InvoiceResponseItem, error) {
	var items []*dto.InvoiceResponseItem

	for _, item := range itemList {
		singleItem, err := buildInvoiceResponseItem(ctx, r, item)

		if err != nil {
			return nil, err
		}

		items = append(items, singleItem)

	}

	return items, nil
}

func buildInvoiceResponseItem(ctx context.Context, r *Resolver, invoice structs.Invoice) (*dto.InvoiceResponseItem, error) {

	response := dto.InvoiceResponseItem{
		ID:                    invoice.ID,
		PassedToInventory:     invoice.PassedToInventory,
		PassedToAccounting:    invoice.PassedToAccounting,
		IsInvoice:             invoice.IsInvoice,
		Issuer:                invoice.Issuer,
		InvoiceNumber:         invoice.InvoiceNumber,
		Type:                  invoice.Type,
		SupplierTitle:         invoice.Supplier,
		DateOfStart:           invoice.DateOfStart,
		Status:                invoice.Status,
		GrossPrice:            invoice.GrossPrice,
		VATPrice:              invoice.VATPrice,
		NetPrice:              invoice.NetPrice,
		OrderID:               invoice.OrderID,
		ProFormaInvoiceDate:   invoice.ProFormaInvoiceDate,
		ProFormaInvoiceNumber: invoice.ProFormaInvoiceNumber,
		DateOfInvoice:         invoice.DateOfInvoice,
		ReceiptDate:           invoice.ReceiptDate,
		DateOfPayment:         invoice.DateOfPayment,
		SSSInvoiceReceiptDate: invoice.SSSInvoiceReceiptDate,
		BankAccount:           invoice.BankAccount,
		Description:           invoice.Description,
		SourceOfFunding:       invoice.SourceOfFunding,
		CreatedAt:             invoice.CreatedAt,
		UpdatedAt:             invoice.UpdatedAt,
	}

	var defaultTime time.Time
	if response.ProFormaInvoiceDate != nil && *response.ProFormaInvoiceDate == defaultTime {
		response.ProFormaInvoiceDate = nil
	}

	if response.DateOfInvoice != nil && *response.DateOfInvoice == defaultTime {
		response.DateOfInvoice = nil
	}

	if response.DateOfPayment != nil && *response.DateOfPayment == defaultTime {
		response.DateOfPayment = nil
	}

	if response.DateOfStart != nil && *response.DateOfStart == defaultTime {
		response.DateOfStart = nil
	}

	if response.ReceiptDate != nil && *response.ReceiptDate == defaultTime {
		response.ReceiptDate = nil
	}

	if invoice.SupplierID != 0 {
		supplier, err := r.Repo.GetSupplier(invoice.SupplierID)

		if err != nil {
			return nil, err
		}

		supplierDropdown := dto.DropdownSimple{
			ID:    supplier.ID,
			Title: supplier.Title,
		}

		response.Supplier = supplierDropdown
	}

	if invoice.MunicipalityID != 0 {
		supplier, err := r.Repo.GetSupplier(invoice.MunicipalityID)

		if err != nil {
			return nil, err
		}

		supplierDropdown := dto.DropdownSimple{
			ID:    supplier.ID,
			Title: supplier.Title,
		}

		response.Municipality = supplierDropdown
	}

	if invoice.OrganizationUnitID != 0 {
		organizationUnit, err := r.Repo.GetOrganizationUnitByID(invoice.OrganizationUnitID)

		if err != nil {
			return nil, err
		}

		OUDropdown := dto.DropdownSimple{
			ID:    organizationUnit.ID,
			Title: organizationUnit.Title,
		}

		response.OrganizationUnit = OUDropdown
	}

	if invoice.TaxAuthorityCodebookID != 0 {
		TaxAuthorityCodebook, err := r.Repo.GetTaxAuthorityCodebookByID(invoice.TaxAuthorityCodebookID)

		if err != nil {
			return nil, err
		}

		OUDropdown := dto.DropdownSimple{
			ID:    TaxAuthorityCodebook.ID,
			Title: TaxAuthorityCodebook.Title,
		}

		response.TaxAuthorityCodebook = OUDropdown
	}

	if invoice.FileID != 0 {
		file, err := r.Repo.GetFileByID(invoice.FileID)

		if err != nil {
			return nil, err
		}

		FileDropdown := dto.FileDropdownSimple{
			ID:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		}

		response.File = FileDropdown
	}

	if invoice.ProFormaInvoiceFileID != 0 {
		file, err := r.Repo.GetFileByID(invoice.ProFormaInvoiceFileID)

		if err != nil {
			return nil, err
		}

		FileDropdown := dto.FileDropdownSimple{
			ID:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		}

		response.ProFormaInvoiceFile = FileDropdown
	}

	if invoice.TypeOfSubject != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(invoice.TypeOfSubject)
		if err != nil {
			return nil, err
		}
		dropdown := dto.DropdownSimple{
			ID:    setting.ID,
			Title: setting.Title,
		}
		response.TypeOfSubject = dropdown
	}

	if invoice.TypeOfDecision != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(invoice.TypeOfDecision)
		if err != nil {
			return nil, err
		}
		dropdown := dto.DropdownSimple{
			ID:    setting.ID,
			Title: setting.Title,
		}
		response.TypeOfDecision = dropdown
	}

	if invoice.TypeOfContract != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(invoice.TypeOfContract)
		if err != nil {
			return nil, err
		}
		dropdown := dto.DropdownSimple{
			ID:    setting.ID,
			Title: setting.Title,
		}
		response.TypeOfContract = dropdown
	}

	if invoice.ActivityID != 0 {
		activity, err := r.Repo.GetActivity(invoice.ActivityID)
		if err != nil {
			return nil, err
		}
		dropdown := dto.DropdownSimple{
			ID:    activity.ID,
			Title: activity.Title,
		}
		response.Activity = dropdown
	}

	if invoice.OrderID != 0 {
		order, err := r.Repo.GetOrderListByID(invoice.OrderID)

		if err != nil {
			return nil, err
		}
		dropdown := dto.DropdownSimple{
			ID: order.ID,
		}

		if order.InvoiceNumber != nil {
			dropdown.Title = *order.InvoiceNumber
		} else {
			dropdown.Title = order.ProFormaInvoiceNumber
		}

		response.Order = dropdown

	}

	articles, err := r.Repo.GetInvoiceArticleList(invoice.ID)

	if err != nil {
		return nil, err
	}

	if len(articles) > 0 {

		response.NetPrice = 0
		response.VATPrice = 0

		for _, article := range articles {
			singleArticle, err := buildInvoiceArtice(r, article)

			if err != nil {
				return nil, err
			}

			response.Articles = append(response.Articles, *singleArticle)
			response.NetPrice += singleArticle.NetPrice
			response.VATPrice += singleArticle.VatPrice
		}
	}
	for _, item := range invoice.AdditionalExpenses {
		builtItem, err := buildAdditionalExpense(r, item)

		if err != nil {
			return nil, err
		}

		response.AdditionalExpenses = append(response.AdditionalExpenses, *builtItem)
	}

	return &response, nil
}

func buildInvoiceArtice(r *Resolver, article structs.InvoiceArticles) (*dto.InvoiceArticleResponse, error) {
	response := dto.InvoiceArticleResponse{
		ID:            article.ID,
		Title:         article.Title,
		NetPrice:      article.NetPrice,
		VatPrice:      article.VatPrice,
		VatPercentage: article.VatPercentage,
		Description:   article.Description,
		Amount:        article.Amount,
		CreatedAt:     article.CreatedAt,
		UpdatedAt:     article.UpdatedAt,
	}

	if article.AccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(article.AccountID)

		if err != nil {
			return nil, err
		}

		accountDropdown := dto.DropdownSimple{
			ID:    account.ID,
			Title: account.Title,
		}
		response.Account = accountDropdown
	}

	if article.CostAccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(article.CostAccountID)

		if err != nil {
			return nil, err
		}

		accountDropdown := dto.DropdownSimple{
			ID:    account.ID,
			Title: account.Title,
		}
		response.CostAccount = accountDropdown
	}

	response.VatPrice = response.NetPrice * float64(response.VatPercentage) / 100

	return &response, nil
}

func buildAdditionalExpenseItemList(ctx context.Context, r *Resolver, itemList []structs.AdditionalExpenses) ([]*dto.AdditionalExpensesResponse, error) {
	var items []*dto.AdditionalExpensesResponse

	for _, item := range itemList {
		singleItem, err := buildAdditionalExpense(r, item)

		if err != nil {
			return nil, err
		}

		items = append(items, singleItem)

	}

	return items, nil
}

func buildAdditionalExpense(r *Resolver, item structs.AdditionalExpenses) (*dto.AdditionalExpensesResponse, error) {
	response := dto.AdditionalExpensesResponse{
		ID:          item.ID,
		Title:       item.Title,
		Price:       item.Price,
		BankAccount: item.BankAccount,
		Status:      item.Status,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}

	if item.AccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(item.AccountID)

		if err != nil {
			return nil, err
		}

		response.Account = dto.DropdownSimple{
			ID:    account.ID,
			Title: account.Title,
		}
	}

	if item.InvoiceID != 0 {
		invoice, err := r.Repo.GetInvoice(item.InvoiceID)
		if err != nil {
			return nil, err
		}

		response.Invoice = dto.DropdownSimple{
			ID:    invoice.ID,
			Title: invoice.InvoiceNumber,
		}
	}

	if item.SubjectID != 0 {
		supplier, err := r.Repo.GetSupplier(item.SubjectID)
		if err != nil {
			return nil, err
		}

		response.Subject = dto.DropdownSimple{
			ID:    supplier.ID,
			Title: supplier.Title,
		}
	}

	if item.OrganizationUnitID != 0 {
		orgUnit, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)
		if err != nil {
			return nil, err
		}

		response.OrganizationUnit = dto.DropdownSimple{
			ID:    orgUnit.ID,
			Title: orgUnit.Title,
		}
	}
	return &response, nil
}
