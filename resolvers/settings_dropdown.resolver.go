package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/graphql-go/graphql"
)

var SettingsDropdownResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	entity := params.Args["entity"].(string)
	page := params.Args["page"]
	size := params.Args["size"]
	search, searchOk := params.Args["search"].(string)

	var (
		items []structs.SettingsDropdown
		total int
	)

	if id != nil && shared.IsInteger(id) && id != 0 {
		setting, err := getDropdownSettingById(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		items = []structs.SettingsDropdown{*setting}
		total = 1
	} else {
		input := dto.GetSettingsInput{}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if searchOk && search != "" {
			input.Search = &search
		}
		input.Entity = entity

		if value, ok := params.Args["value"].(string); ok && value != "" {
			input.Value = &value
		}

		spew.Dump(params.Args["value"])

		res, err := getDropdownSettings(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		items = res.Data
		total = res.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   total,
	}, nil
}

var SettingsDropdownInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.SettingsDropdown
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	response := dto.ResponseSingle{
		Status: "success",
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		itemRes, err := updateDropdownSettings(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = itemRes

	} else {
		itemRes, err := createDropdownSettings(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = itemRes
	}

	return response, nil

}

var SettingsDropdownDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteDropdownSettings(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func createDropdownSettings(data *structs.SettingsDropdown) (*structs.SettingsDropdown, error) {
	res := &dto.GetDropdownTypeResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.SETTINGS_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteDropdownSettings(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.SETTINGS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func updateDropdownSettings(id int, data *structs.SettingsDropdown) (*structs.SettingsDropdown, error) {
	res := &dto.GetDropdownTypeResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.SETTINGS_ENDPOINT+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getDropdownSettings(input *dto.GetSettingsInput) (*dto.GetDropdownTypesResponseMS, error) {
	res := &dto.GetDropdownTypesResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.SETTINGS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getDropdownSettingById(id int) (*structs.SettingsDropdown, error) {
	res := &dto.GetDropdownTypeResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.SETTINGS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
