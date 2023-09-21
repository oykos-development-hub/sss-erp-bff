package dto

import "bff/structs"

type GetSettingsInput struct {
	Entity string  `json:"entity"`
	Page   *int    `json:"page"`
	Size   *int    `json:"size"`
	Search *string `json:"search"`
	Value  *string `json:"value"`
}

type GetOfficesOfOrganizationInput struct {
	Entity string  `json:"entity"`
	Value  *string `json:"value" validate:"omitempty"`
	Page   *int    `json:"page" validate:"omitempty"`
	Size   *int    `json:"size" validate:"omitempty"`
	Search *string `json:"search" validate:"omitempty"`
}

type GetDropdownTypesResponseMS struct {
	Data  []structs.SettingsDropdown `json:"data"`
	Total int                        `json:"total"`
}

type GetDropdownTypeResponseMS struct {
	Data structs.SettingsDropdown `json:"data"`
}
