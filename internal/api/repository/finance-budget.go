package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

func (repo *MicroserviceRepository) CreateBudget(ctx context.Context, budgetItem *structs.Budget) (*structs.Budget, error) {
	res := &dto.GetBudgetResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Budget, budgetItem, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateBudget(ctx context.Context, item *structs.Budget) (*structs.Budget, error) {
	res := &dto.GetBudgetResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetBudgetList(input *dto.GetBudgetListInputMS) ([]structs.Budget, error) {
	res := &dto.GetBudgetListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Budget, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetBudget(id int) (*structs.Budget, error) {
	res := &dto.GetBudgetResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteBudget(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) CreateFinancialBudget(ctx context.Context, financialBudget *structs.FinancialBudget) (*structs.FinancialBudget, error) {
	res := &dto.GetFinancialBudgetResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FinancialBudget, financialBudget, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateFinancialBudget(ctx context.Context, financialBudget *structs.FinancialBudget) (*structs.FinancialBudget, error) {
	res := &dto.GetFinancialBudgetResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FinancialBudget+"/"+strconv.Itoa(financialBudget.ID), financialBudget, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFinancialBudgetByBudgetID(id int) (*structs.FinancialBudget, error) {
	res := &dto.GetFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(id)+"/financial", nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateBudgetLimit(budgetLimit *structs.FinancialBudgetLimit) (*structs.FinancialBudgetLimit, error) {
	res := &dto.GetFinancialBudgetLimitResponseMS{}

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FinancialBudgetLimit, budgetLimit, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateBudgetLimit(budgetLimit *structs.FinancialBudgetLimit) (*structs.FinancialBudgetLimit, error) {
	res := &dto.GetFinancialBudgetLimitResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FinancialBudgetLimit+"/"+strconv.Itoa(budgetLimit.ID), budgetLimit, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteBudgetLimit(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FinancialBudgetLimit+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetBudgetLimits(budgetID int) ([]structs.FinancialBudgetLimit, error) {
	input := dto.GetFinancialBudgetListInputMS{
		BudgetID: budgetID,
	}
	res := &dto.GetFinancialBudgetLimitListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FinancialBudgetLimit, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetBudgetUnitLimit(budgetID, unitID int) (int, error) {
	input := dto.GetFinancialBudgetListInputMS{
		BudgetID: budgetID,
		UnitID:   &unitID,
	}
	res := &dto.GetFinancialBudgetLimitListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FinancialBudgetLimit, input, res)
	if err != nil {
		return 0, errors.Wrap(err, "make api request")
	}

	if len(res.Data) != 1 {
		return 0, nil
	}

	return res.Data[0].Limit, nil
}

func (repo *MicroserviceRepository) GetFilledFinancialBudgetList(input *dto.FilledFinancialBudgetInputMS) ([]structs.FilledFinanceBudget, error) {
	res := &dto.GetFilledFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FilledFinancialBudget, input, res)
	if err != nil {
		return nil, errors.WrapInternalServerError(err, "repo.GetFilledFinancialBudgetList")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetFinancialBudgetByID(id int) (*structs.FinancialBudget, error) {
	res := &dto.GetFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FinancialBudget+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) FillFinancialBudget(ctx context.Context, data *structs.FilledFinanceBudget) (*structs.FilledFinanceBudget, error) {
	res := &dto.GetFilledFinancialBudgetResponseItemMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FilledFinancialBudget, data, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateFilledFinancialBudget(ctx context.Context, id int, data *structs.FilledFinanceBudget) (*structs.FilledFinanceBudget, error) {
	res := &dto.GetFilledFinancialBudgetResponseItemMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FilledFinancialBudget+"/"+strconv.Itoa(id), data, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

// prosledjuje se accountID i budget request ID
func (repo *MicroserviceRepository) FillActualFinancialBudget(ctx context.Context, id int, actual decimal.Decimal, requestID int) (*structs.FilledFinanceBudget, error) {
	data := &dto.FillActualFinanceBudgetInput{Actual: actual}
	res := &dto.GetFilledFinancialBudgetResponseItemMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PATCH", repo.Config.Microservices.Finance.FilledFinancialBudget+"/"+strconv.Itoa(id)+"/actual"+strconv.Itoa(requestID), data, res, header)
	if err != nil {
		return nil, errors.WrapMicroserviceError(err, "FillActualFinancialBudget")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteFilledFinancialBudgetData(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FilledFinancialBudget+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetFinancialFilledSummary(budgetID int, reqType structs.RequestType) ([]structs.FilledFinanceBudget, error) {
	res := &dto.GetFilledFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("GET", fmt.Sprintf("%s/%d/filled-financial-summary/%d", repo.Config.Microservices.Finance.Budget, budgetID, reqType), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetCurrentBudgetByOrganizationUnit(organizationUnitID int) ([]structs.CurrentBudget, error) {
	res := &dto.GetCurrentBudgetListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetCurrentBudgetByOrganizationUnit+"/"+strconv.Itoa(organizationUnitID), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}
