package structs

import "time"

type TaxAuthorityCodebook struct {
	ID                                   int       `json:"id"`
	Title                                string    `json:"title"`
	Code                                 string    `json:"code"`
	Active                               bool      `json:"active"`
	TaxPercentage                        float64   `json:"tax_percentage"`
	TaxSupplierID                        int       `json:"tax_supplier_id"`
	ReleasePercentage                    float64   `json:"release_percentage"`
	PioPercentage                        float64   `json:"pio_percentage"`
	PioSupplierID                        int       `json:"pio_supplier_id"`
	PioPercentageEmployerPercentage      float64   `json:"pio_percentage_employer_percentage"`
	PioEmployerSupplierID                int       `json:"pio_employer_supplier_id"`
	PioPercentageEmployeePercentage      float64   `json:"pio_percentage_employee_percentage"`
	PioEmployeeSupplierID                int       `json:"pio_employee_supplier_id"`
	UnemploymentPercentage               float64   `json:"unemployment_percentage"`
	UnemploymentSupplierID               int       `json:"unemployment_supplier_id"`
	UnemploymentEmployerPercentage       float64   `json:"unemployment_employer_percentage"`
	UnemploymentEmployerSupplierID       int       `json:"unemployment_employer_supplier_id"`
	UnemploymentEmployeePercentage       float64   `json:"unemployment_employee_percentage"`
	UnemploymentEmployeeSupplierID       int       `json:"unemployment_employee_supplier_id"`
	LaborFund                            float64   `json:"labor_fund"`
	LaborFundSupplierID                  int       `json:"labor_fund_supplier_id"`
	PreviousIncomePercentageLessThan700  float64   `json:"previous_income_percentage_less_than_700"`
	PreviousIncomePercentageLessThan1000 float64   `json:"previous_income_percentage_less_than_1000"`
	PreviousIncomePercentageMoreThan1000 float64   `json:"previous_income_percentage_more_than_1000"`
	Coefficient                          float64   `json:"coefficient"`
	CoefficientLess700                   float64   `json:"coefficient_less_700"`
	CoefficientLess1000                  float64   `json:"coefficient_less_1000"`
	CoefficientMore1000                  float64   `json:"coefficient_more_1000"`
	CreatedAt                            time.Time `json:"created_at"`
	UpdatedAt                            time.Time `json:"updated_at"`
}
