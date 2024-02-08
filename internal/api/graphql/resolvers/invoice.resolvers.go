package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"

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
	} else {
		item, err = r.Repo.UpdateInvoice(&data)
		if err != nil {
			return errors.HandleAPIError(err)
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

	err := r.Repo.DeleteInvoice(itemID)
	if err != nil {
		fmt.Printf("Deleting invoice item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
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
		InvoiceNumber:         invoice.InvoiceNumber,
		Status:                invoice.Status,
		GrossPrice:            invoice.GrossPrice,
		VATPrice:              invoice.VATPrice,
		OrderID:               invoice.OrderID,
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

	articles, err := r.Repo.GetInvoiceArticleList(invoice.ID)

	if err != nil {
		return nil, err
	}

	for _, article := range articles {
		singleArticle, err := buildInvoiceArtice(r, article)

		if err != nil {
			return nil, err
		}

		response.Articles = append(response.Articles, *singleArticle)
	}

	return &response, nil
}

func buildInvoiceArtice(r *Resolver, article structs.InvoiceArticles) (*dto.InvoiceArticleResponse, error) {
	response := dto.InvoiceArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		NetPrice:    article.NetPrice,
		VatPrice:    article.VatPrice,
		Description: article.Description,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
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

	return &response, nil
}
