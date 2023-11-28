package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"strconv"

	"github.com/graphql-go/graphql"
)

var NotificationsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	loggedInAccount := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

	if id != nil && id.(int) > 0 {
		notification, err := getNotification(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []structs.Notifications{*notification},
			Total:   1,
		}, nil
	} else {
		input := dto.GetNotificationInputMS{}
		if page != nil && size != nil {
			pageValue := page.(int)
			sizeValue := size.(int)

			input.Size = &sizeValue
			input.Page = &pageValue
		}
		input.ToUserID = &loggedInAccount.Id

		res, err := getNotificationList(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   res.Data,
			Total:   res.Total,
		}, nil
	}
}

var NotificationsReadResolver = func(params graphql.ResolveParams) (interface{}, error) {
	notificationID := params.Args["id"].(int)
	isRead := params.Args["is_read"].(bool)

	notification, err := getNotification(notificationID)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	notification.IsRead = isRead

	res, err := updateNotification(notificationID, notification)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Message: "You updated this item!",
		Item:    res,
		Status:  "success"}, nil
}

var NotificationsDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteNotification(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func createNotification(notification *structs.Notifications) (*structs.Notifications, error) {
	res := &dto.GetNotificationResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.NOTIFICATIONS_ENDPOINT, notification, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateNotification(id int, notification *structs.Notifications) (*structs.Notifications, error) {
	res := &dto.GetNotificationResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.NOTIFICATIONS_ENDPOINT+"/"+strconv.Itoa(id), notification, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteNotification(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.NOTIFICATIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getNotification(id int) (*structs.Notifications, error) {
	res := &dto.GetNotificationResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.NOTIFICATIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getNotificationList(input *dto.GetNotificationInputMS) (*dto.GetNotificationListResponseMS, error) {
	res := &dto.GetNotificationListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.NOTIFICATIONS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
