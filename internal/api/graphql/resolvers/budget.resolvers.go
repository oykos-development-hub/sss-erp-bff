package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/shared"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

func BudgetItemProperties(basicInventoryItems []interface{}, id int, typeBudget string, year string, status string) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if id != 0 && id != mergedItem["id"] {
			continue
		}
		// Filtering by type
		if len(typeBudget) > 0 && typeBudget != mergedItem["type"] {
			continue
		}
		// Filtering by year
		if len(year) > 0 && year != mergedItem["year"] {
			continue
		}
		// Filtering by status
		if len(status) > 0 && status != mergedItem["status"] {
			continue
		}

		items = append(items, mergedItem)
	}

	return items
}

func (r *Resolver) BudgetOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		budget, err := r.Repo.GetBudget(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		budgetResItem, err := buildBudgetResponseItem(params.Context, *budget)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.BudgetResponseItem{budgetResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetBudgetListInputMS{}
	if budgetType, ok := params.Args["budget_type"].(int); ok && budgetType != 0 {
		input.BudgetType = &budgetType
	}
	if year, ok := params.Args["year"].(int); ok && year != 0 {
		input.Year = &year
	}
	if status, ok := params.Args["status"].(string); ok && status != "" {
		input.Status = &status
	}

	budgets, err := r.Repo.GetBudgetList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	budgetResItem, err := buildBudgetResponseItemList(params.Context, budgets)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   budgetResItem,
		Total:   len(budgets),
	}, nil
}

func (r *Resolver) FinancialBudgetOverview(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{Version: &financialBudget.AccountVersion})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	accountResItemlist, err := buildAccountItemResponseItemList(accounts.Data)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	accountResItemlist, err = CreateTree(accountResItemlist)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	financialBudgetOveriew := &dto.FinancialBudgetOverviewResponse{
		AccountVersion: financialBudget.AccountVersion,
		Accounts:       accountResItemlist,
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    financialBudgetOveriew,
	}, nil
}

// TODO: add logic  to determine budget status based on states and logged in user
func buildBudgetStatus(ctx context.Context, budget structs.Budget) string {
	return "Kreiran"
}

func buildBudgetResponseItemList(ctx context.Context, budgetList []structs.Budget) (budgetResItemList []*dto.BudgetResponseItem, err error) {
	for _, budget := range budgetList {
		status := buildBudgetStatus(ctx, budget)
		budgetResItemList = append(budgetResItemList, &dto.BudgetResponseItem{
			ID:         budget.ID,
			Year:       budget.Year,
			BudgetType: budget.BudgetType,
			Status:     status,
		})
	}

	return
}

func buildBudgetResponseItem(ctx context.Context, budget structs.Budget) (*dto.BudgetResponseItem, error) {
	status := buildBudgetStatus(ctx, budget)
	item := &dto.BudgetResponseItem{
		ID:         budget.ID,
		Year:       budget.Year,
		BudgetType: budget.BudgetType,
		Status:     status,
	}

	return item, nil
}

