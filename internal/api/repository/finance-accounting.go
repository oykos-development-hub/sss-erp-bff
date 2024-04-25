package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
)

func (repo *MicroserviceRepository) GetAllObligationsForAccounting(input dto.ObligationsFilter) ([]dto.ObligationForAccounting, int, error) {
	res := &dto.GetObligationsForAccountingResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetObligationsForAccounting, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) BuildAccountingOrderForObligations(data structs.AccountingOrderForObligationsData) (*dto.AccountingOrderForObligations, error) {
	res := &dto.GetAccountingOrderForObligations{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetObligationsForAccounting, data, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
