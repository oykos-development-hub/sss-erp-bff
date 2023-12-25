package resolvers

import (
	"bff/internal/api/dto"
	"bff/shared"
	"bff/structs"
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
	var items []interface{}
	var total int
	var id int
	var status string
	var year string
	var typeBudget string

	if params.Args["id"] != nil && params.Args["id"].(int) > 0 {
		id = params.Args["id"].(int)
	}

	if params.Args["status"] == nil {
		status = ""
	} else {
		status = params.Args["status"].(string)
	}

	if params.Args["year"] == nil {
		year = ""
	} else {
		year = params.Args["year"].(string)
	}

	if params.Args["type_budget"] == nil {
		typeBudget = ""
	} else {
		typeBudget = params.Args["type_budget"].(string)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	BudgetType := &structs.BudgetItem{}
	BudgetData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget.json", BudgetType)

	if err != nil {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
	}

	// Populate data for each Basic Inventory Real Estates
	items = BudgetItemProperties(BudgetData, id, typeBudget, year, status)

	total = len(items)

	// Filtering by Pagination params
	if page != nil && page != 0 && size != nil && size != 0 {
		items = shared.Pagination(items, page.(int), size.(int))
	}

	return dto.Response{
		Status:  "success",
		Message: "You fetched items!",
		Items:   items,
		Total:   total,
	}, nil
}

func (r *Resolver) BudgetInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetItemType := &structs.BudgetItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	BudgetData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget.json", BudgetItemType)

	if err != nil {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
	}

	if itemID != 0 {
		BudgetData = shared.FilterByProperty(BudgetData, "ID", itemID)
	} else {
		data.ID = shared.GetRandomNumber()
		data.Status = "kreiran"
	}

	var updatedData = append(BudgetData, data)

	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/budget.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetItemProperties(sliceData, itemID, "", "", "")

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedData[0],
	}, nil
}

func (r *Resolver) BudgetSendResolver(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemID := params.Args["id"]

	BudgetItemType := &structs.BudgetItem{}
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
	newItem := structs.BudgetItem{}

	for _, item := range budget {
		if updateBudget, ok := item.(*structs.BudgetItem); ok {
			newItem.ID = updateBudget.ID
			newItem.Year = updateBudget.Year
			newItem.Source = updateBudget.Source
			newItem.Type = updateBudget.Type
			newItem.Status = "poslat"
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
	var projectRoot, _ = shared.GetProjectRoot()
	itemID := params.Args["id"]
	BudgetItemType := &structs.BudgetItem{}
	BudgetData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget.json", BudgetItemType)

	if err != nil {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
	}

	if itemID != 0 {
		BudgetData = shared.FilterByProperty(BudgetData, "ID", itemID)
	}

	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/budget.json"), BudgetData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
