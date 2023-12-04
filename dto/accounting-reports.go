package dto

type OveralSpendingFilter struct {
	StartDate          *string  `json:"start_date"`
	EndDate            *string  `json:"end_date"`
	Title              *string  `json:"title"`
	OfficeID           *int     `json:"office_id"`
	Exception          *bool    `json:"exception"`
	OrganizationUnitID *int     `json:"organization_unit_id"`
	Articles           []string `json:"articles"`
}

type ArticleReport struct {
	Year        string `json:"year"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	OfficeID    int    `json:"office_id"`
}

type ArticleReportMS struct {
	Data []ArticleReport `json:"data"`
}
