package dto

import "bff/structs"

type GetEmployeesInOrganizationUnitsResponseMS struct {
	Data *structs.EmployeesInOrganizationUnits `json:"data"`
}

type GetEmployeesInOrganizationUnitsListResponseMS struct {
	Data []*structs.EmployeesInOrganizationUnits `json:"data"`
}
