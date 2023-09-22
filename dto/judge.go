package dto

import (
	"bff/structs"
)

type GetJudgeNormResponseMS struct {
	Data structs.JudgeNorms `json:"data"`
}

type GetJudgeNormListResponseMS struct {
	Data []*structs.JudgeNorms `json:"data"`
}

type GetJudgeResolutionResponseMS struct {
	Data structs.JudgeResolutions `json:"data"`
}

type GetJudgeResolutionListResponseMS struct {
	Data  []*structs.JudgeResolutions `json:"data"`
	Total int                         `json:"total"`
}

type GetJudgeResolutionItemResponseMS struct {
	Data structs.JudgeResolutionItems `json:"data"`
}

type GetJudgeResolutionItemListResponseMS struct {
	Data []*structs.JudgeResolutionItems `json:"data"`
}

type GetJudgeResolutionListInputMS struct {
	Page *int `json:"page"`
	Size *int `json:"size"`
}

type GetJudgeResolutionItemListInputMS struct {
	ResolutionID *int `json:"resolution_id"`
}

type Judges struct {
	ID               int                      `json:"id"`
	OrganizationUnit structs.SettingsDropdown `json:"organization_unit"`
	JobPosition      structs.SettingsDropdown `json:"job_position"`
	IsJudgePresident bool                     `json:"is_judge_president"`
	FirstName        string                   `json:"first_name"`
	LastName         string                   `json:"last_name"`
	CreatedAt        string                   `json:"created_at"`
	UpdatedAt        string                   `json:"updated_at"`
	FolderID         *int                     `json:"folder_id"`
	Norms            []*NormResItem           `json:"norms"`
}

type NormResItem struct {
	Id                       int                 `json:"id"`
	UserProfileId            int                 `json:"user_profile_id"`
	Topic                    string              `json:"topic"`
	Title                    string              `json:"title"`
	PercentageOfNormDecrease float32             `json:"percentage_of_norm_decrease"`
	NumberOfNormDecrease     int                 `json:"number_of_norm_decrease"`
	NumberOfItems            int                 `json:"number_of_items"`
	NumberOfItemsSolved      int                 `json:"number_of_items_solved"`
	Evaluation               *structs.Evaluation `json:"evaluation"`
	DateOfEvaluationValidity string              `json:"date_of_evaluation_validity"`
	FileID                   int                 `json:"file_id"`
	Relocation               *structs.Absent     `json:"relocation,omitempty"`
	CreatedAt                string              `json:"created_at"`
	UpdatedAt                string              `json:"updated_at"`
}

type GetEmployeeNormListResponseMS struct {
	Data []structs.JudgeNorms `json:"data"`
}

type JudgeResolutionItemResponseItem struct {
	Id                       int                      `json:"id"`
	ResolutionId             int                      `json:"resolution_id"`
	OrganizationUnit         structs.SettingsDropdown `json:"organization_unit"`
	NumberOfJudges           int                      `json:"number_of_judges"`
	NumberOfPresidents       int                      `json:"number_of_presidents"`
	AvailableSlotsPredisents int                      `json:"available_slots_presidents"`
	AvailableSlotsJudges     int                      `json:"available_slots_judges"`
	NumberOfEmployees        int                      `json:"number_of_employees"`
	NumberOfRelocatedJudges  int                      `json:"number_of_relocated_judges"`
	NumberOfSuspendedJudges  int                      `json:"number_of_suspended_judges"`
}

type JudgeResolutionsResponseItem struct {
	Id                   int                                `json:"id"`
	SerialNumber         string                             `json:"serial_number"`
	Year                 string                             `json:"year"`
	CreatedAt            string                             `json:"created_at"`
	UpdatedAt            string                             `json:"updated_at"`
	Active               bool                               `json:"active"`
	NumberOfJudges       int                                `json:"number_of_judges"`
	AvailableSlotsJudges int                                `json:"available_slots_judges"`
	Items                []*JudgeResolutionItemResponseItem `json:"items"`
}
