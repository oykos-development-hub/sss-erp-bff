package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

var OrganizationUnitsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []dto.OrganizationUnitsOverviewResponse
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	parent_id := params.Args["parent_id"]
	search, searchOk := params.Args["search"].(string)

	if id != nil && shared.IsInteger(id) && id != 0 {
		organizationUnit, err := getOrganizationUnitById(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}

		organizationUnitItem, err := buildOrganizationUnitOverviewResponse(organizationUnit)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		items = []dto.OrganizationUnitsOverviewResponse{*organizationUnitItem}
		total = 1
	} else {
		input := dto.GetOrganizationUnitsInput{}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if shared.IsInteger(parent_id) && parent_id.(int) > 0 {
			parentID := parent_id.(int)
			input.ParentID = &parentID
		}
		if searchOk && search != "" {
			input.Search = &search
		}

		organizationUnits, err := getOrganizationUnits(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, organizationUnit := range organizationUnits.Data {
			organizationUnitItem, err := buildOrganizationUnitOverviewResponse(&organizationUnit)
			if err != nil {
				return shared.HandleAPIError(err)
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

var OrganizationUnitInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrganizationUnits
	var organizationUnitResponse *dto.GetOrganizationUnitResponseMS
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		organizationUnitResponse, err = updateOrganizationUnits(itemId, &data)
	} else {
		organizationUnitResponse, err = createOrganizationUnits(&data)
	}

	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    organizationUnitResponse.Data,
	}, nil

}

var OrganizationUnitDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	if !shared.IsInteger(itemId) && !(itemId.(int) <= 0) {
		return shared.ErrorResponse("You must pass the item id"), nil
	}

	err := deleteOrganizationUnits(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil

}

func deleteOrganizationUnits(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getOrganizationUnits(input *dto.GetOrganizationUnitsInput) (*dto.GetOrganizationUnitsResponseMS, error) {
	res := &dto.GetOrganizationUnitsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ORGANIZATION_UNITS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getOrganizationUnitById(id int) (*structs.OrganizationUnits, error) {
	res := &dto.GetOrganizationUnitResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateOrganizationUnits(id int, data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error) {
	res := &dto.GetOrganizationUnitResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func createOrganizationUnits(data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error) {
	res := &dto.GetOrganizationUnitResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ORGANIZATION_UNITS_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getOrganizationUnitIdByUserProfile(id int) (*int, error) {
	employeesInOrganizationUnit, err := getEmployeesInOrganizationUnitsByProfileId(id)
	if err != nil {
		return nil, err
	}

	if employeesInOrganizationUnit == nil {
		return nil, nil
	}

	jobPositionInOrganizationUnit, err := getJobPositionsInOrganizationUnitsById(employeesInOrganizationUnit.PositionInOrganizationUnitId)
	if err != nil {
		return nil, err
	}

	systematization, err := getSystematizationById(jobPositionInOrganizationUnit.SystematizationId)
	if err != nil {
		return nil, err
	}

	return &systematization.OrganizationUnitId, nil
}

func buildOrganizationUnitOverviewResponse(
	organizationUnits *structs.OrganizationUnits,
) (*dto.OrganizationUnitsOverviewResponse, error) {
	input := dto.GetOrganizationUnitsInput{}
	input.ParentID = &organizationUnits.Id

	organizationUnitsChildrenResponse, err := getOrganizationUnits(&input)
	if err != nil {
		return nil, err
	}

	return &dto.OrganizationUnitsOverviewResponse{
		Id:             organizationUnits.Id,
		ParentId:       organizationUnits.ParentId,
		NumberOfJudges: organizationUnits.NumberOfJudges,
		Title:          organizationUnits.Title,
		Abbreviation:   organizationUnits.Abbreviation,
		Color:          organizationUnits.Color,
		Description:    organizationUnits.Description,
		Address:        organizationUnits.Address,
		Icon:           organizationUnits.Icon,
		Children:       &organizationUnitsChildrenResponse.Data,
		FolderId:       organizationUnits.FolderId,
	}, nil
}
