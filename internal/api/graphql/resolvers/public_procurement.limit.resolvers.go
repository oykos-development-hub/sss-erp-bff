package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) PublicProcurementPlanItemLimitsResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.ProcurementOULimitResponseItem
	var input dto.GetProcurementOULimitListInputMS
	if itemID, ok := params.Args["procurement_id"].(int); ok && itemID != 0 {
		input.ItemID = &itemID
	}
	limits, err := r.Repo.GetProcurementOULimitList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	for _, limit := range limits {
		resItem, err := buildProcurementOULimitResponseItem(r.Repo, limit)
		if err != nil {
			return errors.HandleAPIError(err)
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

func (r *Resolver) PublicProcurementPlanItemLimitInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PublicProcurementLimit
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		res, err := r.Repo.UpdateProcurementOULimit(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		item, err := buildProcurementOULimitResponseItem(r.Repo, res)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		res, err := r.Repo.CreateProcurementOULimit(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		item, err := buildProcurementOULimitResponseItem(r.Repo, res)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func buildProcurementOULimitResponseItem(r repository.MicroserviceRepositoryInterface, limit *structs.PublicProcurementLimit) (*dto.ProcurementOULimitResponseItem, error) {
	item, err := r.GetProcurementItem(limit.PublicProcurementID)
	if err != nil {
		return nil, err
	}
	itemDropdown := dto.DropdownSimple{
		ID:    item.ID,
		Title: item.Title,
	}

	organization, err := r.GetOrganizationUnitByID(limit.OrganizationUnitID)
	if err != nil {
		return nil, err
	}
	organizationDropdown := dto.DropdownSimple{
		ID:    organization.ID,
		Title: organization.Title,
	}

	res := dto.ProcurementOULimitResponseItem{
		ID:                limit.ID,
		OrganizationUnit:  organizationDropdown,
		PublicProcurement: itemDropdown,
		Limit:             limit.Limit,
	}

	return &res, nil
}
