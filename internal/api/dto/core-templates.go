package dto

import (
	"bff/structs"
)

type TemplatesResponse struct {
	ID               int                `json:"id,omitempty"`
	OrganizationUnit DropdownSimple     `json:"organization_unit"`
	File             FileDropdownSimple `json:"file"`
	Template         DropdownSimple     `json:"template"`
}

type TemplateFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	TemplateID         *int    `json:"template_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
}

type GetTemplateResponseListMS struct {
	Data  []structs.Template `json:"data"`
	Total int                `json:"total"`
}

type GetListOfParametersResponseListMS struct {
	Data  []structs.ListOfParameters `json:"data"`
	Total int                        `json:"total"`
}

type GetTemplateResponseMS struct {
	Data structs.Template `json:"data"`
}

type GetCustomerResponseMS struct {
	Data structs.CustomerSupport `json:"data"`
}

type CustomerSupport struct {
	ID                    int                `json:"id"`
	UserDocumentationFile FileDropdownSimple `json:"user_documentation_file"`
}
