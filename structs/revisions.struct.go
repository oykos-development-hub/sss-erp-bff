package structs

type Revision struct {
	ID                              int      `json:"id"`
	RevisionTypeID                  int      `json:"revision_type_id"`
	RevisorUserProfileID            *int     `json:"revisor_user_profile_id"`
	RevisorUserProfile              *string  `json:"revisor_user_profile"`
	InternalOrganizationUnitID      *int     `json:"internal_organization_unit_id"`
	ExternalOrganizationUnitID      *int     `json:"external_organization_unit_id"`
	ResponsibleUserProfileID        *int     `json:"responsible_user_profile_id"`
	ResponsibleUserProfile          *string  `json:"responsible_user_profile"`
	ImplementationUserProfileID     *int     `json:"implementation_user_profile_id"`
	ImplementationUserProfile       *string  `json:"implementation_user_profile"`
	Title                           string   `json:"title"`
	PlannedYear                     string   `json:"planned_year"`
	PlannedQuarter                  string   `json:"planned_quarter"`
	SerialNumber                    string   `json:"serial_number"`
	Priority                        *string  `json:"priority"`
	DateOfRevision                  JSONDate `json:"date_of_revision"`
	DateOfAcceptance                JSONDate `json:"date_of_acceptance"`
	DateOfRejection                 JSONDate `json:"date_of_rejection"`
	ImplementationSuggestion        *string  `json:"implementation_suggestion"`
	ImplementationMonthSpan         *string  `json:"implementation_month_span"`
	DateOfImplementation            JSONDate `json:"date_of_implementation"`
	StateOfImplementation           *string  `json:"state_of_implementation"`
	ImplementationFailedDescription *string  `json:"implementation_failed_description"`
	SecondImplementationMonthSpan   *string  `json:"second_implementation_month_span"`
	SecondDateOfRevision            JSONDate `json:"second_date_of_revision"`
	FileID                          *int     `json:"file_id"`
	CreatedAt                       string   `json:"created_at"`
	UpdatedAt                       string   `json:"updated_at"`
}
