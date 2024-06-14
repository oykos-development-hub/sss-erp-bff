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

func (r *Resolver) InternalReallocationOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		InternalReallocation, err := r.Repo.GetInternalReallocationByID(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		res, err := buildInternalReallocation(*InternalReallocation, r)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.InternalReallocationResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.InternalReallocationFilter{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	} else {
		input.OrganizationUnitID, _ = params.Context.Value(config.OrganizationUnitIDKey).(*int)
	}

	if value, ok := params.Args["year"].(int); ok && value != 0 {
		input.Year = &value
	}

	if value, ok := params.Args["requested_by"].(int); ok && value != 0 {
		input.RequestedBy = &value
	}

	if value, ok := params.Args["budget_id"].(int); ok && value != 0 {
		input.BudgetID = &value
	}

	items, total, err := r.Repo.GetInternalReallocationList(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var resItems []dto.InternalReallocationResponse
	for _, item := range items {
		resItem, err := buildInternalReallocation(item, r)

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

func (r *Resolver) InternalReallocationInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.InternalReallocation
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

	if data.OrganizationUnitID == 0 {

		organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}

		data.OrganizationUnitID = *organizationUnitID

	}

	if data.RequestedBy == 0 {
		userProfile, ok := params.Context.Value(config.LoggedInProfileKey).(*structs.UserProfiles)
		if !ok || userProfile == nil {
			return apierrors.HandleAPIError(fmt.Errorf("error during checking user profile id"))
		}

		data.RequestedBy = userProfile.ID
	}

	var item *structs.InternalReallocation

	item, err = r.Repo.CreateInternalReallocation(params.Context, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	singleItem, err := buildInternalReallocation(*item, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) InternalReallocationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteInternalReallocation(params.Context, itemID)
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

func buildInternalReallocation(item structs.InternalReallocation, r *Resolver) (*dto.InternalReallocationResponse, error) {

	response := dto.InternalReallocationResponse{
		ID:            item.ID,
		Title:         item.Title,
		DateOfRequest: item.DateOfRequest,
		Sum:           item.Sum,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}

	if item.OrganizationUnitID != 0 {
		value, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)

		if err != nil {
			return nil, err
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: value.Title,
		}

		response.OrganizationUnit = dropdown
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
		builtItem, err := buildInternalReallocationItem(orderItem, r)

		if err != nil {
			return nil, err
		}

		response.Items = append(response.Items, *builtItem)
	}

	return &response, nil
}

func buildInternalReallocationItem(item structs.InternalReallocationItem, r *Resolver) (*dto.InternalReallocationItemResponse, error) {
	response := dto.InternalReallocationItemResponse{
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
