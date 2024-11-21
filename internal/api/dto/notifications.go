package dto

import "bff/structs"

type GetNotificationResponseMS struct {
	Data structs.Notifications `json:"data"`
}

type GetNotificationInputMS struct {
	ToUserID *int `json:"to_user_id"`
}

type GetNotificationListResponseMS struct {
	Data  []structs.Notifications `json:"data"`
	Total int                     `json:"total"`
}

type ProcurementPlanNotification struct {
	ID          int    `json:"plan_id"`
	Year        string `json:"year"`
	IsPreBudget bool   `json:"is_pre_budget"`
}
