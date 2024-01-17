package dto

import (
	"bff/structs"
	"time"
)

type GetAbsentTypeResponseMS struct {
	Data structs.AbsentType `json:"data"`
}

type GetAbsentTypeListResponseMS struct {
	Data  []*structs.AbsentType `json:"data"`
	Total int                   `json:"total"`
}

type GetAbsentResponseMS struct {
	Data structs.Absent `json:"data"`
}

type GetAbsentListResponseMS struct {
	Data []*structs.Absent `json:"data"`
}

type AbsentsSummary struct {
	CurrentAvailableDays int `json:"current_available_days"`
	PastAvailableDays    int `json:"past_available_days"`
	UsedDays             int `json:"used_days"`
}

type EmployeeAbsentsInput struct {
	Date *time.Time `json:"date"`
	From *time.Time `json:"from"`
	To   *time.Time `json:"to"`
}

type Vacation struct {
	ID                int            `json:"id"`
	UserProfile       DropdownSimple `json:"user_profile"`
	ResolutionType    DropdownSimple `json:"resolution_type"`
	ResolutionPurpose string         `json:"resolution_purpose"`
	Year              int            `json:"year"`
	NumberOfDays      int            `json:"number_of_days"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         string         `json:"updated_at"`
	FileID            int            `json:"file_id"`
}

type VacationReportInput struct {
	Type               string `json:"type"`
	OrganizationUnitID int    `json:"organization_unit_id"`
	EmployeeID         *int   `json:"employee_id"`
}

type VacationReportResItem struct {
	FullName         string `json:"full_name"`
	OrganizationUnit string `json:"organization_unit"`
	TotalDays        int    `json:"total_days"`
	UsedDays         int    `json:"used_days"`
	LeftDays         int    `json:"left_days"`
}
