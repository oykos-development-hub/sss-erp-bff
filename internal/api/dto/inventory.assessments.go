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
	ID                   int            `json:"id"`
	Type                 string         `json:"type"`
	InventoryID          int            `json:"inventory_id"`
	EstimatedDuration    int            `json:"estimated_duration"`
	Active               bool           `json:"active"`
	ResidualPrice        *float32       `json:"residual_price"`
	DepreciationRate     string         `json:"depreciation_rate"`
	DepreciationType     DropdownSimple `json:"depreciation_type"`
	UserProfile          DropdownSimple `json:"user_profile"`
	GrossPriceNew        float32        `json:"gross_price_new"`
	GrossPriceDifference float32        `json:"gross_price_difference"`
	DateOfAssessment     *string        `json:"date_of_assessment"`
	CreatedAt            string         `json:"created_at"`
	UpdatedAt            string         `json:"updated_at"`
	FileID               *int           `json:"file_id"`
}
