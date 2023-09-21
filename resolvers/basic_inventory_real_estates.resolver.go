package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"strconv"

	"github.com/graphql-go/graphql"
)

var BasicInventoryRealEstatesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var input dto.GetInventoryRealEstateListInputMS
	if page, ok := params.Args["page"].(int); ok && page != 0 {
		input.Page = &page
	}
	if size, ok := params.Args["size"].(int); ok && size != 0 {
		input.Size = &size
	}

	if id, ok := params.Args["id"].(int); ok && id != 0 {
		realEstate, err := getInventoryRealEstate(id)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*structs.BasicInventoryRealEstatesItem{realEstate},
			Total:   1,
		}, nil
	}

	res, err := getInventoryRealEstatesList(&input)
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

func getInventoryRealEstatesList(input *dto.GetInventoryRealEstateListInputMS) (*dto.GetInventoryRealEstateListResponseMS, error) {
	res := &dto.GetInventoryRealEstateListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.REAL_ESTATES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getInventoryRealEstate(id int) (*structs.BasicInventoryRealEstatesItem, error) {
	res := &dto.GetInventoryRealEstateResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.REAL_ESTATES_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
