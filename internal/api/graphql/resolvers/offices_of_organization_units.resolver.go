package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

func PopulateOfficesOfOrganizationUnitItemProperties(basicInventoryItems []interface{}, id int, organizationUnitID int, search string) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if id != 0 && id != mergedItem["id"] {
			continue
		}

		// Filtering by organizationUnitID
		if organizationUnitID != 0 && organizationUnitID != mergedItem["organization_unit_id"] {
			continue
		}

		// Filtering by status
		if len(search) > 0 && !shared.StringContains(mergedItem["title"].(string), search) {
			continue
		}

		if mergedItem["organization_unit_id"].(int) > 0 {
			var relatedOfficesOrganizationUnit = shared.FetchByProperty(
				"organization_unit",
				"ID",
				mergedItem["organization_unit_id"],
			)
			if len(relatedOfficesOrganizationUnit) > 0 {
				var relatedOrganizationUnit = shared.WriteStructToInterface(relatedOfficesOrganizationUnit[0])

				mergedItem["organization_unit"] = map[string]interface{}{
					"title": relatedOrganizationUnit["title"],
					"id":    relatedOrganizationUnit["id"],
				}
			}
		} else {
			continue
		}

		items = append(items, mergedItem)
	}

	return items
}

func (r *Resolver) OfficesOfOrganizationUnitOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	var items []*dto.OfficesOfOrganizationResponse
	var total int

	if id != nil && id.(int) != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}

		if setting.Entity != "office_types" {

			return errors.HandleAPIError(fmt.Errorf("not found"))
		}

		item, err := buildOfficeOfOrganizationUnit(r.Repo, setting)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		items = append(items, item)

		total = 1
	} else {
		input := dto.GetOfficesOfOrganizationInput{}

		if search, ok := params.Args["search"].(string); ok && search != "" {
			input.Search = &search
		}

		if page, ok := params.Args["page"].(int); ok && page != 0 {
			input.Page = &page
		}

		if size, ok := params.Args["size"].(int); ok && size != 0 {
			input.Size = &size
		}

		if orgUnit, ok := params.Args["organization_unit_id"].(int); ok && orgUnit != 0 {
			organizationUnit := strconv.Itoa(orgUnit)
			input.Value = &organizationUnit
		}

		res, err := r.Repo.GetOfficeDropdownSettings(&input)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		data := res.Data

		for i := 0; i < len(data); i++ {
			item, err := buildOfficeOfOrganizationUnit(r.Repo, &data[i])
			if err != nil {
				return errors.HandleAPIError(err)
			}
			items = append(items, item)
		}
		total = res.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   total,
	}, nil
}

func (r *Resolver) OfficesOfOrganizationUnitInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OfficesOfOrganizationUnitItem
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	response := dto.ResponseSingle{
		Status: "success",
	}

	arg := structs.SettingsDropdown{
		Value:        strconv.Itoa(data.OrganizationUnitID),
		Title:        data.Title,
		Abbreviation: data.Abbreviation,
		Description:  data.Description,
		Color:        data.Color,
		Icon:         data.Icon,
		Entity:       config.OfficeTypes,
	}

	if itemID != 0 {
		itemRes, err := r.Repo.UpdateDropdownSettings(itemID, &arg)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"

		item, err := buildOfficeOfOrganizationUnit(r.Repo, itemRes)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Item = item

	} else {
		itemRes, err := r.Repo.CreateDropdownSettings(&arg)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"

		item, err := buildOfficeOfOrganizationUnit(r.Repo, itemRes)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Item = item
	}

	return response, nil
}

func (r *Resolver) OfficesOfOrganizationUnitDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	setting, err := r.Repo.GetDropdownSettingByID(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	if setting.Entity != config.OfficeTypes {
		return dto.ResponseSingle{
			Status:  "failed",
			Message: "You can not delete this item! (Item entity must be 'office type')",
		}, nil
	}

	err = r.Repo.DeleteDropdownSettings(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildOfficeOfOrganizationUnit(r repository.MicroserviceRepositoryInterface, item *structs.SettingsDropdown) (*dto.OfficesOfOrganizationResponse, error) {
	organizationUnitID, err := strconv.Atoi(item.Value)

	if err != nil {
		return nil, err
	}

	organizationUnitDropdown := dto.DropdownSimple{}
	if organizationUnitID != 0 {
		organizationUnit, err := r.GetOrganizationUnitByID(organizationUnitID)
		if err != nil {
			return nil, err
		}
		if organizationUnit != nil {
			organizationUnitDropdown = dto.DropdownSimple{ID: organizationUnit.ID, Title: organizationUnit.Title}
		}
	}

	data := dto.OfficesOfOrganizationResponse{
		ID:               item.ID,
		OrganizationUnit: organizationUnitDropdown,
		Title:            item.Title,
		Abbreviation:     item.Abbreviation,
		Description:      item.Description,
		Icon:             item.Icon,
		Color:            item.Color,
	}

	return &data, nil
}
