package dto

import (
	"bff/structs"
)

type GetProcurementPlanResponseMS struct {
	Data structs.PublicProcurementPlan `json:"data"`
}

type GetProcurementPlansInput struct {
	Status         *bool   `json:"is_active"`
	Year           *string `json:"year"`
	IsPreBudget    *bool   `json:"is_pre_budget"`
	TargetBudgetID *int    `json:"target_budget_id"`
}

type GetProcurementPlanListResponseMS struct {
	Data []*structs.PublicProcurementPlan `json:"data"`
}

type ProcurementPlanResponseItem struct {
	Id               int                      `json:"id"`
	PreBudgetPlan    structs.SettingsDropdown `json:"pre_budget_plan"`
	IsPreBudget      bool                     `json:"is_pre_budget"`
	Active           bool                     `json:"active"`
	Year             string                   `json:"year"`
	Title            string                   `json:"title"`
	Status           *string                  `json:"status"`
	SerialNumber     *string                  `json:"serial_number"`
	DateOfPublishing *string                  `json:"date_of_publishing"`
	DateOfClosing    *string                  `json:"date_of_closing"`
	PreBudgetId      *int                     `json:"pre_budget_id"`
	FileId           *int                     `json:"file_id"`
	// @TODO
	Items     []*ProcurementItemResponseItem `json:"items"`
	CreatedAt string                         `json:"created_at"`
	UpdatedAt string                         `json:"updated_at"`
}
