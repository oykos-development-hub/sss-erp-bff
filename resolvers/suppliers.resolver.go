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

var SuppliersOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	search := params.Args["search"]
	entity := params.Args["entity"]

	if shared.IsInteger(id) && id.(int) > 0 {
		supplier, err := getSupplier(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []structs.Suppliers{*supplier},
			Total:   1,
		}, nil

	} else {
		input := dto.GetSupplierInputMS{}
		if search != nil {
			searchValue := search.(string)
			input.Search = &searchValue

		}

		if entity != nil {
			entityValue := entity.(string)
			input.Entity = &entityValue

		}
		if page != nil && size != nil {
			pageValue := page.(int)
			sizeValue := size.(int)

			input.Size = &sizeValue
			input.Page = &pageValue

		}

		res, err := getSupplierList(&input)
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
}

var SuppliersInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Suppliers
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
		res, err := updateSupplier(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = res
	} else {
		res, err := createSupplier(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = res
	}

	return response, nil
}

var SuppliersDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteSupplier(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func createSupplier(supplier *structs.Suppliers) (*structs.Suppliers, error) {
	res := &dto.GetSupplierResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.SUPPLIERS_ENDPOINT, supplier, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateSupplier(id int, supplier *structs.Suppliers) (*structs.Suppliers, error) {
	res := &dto.GetSupplierResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.SUPPLIERS_ENDPOINT+"/"+strconv.Itoa(id), supplier, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteSupplier(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.SUPPLIERS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getSupplier(id int) (*structs.Suppliers, error) {
	res := &dto.GetSupplierResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.SUPPLIERS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getSupplierList(input *dto.GetSupplierInputMS) (*dto.GetSupplierListResponseMS, error) {
	res := &dto.GetSupplierListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.SUPPLIERS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
