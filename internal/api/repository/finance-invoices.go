package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateInvoice(ctx context.Context, item *structs.Invoice) (*structs.Invoice, error) {
	res := &dto.GetInvoiceResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Invoice, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateInvoice(ctx context.Context, item *structs.Invoice) (*structs.Invoice, error) {
	res := &dto.GetInvoiceResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Invoice+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetInvoiceList(input *dto.GetInvoiceListInputMS) ([]structs.Invoice, int, error) {
	res := &dto.GetInvoiceListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Invoice, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) GetInvoice(id int) (*structs.Invoice, error) {
	res := &dto.GetInvoiceResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Invoice+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteInvoice(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Invoice+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) DeleteInvoiceArticle(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.InvoiceArticle+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetInvoiceArticleList(id int) ([]structs.InvoiceArticles, error) {
	res := &dto.GetInvoiceArticleListResponseMS{}

	filter := &dto.InvoiceArticleFilterDTO{
		InvoiceID: &id,
	}

	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.InvoiceArticle, filter, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateInvoiceArticle(article *structs.InvoiceArticles) (*structs.InvoiceArticles, error) {
	res := &dto.GetInvoiceArticleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.InvoiceArticle, article, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateInvoiceArticle(article *structs.InvoiceArticles) (*structs.InvoiceArticles, error) {
	res := &dto.GetInvoiceArticleResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.InvoiceArticle+"/"+strconv.Itoa(article.ID), article, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetAdditionalExpenses(input *dto.AdditionalExpensesListInputMS) ([]structs.AdditionalExpenses, int, error) {
	res := &dto.GetAdditionalExpensesListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.AdditionalExpenses, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}
