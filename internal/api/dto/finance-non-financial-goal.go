package dto

import "bff/structs"

type GetNonFinancialGoalResponseMS struct {
	Data structs.NonFinancialGoalItem `json:"data"`
}

type GetNonFinancialGoalListResponseMS struct {
	Data []structs.NonFinancialGoalItem `json:"data"`
}

type GetNonFinancialGoalListInputMS struct {
	ActivityID           *int    `json:"activity_id"`
	NonFinancialBudgetID *int    `json:"non_financial_budget_id"`
	Search               *string `json:"search"`
}

type NonFinancialGoalResItem struct {
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
