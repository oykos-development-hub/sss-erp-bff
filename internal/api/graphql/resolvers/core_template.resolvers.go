package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) TemplateResolver(params graphql.ResolveParams) (interface{}, error) {

	templateID, templateIDOK := params.Args["template_id"].(int)
	organizationUnitID, organizartionUnitIDOK := params.Args["organization_unit_id"].(int)

	input := dto.TemplateFilter{}

	if templateIDOK && templateID != 0 {
		input.TemplateID = &templateID
	}

	if organizartionUnitIDOK && organizationUnitID != 0 {
		input.OrganizationUnitID = &organizationUnitID
	} else {
		input.OrganizationUnitID = params.Context.Value(config.OrganizationUnitIDKey).(*int)
	}

	res, err := r.Repo.GetTemplateList(input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	var responseItems []dto.TemplatesResponse

	for _, item := range res {
		responseItem, err := r.buildTemplateResponse(item)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		responseItems = append(responseItems, *responseItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   responseItems,
	}, nil
}

func (r *Resolver) TemplateInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Template
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	response := dto.ResponseSingle{
		Status: "success",
	}

	err := r.Repo.CreateTemplate(params.Context, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	response.Message = "You created this item!"

	return response, nil

}

func (r *Resolver) TemplateUpdateResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Template
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	response := dto.ResponseSingle{
		Status: "success",
	}

	err := r.Repo.UpdateTemplate(params.Context, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	response.Message = "You updated this item!"

	return response, nil

}

func (r *Resolver) TemplateItemUpdateResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Template
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	response := dto.ResponseSingle{
		Status: "success",
	}

	err := r.Repo.UpdateTemplateItem(params.Context, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	response.Message = "You updated this item!"

	return response, nil

}

func (r *Resolver) TemplateDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteTemplate(params.Context, itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) buildTemplateResponse(item structs.Template) (*dto.TemplatesResponse, error) {
	responseItem := dto.TemplatesResponse{
		ID: item.ID,
	}

	if item.TemplateID != 0 {
		dropdown, err := r.Repo.GetTemplateByID(item.TemplateID)

		if err != nil {
			return nil, err
		}

		responseItem.Template = dto.DropdownSimple{
			ID:    dropdown.ID,
			Title: dropdown.Title,
		}
	}

	if item.OrganizationUnitID != 0 {
		dropdown, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)

		if err != nil {
			return nil, err
		}

		responseItem.OrganizationUnit = dto.DropdownSimple{
			ID:    dropdown.ID,
			Title: dropdown.Title,
		}
	}

	if item.FileID != 0 {
		dropdown, err := r.Repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, err
		}

		responseItem.File = dto.FileDropdownSimple{
			ID:   dropdown.ID,
			Name: dropdown.Name,
			Type: *dropdown.Type,
		}
	}

	return &responseItem, nil
}
