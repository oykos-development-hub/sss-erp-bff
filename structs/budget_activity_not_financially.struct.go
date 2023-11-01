package structs

type BudgetActivityNotFinanciallyItem struct {
	Id                               int    `json:"id"`
	RequestId                        int    `json:"request_id"`
	PersonResponsibleNameSurname     string `json:"person_responsible_name_surname"`
	PersonResponsibleWorkingPlace    string `json:"person_responsible_working_place"`
	PersonResponsibleTelephoneNumber string `json:"person_responsible_telephone_number"`
	PersonResponsibleEmail           string `json:"person_responsible_email"`
	ContactPersonNameSurname         string `json:"contact_person_name_surname"`
	ContactPersonWorkingPlace        string `json:"contact_person_working_place"`
	ContactPersonTelephoneNumber     string `json:"contact_person_telephone_number"`
	ContactPersonEmail               string `json:"contact_person_email"`
}

type BudgetActivityNotFinanciallyProgramItem struct {
	Id                     int    `json:"id"`
	BudgetNotFinanciallyId int    `json:"budget_not_financially_id"` // foreign key BudgetActivityNotFinanciallyItem
	ProgramId              int    `json:"program_id"`                // foreign key ProgramItem
	Description            string `json:"description"`
}

type BudgetActivityNotFinanciallyGoalsItem struct {
	Id              int    `json:"id"`
	BudgetProgramId int    `json:"budget_program_id"` // foreign key BudgetActivityNotFinanciallyProgramItem
	Title           string `json:"title"`
	Description     string `json:"description"`
}

type BudgetActivityNotFinanciallyIndicatorItem struct {
	Id                   int    `json:"id"`
	GoalsId              int    `json:"goals_id"` // foreign key BudgetActivityNotFinanciallyProgramItem
	Code                 string `json:"code"`
	Source               string `json:"source"`
	BaseYear             string `json:"base_year"`
	GenderEquality       string `json:"gender_equality"`
	BaseValue            string `json:"base_value"`
	SourceInformation    string `json:"source_information"`
	Unit                 string `json:"unit"`
	Description          string `json:"description"`
	PlannedCurrentYear   string `json:"planned_current_year"`
	RevisedCurrentYear   string `json:"revised_current_year"`
	ValueCurrentYear     string `json:"value_current_year"`
	PlannedNextYear      string `json:"planned_next_year"`
	RevisedNextYear      string `json:"revised_next_year"`
	ValueNextYear        string `json:"value_next_year"`
	PlannedAfterNextYear string `json:"planned_after_next_year"`
	RevisedAfterNextYear string `json:"revised_after_next_year"`
	ValueAfterNextYear   string `json:"value_after_next_year"`
}
