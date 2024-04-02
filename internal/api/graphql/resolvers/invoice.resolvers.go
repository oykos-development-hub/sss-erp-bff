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

		var defaultTime time.Time

		if invoice.DateOfInvoice == defaultTime && data.DateOfInvoice != defaultTime && invoice.InvoiceNumber == "" && data.InvoiceNumber != "" && data.OrderID != 0 {
			order, err := r.Repo.GetOrderListByID(data.OrderID)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			invoiceDate := data.DateOfInvoice.Format("2006-01-02T15:04:05Z")
			order.InvoiceDate = &invoiceDate
			order.InvoiceNumber = &data.InvoiceNumber

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
	price := params.Args["price"].(float64)
	previousIncome, previousIncomeOK := params.Args["previous_income"].(float64)

	if !previousIncomeOK {
		return errors.HandleAPIError(err)
	}

	fullPrice := price + previousIncome

	//ceka se uputstvo kolege lalovica za ovo
	taxPrice := fullPrice * float64(taxAuthorityCodebook.Percentage) / 100

	var additionalExpenses []dto.AdditionalExpensesResponse

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return errors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	organizationUnit, err := r.Repo.GetOrganizationUnitByID(*organizationUnitID)

	if !previousIncomeOK {
		return errors.HandleAPIError(err)
	}

	supplier, err := getTaxAuthority(r)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	additionalExpenseTax := dto.AdditionalExpensesResponse{
		Title:  "Porez",
		Price:  float32(taxPrice),
		Status: structs.AdditionalExpenseStatusCreated,
		OrganizationUnit: dto.DropdownSimple{
			ID:    organizationUnit.ID,
			Title: organizationUnit.Title,
		},
		Subject: *supplier,
	}

	additionalExpenses = append(additionalExpenses, additionalExpenseTax)

	municipality, err := r.Repo.GetDropdownSettingByID(municipalityID)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	additionalPercentage, err := strconv.Atoi(municipality.Value)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	additionalExpenseAdditionalTax := dto.AdditionalExpensesResponse{
		Title:  "Prirez",
		Price:  float32(taxPrice) * float32(additionalPercentage) / 100,
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

	additionalExpenses[0].Price = additionalExpenses[0].Price - float32(taxPrice)*float32(additionalPercentage)/100

	additionalExpenses = append(additionalExpenses, additionalExpenseAdditionalTax)

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   additionalExpenses,
		//	Total:   total,
	}, nil
}

func getTaxAuthority(r *Resolver) (*dto.DropdownSimple, error) {
	taxAuthority := "Poreska uprava"
	supplier, err := r.Repo.GetSupplierList(&dto.GetSupplierInputMS{Search: &taxAuthority})

	if err != nil {
		return nil, err
	}

	if len(supplier.Data) > 0 {
		return &dto.DropdownSimple{
			ID:    supplier.Data[0].ID,
			Title: supplier.Data[0].Title,
		}, err
	}

	taxAuthority = "Uprava prihoda"
	supplier, err = r.Repo.GetSupplierList(&dto.GetSupplierInputMS{Search: &taxAuthority})

	if err != nil {
		return nil, err
	}

	if len(supplier.Data) > 0 {
		return &dto.DropdownSimple{
			ID:    supplier.Data[0].ID,
			Title: supplier.Data[0].Title,
		}, err
	}

	return nil, errors.APIError{Message: "you_must_add_tax_service_in_codebook"}
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
