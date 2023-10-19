package dto

import (
	"bff/structs"
)

type GetProcurementPlanResponseMS struct {
	Data structs.PublicProcurementPlan `json:"data"`
}

type GetProcurementPlansInput struct {
	Year           *string `json:"year"`
	IsPreBudget    *bool   `json:"is_pre_budget"`
	TargetBudgetID *int    `json:"target_budget_id"`
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
	Id                  int                            `json:"id"`
	PreBudgetPlan       *DropdownSimple                `json:"pre_budget_plan"`
	IsPreBudget         bool                           `json:"is_pre_budget"`
	Active              bool                           `json:"active"`
	Year                string                         `json:"year"`
	Title               string                         `json:"title"`
	Status              PlanStatus                     `json:"status"`
	SerialNumber        *string                        `json:"serial_number"`
	DateOfPublishing    *string                        `json:"date_of_publishing"`
	DateOfClosing       *string                        `json:"date_of_closing"`
	PreBudgetId         *int                           `json:"pre_budget_id"`
	FileId              *int                           `json:"file_id"`
	Requests            int                            `json:"requests"`
	RejectedDescription *string                        `json:"rejected_description"`
	Items               []*ProcurementItemResponseItem `json:"items"`
	CreatedAt           string                         `json:"created_at"`
	UpdatedAt           string                         `json:"updated_at"`
}
