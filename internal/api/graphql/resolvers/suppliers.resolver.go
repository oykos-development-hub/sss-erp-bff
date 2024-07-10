package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) SuppliersOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	search := params.Args["search"]
	entity := params.Args["entity"]
	parentID := params.Args["parent_id"]

	if id != nil && id.(int) > 0 {
		supplier, err := r.Repo.GetSupplier(id.(int))
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []structs.Suppliers{*supplier},
			Total:   1,
		}, nil

	}
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

	if parentID != nil {
		value := parentID.(int)
		input.ParentID = &value

	}

	res, err := r.Repo.GetSupplierList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   res.Data,
		Total:   res.Total,
	}, nil
}

func (r *Resolver) SuppliersInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Suppliers
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		res, err := r.Repo.UpdateSupplier(itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Message = "You updated this item!"
		response.Item = res
	} else {
		res, err := r.Repo.CreateSupplier(&data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Message = "You created this item!"
		response.Item = res
	}

	return response, nil
}

func (r *Resolver) SuppliersDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteSupplier(itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
