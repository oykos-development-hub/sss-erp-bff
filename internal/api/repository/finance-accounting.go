package repository

import (
	"bff/internal/api/dto"
)

func (repo *MicroserviceRepository) GetAllObligationsForAccounting(input dto.ObligationsFilter) ([]dto.ObligationForAccounting, int, error) {
	res := &dto.GetObligationsForAccountingResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetObligationsForAccounting, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}
