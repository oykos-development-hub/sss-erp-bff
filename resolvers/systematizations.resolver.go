package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

var SystematizationsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []dto.SystematizationOverviewResponse
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	organizationUnitId := params.Args["organization_unit_id"]

	if id != nil && shared.IsInteger(id) && id != 0 {
		systematization, err := getSystematizationById(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		systematizationResItem, err := buildSystematizationOverviewResponse(systematization)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		items = []dto.SystematizationOverviewResponse{systematizationResItem}
		total = 1
	} else {
		input := dto.GetSystematizationsInput{}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if shared.IsInteger(organizationUnitId) && organizationUnitId.(int) > 0 {
			organizationUnitId := organizationUnitId.(int)
			input.OrganizationUnitID = &organizationUnitId
		}

		systematizationsResponse, err := getSystematizations(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		for _, systematization := range systematizationsResponse.Data {
			systematizationResItem, err := buildSystematizationOverviewResponse(&systematization)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			items = append(items, systematizationResItem)
		}
		total = systematizationsResponse.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}

var SystematizationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	systematization, err := getSystematizationById(id.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}
	systematizationResItem, err := buildSystematizationOverviewResponse(systematization)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    systematizationResItem,
	}, nil
}

var SystematizationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Systematization
	var systematization *structs.Systematization
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		systematization, err = updateSystematization(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	} else {
		systematization, err = createSystematization(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	systematizationResItem, err := buildSystematizationOverviewResponse(systematization)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    systematizationResItem,
	}, nil
}

var SystematizationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	fmt.Println("Brisanje sistematizacije")
	itemId := params.Args["id"]

	if !shared.IsInteger(itemId) && !(itemId.(int) <= 0) {
		return shared.ErrorResponse("You must pass the item id"), nil
	}

	err := deleteSystematization(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

func getSystematizationById(id int) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.SYSTEMATIZATIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getSystematizations(input *dto.GetSystematizationsInput) (*dto.GetSystematizationsResponseMS, error) {
	res := &dto.GetSystematizationsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.SYSTEMATIZATIONS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func updateSystematization(id int, data *structs.Systematization) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.SYSTEMATIZATIONS_ENDPOINT+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createSystematization(data *structs.Systematization) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.SYSTEMATIZATIONS_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteSystematization(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.SYSTEMATIZATIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func buildSystematizationOverviewResponse(systematization *structs.Systematization) (dto.SystematizationOverviewResponse, error) {
	result := dto.SystematizationOverviewResponse{
		Id:                 systematization.Id,
		UserProfileId:      systematization.UserProfileId,
		OrganizationUnitId: systematization.OrganizationUnitId,
		Description:        systematization.Description,
		SerialNumber:       systematization.SerialNumber,
		Active:             systematization.Active,
		FileId:             systematization.FileId,
		DateOfActivation:   systematization.DateOfActivation,
		Sectors:            &[]dto.OrganizationUnitsSectorResponse{},
		CreatedAt:          systematization.CreatedAt,
		UpdatedAt:          systematization.UpdatedAt,
	}

	// Getting Organization Unit
	var relatedOrganizationUnit, err = getOrganizationUnitById(systematization.OrganizationUnitId)
	if err != nil {
		return result, err
	}
	result.OrganizationUnit = relatedOrganizationUnit

	// Getting Sectors
	inputOrganizationUnits := dto.GetOrganizationUnitsInput{
		ParentID: &systematization.OrganizationUnitId,
	}
	organizationUnitsResponse, err := getOrganizationUnits(&inputOrganizationUnits)
	if err != nil {
		return result, err
	}
	for _, organizationUnit := range organizationUnitsResponse.Data {
		*result.Sectors = append(*result.Sectors, *dto.ToOrganizationUnitsSectorResponse(organizationUnit))
	}

	// Getting Job positions
	if result.Sectors != nil {
		for i, sector := range *result.Sectors {
			(*result.Sectors)[i].JobPositions, err = getJobPositionsForSector(sector.Id)
			if err != nil {
				return result, err
			}
		}
	}

	return result, err
}

func getJobPositionsForSector(sectorId int) (*[]structs.JobPositions, error) {
	result := []structs.JobPositions{}
	input := dto.GetJobPositionInOrganizationUnitsInput{
		OrganizationUnitID: &sectorId,
	}
	jobPositionsInOrganizationUnitsResponse, err := getJobPositionsInOrganizationUnits(&input)
	if err != nil {
		return &result, err
	}

	for _, jobPositionsInOrganizationUnits := range jobPositionsInOrganizationUnitsResponse.Data {
		getJobPositionResponse, err := getJobPositionById(jobPositionsInOrganizationUnits.JobPositionId)
		if err != nil {
			return &result, err
		}
		result = append(result, *getJobPositionResponse)
	}

	return &result, nil
}
