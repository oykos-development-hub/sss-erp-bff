package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/shopspring/decimal"
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
	var message string
	if total == 0 {
		message = "There aren`t item!"
	} else {
		message = "Here's the list you asked for!"
	}

	return dto.Response{
		Status:  "success",
		Message: message,
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

	isCreated := true
	for _, article := range data.Articles {
		if article.AccountID == 0 {
			isCreated = false
		}
	}

	if isCreated {
		data.Status = "Kreiran"
	} else {
		data.Status = "Nepotpun"
	}

	var orderID int
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

			data.OrderID = order.ID

			var orderArticles []structs.OrderArticleInsertItem

			for _, article := range data.Articles {

				float32Value := float32(article.NetPrice.InexactFloat64())

				orderArticles = append(orderArticles, structs.OrderArticleInsertItem{
					Amount:        article.Amount,
					Title:         article.Title,
					Description:   article.Description,
					NetPrice:      float32Value,
					VatPercentage: article.VatPercentage,
				})
			}

			err = r.Repo.CreateOrderListProcurementArticles(orderID, structs.OrderListInsertItem{
				ID:       orderID,
				Articles: orderArticles,
			})
			if err != nil {
				return errors.HandleAPIError(err)
			}

		}

		item, err = r.Repo.CreateInvoice(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {

		invoice, err := r.Repo.GetInvoice(data.ID)
		//orderID = invoice.OrderID
		if err != nil {
			return errors.HandleAPIError(err)
		}

		item, err = r.Repo.UpdateInvoice(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		if invoice.OrderID != 0 {
			order, err := r.Repo.GetOrderListByID(invoice.OrderID)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			var invoiceDate string
			if data.DateOfInvoice != nil {
				invoiceDate = data.DateOfInvoice.Format("2006-01-02T15:04:05Z")
				order.InvoiceDate = &invoiceDate
			}
			order.InvoiceNumber = &data.InvoiceNumber
			order.ReceiveFile = []int{data.ProFormaInvoiceFileID}

			_, err = r.Repo.UpdateOrderListItem(order.ID, order)

			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	}

	/*if orderID != 0 {
		order, err := r.Repo.GetOrderListByID(orderID)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		_, err = r.Repo.UpdateOrderListItem(orderID, &structs.OrderListItem{
			ID:                    order.ID,
			DateOrder:             order.DateOrder,
			TotalPrice:            order.TotalPrice,
			PublicProcurementID:   order.PublicProcurementID,
			GroupOfArticlesID:     order.GroupOfArticlesID,
			SupplierID:            order.SupplierID,
			Status:                order.Status,
			PassedToFinance:       true,
			UsedInFinance:         true,
			DateSystem:            order.DateSystem,
			InvoiceDate:           order.InvoiceDate,
			InvoiceNumber:         order.InvoiceNumber,
			OrganizationUnitID:    order.OrganizationUnitID,
			OfficeID:              order.OfficeID,
			RecipientUserID:       order.RecipientUserID,
			Description:           order.Description,
			IsUsed:                order.IsUsed,
			OrderFile:             order.OrderFile,
			ReceiveFile:           order.ReceiveFile,
			MovementFile:          order.MovementFile,
			IsProFormaInvoice:     order.IsProFormaInvoice,
			ProFormaInvoiceDate:   order.ProFormaInvoiceDate,
			ProFormaInvoiceNumber: order.ProFormaInvoiceNumber,
			AccountID:             order.AccountID,
		})

		if err != nil {
			return errors.HandleAPIError(err)
		}

		//nije dobro brisanje artikala zbog javnih nabavki i racunanja preostalih kolicina
		/*getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
			OrderID: &order.ID,
		}

		articles, err := r.Repo.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
		if err != nil {
			return nil, err
		}

		for _, article := range articles.Data {
			err := r.Repo.DeleteOrderProcurementArticle(article.ID)
			if err != nil {
				return errors.HandleAPIError(err)

			}
		}

	}*/

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

func calculateCoefficient(item structs.TaxAuthorityCodebook, subTax decimal.Decimal, netPrice decimal.Decimal) decimal.Decimal {
	var coefficient decimal.Decimal

	coefficient = decimal.NewFromInt(1)
	release := item.ReleasePercentage

	if item.ReleaseAmount.Cmp(decimal.NewFromInt(0)) != 0 {
		release = item.ReleaseAmount.Div(netPrice)
	}

	zero := decimal.NewFromInt(0)
	hundred := decimal.NewFromInt(100)

	if release.Cmp(zero) != 0 {
		if item.TaxPercentage.Cmp(zero) != 0 {
			coefficient = coefficient.Sub(
				decimal.NewFromFloat(1).Sub(release.Div(hundred)).Mul(item.TaxPercentage.Div(hundred)),
			)
			if subTax.Cmp(zero) != 0 {
				coefficient = coefficient.Sub(
					decimal.NewFromFloat(1).Sub(release.Div(hundred)).Mul(item.TaxPercentage.Div(hundred)).Mul(subTax.Div(hundred)),
				)
			}
		}
		if item.PioPercentage.Cmp(zero) != 0 {
			coefficient = coefficient.Sub(
				decimal.NewFromFloat(1).Sub(release.Div(hundred)).Mul(item.PioPercentage.Div(hundred)),
			)
		}
	} else {
		if item.TaxPercentage.Cmp(zero) != 0 {
			coefficient = coefficient.Sub(
				decimal.NewFromFloat(1).Sub(item.TaxPercentage.Div(hundred)),
			)
			if subTax.Cmp(zero) != 0 {
				coefficient = coefficient.Sub(
					decimal.NewFromFloat(1).Sub(item.TaxPercentage.Div(hundred).Mul(subTax.Div(hundred))),
				)
			}
		}
		if item.PioPercentage.Cmp(zero) != 0 {
			coefficient = coefficient.Sub(
				decimal.NewFromFloat(1).Sub(item.PioPercentage.Div(hundred)),
			)
		}
	}

	return coefficient
}

func calculateCoefficientLess700(item structs.TaxAuthorityCodebook, previousIncomeGross decimal.Decimal, r *Resolver, organizationUnit *structs.OrganizationUnits, municipality structs.Suppliers) (decimal.Decimal, decimal.Decimal, error) {

	additionalExpenses, err := calculateAdditionalExpenses(item, decimal.NewFromInt(700), previousIncomeGross, r, organizationUnit, municipality)

	if err != nil {
		return decimal.NewFromInt(0), decimal.NewFromInt(0), err
	}

	maxNetAmount := additionalExpenses[len(additionalExpenses)-1].Price

	coefficient := maxNetAmount.Div(decimal.NewFromInt(700))

	return decimal.Decimal(coefficient), decimal.Decimal(maxNetAmount), nil
}

func calculateCoefficientLess1000(item structs.TaxAuthorityCodebook, previousIncomeGross decimal.Decimal, r *Resolver, organizationUnit *structs.OrganizationUnits, municipality structs.Suppliers) (decimal.Decimal, decimal.Decimal, error) {

	additionalExpenses, err := calculateAdditionalExpenses(item, decimal.NewFromInt(1000), previousIncomeGross, r, organizationUnit, municipality)

	if err != nil {
		return decimal.NewFromInt(0), decimal.NewFromInt(0), err
	}

	maxNetAmount := additionalExpenses[len(additionalExpenses)-1].Price
	//if item.IncludeSubtax {
	maxNetAmount = maxNetAmount.Add(additionalExpenses[1].Price)
	//}

	_, amount, err := calculateCoefficientLess700(item, previousIncomeGross, r, organizationUnit, municipality)

	if err != nil {
		return decimal.NewFromInt(0), decimal.NewFromInt(0), err
	}

	coefficient := maxNetAmount.Sub(amount).Div(decimal.NewFromFloat(300))

	return decimal.Decimal(coefficient), decimal.Decimal(maxNetAmount), nil
}

func calculateCoefficientMore1000(item structs.TaxAuthorityCodebook, previousIncomeGross decimal.Decimal, r *Resolver, organizationUnit *structs.OrganizationUnits, municipality structs.Suppliers) (decimal.Decimal, decimal.Decimal, error) {

	coefficient := item.PreviousIncomePercentageMoreThan1000.Add(item.PioPercentage).Add(item.PioPercentageEmployerPercentage).Add(
		item.UnemploymentEmployeePercentage).Add(item.UnemploymentPercentage)

	return decimal.NewFromInt(1).Sub(coefficient.Div(decimal.NewFromInt(100))), decimal.NewFromInt(0), nil
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

	previousIncomeGross, previousIncomeGrossOK := params.Args["previous_income_gross"].(decimal.Decimal)
	previousIncomeNet, previousIncomeNetOK := params.Args["previous_income_net"].(decimal.Decimal)

	taxAuthorityCodebook.CoefficientLess700, taxAuthorityCodebook.AmountLess700, err = calculateCoefficientLess700(*taxAuthorityCodebook, decimal.NewFromInt(0), r, organizationUnit, *municipality)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	taxAuthorityCodebook.CoefficientLess1000, taxAuthorityCodebook.AmountLess1000, err = calculateCoefficientLess1000(*taxAuthorityCodebook, decimal.NewFromInt(0), r, organizationUnit, *municipality)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	taxAuthorityCodebook.CoefficientMore1000, taxAuthorityCodebook.AmountMore1000, err = calculateCoefficientMore1000(*taxAuthorityCodebook, decimal.NewFromInt(0), r, organizationUnit, *municipality)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	if !previousIncomeGrossOK {

		//konvertuje neto u bruto
		if previousIncomeNetOK && taxAuthorityCodebook.TaxPercentage.Cmp(decimal.NewFromInt(0)) != 0 {
			taxAuthorityCodebook.Coefficient = calculateCoefficient(*taxAuthorityCodebook, decimal.NewFromFloat32(municipality.TaxPercentage), previousIncomeNet)
			previousIncomeGross := previousIncomeNet.Div(taxAuthorityCodebook.Coefficient)
			helper := previousIncomeGross.Round(2)
			previousIncomeGross = helper

		} else if previousIncomeNetOK && taxAuthorityCodebook.TaxPercentage.Cmp(decimal.NewFromInt(0)) == 0 {
			if previousIncomeNet.Cmp(taxAuthorityCodebook.AmountLess700) < 0 {
				previousIncomeGross := previousIncomeNet.Div(taxAuthorityCodebook.CoefficientLess700)
				helper := previousIncomeGross.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
				previousIncomeGross = helper
			} else if previousIncomeNet.Cmp(taxAuthorityCodebook.AmountLess1000) > 0 {
				previousIncomeGross := taxAuthorityCodebook.AmountLess700.Div(taxAuthorityCodebook.CoefficientLess700).Add(
					taxAuthorityCodebook.AmountLess1000.Sub(taxAuthorityCodebook.AmountLess700).Div(taxAuthorityCodebook.CoefficientLess1000)).Add(
					previousIncomeNet.Sub(taxAuthorityCodebook.AmountLess1000).Div(taxAuthorityCodebook.CoefficientMore1000))
				helper := previousIncomeGross.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
				previousIncomeGross = helper
			} else {
				previousIncomeGross := taxAuthorityCodebook.AmountLess700.Div(taxAuthorityCodebook.CoefficientLess700).Add(
					previousIncomeNet.Sub(taxAuthorityCodebook.AmountLess700).Div(taxAuthorityCodebook.CoefficientLess1000))
				helper := previousIncomeGross.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
				previousIncomeGross = helper
			}
		}

	}

	grossPrice, grossPriceOK := params.Args["gross_price"].(decimal.Decimal)

	if !grossPriceOK {
		netPrice, netPriceOK := params.Args["net_price"].(decimal.Decimal)

		//konvertuje neto u bruto
		if !netPriceOK {
			err := &errors.APIError{Message: "you must provide price"}
			return errors.HandleAPIError(err)
		}

		if taxAuthorityCodebook.TaxPercentage.Cmp(decimal.NewFromInt(0)) != 0 {
			taxAuthorityCodebook.Coefficient = calculateCoefficient(*taxAuthorityCodebook, decimal.NewFromFloat32(municipality.TaxPercentage), netPrice)
			grossPrice := netPrice.Div(taxAuthorityCodebook.Coefficient)
			helper := grossPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
			grossPrice = helper

			fmt.Println("Gross Price:", grossPrice)
		} else if previousIncomeNetOK {
			if !previousIncomeGrossOK {
				sumNetPrice := previousIncomeNet.Add(netPrice)

				// konvertuje neto u bruto
				if sumNetPrice.Cmp(taxAuthorityCodebook.AmountLess700) < 0 {
					sumNetPrice = sumNetPrice.Div(taxAuthorityCodebook.CoefficientLess700)
					helper := sumNetPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
					sumNetPrice = helper
				} else if sumNetPrice.Cmp(taxAuthorityCodebook.AmountLess1000) > 0 {
					sumNetPrice = taxAuthorityCodebook.AmountLess700.Div(taxAuthorityCodebook.CoefficientLess700).
						Add(taxAuthorityCodebook.AmountLess1000.Sub(taxAuthorityCodebook.AmountLess700).Div(taxAuthorityCodebook.CoefficientLess1000)).
						Add(sumNetPrice.Sub(taxAuthorityCodebook.AmountLess1000).Div(taxAuthorityCodebook.CoefficientMore1000))
					helper := sumNetPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
					sumNetPrice = helper
				} else {
					sumNetPrice = taxAuthorityCodebook.AmountLess700.Div(taxAuthorityCodebook.CoefficientLess700).
						Add(sumNetPrice.Sub(taxAuthorityCodebook.AmountLess700).Div(taxAuthorityCodebook.CoefficientLess1000))
					helper := sumNetPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
					sumNetPrice = helper
				}
				grossPrice = sumNetPrice.Sub(previousIncomeGross)
			} else {
				if netPrice.Cmp(taxAuthorityCodebook.AmountLess700) < 0 {
					grossPrice := netPrice.Div(taxAuthorityCodebook.CoefficientLess700)
					helper := grossPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
					grossPrice = helper
				} else if netPrice.Cmp(taxAuthorityCodebook.AmountLess1000) > 0 {
					grossPrice := taxAuthorityCodebook.AmountLess700.Div(taxAuthorityCodebook.CoefficientLess700).
						Add(taxAuthorityCodebook.AmountLess1000.Sub(taxAuthorityCodebook.AmountLess700).Div(taxAuthorityCodebook.CoefficientLess1000)).
						Add(netPrice.Sub(taxAuthorityCodebook.AmountLess1000).Div(taxAuthorityCodebook.CoefficientMore1000))
					helper := grossPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
					grossPrice = helper
				} else {
					grossPrice := taxAuthorityCodebook.AmountLess700.Div(taxAuthorityCodebook.CoefficientLess700).
						Add(netPrice.Sub(taxAuthorityCodebook.AmountLess700).Div(taxAuthorityCodebook.CoefficientLess1000))
					helper := grossPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
					grossPrice = helper
				}

			}
		}

	}

	additionalExpenses, err := calculateAdditionalExpenses(*taxAuthorityCodebook, grossPrice, previousIncomeGross, r, organizationUnit, *municipality)

	if taxAuthorityCodebook.IncludeSubtax {
		additionalExpenses[len(additionalExpenses)-1].Price = additionalExpenses[len(additionalExpenses)-1].Price.Add(additionalExpenses[1].Price)
	}

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

func calculateAdditionalExpenses(taxAuthorityCodebook structs.TaxAuthorityCodebook, grossPrice decimal.Decimal, previousIncomeGross decimal.Decimal, r *Resolver, organizationUnit *structs.OrganizationUnits, municipality structs.Suppliers) ([]dto.AdditionalExpensesResponse, error) {
	var additionalExpenses []dto.AdditionalExpensesResponse

	nonReleasedGrossPrice := grossPrice

	//oslobodjenje
	if taxAuthorityCodebook.ReleasePercentage.Cmp(decimal.NewFromInt(0)) != 0 {
		grossPrice = grossPrice.Sub(grossPrice.Mul(taxAuthorityCodebook.ReleasePercentage).Div(decimal.NewFromInt(100)))
		helper := grossPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
		grossPrice = helper
	} else if taxAuthorityCodebook.ReleaseAmount.Cmp(decimal.NewFromInt(0)) != 0 {
		grossPrice = grossPrice.Sub(taxAuthorityCodebook.ReleaseAmount)
		helper := grossPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
		grossPrice = helper

		if grossPrice.Cmp(decimal.NewFromInt(0)) < 0 {
			grossPrice = decimal.NewFromInt(0)
		}
	}

	var taxPrice decimal.Decimal
	//porez
	if taxAuthorityCodebook.TaxPercentage.Cmp(decimal.NewFromInt(0)) != 0 {
		taxPrice := grossPrice.Mul(taxAuthorityCodebook.TaxPercentage).Div(decimal.NewFromInt(100))
		helper := taxPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
		taxPrice = helper

		taxSupplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.TaxSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  structs.ObligationTaxTitle,
			Price:  decimal.Decimal(taxPrice),
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
		if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpenses = append(additionalExpenses, additionalExpenseTax)
		}
	}

	if taxAuthorityCodebook.PreviousIncomePercentageLessThan700.Cmp(decimal.NewFromInt(0)) != 0 || taxAuthorityCodebook.PreviousIncomePercentageLessThan1000.Cmp(decimal.NewFromInt(0)) != 0 || taxAuthorityCodebook.PreviousIncomePercentageMoreThan1000.Cmp(decimal.NewFromInt(0)) != 0 {

		taxSupplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.TaxSupplierID)

		if err != nil {
			return nil, err
		}

		remainGross := grossPrice
		helpGross := grossPrice.Add(previousIncomeGross)

		firstGross := helpGross.Sub(decimal.NewFromInt(1000))

		if firstGross.Cmp(decimal.NewFromInt(0)) > 0 {
			taxPrice := firstGross.Mul(taxAuthorityCodebook.PreviousIncomePercentageMoreThan1000).Div(decimal.NewFromInt(100))
			remainGross = remainGross.Sub(firstGross)
			helper := taxPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
			taxPrice = helper
		}

		secondGross := remainGross.Sub(decimal.NewFromInt(700)).Add(previousIncomeGross)

		if firstGross.Cmp(decimal.NewFromInt(0)) > 0 {
			taxPrice := firstGross.Mul(taxAuthorityCodebook.PreviousIncomePercentageMoreThan1000).Div(decimal.NewFromInt(100))
			remainGross = remainGross.Sub(firstGross)
			helper := taxPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
			taxPrice = helper
		}

		if secondGross.Cmp(decimal.NewFromInt(0)) < 0 {
			secondGross = decimal.NewFromInt(0)
		} else {
			remainGross = remainGross.Sub(secondGross)
		}

		taxPrice = taxPrice.Add(secondGross.Mul(taxAuthorityCodebook.PreviousIncomePercentageLessThan1000).Div(decimal.NewFromInt(100)))

		helper := taxPrice.Mul(decimal.NewFromInt(100)).Round(0).Div(decimal.NewFromInt(100))
		taxPrice = helper

		taxPrice = taxPrice.Add(remainGross.Mul(taxAuthorityCodebook.PreviousIncomePercentageLessThan700).Div(decimal.NewFromInt(100)))

		if previousIncomeGross.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpensesForTax, err := calculateAdditionalExpenses(taxAuthorityCodebook, previousIncomeGross, decimal.NewFromInt(0), r, organizationUnit, municipality)

			if err != nil {
				return nil, err
			}

			if len(additionalExpensesForTax) > 0 {
				taxPrice = taxPrice.Sub(additionalExpensesForTax[0].Price)
			}
		}

		helper = taxPrice.Round(2)
		taxPrice = decimal.Decimal(helper)

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  structs.ObligationTaxTitle,
			Price:  decimal.Decimal(taxPrice),
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

		if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpenses = append(additionalExpenses, additionalExpenseTax)
		}
	}

	subTaxPrice := decimal.Decimal(taxPrice).Mul(decimal.NewFromFloat32(municipality.TaxPercentage / 100))
	subTaxPrice = subTaxPrice.Round(2)

	additionalExpenseTax := dto.AdditionalExpensesResponse{
		Title:  structs.ObligationSubTaxTitle,
		Price:  decimal.Decimal(subTaxPrice),
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

	if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

	//fond rada
	if taxAuthorityCodebook.LaborFund.Cmp(decimal.NewFromInt(0)) > 0 {
		taxPrice := grossPrice.Mul(taxAuthorityCodebook.LaborFund).Div(decimal.NewFromInt(100))

		taxPrice = taxPrice.Round(2)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.LaborFundSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  structs.LaborFundTitle,
			Price:  decimal.Decimal(taxPrice),
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

		if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpenses = append(additionalExpenses, additionalExpenseTax)
		}
	}

	//pio
	if taxAuthorityCodebook.PioPercentage.Cmp(decimal.NewFromInt(0)) > 0 {
		taxPrice := grossPrice.Mul(taxAuthorityCodebook.PioPercentage).Div(decimal.NewFromInt(100))

		taxPrice = taxPrice.Round(2)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.PioSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  structs.ContributionForPIOTitle,
			Price:  decimal.Decimal(taxPrice),
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

		if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpenses = append(additionalExpenses, additionalExpenseTax)
		}
	}

	//pio na teret zaposlenog
	if taxAuthorityCodebook.PioPercentageEmployeePercentage.Cmp(decimal.NewFromInt(0)) > 0 {
		taxPrice := grossPrice.Mul(taxAuthorityCodebook.PioPercentageEmployerPercentage).Div(decimal.NewFromInt(100))

		taxPrice = taxPrice.Round(2)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.PioSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  structs.ContributionForPIOEmployeeTitle,
			Price:  decimal.Decimal(taxPrice),
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

		if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpenses = append(additionalExpenses, additionalExpenseTax)
		}
	}

	//pio na teret poslodavca
	if taxAuthorityCodebook.PioPercentageEmployerPercentage.Cmp(decimal.NewFromInt(0)) > 0 {
		taxPrice := grossPrice.Mul(taxAuthorityCodebook.PioPercentageEmployeePercentage).Div(decimal.NewFromInt(100))

		taxPrice = taxPrice.Round(2)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.PioSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  structs.ContributionForPIOEmployerTitle,
			Price:  decimal.Decimal(taxPrice),
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

		if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpenses = append(additionalExpenses, additionalExpenseTax)
		}
	}

	//za nezaposlenost
	if taxAuthorityCodebook.UnemploymentPercentage.Cmp(decimal.NewFromInt(0)) > 0 {
		taxPrice := grossPrice.Mul(taxAuthorityCodebook.UnemploymentPercentage).Div(decimal.NewFromInt(100))

		taxPrice = taxPrice.Round(2)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.UnemploymentSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  structs.ContributionForUnemploymentTitle,
			Price:  decimal.Decimal(taxPrice),
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

		if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpenses = append(additionalExpenses, additionalExpenseTax)
		}
	}

	//za nezaposlenost na teret poslodavca
	if taxAuthorityCodebook.UnemploymentEmployerPercentage.Cmp(decimal.NewFromInt(0)) > 0 {
		taxPrice := grossPrice.Mul(taxAuthorityCodebook.UnemploymentEmployerPercentage).Div(decimal.NewFromInt(100))

		taxPrice = taxPrice.Round(2)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.UnemploymentSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  structs.ContributionForUnemploymentEmployerTitle,
			Price:  decimal.Decimal(taxPrice),
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

		if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpenses = append(additionalExpenses, additionalExpenseTax)
		}
	}

	//za nezaposlenost na teret zaposlenog
	if taxAuthorityCodebook.UnemploymentEmployeePercentage.Cmp(decimal.NewFromInt(0)) > 0 {
		taxPrice := grossPrice.Mul(taxAuthorityCodebook.UnemploymentEmployeePercentage).Div(decimal.NewFromInt(100))

		taxPrice = taxPrice.Round(100)

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.UnemploymentSupplierID)

		if err != nil {
			return nil, err
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  structs.ContributionForUnemploymentEmployeeTitle,
			Price:  decimal.Decimal(taxPrice),
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

		if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
			additionalExpenses = append(additionalExpenses, additionalExpenseTax)
		}
	}

	for _, item := range additionalExpenses {
		//ako prirez ne ide na teret poslodavca, ili je nesto na teret poslodavca, ne pravi razliku izmedju bruto i neto
		if /*(item.Title == "Prirez" && !taxAuthorityCodebook.IncludeSubtax) ||*/
		(item.Title == structs.ContributionForUnemploymentEmployerTitle) || (item.Title == structs.ContributionForPIOEmployerTitle) ||
			(item.Title == structs.LaborFundTitle) {
			nonReleasedGrossPrice = nonReleasedGrossPrice.Sub(decimal.NewFromInt(0))
		} else {
			nonReleasedGrossPrice = nonReleasedGrossPrice.Sub(item.Price)
		}
	}

	additionalExpenseTax = dto.AdditionalExpensesResponse{
		Title:  structs.NetTitle,
		Price:  decimal.Decimal(nonReleasedGrossPrice),
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

	if additionalExpenseTax.Price.Cmp(decimal.NewFromInt(0)) > 0 {
		additionalExpenses = append(additionalExpenses, additionalExpenseTax)
	}

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
		ID:                            invoice.ID,
		PassedToInventory:             invoice.PassedToInventory,
		PassedToAccounting:            invoice.PassedToAccounting,
		IsInvoice:                     invoice.IsInvoice,
		Issuer:                        invoice.Issuer,
		InvoiceNumber:                 invoice.InvoiceNumber,
		Type:                          invoice.Type,
		SupplierTitle:                 invoice.Supplier,
		DateOfStart:                   invoice.DateOfStart,
		Status:                        invoice.Status,
		GrossPrice:                    invoice.GrossPrice,
		VATPrice:                      invoice.VATPrice,
		NetPrice:                      invoice.NetPrice,
		OrderID:                       invoice.OrderID,
		ProFormaInvoiceDate:           invoice.ProFormaInvoiceDate,
		ProFormaInvoiceNumber:         invoice.ProFormaInvoiceNumber,
		DateOfInvoice:                 invoice.DateOfInvoice,
		ReceiptDate:                   invoice.ReceiptDate,
		DateOfPayment:                 invoice.DateOfPayment,
		SSSInvoiceReceiptDate:         invoice.SSSInvoiceReceiptDate,
		SSSProFormaInvoiceReceiptDate: invoice.SSSProFormaInvoiceReceiptDate,
		BankAccount:                   invoice.BankAccount,
		Description:                   invoice.Description,
		SourceOfFunding:               invoice.SourceOfFunding,
		Registred:                     invoice.Registred,
		CreatedAt:                     invoice.CreatedAt,
		UpdatedAt:                     invoice.UpdatedAt,
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
		setting, err := r.Repo.GetSupplier(invoice.TypeOfDecision)
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

		response.NetPrice = decimal.NewFromInt(0)
		response.VATPrice = decimal.NewFromInt(0)

		for _, article := range articles {
			singleArticle, err := buildInvoiceArtice(r, article)

			if err != nil {
				return nil, err
			}

			response.Articles = append(response.Articles, *singleArticle)

			amountDecimal := decimal.NewFromInt(int64(singleArticle.Amount))
			response.NetPrice = response.NetPrice.Add(singleArticle.NetPrice.Mul(amountDecimal))
			response.VATPrice = response.VATPrice.Add(singleArticle.VatPrice.Mul(amountDecimal))
		}

		accountMap := make(map[string]decimal.Decimal)

		for _, item := range response.Articles {
			amountDecimal := decimal.NewFromInt(int64(item.Amount))
			vatPercentageDecimal := decimal.NewFromInt(int64(item.VatPercentage))
			netPriceWithVAT := item.NetPrice.Add(item.NetPrice.Mul(vatPercentageDecimal).Div(decimal.NewFromInt(100)))
			if currentAmount, exists := accountMap[item.Account.Title]; exists {
				accountMap[item.Account.Title] = currentAmount.Add(amountDecimal.Mul(netPriceWithVAT))
			} else {
				accountMap[item.Account.Title] = amountDecimal.Mul(netPriceWithVAT)
			}
		}

		for title, amount := range accountMap {
			response.AccountAmounts = append(response.AccountAmounts, dto.AccountAmounts{
				Account: title,
				Amount:  amount,
			})
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

	response.VatPrice = response.NetPrice.Mul(decimal.NewFromInt32(int32(response.VatPercentage))).Div(decimal.NewFromInt(100))

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
		ID:               item.ID,
		Title:            item.Title,
		Price:            item.Price,
		BankAccount:      item.BankAccount,
		Status:           item.Status,
		ObligationType:   item.ObligationType,
		ObligationNumber: item.ObligationNumber,
		CreatedAt:        item.CreatedAt,
		UpdatedAt:        item.UpdatedAt,
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

	if item.ObligationSupplierID != 0 {
		supplier, err := r.Repo.GetSupplier(item.ObligationSupplierID)
		if err != nil {
			return nil, err
		}

		response.ObligationSupplier = dto.DropdownSimple{
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
