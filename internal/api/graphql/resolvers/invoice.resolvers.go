package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"
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

	if value, ok := params.Args["status"].(int); ok && value != 0 {
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

func (r *Resolver) CalculateAdditionalExpensesResolver(params graphql.ResolveParams) (interface{}, error) {
	taxAuthorityCodebookID := params.Args["tax_authority_codebook_id"].(int)

	taxAuthorityCodebook, err := r.Repo.GetTaxAuthorityCodebookByID(taxAuthorityCodebookID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	municipalityID := params.Args["municipality_id"].(int)

	previousIncomeGross, previousIncomeGrossOK := params.Args["previous_income_gross"].(float64)

	if !previousIncomeGrossOK {
		previousIncomeNet, previousIncomeNetOK := params.Args["previous_income_net"].(float64)

		//konvertuje neto u bruto
		if previousIncomeNetOK && taxAuthorityCodebook.Coefficient != 0 {
			previousIncomeGross = previousIncomeNet / taxAuthorityCodebook.Coefficient
		} else if previousIncomeNetOK && taxAuthorityCodebook.Coefficient == 0 {
			if previousIncomeNet > 818 {
				previousIncomeGross = (previousIncomeNet - 123) * taxAuthorityCodebook.CoefficientLess700
			} else if previousIncomeNet > 591.51 {
				previousIncomeGross = (previousIncomeNet - 63) * taxAuthorityCodebook.CoefficientLess1000
			} else {
				previousIncomeGross = previousIncomeNet * taxAuthorityCodebook.CoefficientMore1000
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

		if taxAuthorityCodebook.Coefficient != 0 {
			grossPrice = netPrice / taxAuthorityCodebook.Coefficient
		} else {
			if netPrice > 818 {
				grossPrice = (netPrice - 123) * taxAuthorityCodebook.CoefficientLess700
			} else if netPrice > 591.51 {
				grossPrice = (netPrice - 63) * taxAuthorityCodebook.CoefficientLess1000
			} else {
				grossPrice = netPrice * taxAuthorityCodebook.CoefficientMore1000
			}
		}

	}

	municipality, err := r.Repo.GetSupplier(municipalityID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

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

	var additionalExpenses []dto.AdditionalExpensesResponse

	//porez
	if taxAuthorityCodebook.TaxPercentage != 0 {

		taxPrice := grossPrice * taxAuthorityCodebook.TaxPercentage / 100

		taxSupplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.TaxSupplierID)

		if err != nil {
			return errors.HandleAPIError(errors.APIError{Message: "tax supplier is not valid"})
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Porez",
			Price:  float32(taxPrice),
			Status: structs.AdditionalExpenseStatusCreated,
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
			return errors.HandleAPIError(errors.APIError{Message: "tax supplier is not valid"})
		}

		if previousIncomeGross < 700 {
			taxPrice := grossPrice * taxAuthorityCodebook.PreviousIncomePercentageLessThan700 / 100

			additionalExpenseTax := dto.AdditionalExpensesResponse{
				Title:  "Porez",
				Price:  float32(taxPrice),
				Status: structs.AdditionalExpenseStatusCreated,
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
		} else if previousIncomeGross > 700 && previousIncomeGross < 1000 {
			taxPrice := grossPrice * taxAuthorityCodebook.PreviousIncomePercentageLessThan1000 / 100

			additionalExpenseTax := dto.AdditionalExpensesResponse{
				Title:  "Porez",
				Price:  float32(taxPrice),
				Status: structs.AdditionalExpenseStatusCreated,
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
		} else {
			taxPrice := grossPrice * taxAuthorityCodebook.PreviousIncomePercentageMoreThan1000 / 100

			additionalExpenseTax := dto.AdditionalExpensesResponse{
				Title:  "Porez",
				Price:  float32(taxPrice),
				Status: structs.AdditionalExpenseStatusCreated,
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

	}

	//fond rada
	if taxAuthorityCodebook.LaborFund != 0 {
		taxPrice := grossPrice * taxAuthorityCodebook.LaborFund / 100

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.LaborFundSupplierID)

		if err != nil {
			return errors.HandleAPIError(errors.APIError{Message: "labor fund supplier is not valid"})
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Fond rada",
			Price:  float32(taxPrice),
			Status: structs.AdditionalExpenseStatusCreated,
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

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.PioSupplierID)

		if err != nil {
			return errors.HandleAPIError(errors.APIError{Message: "pio supplier is not valid"})
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "PIO",
			Price:  float32(taxPrice),
			Status: structs.AdditionalExpenseStatusCreated,
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
		taxPrice := grossPrice * taxAuthorityCodebook.PioPercentageEmployeePercentage / 100

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.PioSupplierID)

		if err != nil {
			return errors.HandleAPIError(errors.APIError{Message: "pio supplier is not valid"})
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "PIO na teret zaposlenog",
			Price:  float32(taxPrice),
			Status: structs.AdditionalExpenseStatusCreated,
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
		taxPrice := grossPrice * taxAuthorityCodebook.PioPercentageEmployerPercentage / 100

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.PioSupplierID)

		if err != nil {
			return errors.HandleAPIError(errors.APIError{Message: "pio supplier is not valid"})
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "PIO na teret poslodavca",
			Price:  float32(taxPrice),
			Status: structs.AdditionalExpenseStatusCreated,
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

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.UnemploymentSupplierID)

		if err != nil {
			return errors.HandleAPIError(errors.APIError{Message: "unemployment supplier is not valid"})
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Nezaposlenost",
			Price:  float32(taxPrice),
			Status: structs.AdditionalExpenseStatusCreated,
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

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.UnemploymentSupplierID)

		if err != nil {
			return errors.HandleAPIError(errors.APIError{Message: "unemployment supplier is not valid"})
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Nezaposlenost na teret poslodavca",
			Price:  float32(taxPrice),
			Status: structs.AdditionalExpenseStatusCreated,
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

		supplier, err := r.Repo.GetSupplier(taxAuthorityCodebook.UnemploymentSupplierID)

		if err != nil {
			return errors.HandleAPIError(errors.APIError{Message: "unemployment supplier is not valid"})
		}

		additionalExpenseTax := dto.AdditionalExpensesResponse{
			Title:  "Nezaposlenost na teret zaposlenog",
			Price:  float32(taxPrice),
			Status: structs.AdditionalExpenseStatusCreated,
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

	var taxValue float32
	for _, item := range additionalExpenses {
		taxValue += item.Price
	}

	taxPrice := taxValue * municipality.TaxPercentage / 100

	additionalExpenseTax := dto.AdditionalExpensesResponse{
		Title:  "Prirez",
		Price:  float32(taxPrice),
		Status: structs.AdditionalExpenseStatusCreated,
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

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   additionalExpenses,
		//	Total:   total,
	}, nil
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
		InvoiceNumber:         invoice.InvoiceNumber,
		Type:                  invoice.Type,
		SupplierTitle:         invoice.Supplier,
		DateOfStart:           invoice.DateOfStart,
		Status:                invoice.Status,
		GrossPrice:            invoice.GrossPrice,
		VATPrice:              invoice.VATPrice,
		OrderID:               invoice.OrderID,
		ProFormaInvoiceDate:   invoice.ProFormaInvoiceDate,
		ProFormaInvoiceNumber: invoice.ProFormaInvoiceNumber,
		DateOfInvoice:         invoice.DateOfInvoice,
		ReceiptDate:           invoice.ReceiptDate,
		DateOfPayment:         invoice.DateOfPayment,
		SSSInvoiceReceiptDate: invoice.SSSInvoiceReceiptDate,
		BankAccount:           invoice.BankAccount,
		Description:           invoice.Description,
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
			Title: TaxAuthorityCodebook.Code,
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

	if invoice.SourceOfFunding != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(invoice.SourceOfFunding)
		if err != nil {
			return nil, err
		}
		dropdown := dto.DropdownSimple{
			ID:    setting.ID,
			Title: setting.Entity,
		}
		response.SourceOfFunding = dropdown
	}

	if invoice.TypeOfSubject != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(invoice.TypeOfSubject)
		if err != nil {
			return nil, err
		}
		dropdown := dto.DropdownSimple{
			ID:    setting.ID,
			Title: setting.Entity,
		}
		response.TypeOfSubject = dropdown
	}

	if invoice.TypeOfContract != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(invoice.TypeOfContract)
		if err != nil {
			return nil, err
		}
		dropdown := dto.DropdownSimple{
			ID:    setting.ID,
			Title: setting.Entity,
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

	articles, err := r.Repo.GetInvoiceArticleList(invoice.ID)

	if err != nil {
		return nil, err
	}

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
