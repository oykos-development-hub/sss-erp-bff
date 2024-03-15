package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) OrganizationUnitsResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []dto.OrganizationUnitsOverviewResponse
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	parentID := params.Args["parent_id"]
	search, searchOk := params.Args["search"].(string)
	settings := params.Args["settings"].(bool)
	disableFilters := params.Args["disable_filters"].(bool)

	if id != nil && id != 0 {
		organizationUnit, err := r.Repo.GetOrganizationUnitByID(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}

		organizationUnitItem, err := buildOrganizationUnitOverviewResponse(r.Repo, organizationUnit)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		items = []dto.OrganizationUnitsOverviewResponse{*organizationUnitItem}
		total = 1
	} else {
		input := dto.GetOrganizationUnitsInput{}
		if page != nil && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if size != nil && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if parentID != nil && parentID.(int) > 0 {
			parentID := parentID.(int)
			input.ParentID = &parentID
		}
		if searchOk && search != "" {
			input.Search = &search
		}

		organizationUnits, err := r.Repo.GetOrganizationUnits(&input)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		loggedInAccount := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
		profileOrganizationUnit := params.Context.Value(config.OrganizationUnitIDKey).(*int)

		active := true
		resolution, err := r.Repo.GetJudgeResolutionList(&dto.GetJudgeResolutionListInputMS{Active: &active})
		if err != nil {
			return errors.HandleAPIError(err)
		}

		organizationUnitsWithPresident := make(map[int]string)
		if len(resolution.Data) > 0 {

			for _, item := range organizationUnits.Data {
				_, numberOfPresidents, _, _, err := calculateEmployeeStats(r.Repo, item.ID, resolution.Data[0].ID)
				if err != nil {
					return errors.HandleAPIError(err)
				}

				if numberOfPresidents == 1 {
					organizationUnitsWithPresident[item.ID] = item.Title
				}
			}

		}

		if err != nil {
			return dto.ErrorResponse(err), nil
		}

		for _, organizationUnit := range organizationUnits.Data {
			organizationUnitItem, err := buildOrganizationUnitOverviewResponse(r.Repo, &organizationUnit)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			if !disableFilters {
				hasGeneralPermission := loggedInAccount.HasPermission(structs.PermissionManageOrganizationUnits)

				// Initialize isOwnOrChildUnit as false
				isOwnOrChildUnit := false

				// Check if the current unit is the user's own unit
				if *profileOrganizationUnit == organizationUnitItem.ID {
					isOwnOrChildUnit = true
				}

				// Check if the current unit is a child of the user's unit
				if organizationUnitItem.ParentID != nil && *profileOrganizationUnit == *organizationUnitItem.ParentID {
					isOwnOrChildUnit = true
				}

				if !hasGeneralPermission && !isOwnOrChildUnit && !settings {
					continue
				}

				hasPresident, ok := params.Args["has_president"].(bool)
				if ok {
					_, exists := organizationUnitsWithPresident[organizationUnit.ID]

					if hasPresident && !exists || !hasPresident && exists {
						continue
					}
				}
			}
			items = append(items, *organizationUnitItem)
		}
		total = organizationUnits.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}

func (r *Resolver) OrganizationUnitInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrganizationUnits
	var organizationUnitResponse *dto.GetOrganizationUnitResponseMS
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		organizationUnitResponse, err = r.Repo.UpdateOrganizationUnits(itemID, &data)
	} else {
		organizationUnitResponse, err = r.Repo.CreateOrganizationUnits(&data)
	}

	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    organizationUnitResponse.Data,
	}, nil
}

func (r *Resolver) OrganizationUnitOrderResolver(params graphql.ResolveParams) (interface{}, error) {
	var data []structs.OrganizationUnits
	var organizationUnitResponse []dto.GetOrganizationUnitResponseMS
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	for _, item := range data {
		organizationUnit, err := r.Repo.UpdateOrganizationUnits(item.ID, &item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		organizationUnitResponse = append(organizationUnitResponse, *organizationUnit)
	}

	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You updated this items!",
		Items:   organizationUnitResponse,
	}, nil

}

func (r *Resolver) OrganizationUnitDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteOrganizationUnits(itemID.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil

}

func buildOrganizationUnitOverviewResponse(
	r repository.MicroserviceRepositoryInterface,
	organizationUnits *structs.OrganizationUnits,
) (*dto.OrganizationUnitsOverviewResponse, error) {
	input := dto.GetOrganizationUnitsInput{}
	input.ParentID = &organizationUnits.ID

	organizationUnitsChildrenResponse, err := r.GetOrganizationUnits(&input)
	if err != nil {
		return nil, err
	}

	return &dto.OrganizationUnitsOverviewResponse{
		ID:             organizationUnits.ID,
		ParentID:       organizationUnits.ParentID,
		NumberOfJudges: organizationUnits.NumberOfJudges,
		Title:          organizationUnits.Title,
		Pib:            organizationUnits.Pib,
		Abbreviation:   organizationUnits.Abbreviation,
		Color:          organizationUnits.Color,
		City:           organizationUnits.City,
		Description:    organizationUnits.Description,
		Address:        organizationUnits.Address,
		Icon:           organizationUnits.Icon,
		Children:       &organizationUnitsChildrenResponse.Data,
		FolderID:       organizationUnits.FolderID,
	}, nil
}
