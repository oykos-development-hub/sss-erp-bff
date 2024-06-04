package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ExternalReallocationOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		ExternalReallocation, err := r.Repo.GetExternalReallocationByID(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		res, err := buildExternalReallocation(*ExternalReallocation, r)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.ExternalReallocationResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.ExternalReallocationFilter{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["source_organization_unit_id"].(int); ok && value != 0 {
		input.SourceOrganizationUnitID = &value
	}

	if value, ok := params.Args["destination_organization_unit_id"].(int); ok && value != 0 {
		input.DestinationOrganizationUnitID = &value
	}

	if value, ok := params.Args["status"].(string); ok && value != "" {
		input.Status = &value
	}

	if value, ok := params.Args["requested_by"].(int); ok && value != 0 {
		input.RequestedBy = &value
	}

	if value, ok := params.Args["budget_id"].(int); ok && value != 0 {
		input.BudgetID = &value
	}

	items, total, err := r.Repo.GetExternalReallocationList(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var resItems []dto.ExternalReallocationResponse
	for _, item := range items {
		resItem, err := buildExternalReallocation(item, r)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		resItems = append(resItems, *resItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   resItems,
		Total:   total,
	}, nil
}

func (r *Resolver) ExternalReallocationInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.ExternalReallocation
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		return apierrors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if data.SourceOrganizationUnitID == 0 {

		organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}

		data.SourceOrganizationUnitID = *organizationUnitID

	}

	if data.RequestedBy == 0 {
		userProfileID, ok := params.Context.Value(config.LoggedInProfileKey).(*int)
		if !ok || userProfileID == nil {
			return apierrors.HandleAPIError(fmt.Errorf("error during checking user profile id"))
		}

		data.RequestedBy = *userProfileID
	}

	var item *structs.ExternalReallocation

	item, err = r.Repo.CreateExternalReallocation(&data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	singleItem, err := buildExternalReallocation(*item, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) ExternalReallocationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteExternalReallocation(itemID)
	if err != nil {
		fmt.Printf("Deleting internal reallocation failed because of this error - %s.\n", err)
		return dto.ResponseSingle{
			Status: "failed",
		}, nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildExternalReallocation(item structs.ExternalReallocation, r *Resolver) (*dto.ExternalReallocationResponse, error) {

	response := dto.ExternalReallocationResponse{
		ID:                      item.ID,
		Title:                   item.Title,
		DateOfRequest:           item.DateOfRequest,
		Status:                  item.Status,
		DateOfActionDestOrgUnit: item.DateOfActionDestOrgUnit,
		DateOfActionSSS:         item.DateOfActionSSS,
		CreatedAt:               item.CreatedAt,
		UpdatedAt:               item.UpdatedAt,
	}

	if item.SourceOrganizationUnitID != 0 {
		value, err := r.Repo.GetOrganizationUnitByID(item.SourceOrganizationUnitID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.SourceOrganizationUnit = dropdown
	}

	if item.DestinationOrganizationUnitID != 0 {
		value, err := r.Repo.GetOrganizationUnitByID(item.DestinationOrganizationUnitID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.DestinationOrganizationUnit = dropdown
	}

	if item.RequestedBy != 0 {
		value, err := r.Repo.GetUserProfileByID(item.RequestedBy)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.FirstName + " " + value.LastName,
		}

		response.RequestedBy = dropdown
	}

	if item.AcceptedBy != 0 {
		value, err := r.Repo.GetUserProfileByID(item.AcceptedBy)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.FirstName + " " + value.LastName,
		}

		response.AcceptedBy = dropdown
	}

	if item.FileID != 0 {
		value, err := r.Repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.FileDropdownSimple{
			ID:   value.ID,
			Name: value.Name,
			Type: *value.Type,
		}

		response.File = dropdown
	}

	if item.DestinationOrgUnitFileID != 0 {
		value, err := r.Repo.GetFileByID(item.DestinationOrgUnitFileID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.FileDropdownSimple{
			ID:   value.ID,
			Name: value.Name,
			Type: *value.Type,
		}

		response.DestinationOrgUnitFile = dropdown
	}

	if item.SSSFileID != 0 {
		value, err := r.Repo.GetFileByID(item.SSSFileID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.FileDropdownSimple{
			ID:   value.ID,
			Name: value.Name,
			Type: *value.Type,
		}

		response.SSSFile = dropdown
	}

	if item.BudgetID != 0 {
		value, err := r.Repo.GetBudget(item.BudgetID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: strconv.Itoa(value.Year),
		}

		response.Budget = dropdown
	}

	for _, orderItem := range item.Items {
		builtItem, err := buildExternalReallocationItem(orderItem, r)

		if err != nil {
			return nil, err
		}

		response.Items = append(response.Items, *builtItem)
	}

	return &response, nil
}

func buildExternalReallocationItem(item structs.ExternalReallocationItem, r *Resolver) (*dto.ExternalReallocationItemResponse, error) {
	response := dto.ExternalReallocationItemResponse{
		ID:     item.ID,
		Amount: item.Amount,
	}

	if item.SourceAccountID != 0 {
		value, err := r.Repo.GetAccountItemByID(item.SourceAccountID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.SourceAccount = dropdown
	}

	if item.DestinationAccountID != 0 {
		value, err := r.Repo.GetAccountItemByID(item.DestinationAccountID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.DestinationAccount = dropdown
	}

	return &response, nil
}
