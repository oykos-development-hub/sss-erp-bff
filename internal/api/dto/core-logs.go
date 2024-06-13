package dto

import (
	"bff/config"
	"bff/structs"
	"encoding/json"
	"time"
)

type LogResponse struct {
	ID          int             `json:"id,omitempty"`
	ChangedAt   time.Time       `json:"changed_at"`
	User        DropdownSimple  `json:"user"`
	UserProfile DropdownSimple  `json:"user_profile"`
	ItemID      int             `json:"item_id"`
	Operation   string          `json:"operation"`
	Entity      string          `json:"entity"`
	OldState    json.RawMessage `json:"old_state"`
	NewState    json.RawMessage `json:"new_state"`
}

type LogFilterDTO struct {
	Page        *int          `json:"page"`
	Size        *int          `json:"size"`
	SortByTitle *string       `json:"sort_by_title"`
	Entity      *string       `json:"entity"`
	UserID      *int          `json:"user_id"`
	Search      *string       `json:"search"`
	ItemID      *int          `json:"item_id"`
	Operation   *string       `json:"operation"`
	Module      config.Module `json:"module"`
}

type GetLogResponseListMS struct {
	Data []structs.Logs `json:"data"`
}

type GetLogResponseMS struct {
	Data structs.Logs `json:"data"`
}
