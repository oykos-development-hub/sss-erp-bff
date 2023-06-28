package structs

type Revision struct {
	Id                              int    `json:"id"`
	RevisionTypeId                  int    `json:"revision_type_id"`
	RevisorUserProfileId            int    `json:"revisor_user_profile_id"`
	RevisionOrganizationUnitId      int    `json:"revision_organization_unit_id"`
	ResponsibleUserProfileId        int    `json:"responsible_user_profile_id"`
	ResponsibleUserProfile          string `json:"responsible_user_profile"`
	ImplementationUserProfileId     int    `json:"implementation_user_profile_id"`
	ImplementationUserProfile       string `json:"implementation_user_profile"`
	PlannedYear                     string `json:"planned_year"`
	PlannedQuarter                  string `json:"planned_quarter"`
	Title                           string `json:"title"`
	SerialNumber                    string `json:"serial_number"`
	Priority                        string `json:"priority"`
	DateOfRevision                  string `json:"date_of_revision"`
	DateOfAcceptance                string `json:"date_of_acceptance"`
	DateOfRejection                 string `json:"date_of_rejection"`
	ImplementationSuggestion        string `json:"implementation_suggestion"`
	ImplementationMonthSpan         string `json:"implementation_month_span"`
	DateOfImplementation            string `json:"date_of_implementation"`
	StateOfImplementation           string `json:"state_of_implementation"`
	ImplementationFailedDescription string `json:"implementation_failed_description"`
	SecondImplementationMonthSpan   string `json:"second_implementation_month_span"`
	SecondDateOfRevision            string `json:"second_date_of_revision"`
	CreatedAt                       string `json:"created_at"`
	UpdatedAt                       string `json:"updated_at"`
	FileId                          int    `json:"file_id"`
}

type RevisionType struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
}
