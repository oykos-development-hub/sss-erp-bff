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

var PublicProcurementPlanItemLimitsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.ProcurementOULimitResponseItem
	var input dto.GetProcurementOULimitListInputMS
	if itemID, ok := params.Args["procurement_id"].(int); ok && itemID != 0 {
		input.ItemID = &itemID
	}
	limits, err := getProcurementOULimitList(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	for _, limit := range limits {
		resItem, err := buildProcurementOULimitResponseItem(limit)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		items = append(items, resItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   len(items),
	}, nil
}

var PublicProcurementPlanItemLimitInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementLimit
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateProcurementOULimit(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementOULimitResponseItem(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := createProcurementOULimit(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildProcurementOULimitResponseItem(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func buildProcurementOULimitResponseItem(limit *structs.PublicProcurementLimit) (*dto.ProcurementOULimitResponseItem, error) {
	item, err := getProcurementItem(limit.PublicProcurementId)
	if err != nil {
		return nil, err
	}
	itemDropdown := dto.DropdownSimple{
		Id:    item.Id,
		Title: item.Title,
	}

	organization, err := getOrganizationUnitById(limit.OrganizationUnitId)
	if err != nil {
		return nil, err
	}
	organizationDropdown := dto.DropdownSimple{
		Id:    organization.Id,
		Title: organization.Title,
	}

	res := dto.ProcurementOULimitResponseItem{
		Id:                limit.Id,
		OrganizationUnit:  organizationDropdown,
		PublicProcurement: itemDropdown,
		Limit:             limit.Limit,
	}

	return &res, nil
}

func getProcurementOULimitList(input *dto.GetProcurementOULimitListInputMS) ([]*structs.PublicProcurementLimit, error) {
	res := &dto.GetProcurementOULimitListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.OU_LIMITS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func createProcurementOULimit(limit *structs.PublicProcurementLimit) (*structs.PublicProcurementLimit, error) {
	res := &dto.GetProcurementOULimitResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.OU_LIMITS_ENDPOINT, limit, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateProcurementOULimit(id int, limit *structs.PublicProcurementLimit) (*structs.PublicProcurementLimit, error) {
	res := &dto.GetProcurementOULimitResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.OU_LIMITS_ENDPOINT+"/"+strconv.Itoa(id), limit, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
