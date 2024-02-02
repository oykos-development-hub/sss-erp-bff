package dto

import "bff/structs"

type GetNonFinancialGoalIndicatorResponseMS struct {
	Data structs.NonFinancialGoalIndicatorItem `json:"data"`
}

type GetNonFinancialGoalIndicatorListResponseMS struct {
	Data []structs.NonFinancialGoalIndicatorItem `json:"data"`
}

type GetNonFinancialGoalIndicatorListInputMS struct {
	GoalID *int    `json:"activity_id"`
	Search *string `json:"search"`
}

type NonFinancialGoalIndicatorResItem struct {
	ID               int            `json:"id"`
	Budet            DropdownSimple `json:"budget"`
	OrganizationUnit DropdownSimple `json:"organization_unit"`

	ImplContactFullName     string `json:"impl_contact_fullname"`
	ImplContactWorkingPlace string `json:"impl_contact_working_place"`
	ImplContactPhone        string `json:"impl_contact_phone"`
	ImplContactEmail        string `json:"impl_contact_email"`

	ContactFullName     string `json:"contact_fullname"`
	ContactWorkingPlace string `json:"contact_working_place"`
	ContactPhone        string `json:"contact_phone"`
	ContactEmail        string `json:"contact_email"`
}
