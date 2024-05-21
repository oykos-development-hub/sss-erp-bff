package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	goerrors "errors"

	"github.com/graphql-go/graphql"
	"github.com/shopspring/decimal"
)

func (r *Resolver) SpendingDynamicOverview(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)
	history := params.Args["history"].(bool)

	if unitID == 0 {
		loggedInOrganizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok {
			return errors.HandleAPPError(errors.NewBadRequestError("Error getting logged in unit"))
		}

		unitID = *loggedInOrganizationUnitID
	}

	var (
		spendingDynamic *structs.SpendingDynamic
		err             error
	)

	if history {
		spendingDynamic, err = r.Repo.GetSpendingDynamicHistory(budgetID, unitID)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic"))
		}
	} else {
		spendingDynamic, err = r.Repo.GetSpendingDynamic(budgetID, unitID)
		if err != nil {
			var apiErr *errors.APIError
			if goerrors.As(err, &apiErr) {
				if apiErr.StatusCode == 404 {
					spendingDynamic, err = r.generatePreparedSpendingDynamicData(budgetID, unitID)
					if err != nil {
						return errors.HandleAPPError(err)
					}
				} else {
					return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic history"))
				}
			}
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the data you asked for!",
		Item:    spendingDynamic,
	}, nil
}

func (r *Resolver) generatePreparedSpendingDynamicData(budgetID, unitID int) (*structs.SpendingDynamic, error) {
	actual, err := r.Repo.GetSpendingDynamicActual(budgetID, unitID)
	if err != nil {
		return nil, errors.WrapBadRequestError(err, "budget has no actual yet")
	}
	if !actual.Valid {
		return nil, errors.NewBadRequestError("budget has no actual yet")
	}

	monthlyAmount := actual.Decimal.Div(decimal.NewFromInt(12)).Round(2)

	// Sum of the first 11 rounded months
	totalForFirst11Months := monthlyAmount.Mul(decimal.NewFromInt(11))

	// Adjust the December amount to account for rounding differences
	decemberAmount := actual.Decimal.Sub(totalForFirst11Months).Round(2)

	spendingDynamic := structs.SpendingDynamic{
		BudgetID: budgetID,
		UnitID:   unitID,
		Actual:   actual.Decimal,
		Entries: []structs.SpendingDynamicEntry{
			{
				January:   monthlyAmount,
				February:  monthlyAmount,
				March:     monthlyAmount,
				April:     monthlyAmount,
				May:       monthlyAmount,
				June:      monthlyAmount,
				July:      monthlyAmount,
				August:    monthlyAmount,
				September: monthlyAmount,
				October:   monthlyAmount,
				November:  monthlyAmount,
				December:  decemberAmount,
			},
		},
	}

	return &spendingDynamic, nil
}
