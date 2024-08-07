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
	active := params.Args["active"].(bool)

	if id != nil && id != 0 {
		organizationUnit, err := r.Repo.GetOrganizationUnitByID(id.(int))
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		organizationUnitItem, err := buildOrganizationUnitOverviewResponse(r.Repo, organizationUnit)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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

		input.Active = &active

		organizationUnits, err := r.Repo.GetOrganizationUnits(&input)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		profileOrganizationUnit := params.Context.Value(config.OrganizationUnitIDKey).(*int)

		hasPresident, hasPresidentOk := params.Args["has_president"].(bool)
		organizationUnitsWithPresident := make(map[int]string)

		if hasPresidentOk {
			active := true
			resolution, err := r.Repo.GetJudgeResolutionList(&dto.GetJudgeResolutionListInputMS{Active: &active})
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			if len(resolution.Data) > 0 {

				for _, item := range organizationUnits.Data {
					_, numberOfPresidents, _, _, err := calculateEmployeeStats(r.Repo, item.ID, resolution.Data[0].ID)
					if err != nil {
						_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
						return errors.HandleAPPError(err)
					}

					if numberOfPresidents == 1 {
						organizationUnitsWithPresident[item.ID] = item.Title
					}
				}

			}
		}

		for _, organizationUnit := range organizationUnits.Data {
			organizationUnitItem, err := buildOrganizationUnitOverviewResponse(r.Repo, &organizationUnit)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			if !disableFilters {

				loggedInAccount := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

				hasPermission, err := r.HasPermission(*loggedInAccount, string(config.HR), config.OperationFullAccess)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}

				// Initialize isOwnOrChildUnit as false
				isOwnOrChildUnit := false

				// Check if the current unit is the user's own unit
				if hasPermission || *profileOrganizationUnit == organizationUnitItem.ID {
					isOwnOrChildUnit = true
				}

				// Check if the current unit is a child of the user's unit
				if hasPermission || organizationUnitItem.ParentID != nil && *profileOrganizationUnit == *organizationUnitItem.ParentID {
					isOwnOrChildUnit = true
				}

				if !hasPermission && !isOwnOrChildUnit && !settings {
					continue
				}

				if hasPresidentOk {
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

	if !data.Active {
		data.Active = true
	}

	itemID := data.ID
	if itemID != 0 {
		organizationUnitResponse, err = r.Repo.UpdateOrganizationUnits(params.Context, itemID, &data)
	} else {
		organizationUnitResponse, err = r.Repo.CreateOrganizationUnits(params.Context, &data)
	}

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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

	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	for _, item := range data {
		if !item.Active {
			item.Active = true
		}

		organizationUnit, err := r.Repo.UpdateOrganizationUnits(params.Context, item.ID, &item)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		organizationUnitResponse = append(organizationUnitResponse, *organizationUnit)
	}

	return dto.Response{
		Status:  "success",
		Message: "You updated this items!",
		Items:   organizationUnitResponse,
	}, nil

}

func (r *Resolver) OrganizationUnitDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	organizationUnit, err := r.Repo.GetOrganizationUnitByID(itemID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if organizationUnit != nil {

		if organizationUnit.ParentID == nil && *organizationUnit.ParentID == 0 {
			return map[string]interface{}{
				"status":  "failed",
				"message": "You can not delete this item!",
			}, nil
		}

		jobPositions, err := r.Repo.GetJobPositionsInOrganizationUnits(&dto.GetJobPositionInOrganizationUnitsInput{
			OrganizationUnitID: &organizationUnit.ID,
		})

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		isDeletable := true

		for _, item := range jobPositions.Data {
			systematization, err := r.Repo.GetSystematizationByID(item.SystematizationID)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			if systematization.Active != 1 {
				isDeletable = false
			}
		}

		if isDeletable {
			organizationUnit.Active = false

			_, err = r.Repo.UpdateOrganizationUnits(params.Context, itemID, organizationUnit)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		} else {
			return map[string]interface{}{
				"status":  "failed",
				"message": "You must first delete job positions in organization unit!",
			}, nil
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deactivated this item!",
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
		return nil, errors.Wrap(err, "repo get organization units")
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
		BankAccounts:   organizationUnits.BankAccounts,
		Code:           organizationUnits.Code,
		Active:         organizationUnits.Active,
	}, nil
}
