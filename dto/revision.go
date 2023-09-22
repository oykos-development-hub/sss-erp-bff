package dto

import (
	"bff/structs"
)

type GetPlansInput struct {
	Size *int    `json:"size"`
	Page *int    `json:"page"`
	Year *string `json:"year"`
}

type GetPlanResponseMS struct {
	Data RevisionPlanItem `json:"data"`
}

type GetRevisionPlanResponseMS struct {
	Data  []RevisionPlanItem `json:"data"`
	Total int                `json:"total"`
}

type RevisionPlanOverviewResponse struct {
	Message string             `json:"message"`
	Status  string             `json:"status"`
	Total   int                `json:"total"`
	Items   []RevisionPlanItem `json:"items"`
}

type RevisionPlanItem struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Year      string  `json:"year"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

//----------------------------------------------------------------------

type GetRevisionFilter struct {
	Page                    *int `json:"page"`
	Size                    *int `json:"size"`
	Revisor                 *int `json:"revisor"`
	RevisionType            *int `json:"revision_type"`
	InternalRevisionSubject *int `json:"internal_revision_subject"`
	PlanID                  *int `json:"plan_id"`
}

type RevisionsOverviewItem struct {
	ID                      int             `json:"id"`
	Title                   string          `json:"title"`
	PlanID                  int             `json:"plan_id"`
	SerialNumber            string          `json:"serial_number"`
	DateOfRevision          string          `json:"date_of_revision"`
	RevisionPriority        string          `json:"revision_priority"`
	RevisionQuartal         string          `json:"revision_quartal"`
	InternalRevisionSubject *DropdownSimple `json:"internal_revision_subject"`
	ExternalRevisionSubject *DropdownSimple `json:"external_revision_subject"`
	Revisor                 DropdownSimple  `json:"revisor"`
	RevisionType            DropdownSimple  `json:"revision_type"`
	FileID                  *int            `json:"file_id"`
	CreatedAt               string          `json:"created_at"`
	UpdatedAt               string          `json:"updated_at"`
}

type GetRevisionMS struct {
	Data structs.Revisions `json:"data"`
}

type GetRevisionsResponseMS struct {
	Data  []*structs.Revisions `json:"data"`
	Total int                  `json:"total"`
}

type RevisionsOverviewResponse struct {
	Revisors []*structs.SettingsDropdown `json:"revisors"`
	Message  string                      `json:"message"`
	Status   string                      `json:"status"`
	Total    int                         `json:"total"`
	Items    []RevisionsOverviewItem     `json:"items"`
}

type RevisionsDetailsResponse struct {
	Revisors []*structs.SettingsDropdown `json:"revisors"`
	Message  string                      `json:"message"`
	Status   string                      `json:"status"`
	Total    int                         `json:"total"`
	Item     RevisionsOverviewItem       `json:"item"`
}

//---------------------------------------------------------------------------

type GetRevisionTipFilter struct {
	Page       *int `json:"page"`
	Size       *int `json:"size"`
	RevisionID *int `json:"revision_id"`
}

type RevisionTipsOverviewItem struct {
	ID                     int                      `json:"id"`
	RevisionID             int                      `json:"revision_id"`
	UserProfile            structs.SettingsDropdown `json:"user_profile"`
	DateOfAccept           *string                  `json:"date_of_accept"`
	DueDate                int                      `json:"due_date"`
	DateOfReject           *string                  `json:"date_of_reject"`
	DateOfExecution        *string                  `json:"date_of_execution"`
	Recommendation         string                   `json:"recommendation"`
	Status                 *string                  `json:"status"`
	Documents              *string                  `json:"documents"`
	ReasonsForNonExecuting *string                  `json:"reasons_for_non_executing"`
	FileID                 *int                     `json:"file_id"`
	CreatedAt              string                   `json:"created_at"`
	UpdatedAt              string                   `json:"updated_at"`
}

type GetRevisionTipMS struct {
	Data structs.RevisionTips `json:"data"`
}

type GetRevisionTipsResponseMS struct {
	Data  []*structs.RevisionTips `json:"data"`
	Total int                     `json:"total"`
}

type RevisionTipsOverviewResponse struct {
	Revisors []*structs.SettingsDropdown `json:"revisors"`
	Message  string                      `json:"message"`
	Status   string                      `json:"status"`
	Total    int                         `json:"total"`
	Items    []RevisionTipsOverviewItem  `json:"items"`
}
