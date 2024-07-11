package dto

import (
	"bff/structs"
	"time"

	"github.com/shopspring/decimal"
)

type GetSpendingReleaseResponseMS struct {
	Data structs.SpendingRelease `json:"data"`
}

type GetSpendingReleaseListResponseMS struct {
	Data []structs.SpendingRelease `json:"data"`
}

type GetSpendingReleaseRequestListResponseMS struct {
	Data []structs.SpendingReleaseRequest `json:"data"`
}

type GetSpendingReleaseListInput struct {
	Year     int `json:"year"`
	Month    int `json:"month"`
	UnitID   int `json:"unit_id"`
	BudgetID int `json:"budget_id"`
}

type DeleteSpendingReleaseInput struct {
	Month    int `json:"month"`
	UnitID   int `json:"unit_id"`
	BudgetID int `json:"budget_id"`
	Year     int `json:"year"`
}

type SpendingReleaseOverviewFilterDTO struct {
	Month    int `json:"month"`
	Year     int `json:"year" validate:"required"`
	BudgetID int `json:"budget_id" validate:"required"`
	UnitID   int `json:"unit_id" validate:"required"`
}

type SpendingReleaseOverviewRequestFilter struct {
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Status             *string `json:"status"`
}

type GetSpendingReleaseOverviewResponseMS struct {
	Data []SpendingReleaseOverviewItem `json:"data"`
}

type SpendingReleaseOverviewItem struct {
	Month                int                `json:"month"`
	Year                 int                `json:"year"`
	CreatedAt            time.Time          `json:"created_at"`
	Value                decimal.Decimal    `json:"value"`
	OrganizationUnitFile FileDropdownSimple `json:"organization_unit_file"`
	SSSFile              FileDropdownSimple `json:"sss_file"`
	OrganizationUnit     DropdownSimple     `json:"organization_unit"`
	Status               string             `json:"status"`
}

type SpendingReleaseDTO struct {
	ID                  int                   `json:"id"`
	AccountID           int                   `json:"account_id"`
	BudgetID            int                   `json:"budget_id"`
	UnitID              int                   `json:"unit_id"`
	CurrentBudgetID     int                   `json:"current_budget_id"`
	Value               decimal.Decimal       `json:"value"`
	AccountSerialNumber string                `json:"account_serial_number"`
	AccountTitle        string                `json:"account_title"`
	CreatedAt           time.Time             `json:"created_at"`
	Username            string                `json:"username"`
	Planned             decimal.Decimal       `json:"planned"`
	Children            []*SpendingReleaseDTO `json:"children"`
}