func (r *Resolver) BudgetInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Budget
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		return errors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		item, err := r.Repo.UpdateBudget(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := r.Repo.CreateBudget(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		accountLatestVersion, err := r.Repo.GetLatestVersionOfAccounts()
		if err != nil {
			return errors.HandleAPIError(err)
		}

		_, err = r.Repo.CreateFinancialBudget(&structs.FinancialBudget{
			AccountVersion: accountLatestVersion,
			BudgetID:       item.ID,
		})
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) BudgetSendResolver(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemID := params.Args["id"]

	BudgetItemType := &structs.Budget{}
	BudgetData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget.json", BudgetItemType)

	if err != nil {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
		return nil, err
	}
	budget := shared.FindByProperty(BudgetData, "ID", itemID)
	if len(budget) == 0 {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
		return nil, err
	}
	BudgetData = shared.FilterByProperty(BudgetData, "ID", itemID)
	newItem := structs.Budget{}

	for _, item := range budget {
		if updateBudget, ok := item.(*structs.Budget); ok {
			newItem.ID = updateBudget.ID
			newItem.Year = updateBudget.Year
			newItem.BudgetType = updateBudget.BudgetType
			// newItem.Status = "poslat"
		}
	}

	var updatedData = append(BudgetData, newItem)
	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/budget.json"), updatedData)

	ActivatesItemType := &structs.ProgramItem{}
	ActivatesData, err := shared.ReadJSON(shared.GetDataRoot()+"/program.json", ActivatesItemType)
	ActivatesData = shared.FilterByProperty(ActivatesData, "OrganizationUnitID", 0)

	if err != nil {
		fmt.Printf("Fetching Activates failed because of this error - %s.\n", err)
		return nil, err
	}

	RequestItemType := &structs.RequestBudgetType{}
	RequestsData, err := shared.ReadJSON(shared.GetDataRoot()+"/request_budget.json", RequestItemType)

	if err != nil {
		fmt.Printf("Fetching Requests failed because of this error - %s.\n", err)
		return nil, err
	}

	for _, activityData := range ActivatesData {
		if activity, ok := activityData.(*structs.ProgramItem); ok {
			currentTime := time.Now().UTC()
			timeString := currentTime.Format("2006-01-02 15:04:05")
			var newRequest = structs.RequestBudgetType{
				ID:                   shared.GetRandomNumber(),
				OrganizationUnitID:   activity.OrganizationUnitID,
				ActivityID:           activity.ID,
				BudgetID:             newItem.ID,
				DateCreate:           timeString,
				StatusNotFinancially: "U toku",
				StatusFinancially:    "U toku",
			}
			RequestsData = append(RequestsData, newRequest)

			BudgetActivityNotFinanciallyType := &structs.BudgetActivityNotFinanciallyItem{}
			BudgetActivityNotFinanciallyData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_activity_not_financially.json", BudgetActivityNotFinanciallyType)

			if err != nil {
				fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
			}

			var newNotFinancially = structs.BudgetActivityNotFinanciallyItem{
				ID:        shared.GetRandomNumber(),
				RequestID: newRequest.ID,
			}
			var NotFinanciallyData = append(BudgetActivityNotFinanciallyData, newNotFinancially)

			_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/budget_activity_not_financially.json"), NotFinanciallyData)

			_ = CreateProgramsForRequestResolver(activity.ID, newNotFinancially.ID)
		}
	}

	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/request_budget.json"), RequestsData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You send budget to OJ!",
	}, nil
}

var CreateProgramsForRequestResolver = func(activityID int, notFinanciallyID int) error {
	var projectRoot, _ = shared.GetProjectRoot()
	ProgramsItemType := &structs.ProgramItem{}
	ProgramsData, err := shared.ReadJSON(shared.GetDataRoot()+"/program.json", ProgramsItemType)

	if err != nil {
		fmt.Printf("Fetching Programs failed because of this error - %s.\n", err)
		return err
	}

	ActivityData := shared.FindByProperty(ProgramsData, "ID", activityID)

	NotFinanciallyProgramType := &structs.BudgetActivityNotFinanciallyProgramItem{}
	NotFinanciallyProgramData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_activity_not_financially.json", NotFinanciallyProgramType)

	if err != nil {
		fmt.Printf("Fetching not financially failed because of this error - %s.\n", err)
		return err
	}

	if len(ActivityData) > 0 {
		for _, activityItem := range ActivityData {
			if activity, ok := activityItem.(*structs.ProgramItem); ok {
				newActivity := structs.BudgetActivityNotFinanciallyProgramItem{
					ID:                     shared.GetRandomNumber(),
					BudgetNotFinanciallyID: notFinanciallyID,
					ProgramID:              activity.ID,
					Description:            "",
				}

				NotFinanciallyProgramData = append(NotFinanciallyProgramData, newActivity)

				SubprogramData := shared.FindByProperty(ProgramsData, "ID", activity.ParentID)
				if len(SubprogramData) > 0 {
					for _, subprogramItem := range SubprogramData {
						if subprogram, ok := subprogramItem.(*structs.ProgramItem); ok {
							newSubprogram := structs.BudgetActivityNotFinanciallyProgramItem{
								ID:                     shared.GetRandomNumber(),
								BudgetNotFinanciallyID: notFinanciallyID,
								ProgramID:              subprogram.ID,
								Description:            "",
							}

							NotFinanciallyProgramData = append(NotFinanciallyProgramData, newSubprogram)

							ProgramData := shared.FindByProperty(ProgramsData, "ID", subprogram.ParentID)

							if len(ProgramData) > 0 {
								for _, programItem := range ProgramData {
									if program, ok := programItem.(*structs.ProgramItem); ok {
										newProgram := structs.BudgetActivityNotFinanciallyProgramItem{
											ID:                     shared.GetRandomNumber(),
											BudgetNotFinanciallyID: notFinanciallyID,
											ProgramID:              program.ID,
											Description:            "",
										}

										NotFinanciallyProgramData = append(NotFinanciallyProgramData, newProgram)

										_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/budget_program_not_financially.json"), NotFinanciallyProgramData)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func (r *Resolver) BudgetDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteBudget(itemID)
	if err != nil {
		fmt.Printf("Deleting budget item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
