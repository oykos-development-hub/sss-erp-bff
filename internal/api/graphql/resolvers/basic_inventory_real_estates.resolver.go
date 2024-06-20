package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) BasicInventoryRealEstatesOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var input dto.GetInventoryRealEstateListInputMS
	if page, ok := params.Args["page"].(int); ok && page != 0 {
		input.Page = &page
	}
	if size, ok := params.Args["size"].(int); ok && size != 0 {
		input.Size = &size
	}

	if id, ok := params.Args["id"].(int); ok && id != 0 {
		realEstate, err := r.Repo.GetInventoryRealEstate(id)
		if err != nil {
			return errors.HandleAPPError(err)
		}
		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*structs.BasicInventoryRealEstatesItem{realEstate},
			Total:   1,
		}, nil
	}

	res, err := r.Repo.GetInventoryRealEstatesList(&input)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   res.Data,
		Total:   res.Total,
	}, nil
}
