package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) GetAllObligationsForAccounting(input dto.ObligationsFilter) ([]dto.ObligationForAccounting, int, error) {
	res := &dto.GetObligationsForAccountingResponseMS{}

	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetObligationsForAccounting, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) GetAllPaymentOrdersForAccounting(input dto.ObligationsFilter) ([]dto.PaymentOrdersForAccounting, int, error) {
	res := &dto.GetPaymentOrdersForAccountingResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetPaymentOrdersForAccounting, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) GetAllEnforcedPaymentsForAccounting(input dto.ObligationsFilter) ([]dto.PaymentOrdersForAccounting, int, error) {
	res := &dto.GetPaymentOrdersForAccountingResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetEnforcedPaymentsForAccounting, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) GetAllReturnedEnforcedPaymentsForAccounting(input dto.ObligationsFilter) ([]dto.PaymentOrdersForAccounting, int, error) {
	res := &dto.GetPaymentOrdersForAccountingResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetReturnedEnforcedPaymentsForAccounting, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) BuildAccountingOrderForObligations(data structs.AccountingOrderForObligationsData) (*dto.AccountingOrderForObligations, error) {
	res := &dto.GetAccountingOrderForObligations{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.BuildAccountingOrderForObligations, data, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateAccountingEntry(ctx context.Context, item *structs.AccountingEntry) (*structs.AccountingEntry, error) {
	res := &dto.GetAccountingEntryResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.AccountingEntry, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetAccountingEntryByID(id int) (*structs.AccountingEntry, error) {
	res := &dto.GetAccountingEntryResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.AccountingEntry+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetAccountingEntryList(filter dto.AccountingEntryFilter) ([]structs.AccountingEntry, int, error) {
	res := &dto.GetAccountingEntryListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.AccountingEntry, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteAccountingEntry(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.AccountingEntry+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetAnalyticalCard(input dto.AnalyticalCardFilter) ([]structs.AnalyticalCard, error) {
	res := &dto.GetAnalyticalCardListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.AnalyticalCard, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
