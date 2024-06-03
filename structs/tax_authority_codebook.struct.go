package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type TaxAuthorityCodebook struct {
	ID                                   int             `json:"id"`
	Title                                string          `json:"title"`
	Code                                 string          `json:"code"`
	Active                               bool            `json:"active"`
	TaxPercentage                        decimal.Decimal `json:"tax_percentage"`
	TaxSupplierID                        int             `json:"tax_supplier_id"`
	ReleasePercentage                    decimal.Decimal `json:"release_percentage"`
	ReleaseAmount                        decimal.Decimal `json:"release_amount"`
	PioPercentage                        decimal.Decimal `json:"pio_percentage"`
	PioSupplierID                        int             `json:"pio_supplier_id"`
	PioPercentageEmployerPercentage      decimal.Decimal `json:"pio_percentage_employer_percentage"`
	PioEmployerSupplierID                int             `json:"pio_employer_supplier_id"`
	PioPercentageEmployeePercentage      decimal.Decimal `json:"pio_percentage_employee_percentage"`
	PioEmployeeSupplierID                int             `json:"pio_employee_supplier_id"`
	UnemploymentPercentage               decimal.Decimal `json:"unemployment_percentage"`
	UnemploymentSupplierID               int             `json:"unemployment_supplier_id"`
	UnemploymentEmployerPercentage       decimal.Decimal `json:"unemployment_employer_percentage"`
	UnemploymentEmployerSupplierID       int             `json:"unemployment_employer_supplier_id"`
	UnemploymentEmployeePercentage       decimal.Decimal `json:"unemployment_employee_percentage"`
	UnemploymentEmployeeSupplierID       int             `json:"unemployment_employee_supplier_id"`
	LaborFund                            decimal.Decimal `json:"labor_fund"`
	LaborFundSupplierID                  int             `json:"labor_fund_supplier_id"`
	PreviousIncomePercentageLessThan700  decimal.Decimal `json:"previous_income_percentage_less_than_700"`
	PreviousIncomePercentageLessThan1000 decimal.Decimal `json:"previous_income_percentage_less_than_1000"`
	PreviousIncomePercentageMoreThan1000 decimal.Decimal `json:"previous_income_percentage_more_than_1000"`
	Coefficient                          decimal.Decimal `json:"coefficient"`
	CoefficientLess700                   decimal.Decimal `json:"coefficient_less_700"`
	CoefficientLess1000                  decimal.Decimal `json:"coefficient_less_1000"`
	CoefficientMore1000                  decimal.Decimal `json:"coefficient_more_1000"`
	AmountLess700                        decimal.Decimal `json:"amount_less_700"`
	AmountLess1000                       decimal.Decimal `json:"amount_less_1000"`
	AmountMore1000                       decimal.Decimal `json:"amount_more_1000"`
	IncludeSubtax                        bool            `json:"include_subtax"`
	CreatedAt                            time.Time       `json:"created_at"`
	UpdatedAt                            time.Time       `json:"updated_at"`
}
