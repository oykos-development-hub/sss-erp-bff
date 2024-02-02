package dto

import "bff/structs"

type GetNonFinancialBudgetResponseMS struct {
	Data structs.NonFinancialBudgetItem `json:"data"`
}

type GetNonFinancialBudgetListResponseMS struct {
	Data []structs.NonFinancialBudgetItem `json:"data"`
}

type GetNonFinancialBudgetListInputMS struct {
	OrganizationUnitID *int    `json:"organization_unit_id"`
	BudgetID           *int    `json:"budget_id"`
	Search             *string `json:"search"`
}

type NonFinancialBudgetResItem struct {
	ID               int                    `json:"id"`
	Budet            DropdownSimple         `json:"budget"`
	OrganizationUnit DropdownSimple         `json:"organization_unit"`
	ActivityRequest  ActivityRequestResItem `json:"activity"`

	ImplContactFullName     string `json:"impl_contact_fullname"`
	ImplContactWorkingPlace string `json:"impl_contact_working_place"`
	ImplContactPhone        string `json:"impl_contact_phone"`
	ImplContactEmail        string `json:"impl_contact_email"`

	ContactFullName     string `json:"contact_fullname"`
	ContactWorkingPlace string `json:"contact_working_place"`
	ContactPhone        string `json:"contact_phone"`
	ContactEmail        string `json:"contact_email"`
}

type ActivityRequestResItem struct {
	ID               int                           `json:"id"`
	SubProgram       DropdownSimple                `json:"sub_program"`
	OrganizationUnit DropdownSimple                `json:"organization_unit"`
	Title            string                        `json:"title"`
	Description      string                        `json:"description"`
	Code             string                        `json:"code"`
	Goals            []*ActivityGoalRequestResItem `json:"goals"`
}

type ActivityGoalRequestResItem struct {
	ID          int                                   `json:"id"`
	Title       string                                `json:"title"`
	Description string                                `json:"description"`
	Indicators  []*BudgetActivityGoalIndicatorResItem `json:"indicators"`
}

type BudgetActivityGoalIndicatorResItem struct {
	ID                       int    `json:"id"`
	PerformanceIndicatorCode string `json:"performance_indicator_code"`
	IndicatorSource          string `json:"indicator_source"`
	BaseYear                 string `json:"base_year"`
	GenderEquality           string `json:"gender_equality"`
	BaseValue                string `json:"base_value"`
	SourceOfInformation      string `json:"source_of_information"`
	UnitOfMeasure            string `json:"unit_of_measure"`
	IndicatorDescription     string `json:"indicator_description"`
	PlannedValue1            string `json:"planned_value_1"`
	RevisedValue1            string `json:"revised_value_1"`
	AchievedValue1           string `json:"achieved_value_1"`
	PlannedValue2            string `json:"planned_value_2"`
	RevisedValue2            string `json:"revised_value_2"`
	AchievedValue2           string `json:"achieved_value_2"`
	PlannedValue3            string `json:"planned_value_3"`
	RevisedValue3            string `json:"revised_value_3"`
	AchievedValue3           string `json:"achieved_value_3"`
}
