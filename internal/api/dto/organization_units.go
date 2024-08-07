package dto

import "bff/structs"

type GetOrganizationUnitsResponseMS struct {
	Data  []structs.OrganizationUnits `json:"data"`
	Total int                         `json:"total"`
}

type GetOrganizationUnitResponseMS struct {
	Data structs.OrganizationUnits `json:"data"`
}

type GetOrganizationUnitsInput struct {
	Page     *int    `json:"page"`
	Size     *int    `json:"page_size"`
	ParentID *int    `json:"parent_id"`
	IsParent *bool   `json:"is_parent"`
	Search   *string `json:"search"`
	Active   *bool   `json:"active"`
}

type OrganizationUnitsOverviewResponse struct {
	ID             int                          `json:"id"`
	ParentID       *int                         `json:"parent_id"`
	NumberOfJudges int                          `json:"number_of_judges"`
	Title          string                       `json:"title"`
	Description    string                       `json:"description"`
	Abbreviation   string                       `json:"abbreviation"`
	City           string                       `json:"city"`
	Address        string                       `json:"address"`
	Pib            string                       `json:"pib"`
	Color          string                       `json:"color"`
	Icon           string                       `json:"icon"`
	FolderID       int                          `json:"folder_id"`
	CreatedAt      string                       `json:"created_at"`
	UpdatedAt      string                       `json:"updated_at"`
	BankAccounts   []string                     `json:"bank_accounts"`
	Active         bool                         `json:"active"`
	Children       *[]structs.OrganizationUnits `json:"children"`
	Code           string                       `json:"code"`
}

type OrganizationUnitsSectorResponse struct {
	ID                            int                             `json:"id"`
	ParentID                      *int                            `json:"parent_id"`
	Title                         string                          `json:"title"`
	Abbreviation                  string                          `json:"abbreviation"`
	Description                   string                          `json:"description"`
	Address                       string                          `json:"address"`
	Color                         string                          `json:"color"`
	Icon                          string                          `json:"icon"`
	FolderID                      int                             `json:"folder_id"`
	Active                        bool                            `json:"active"`
	CreatedAt                     string                          `json:"created_at"`
	UpdatedAt                     string                          `json:"updated_at"`
	JobPositionsOrganizationUnits []JobPositionsOrganizationUnits `json:"job_positions_organization_units"`
}
type JobPositionsOrganizationUnits struct {
	ID             int              `json:"id"`
	JobPositions   DropdownSimple   `json:"job_positions"`
	Description    *string          `json:"description"`
	SerialNumber   string           `json:"serial_number"`
	Requirements   *string          `json:"requirements"`
	AvailableSlots int              `json:"available_slots"`
	Employees      []DropdownSimple `json:"employees"`
}

func ToOrganizationUnitsSectorResponse(organizationUnit structs.OrganizationUnits) *OrganizationUnitsSectorResponse {
	return &OrganizationUnitsSectorResponse{
		ID:           organizationUnit.ID,
		ParentID:     organizationUnit.ParentID,
		Title:        organizationUnit.Title,
		Abbreviation: organizationUnit.Abbreviation,
		Description:  organizationUnit.Description,
		Address:      organizationUnit.Address,
		Color:        organizationUnit.Color,
		Icon:         organizationUnit.Icon,
		Active:       organizationUnit.Active,
		FolderID:     organizationUnit.FolderID,
		UpdatedAt:    organizationUnit.UpdatedAt,
		CreatedAt:    organizationUnit.CreatedAt,
	}
}
