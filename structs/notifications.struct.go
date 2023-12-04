package structs

import "encoding/json"

type Notifications struct {
	Id          int             `json:"id"`
	Content     string          `json:"content"`
	Module      string          `json:"module"`
	FromContent string          `json:"from_content"`
	FromUserID  int             `json:"from_user_id"`
	Path        string          `json:"path"`
	Data        json.RawMessage `json:"data"`
	ToUserID    int             `json:"to_user_id"`
	IsRead      bool            `json:"is_read"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}
