package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	goerrors "errors"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/shopspring/decimal"
)

func (r *Resolver) SpendingReleaseInsert(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)

	loggedInOrganizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok {
		return errors.HandleAPPError(errors.NewBadRequestError("Error getting logged in unit"))
	}

	if unitID == 0 {
		unitID = *loggedInOrganizationUnitID
	}

	if budgetID == 0 {
		currentYear := time.Now().Year()
		//TODO: after planning budget is done on FE, add status filter Done
		budget, err := r.Repo.GetBudgetList(&dto.GetBudgetListInputMS{
			Year: &currentYear,
		})
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting budget for current year"))
		}
		if len(budget) != 1 {
			return errors.HandleAPPError(errors.NewBadRequestError("Budget for current year not found"))
		}
		budgetID = budget[0].ID
	}

	var data []structs.SpendingReleaseInsert

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	items, err := r.Repo.CreateSpendingRelease(params.Context, data, budgetID, unitID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You created this item!",
		Items:   items,
	}, nil
}

func (r *Resolver) SpendingReleaseRequestInsert(params graphql.ResolveParams) (interface{}, error) {
	fileID := params.Args["file_id"].(int)

	var data structs.SpendingReleaseRequest
	data.OrganizationUnitFileID = fileID

	loggedInOrganizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok {
		return errors.HandleAPPError(errors.NewBadRequestError("Error getting logged in unit"))
	}

	data.OrganizationUnitID = *loggedInOrganizationUnitID

	currentYear := time.Now().Year()
	//TODO: after planning budget is done on FE, add status filter Done
	budget, err := r.Repo.GetBudgetList(&dto.GetBudgetListInputMS{
		Year: &currentYear,
	})
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting budget for current year"))
	}
	if len(budget) != 1 {
		return errors.HandleAPPError(errors.NewBadRequestError("Budget for current year not found"))
	}

	data.Year = currentYear
	data.Month = int(time.Now().Month())

	err = r.Repo.CreateSpendingReleaseRequest(params.Context, data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetUsers, err := r.Repo.GetUsersByPermission(config.FinanceBudget, config.OperationFullAccess)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, targetUser := range targetUsers {
		if targetUser.ID != loggedInUser.ID {
			_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
				Content:     "Podnijet je novi zahtjev za otpuštanje sredstava.",
				Module:      "Finansije",
				FromUserID:  loggedInUser.ID,
				ToUserID:    targetUser.ID,
				FromContent: "Službenik za budžet",
				Path:        "/finance/budget/current/fund-release/requests",
				IsRead:      false,
			})
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

		}
	}

	return dto.Response{
		Status:  "success",
		Message: "You created this item!",
	}, nil
}

