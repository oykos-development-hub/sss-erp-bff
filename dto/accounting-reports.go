package dto

type OveralSpendingFilter struct {
	Year               *string `json:"year"`
	Title              *string `json:"title"`
	OfficeID           *int    `json:"office_id"`
	Exception          *bool   `json:"exception"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
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
