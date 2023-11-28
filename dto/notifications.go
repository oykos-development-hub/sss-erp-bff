package dto

import "bff/structs"

type GetNotificationResponseMS struct {
	Data structs.Notifications `json:"data"`
}

type GetNotificationInputMS struct {
	Page     *int `json:"page"`
	Size     *int `json:"size"`
	ToUserID *int `json:"to_user_id"`
}

type GetNotificationListResponseMS struct {
	Data  []*structs.Notifications `json:"data"`
	Total int                      `json:"total"`
}
