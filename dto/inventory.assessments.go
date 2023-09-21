package dto

import "bff/structs"

type AssessmentResponseMS struct {
	Data structs.BasicInventoryAssessmentsTypesItem `json:"data"`
}

type AssessmentResponseArrayMS struct {
	Data []structs.BasicInventoryAssessmentsTypesItem `json:"data"`
}

type GetAssessmentResponseMS struct {
	Data *BasicInventoryResponseAssessment `json:"data"`
}

type BasicInventoryResponseAssessment struct {
	Id                   int            `json:"id"`
	InventoryId          int            `json:"inventory_id"`
	Active               bool           `json:"active"`
	DepreciationType     DropdownSimple `json:"depreciation_type"`
	UserProfile          DropdownSimple `json:"user_profile"`
	GrossPriceNew        int            `json:"gross_price_new"`
	GrossPriceDifference int            `json:"gross_price_difference"`
	DateOfAssessment     *string        `json:"date_of_assessment"`
	CreatedAt            string         `json:"created_at"`
	UpdatedAt            string         `json:"updated_at"`
	FileId               *int           `json:"file_id"`
}
