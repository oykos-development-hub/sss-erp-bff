package dto

import (
	"bff/structs"
	"time"
)

type GetFixedDepositResponseMS struct {
	Data structs.FixedDeposit `json:"data"`
}

type GetFixedDepositWillDispatchResponseMS struct {
	Data structs.FixedDepositWillDispatch `json:"data"`
}

type GetFixedDepositListResponseMS struct {
	Data  []structs.FixedDeposit `json:"data"`
	Total int                    `json:"total"`
}

type FixedDepositFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	Type               *string `json:"type"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Search             *string `json:"search"`
	JudgeID            *int    `json:"judge_id"`
	Status             *string `json:"status"`
}

type FixedDepositWillFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	Search             *string `json:"search"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Status             *string `json:"status"`
}

type FixedDepositResponse struct {
	ID                   int                            `json:"id"`
	OrganizationUnit     DropdownSimple                 `json:"organization_unit"`
	Subject              string                         `json:"subject"`
	Judge                DropdownSimple                 `json:"judge"`
	CaseNumber           string                         `json:"case_number"`
	DateOfRecipiet       *time.Time                     `json:"date_of_recipiet"`
	DateOfCase           *time.Time                     `json:"date_of_case"`
	DateOfFinality       *time.Time                     `json:"date_of_finality"`
	DateOfEnforceability *time.Time                     `json:"date_of_enforceability"`
	DateOfEnd            *time.Time                     `json:"date_of_end"`
	Account              DropdownSimple                 `json:"account"`
	File                 FileDropdownSimple             `json:"file"`
	Status               string                         `json:"status"`
	Type                 string                         `json:"type"`
	Items                []FixedDepositItemResponse     `json:"items"`
	Dispatches           []FixedDepositDispatchResponse `json:"dispatches"`
	Judges               []FixedDepositJudgeResponse    `json:"judges"`
	CreatedAt            time.Time                      `json:"created_at"`
	UpdatedAt            time.Time                      `json:"updated_at"`
}

type FixedDepositItemResponse struct {
	ID                 int                `json:"id"`
	DepositID          int                `json:"deposit_id"`
	Category           DropdownSimple     `json:"category"`
	Type               DropdownSimple     `json:"type"`
	Unit               string             `json:"unit"`
	Currency           string             `json:"currency"`
	Amount             float32            `json:"amount"`
	SerialNumber       string             `json:"serial_number"`
	DateOfConfiscation *time.Time         `json:"date_of_confiscation"`
	CaseNumber         string             `json:"case_number"`
	Judge              DropdownSimple     `json:"judge"`
	File               FileDropdownSimple `json:"file"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}

type FixedDepositDispatchResponse struct {
	ID           int                `json:"id"`
	DepositID    int                `json:"deposit_id"`
	Category     DropdownSimple     `json:"category"`
	Type         DropdownSimple     `json:"type"`
	Unit         string             `json:"unit"`
	Currency     string             `json:"currency"`
	Amount       float32            `json:"amount"`
	SerialNumber string             `json:"serial_number"`
	DateOfAction *time.Time         `json:"date_of_action"`
	Subject      string             `json:"subject"`
	Action       string             `json:"action"`
	CaseNumber   string             `json:"case_number"`
	Judge        DropdownSimple     `json:"judge"`
	File         FileDropdownSimple `json:"file"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

type FixedDepositJudgeResponse struct {
	ID          int                `json:"id"`
	Judge       DropdownSimple     `json:"judge"`
	DepositID   int                `json:"deposit_id"`
	WillID      int                `json:"will_id"`
	DateOfStart time.Time          `json:"date_of_start"`
	DateOfEnd   *time.Time         `json:"date_of_end"`
	File        FileDropdownSimple `json:"file"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type GetFixedDepositWillResponseMS struct {
	Data structs.FixedDepositWill `json:"data"`
}

type GetFixedDepositWillListResponseMS struct {
	Data  []structs.FixedDepositWill `json:"data"`
	Total int                        `json:"total"`
}

type FixedDepositWillResponse struct {
	ID               int                                `json:"id"`
	OrganizationUnit DropdownSimple                     `json:"organization_unit"`
	Subject          string                             `json:"subject"`
	FatherName       string                             `json:"father_name"`
	DateOfBirth      time.Time                          `json:"date_of_birth"`
	JMBG             string                             `json:"jmbg"`
	CaseNumberSI     string                             `json:"case_number_si"`
	CaseNumberRS     string                             `json:"case_number_rs"`
	DateOfReceiptSI  *time.Time                         `json:"date_of_receipt_si"`
	DateOfReceiptRS  *time.Time                         `json:"date_of_receipt_rs"`
	DateOfEnd        *time.Time                         `json:"date_of_end"`
	Status           string                             `json:"status"`
	File             FileDropdownSimple                 `json:"file"`
	Judges           []FixedDepositJudgeResponse        `json:"judges"`
	Dispatches       []FixedDepositWillDispatchResponse `json:"dispatches"`
	CreatedAt        time.Time                          `json:"created_at"`
	UpdatedAt        time.Time                          `json:"updated_at"`
}

type FixedDepositWillDispatchResponse struct {
	ID             int                `json:"id"`
	WillID         int                `json:"will_id"`
	DispatchType   string             `json:"dispatch_type"`
	Judge          DropdownSimple     `json:"judge"`
	CaseNumber     string             `json:"case_number"`
	DateOfDispatch time.Time          `json:"date_of_dispatch"`
	File           FileDropdownSimple `json:"file"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
