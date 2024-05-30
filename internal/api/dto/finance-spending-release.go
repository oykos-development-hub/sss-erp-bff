package dto

import (
	"bff/structs"
)

type GetSpendingReleaseResponseMS struct {
	Data structs.SpendingRelease `json:"data"`
}

type GetSpendingReleaseListInput struct {
	Year      *int `json:"year"`
	Month     *int `json:"month"`
	UnitID    int  `json:"unit_id"`
	AccountID int  `json:"account_id"`
}
