package dto

import (
	"bff/structs"

	"github.com/shopspring/decimal"
)

type GetNonFinancialBudgetResponseMS struct {
	Data structs.NonFinancialBudgetItem `json:"data"`
}

type GetNonFinancialBudgetListResponseMS struct {
	Data []structs.NonFinancialBudgetItem `json:"data"`
}

type GetNonFinancialBudgetListInputMS struct {
	RequestIDList *[]int `json:"request_id_list"`
}

type NonFinancialBudgetResItem struct {
	ID              int                    `json:"id"`
	RequestID       int                    `json:"request_id"`
	Status          DropdownSimple         `json:"status"`
	ActivityRequest ActivityRequestResItem `json:"activity"`

	ImplContactFullName     string `json:"impl_contact_fullname"`
	ImplContactWorkingPlace string `json:"impl_contact_working_place"`
	ImplContactPhone        string `json:"impl_contact_phone"`
	ImplContactEmail        string `json:"impl_contact_email"`

	ContactFullName     string `json:"contact_fullname"`
	ContactWorkingPlace string `json:"contact_working_place"`
	ContactPhone        string `json:"contact_phone"`
	ContactEmail        string `json:"contact_email"`

	Statement string `json:"statement"`

	OfficialComment string `json:"official_comment"`
}

type ActivityRequestResItem struct {
	ID               int                           `json:"id"`
	SubProgram       DropdownSimple                `json:"sub_program"`
	OrganizationUnit structs.OrganizationUnits     `json:"organization_unit"`
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

type FilledFinancialBudgetResItem struct {
	ID               int            `json:"id"`
	RequestID        int            `json:"request_id"`
	OrganizationUnit DropdownSimple `json:"organization_unit"`
	Account          DropdownSimple `json:"account"`

	CurrentYear   decimal.Decimal     `json:"current_year"`
	NextYear      decimal.Decimal     `json:"next_year"`
	YearAfterNext decimal.Decimal     `json:"year_after_next"`
	Actual        decimal.NullDecimal `json:"actual"`

	Description string `json:"description"`
}
