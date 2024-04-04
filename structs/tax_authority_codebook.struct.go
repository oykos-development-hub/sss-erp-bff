package structs

import "time"

type TaxAuthorityCodebook struct {
	ID                                   int       `json:"id"`
	Title                                string    `json:"title"`
	Code                                 string    `json:"code"`
	Active                               bool      `json:"active"`
	TaxPercentage                        float64   `json:"tax_percentage"`
	PioPercentage                        float64   `json:"pio_percentage"`
	ReleasePercentage                    float64   `json:"release_percentage"`
	PreviousIncomePercentageLessThan700  float64   `json:"previous_income_percentage_less_than_700"`
	PreviousIncomePercentageLessThan1000 float64   `json:"previous_income_percentage_less_than_1000"`
	PreviousIncomePercentageMoreThan1000 float64   `json:"previous_income_percentage_more_than_1000"`
	CreatedAt                            time.Time `json:"created_at"`
	UpdatedAt                            time.Time `json:"updated_at"`
}
