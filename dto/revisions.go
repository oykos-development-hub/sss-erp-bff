package dto

import "bff/structs"

type GetRevisionResponseMS struct {
	Data structs.Revision `json:"data"`
}

type GetRevisionListResponseMS struct {
	Data  []*structs.Revision `json:"data"`
	Total int                 `json:"total"`
}

type RevisionOverviewResponse struct {
	Revisors []*structs.SettingsDropdown `json:"revisors"`
	Message  string                      `json:"message"`
	Status   string                      `json:"status"`
	Total    int                         `json:"total"`
	Items    []RevisionOverviewItem      `json:"items"`
}

type RevisionOverviewItem struct {
	Id                       int                      `json:"id"`
	RevisionType             structs.SettingsDropdown `json:"revision_type"`
	RevisorUserProfile       structs.SettingsDropdown `json:"revisor_user_profile"`
	RevisionOrganizationUnit structs.SettingsDropdown `json:"revision_organization_unit"`
	Title                    string                   `json:"title"`
	PlannedYear              string                   `json:"planned_year"`
	PlannedQuarter           string                   `json:"planned_quarter"`
	CreatedAt                string                   `json:"created_at"`
	UpdatedAt                string                   `json:"updated_at"`
}

type RevisionDetailsItem struct {
	ID                              int                      `json:"id"`
	RevisionType                    structs.SettingsDropdown `json:"revision_type"`
	RevisorUserProfile              structs.SettingsDropdown `json:"revisor_user_profile"`
	RevisionOrganizationUnit        structs.SettingsDropdown `json:"revision_organization_unit"`
	ResponsibleUserProfile          structs.SettingsDropdown `json:"responsible_user_profile"`
	ImplementationUserProfile       structs.SettingsDropdown `json:"implementation_user_profile"`
	Title                           string                   `json:"title"`
	PlannedYear                     string                   `json:"planned_year"`
	PlannedQuarter                  string                   `json:"planned_quarter"`
	SerialNumber                    string                   `json:"serial_number"`
	Priority                        *string                  `json:"priority"`
	DateOfRevision                  structs.JSONDate         `json:"date_of_revision"`
	DateOfAcceptance                structs.JSONDate         `json:"date_of_acceptance"`
	DateOfRejection                 structs.JSONDate         `json:"date_of_rejection"`
	ImplementationSuggestion        *string                  `json:"implementation_suggestion"`
	ImplementationMonthSpan         *string                  `json:"implementation_month_span"`
	DateOfImplementation            structs.JSONDate         `json:"date_of_implementation"`
	StateOfImplementation           *string                  `json:"state_of_implementation"`
	ImplementationFailedDescription *string                  `json:"implementation_failed_description"`
	SecondImplementationMonthSpan   *string                  `json:"second_implementation_month_span"`
	SecondDateOfRevision            structs.JSONDate         `json:"second_date_of_revision"`
	FileID                          *int                     `json:"file_id"`
	RefDocument                     string                   `json:"ref_document"`
	CreatedAt                       string                   `json:"created_at"`
	UpdatedAt                       string                   `json:"updated_at"`
}

type GetRevisionsInput struct {
	Page                       *int `json:"page"`
	Size                       *int `json:"size"`
	InternalOrganizationUnitID *int `json:"internal_organization_unit_id"`
	ExternalOrganizationUnitID *int `json:"external_organization_unit_id"`
	RevisorUserProfileID       *int `json:"revisor_user_profile_id"`
}
