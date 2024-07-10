package structs

import (
	"encoding/json"
	"time"
)

type Logs struct {
	ID        int             `json:"id,omitempty"`
	ChangedAt time.Time       `json:"changed_at"`
	UserID    int             `json:"user_id"`
	ItemID    int             `json:"item_id"`
	Operation string          `json:"operation"`
	Entity    string          `json:"entity"`
	OldState  json.RawMessage `json:"old_state"`
	NewState  json.RawMessage `json:"new_state"`
}

type ErrorLogs struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Error     string    `json:"error"`
	Entity    string    `json:"entity"`
	Code      int       `json:"code"`
}