func (r *Resolver) SpendingReleaseAcceptSSS(params graphql.ResolveParams) (interface{}, error) {
	fileID := params.Args["file_id"].(int)
	id := params.Args["id"].(int)

	err := r.Repo.SpendingReleaseAcceptSSS(id, fileID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	request, err := r.Repo.GetSpendingReleaseRequestByID(id)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetUsers, err := r.Repo.GetUsersByPermission(config.FinanceBudget, config.OperationRead)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	employees, err := GetEmployeesOfOrganizationUnit(r.Repo, request.OrganizationUnitID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, targetUser := range targetUsers {
		for _, employee := range employees {
			if targetUser.ID != loggedInUser.ID && employee.UserAccountID == targetUser.ID {
				_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
					Content:     "Vaš zahtjev za otpuštanje je prihvaćen.",
					Module:      "Finansije",
					FromUserID:  loggedInUser.ID,
					ToUserID:    targetUser.ID,
					FromContent: "Službenik za budžet",
					Path:        "/finance/budget/current/fund-release",
					IsRead:      false,
				})
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "You accepted this item!",
	}, nil
}

func (r *Resolver) SpendingReleaseOverview(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)
	month := params.Args["month"].(int)
	year := params.Args["year"].(int)
	status, statusOK := params.Args["status"].(string)

	requestFilter := dto.SpendingReleaseOverviewRequestFilter{}

	if unitID != 0 {
		requestFilter.OrganizationUnitID = &unitID
	}

	if month != 0 {
		requestFilter.Month = &month
	}

	if year != 0 {
		requestFilter.Year = &year
	}

	if statusOK && status != "" {
		requestFilter.Status = &status
	}

	spendingReleaseRequests, err := r.Repo.GetSpendingReleaseRequests(requestFilter)

	if err != nil {
		var apiErr *errors.APIError
		if goerrors.As(err, &apiErr) {
			if apiErr.StatusCode != 404 {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending release requests"))
			}
		}
	}

	if budgetID == 0 {
		currentYear := time.Now().Year()
		//TODO: after planning budget is done on FE, add status filter Done
		budget, err := r.Repo.GetBudgetList(&dto.GetBudgetListInputMS{
			Year: &currentYear,
		})
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting budget for current year"))
		}
		if len(budget) != 1 {
			return errors.HandleAPPError(errors.NewBadRequestError("Budget for current year not found"))
		}
		budgetID = budget[0].ID
	}

	var spendingReleaseOverview []dto.SpendingReleaseOverviewItem

	for i := 0; i < len(spendingReleaseRequests); i++ {
		input := &dto.GetSpendingReleaseListInput{
			Year:     spendingReleaseRequests[i].Year,
			BudgetID: budgetID,
			UnitID:   spendingReleaseRequests[i].OrganizationUnitID,
			Month:    spendingReleaseRequests[i].Month,
		}

		spendingReleaseOverviewItems, err := r.Repo.GetSpendingReleaseList(params.Context, input)
		if err != nil {
			var apiErr *errors.APIError
			if goerrors.As(err, &apiErr) {
				if apiErr.StatusCode != 404 {
					return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic"))
				}
			}
		}

		var SSSFile dto.FileDropdownSimple
		var OUFile dto.FileDropdownSimple
		var OU dto.DropdownSimple
		if spendingReleaseRequests[i].SSSFileID != 0 {
			file, err := r.Repo.GetFileByID(spendingReleaseRequests[i].SSSFileID)

			if err != nil {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting file by id"))
			}

			SSSFile = dto.FileDropdownSimple{
				ID:   file.ID,
				Name: file.Name,
				Type: *file.Type,
			}
		}

		if spendingReleaseRequests[i].OrganizationUnitFileID != 0 {
			file, err := r.Repo.GetFileByID(spendingReleaseRequests[i].OrganizationUnitFileID)

			if err != nil {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting file by id"))
			}

			OUFile = dto.FileDropdownSimple{
				ID:   file.ID,
				Name: file.Name,
				Type: *file.Type,
			}
		}

		if spendingReleaseRequests[i].OrganizationUnitID != 0 {
			value, err := r.Repo.GetOrganizationUnitByID(spendingReleaseRequests[i].OrganizationUnitID)

			if err != nil {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting file by id"))
			}

			OU = dto.DropdownSimple{
				ID:    value.ID,
				Title: value.Title,
			}
		}

		value := decimal.NewFromInt(0)

		for _, item := range spendingReleaseOverviewItems {
			value = value.Add(item.Value)
		}

		spendingReleaseOverview = append(spendingReleaseOverview, dto.SpendingReleaseOverviewItem{
			ID:                   spendingReleaseRequests[i].ID,
			Month:                spendingReleaseRequests[i].Month,
			Year:                 spendingReleaseRequests[i].Year,
			Value:                value,
			Status:               spendingReleaseRequests[i].Status,
			OrganizationUnitFile: OUFile,
			SSSFile:              SSSFile,
			OrganizationUnit:     OU,
			CreatedAt:            spendingReleaseRequests[i].CreatedAt,
		})

	}

	hide := params.Args["hide"].(bool)

	if hide && statusOK && status != "" {
		var responseItems []dto.SpendingReleaseOverviewItem
		for _, item := range spendingReleaseOverview {
			if item.Status == status {
				responseItems = append(responseItems, item)
			}
		}
		spendingReleaseOverview = responseItems
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the data you asked for!",
		Items:   spendingReleaseOverview,
	}, nil
}

func (r *Resolver) SpendingReleaseDelete(params graphql.ResolveParams) (interface{}, error) {
	unitID := params.Args["unit_id"].(int)

	input := &dto.DeleteSpendingReleaseInput{
		Month: int(time.Now().Month()),
	}

	loggedInOrganizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok {
		return errors.HandleAPPError(errors.NewBadRequestError("get logged in unit from context"))
	}

	if unitID == 0 {
		unitID = *loggedInOrganizationUnitID
	}

	input.UnitID = unitID

	currentYear := time.Now().Year()

	//TODO: after planning budget is done on FE, add status filter Done
	budget, err := r.Repo.GetBudgetList(&dto.GetBudgetListInputMS{
		Year: &currentYear,
	})
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "repo get budget list"))
	}

	if len(budget) != 1 {
		return errors.HandleAPPError(errors.NewBadRequestError("budget for current year not found"))
	}

	input.BudgetID = budget[0].ID
	input.Year = currentYear

	err = r.Repo.DeleteSpendingRelease(params.Context, input)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "repo delete spending release"))
	}

	return dto.Response{
		Status:  "success",
		Message: "Deleted!",
	}, nil
}

