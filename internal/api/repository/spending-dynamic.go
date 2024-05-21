package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"fmt"

	"github.com/shopspring/decimal"
)

func (repo *MicroserviceRepository) GetSpendingDynamic(BudgetID, unitID int) (*structs.SpendingDynamic, error) {
	res := dto.GetSpendingDynamicResponseMS{}
	_, err := makeAPIRequest("GET", fmt.Sprintf(repo.Config.Microservices.Finance.SpendingDynamicGet, BudgetID, unitID), nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetSpendingDynamicHistory(BudgetID, unitID int) (*structs.SpendingDynamic, error) {
	res := dto.GetSpendingDynamicResponseMS{}
	_, err := makeAPIRequest("GET", fmt.Sprintf(repo.Config.Microservices.Finance.SpendingDynamicGetHistory, BudgetID, unitID), nil, &res)
	path := fmt.Sprintf(repo.Config.Microservices.Finance.SpendingDynamicGetHistory, BudgetID, unitID)
	fmt.Println(path)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateSpendingDynamic(resolution *structs.PublicProcurementPlan) (*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.Plans, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetSpendingDynamicActual(BudgetID, unitID int) (decimal.NullDecimal, error) {
	res := dto.GetSpendingDynamicActualResponseMS{}
	_, err := makeAPIRequest("GET", fmt.Sprintf(repo.Config.Microservices.Finance.SpendingDynamicActual, BudgetID, unitID), nil, &res)
	if err != nil {
		return res.Data, err
	}

	return res.Data, nil
}
