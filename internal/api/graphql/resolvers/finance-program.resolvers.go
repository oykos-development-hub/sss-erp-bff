package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func ProgramItemProperties(basicInventoryItems []interface{}, id int, typeProgram string) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if id != 0 && id != mergedItem["id"] {
			continue
		}

		// Filtering by program
		if len(typeProgram) > 0 && typeProgram == "program" && mergedItem["parent_id"].(int) > 0 {
			continue
		}

		// Filtering by activities
		if len(typeProgram) > 0 && typeProgram == "activities" && mergedItem["organization_unit_id"].(int) == 0 {
			continue
		}

		// Filtering by subprogram
		if len(typeProgram) > 0 && typeProgram == "subprogram" && mergedItem["organization_unit_id"].(int) > 0 || mergedItem["parent_id"].(int) == 0 {
			continue
		}

		if shared.IsInteger(mergedItem["parent_id"]) && mergedItem["parent_id"].(int) > 0 {
			var relatedProgram = shared.FetchByProperty(
				"program",
				"ID",
				mergedItem["parent_id"],
			)
			if len(relatedProgram) > 0 {
				var relatedProgram = shared.WriteStructToInterface(relatedProgram[0])

				mergedItem["parent"] = map[string]interface{}{
					"title": relatedProgram["title"],
					"id":    relatedProgram["id"],
				}
			}
		}

		if shared.IsInteger(mergedItem["organization_unit_id"]) && mergedItem["organization_unit_id"].(int) > 0 {
			var relatedOfficesOrganizationUnit = shared.FetchByProperty(
				"organization_unit",
				"ID",
				mergedItem["organization_unit_id"],
			)
			if len(relatedOfficesOrganizationUnit) > 0 {
				var relatedOrganizationUnit = shared.WriteStructToInterface(relatedOfficesOrganizationUnit[0])

				mergedItem["organization_unit"] = map[string]interface{}{
					"title": relatedOrganizationUnit["title"],
					"id":    relatedOrganizationUnit["id"],
				}
			}
		}

		items = append(items, mergedItem)
	}

	return items
}

func (r *Resolver) ProgramOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		program, err := r.Repo.GetProgram(id)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*structs.ProgramItem{program},
			Total:   1,
		}, nil
	}

	input := dto.GetFinanceProgramListInputMS{}
	if isProgram, ok := params.Args["is_program"].(bool); ok {
		input.IsProgram = &isProgram
	}
	if search, ok := params.Args["search"].(string); ok {
		input.Search = &search
	}

	programs, err := r.Repo.GetProgramList(&input)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	programResItem, err := buildProgramResItemList(r.Repo, programs)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   programResItem,
		Total:   len(programResItem),
	}, nil
}

func buildProgramResItemList(r repository.MicroserviceRepositoryInterface, programs []structs.ProgramItem) (programResItemList []*dto.ProgramResItem, err error) {
	for _, program := range programs {
		program, err := buildProgramResItem(r, program)
		if err != nil {
			return nil, errors.Wrap(err, "build program res item")
		}
		programResItemList = append(programResItemList, program)
	}

	return
}

func buildProgramResItem(r repository.MicroserviceRepositoryInterface, program structs.ProgramItem) (*dto.ProgramResItem, error) {
	resItem := &dto.ProgramResItem{
		ID:          program.ID,
		Title:       program.Title,
		Description: program.Description,
		Code:        program.Code,
	}
	if program.ParentID != nil {
		subProgram, err := r.GetProgram(*program.ParentID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get prorgam")
		}
		resItem.Parent = &dto.DropdownSimple{ID: subProgram.ID, Title: subProgram.Title}
	}
	return resItem, nil
}

func (r *Resolver) ProgramInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.ProgramItem
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		item, err := r.Repo.UpdateProgram(params.Context, itemID, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		resItem, err := buildProgramResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		item, err := r.Repo.CreateProgram(params.Context, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		resItem, err := buildProgramResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil
}

func (r *Resolver) ProgramDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteProgram(params.Context, itemID)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func DeleteProgramChildren(itemID int, data []interface{}) []interface{} {
	toDelete := map[int]bool{itemID: true}
	prevLen := 0

	for len(toDelete) != prevLen {
		prevLen = len(toDelete)
		for _, item := range data {
			if item, ok := item.(*structs.ProgramItem); ok {
				id := item.ID
				parentID := item.ParentID
				if _, exists := toDelete[*parentID]; exists {
					toDelete[id] = true
				}
			}
		}
	}

	var result []interface{}
	for _, item := range data {
		if item, ok := item.(*structs.ProgramItem); ok {
			id := item.ID
			if _, exists := toDelete[id]; !exists {
				result = append(result, item)
			}
		}
	}

	return result
}
