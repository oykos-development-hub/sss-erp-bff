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
		return errors.HandleAPIError(err)
	}

	units, err := r.Repo.GetOrganizationUnits(nil)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	accounts, err := r.Repo.GetAccountItems(nil)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, budget := range budgets {
		for _, unit := range units.Data {
			for _, account := range accounts.Data {
				_, err := r.Repo.CreateCurrentBudget(params.Context, &structs.CurrentBudget{
					BudgetID:       budget.ID,
					UnitID:         unit.ID,
					AccountID:      account.ID,
					InititalActual: decimal.NewFromFloat(1200),
					Actual:         decimal.NewFromFloat(1200),
					Balance:        decimal.Zero,
				})
				if err != nil {
					return errors.HandleAPIError(err)
				}
			}
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}, nil
}
