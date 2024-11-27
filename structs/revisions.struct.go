package structs

const (
	RevisionConducted string = "Sprovedena"
)

type Revision struct {
	ID                              int     `json:"id"`
	Name                            *string `json:"name"`
	RevisionTypeID                  *int    `json:"revision_type_id"`
	RevisorUserProfileID            *int    `json:"revisor_user_profile_id"`
	RevisorUserProfile              *string `json:"revisor_user_profile"`
	InternalOrganizationUnitID      *int    `json:"internal_organization_unit_id"`
	ExternalOrganizationUnitID      *int    `json:"external_organization_unit_id"`
	ResponsibleUserProfileID        *int    `json:"responsible_user_profile_id"`
	ResponsibleUserProfile          *string `json:"responsible_user_profile"`
	ImplementationUserProfileID     *int    `json:"implementation_user_profile_id"`
	ImplementationUserProfile       *string `json:"implementation_user_profile"`
	Title                           *string `json:"title"`
	PlannedYear                     *string `json:"planned_year"`
	PlannedQuarter                  *string `json:"planned_quarter"`
	SerialNumber                    *string `json:"serial_number"`
	Priority                        *string `json:"priority"`
	DateOfRevision                  *string `json:"date_of_revision"`
	DateOfAcceptance                *string `json:"date_of_acceptance"`
	DateOfRejection                 *string `json:"date_of_rejection"`
	ImplementationSuggestion        *string `json:"implementation_suggestion"`
	ImplementationMonthSpan         *string `json:"implementation_month_span"`
	DateOfImplementation            *string `json:"date_of_implementation"`
	StateOfImplementation           *string `json:"state_of_implementation"`
	ImplementationFailedDescription *string `json:"implementation_failed_description"`
	SecondImplementationMonthSpan   *string `json:"second_implementation_month_span"`
	SecondDateOfRevision            *string `json:"second_date_of_revision"`
	FileID                          *int    `json:"file_id"`
	TipsFileID                      *int    `json:"tips_file_id"`
	RefDocument                     *string `json:"ref_document"`
	CreatedAt                       string  `json:"created_at"`
	UpdatedAt                       string  `json:"updated_at"`
}

type Revisions struct {
	ID                      int    `json:"id"`
	Title                   string `json:"title"`
	PlanID                  int    `json:"plan_id"`
	SerialNumber            string `json:"serial_number"`
	DateOfRevision          string `json:"date_of_revision"`
	RevisionQuartal         string `json:"revision_quartal"`
	InternalRevisionsubject []int  `json:"internal_revision_subject_id"`
	ExternalRevisionsubject *int   `json:"external_revision_subject_id"`
	Revisor                 []int  `json:"revisor_id"`
	RevisionType            int    `json:"revision_type_id"`
	TipsFileIDs             []int  `json:"tips_file_ids"`
	FileIDs                 []int  `json:"file_ids"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
}

type RevisionTips struct {
	ID                int     `json:"id"`
	RevisionID        int     `json:"revision_id"`
	UserProfileID     *int    `json:"user_profile_id"`
	DateOfAccept      *string `json:"date_of_accept"`
	DueDate           int     `json:"due_date"`
	RevisionPriority  *string `json:"revision_priority"`
	EndDate           *string `json:"end_date"`
	DateOfReject      *string `json:"date_of_reject"`
	DateOfExecution   *string `json:"date_of_execution"`
	Recommendation    string  `json:"recommendation"`
	Status            string  `json:"status"`
	ResponsiblePerson *string `json:"responsible_person"`
	FileIDs           []int   `json:"file_ids"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

type RevisionTipImplementations struct {
	ID                     int     `json:"id,omitempty"`
	TipID                  int     `json:"tip_id"`
	Status                 string  `json:"status"`
	NewDueDate             *int    `json:"new_due_date"`
	NewDateOfExecution     *string `json:"new_date_of_execution"`
	ReasonsForNonExecuting *string `json:"reasons_for_non_executing"`
	RevisorID              *int    `json:"revisor_id"`
	Documents              string  `json:"documents"`
	FileIDs                []int   `json:"file_ids"`
	CreatedAt              string  `json:"created_at"`
	UpdatedAt              string  `json:"updated_at"`
}
