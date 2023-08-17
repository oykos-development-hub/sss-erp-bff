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
	Search   *string `json:"search"`
}

type OrganizationUnitsOverviewResponse struct {
	Id             int                          `json:"id"`
	ParentId       *int                         `json:"parent_id"`
	NumberOfJudges int                          `json:"number_of_judges"`
	Title          string                       `json:"title"`
	Description    string                       `json:"description"`
	Abbreviation   string                       `json:"abbreviation"`
	Address        string                       `json:"address"`
	Color          string                       `json:"color"`
	Icon           string                       `json:"icon"`
	FolderId       int                          `json:"folder_id"`
	CreatedAt      string                       `json:"created_at"`
	UpdatedAt      string                       `json:"updated_at"`
	Children       *[]structs.OrganizationUnits `json:"children"`
}

type OrganizationUnitsSectorResponse struct {
	Id                            int                             `json:"id"`
	ParentId                      *int                            `json:"parent_id"`
	Title                         string                          `json:"title"`
	Abbreviation                  string                          `json:"abbreviation"`
	Color                         string                          `json:"color"`
	Icon                          string                          `json:"icon"`
	FolderId                      int                             `json:"folder_id"`
	CreatedAt                     string                          `json:"created_at"`
	UpdatedAt                     string                          `json:"updated_at"`
	JobPositionsOrganizationUnits []JobPositionsOrganizationUnits `json:"job_positions_organization_units"`
}
type JobPositionsOrganizationUnits struct {
	Id             int              `json:"id"`
	JobPositions   DropdownSimple   `json:"job_positions"`
	Description    string           `json:"description"`
	SerialNumber   string           `json:"serial_number"`
	Requirements   string           `json:"requirements"`
	AvailableSlots int              `json:"available_slots"`
	Employees      []DropdownSimple `json:"employees"`
}

func ToOrganizationUnitsSectorResponse(organizationUnit structs.OrganizationUnits) *OrganizationUnitsSectorResponse {
	return &OrganizationUnitsSectorResponse{
		Id:           organizationUnit.Id,
		ParentId:     organizationUnit.ParentId,
		Title:        organizationUnit.Title,
		Abbreviation: organizationUnit.Abbreviation,
		Color:        organizationUnit.Color,
		Icon:         organizationUnit.Icon,
		FolderId:     organizationUnit.FolderId,
		UpdatedAt:    organizationUnit.UpdatedAt,
		CreatedAt:    organizationUnit.CreatedAt,
	}
}
