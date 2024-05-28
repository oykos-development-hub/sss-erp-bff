package dto

import (
	"bff/structs"
)

type GetSpendingReleaseResponseMS struct {
	Data structs.SpendingRelease `json:"data"`
}
