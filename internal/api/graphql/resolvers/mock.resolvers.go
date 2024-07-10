package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"

	"github.com/graphql-go/graphql"
	"github.com/shopspring/decimal"
)

func (r *Resolver) CurrentBudgetMockResolver(params graphql.ResolveParams) (interface{}, error) {
	budgets, err := r.Repo.GetBudgetList(nil)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	isParent := true
	units, err := r.Repo.GetOrganizationUnits(&dto.GetOrganizationUnitsInput{IsParent: &isParent})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, budget := range budgets {
		financeBudgetDetails, err := r.Repo.GetFinancialBudgetByBudgetID(budget.ID)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "get financial budget by budget id"))
		}

		accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{
			Leaf:    true,
			Version: &financeBudgetDetails.AccountVersion,
		})
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "get account items"))
		}

		for _, unit := range units.Data {
			for _, account := range accounts.Data {
				_, err := r.Repo.CreateCurrentBudget(params.Context, &structs.CurrentBudget{
					BudgetID:      budget.ID,
					UnitID:        unit.ID,
					AccountID:     account.ID,
					InitialActual: decimal.NewFromFloat(1200),
					Actual:        decimal.NewFromFloat(1200),
					Balance:       decimal.Zero,
				})
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}, nil
}