func (r *Resolver) SpendingReleaseGet(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)
	month := params.Args["month"].(int)
	year := params.Args["year"].(int)

	input := &dto.GetSpendingReleaseListInput{
		Month: month,
		Year:  year,
	}

	loggedInOrganizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok {
		return errors.HandleAPPError(errors.NewBadRequestError("Error getting logged in unit"))
	}

	if unitID == 0 {
		unitID = *loggedInOrganizationUnitID
	}

	if budgetID == 0 {
		currentYear := time.Now().Year()
		//TODO: after planning budget is done on FE, add status filter Done
		budget, err := r.Repo.GetBudgetList(&dto.GetBudgetListInputMS{
			Year: &currentYear,
		})
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting budget for current year"))
		}
		if len(budget) != 1 {
			return errors.HandleAPPError(errors.NewBadRequestError("Budget for current year not found"))
		}
		budgetID = budget[0].ID
	}

	input.UnitID = unitID
	input.BudgetID = budgetID

	spendingReleases, err := r.Repo.GetSpendingReleaseList(params.Context, input)
	if err != nil {
		var apiErr *errors.APIError
		if goerrors.As(err, &apiErr) {
			if apiErr.StatusCode != 404 {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic"))
			}
		}
	}

	spendingDynamic, err := r.Repo.GetSpendingDynamic(budgetID, unitID, nil)
	if err != nil {
		var apiErr *errors.APIError
		if goerrors.As(err, &apiErr) {
			if apiErr.StatusCode != 404 {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic"))
			}
		}
	}

	financeBudgetDetails, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "get financial budget by budget id"))
	}

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{
		Version: &financeBudgetDetails.AccountVersion,
	})
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "get account items"))
	}

	tree := r.buildSpendingReleaseTree(month, accounts.Data, spendingReleases, spendingDynamic)

	return dto.Response{
		Status:  "success",
		Message: "Here's the data you asked for!",
		Items:   tree,
	}, nil
}

func populateSpendingReleaseData(spendingData []structs.SpendingRelease) map[int]*dto.SpendingReleaseDTO {
	spendingMap := make(map[int]*dto.SpendingReleaseDTO)
	for _, data := range spendingData {
		data := data // capture range variable
		spendingMap[data.AccountID] = &dto.SpendingReleaseDTO{
			ID:              data.ID,
			AccountID:       data.AccountID,
			BudgetID:        data.BudgetID,
			UnitID:          data.UnitID,
			CurrentBudgetID: data.CurrentBudgetID,
			Value:           data.Value,
			CreatedAt:       data.CreatedAt,
			Username:        data.Username,
		}
	}
	return spendingMap
}

func (r *Resolver) buildSpendingReleaseTree(month int, accounts []*structs.AccountItem, spendingData []structs.SpendingRelease, dynamicData []dto.SpendingDynamicDTO) []*dto.SpendingReleaseDTO {
	accountTree := r.buildAccountTree(accounts)
	spendingMap := populateSpendingReleaseData(spendingData)
	spendingDynamicMap := r.populateSpendingData(dynamicData)

	var roots []*dto.SpendingReleaseDTO

	for _, account := range accountTree[0] {
		root := &dto.SpendingReleaseDTO{
			AccountID:           account.ID,
			AccountSerialNumber: account.SerialNumber,
			AccountTitle:        account.Title,
		}

		buildSpendingReleaseTreeRecursively(month, account.ID, root, accountTree, spendingMap, spendingDynamicMap)
		calculateReleaseSums(root)
		roots = append(roots, root)
	}

	return roots
}

func buildSpendingReleaseTreeRecursively(month, accountID int, parent *dto.SpendingReleaseDTO, accountTree map[int][]*structs.AccountItem, spendingMap map[int]*dto.SpendingReleaseDTO, dynamicMap map[int]*dto.SpendingDynamicDTO) {
	for _, childAccount := range accountTree[accountID] {
		child := &dto.SpendingReleaseDTO{
			AccountID:           childAccount.ID,
			AccountSerialNumber: childAccount.SerialNumber,
			AccountTitle:        childAccount.Title,
		}

		if data, exists := spendingMap[childAccount.ID]; exists {
			child.Value = data.Value
			child.CurrentBudgetID = data.CurrentBudgetID
			child.UnitID = data.UnitID
			child.BudgetID = data.BudgetID
			child.CreatedAt = data.CreatedAt
			child.Username = data.Username
		}
		if data, exists := dynamicMap[childAccount.ID]; exists {
			switch month {
			case int(time.January):
				child.Planned = data.January.Value
			case int(time.February):
				child.Planned = data.February.Value
			case int(time.March):
				child.Planned = data.March.Value
			case int(time.April):
				child.Planned = data.April.Value
			case int(time.May):
				child.Planned = data.May.Value
			case int(time.June):
				child.Planned = data.June.Value
			case int(time.July):
				child.Planned = data.July.Value
			case int(time.August):
				child.Planned = data.August.Value
			case int(time.September):
				child.Planned = data.September.Value
			case int(time.October):
				child.Planned = data.October.Value
			case int(time.November):
				child.Planned = data.November.Value
			case int(time.December):
				child.Planned = data.December.Value
			}
		}

		buildSpendingReleaseTreeRecursively(month, childAccount.ID, child, accountTree, spendingMap, dynamicMap)
		parent.Children = append(parent.Children, child)
	}
}

func calculateReleaseSums(node *dto.SpendingReleaseDTO) {
	for _, child := range node.Children {
		calculateReleaseSums(child)
		node.Value = node.Value.Add(child.Value)
		node.Planned = node.Planned.Add(child.Planned)

	}
}
