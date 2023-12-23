package dto

import (
	"bff/structs"
)

type GetProcurementPlanResponseMS struct {
	Data structs.PublicProcurementPlan `json:"data"`
}

type GetProcurementPlansInput struct {
	Year                   *string `json:"year"`
	IsPreBudget            *bool   `json:"is_pre_budget"`
	TargetBudgetID         *int    `json:"target_budget_id"`
	SortByYear             *string `json:"sort_by_year"`
	SortByTitle            *string `json:"sort_by_title"`
	SortByDateOfPublishing *string `json:"sort_by_date_of_publishing"`
}

type GetProcurementPlanListResponseMS struct {
	Data []*structs.PublicProcurementPlan `json:"data"`
}

type PlanStatus string

const (
	PlanStatusNotAccessible      PlanStatus = "Nedostupan"
	PlanStatusAdminInProggress   PlanStatus = "U toku"
	PlanStatusAdminPublished     PlanStatus = "Poslat"
	PlanStatusUserPublished      PlanStatus = "Obradi"
	PlanStatusUserRequested      PlanStatus = "Na čekanju"
	PlanStatusUserAccepted       PlanStatus = "Odobren"
	PlanStatusUserRejected       PlanStatus = "Odbijen"
	PlanStatusPreBudgetClosed    PlanStatus = "Zaključen"
	PlanStatusPreBudgetConverted PlanStatus = "Konvertovan"
	PlanStatusPostBudgetClosed   PlanStatus = "Objavljen"
)

type ProcurementPlanResponseItem struct {
	ID                  int                            `json:"id"`
	PreBudgetPlan       *DropdownSimple                `json:"pre_budget_plan"`
	IsPreBudget         bool                           `json:"is_pre_budget"`
	Active              bool                           `json:"active"`
	Year                string                         `json:"year"`
	Title               string                         `json:"title"`
	Status              PlanStatus                     `json:"status"`
	SerialNumber        *string                        `json:"serial_number"`
	DateOfPublishing    *string                        `json:"date_of_publishing"`
	DateOfClosing       *string                        `json:"date_of_closing"`
	PreBudgetID         *int                           `json:"pre_budget_id"`
	FileID              *int                           `json:"file_id"`
	Requests            int                            `json:"requests"`
	ApprovedRequests    int                            `json:"approved_requests"`
	RejectedDescription *string                        `json:"rejected_description"`
	Items               []*ProcurementItemResponseItem `json:"items"`
	TotalNet            float32                        `json:"total_net"`
	TotalGross          float32                        `json:"total_gross"`
	CreatedAt           string                         `json:"created_at"`
	UpdatedAt           string                         `json:"updated_at"`
}

type PlanPDFResponse struct {
	PlanID        string                `json:"plan_id"`
	Year          string                `json:"year"`
	PublishedDate string                `json:"published_date"`
	TotalGross    string                `json:"total_gross"`
	TotalVAT      string                `json:"total_vat"`
	TableData     []PlanPDFTableDataRow `json:"table_data"`
}

type PlanPDFTableDataRow struct {
	ID              string `json:"id"`
	ArticleType     string `json:"article_type"`
	Title           string `json:"title"`
	TotalGross      string `json:"total_gross"`
	TotalVAT        string `json:"total_vat"`
	TypeOfProcedure string `json:"type_of_procedure"`
	BudgetIndent    string `json:"budget_indent"`
	FundingSource   string `json:"funding_source"`
}
