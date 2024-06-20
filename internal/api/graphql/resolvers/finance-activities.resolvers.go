package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ActivitiesOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		activity, err := r.Repo.GetActivity(id)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*structs.ActivitiesItem{activity},
			Total:   1,
		}, nil
	}

	input := dto.GetFinanceActivityListInputMS{}
	if organizationUnitID, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitID != 0 {
		input.OrganizationUnitID = &organizationUnitID
	}
	if subProgramID, ok := params.Args["sub_program_id"].(int); ok && subProgramID != 0 {
		input.SubProgramID = &subProgramID
	}
	if search, ok := params.Args["search"].(string); ok {
		input.Search = &search
	}

	activities, err := r.Repo.GetActivityList(&input)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	activityResItemList, err := buildActivityResItemList(r.Repo, activities)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   activityResItemList,
		Total:   len(activityResItemList),
	}, nil
}

func buildActivityResItemList(r repository.MicroserviceRepositoryInterface, activitys []structs.ActivitiesItem) (activityResItemList []*dto.ActivityResItem, err error) {
	for _, activity := range activitys {
		activity, err := buildActivityResItem(r, activity)
		if err != nil {
			return nil, errors.Wrap(err, "build activity res item")
		}
		activityResItemList = append(activityResItemList, activity)
	}

	return
}

func buildActivityResItem(r repository.MicroserviceRepositoryInterface, activity structs.ActivitiesItem) (*dto.ActivityResItem, error) {
	resItem := &dto.ActivityResItem{
		ID:          activity.ID,
		Title:       activity.Title,
		Description: activity.Description,
		Code:        activity.Code,
	}
	subProgram, err := r.GetProgram(activity.SubProgramID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get program")
	}
	resItem.SubProgram = dto.DropdownSimple{ID: subProgram.ID, Title: subProgram.Title}

	program, err := r.GetProgram(*subProgram.ParentID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get program")
	}
	resItem.Program = dto.DropdownSimple{ID: program.ID, Title: program.Title}

	organizationUnit, err := r.GetOrganizationUnitByID(activity.OrganizationUnitID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get organization unit by id")
	}
	resItem.OrganizationUnit = *organizationUnit

	return resItem, nil
}

func (r *Resolver) ActivityInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.ActivitiesItem
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
		item, err := r.Repo.UpdateActivity(params.Context, itemID, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		resItem, err := buildActivityResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		item, err := r.Repo.CreateActivity(params.Context, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		resItem, err := buildActivityResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil
}

func (r *Resolver) ActivityDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteActivity(params.Context, itemID)
	if err != nil {
		fmt.Printf("Deleting activity item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
