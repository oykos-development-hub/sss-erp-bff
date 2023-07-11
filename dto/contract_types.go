package dto

import "bff/structs"

type GetDropdownTypesResponseMS struct {
	Data []structs.SettingsDropdown `json:"data"`
}

type GetDropdownTypeResponseMS struct {
	Data structs.SettingsDropdown `json:"data"`
}
