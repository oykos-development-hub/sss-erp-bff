package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) NotificationOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

	notificiations, err := r.Repo.FetchNotifications(loggedInUser.ID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   notificiations,
		Total:   len(notificiations),
	}, nil
}

func (r *Resolver) NotificationReadResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.MarkNotificationRead(itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You read this item!",
	}, nil
}

func (r *Resolver) NotificationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteNotification(itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You delete this item!",
	}, nil
}
