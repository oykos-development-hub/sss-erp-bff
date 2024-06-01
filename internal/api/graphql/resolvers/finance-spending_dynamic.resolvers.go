package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	goerrors "errors"
	"sort"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) SpendingDynamicInsert(params graphql.ResolveParams) (interface{}, error) {
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

	var data []structs.SpendingDynamicInsert

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	items, err := r.Repo.CreateSpendingDynamic(params.Context, budgetID, unitID, data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You created this item!",
		Items:   items,
	}, nil
}

func (r *Resolver) SpendingDynamicOverview(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)
	version := params.Args["version"].(int)
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

	input := &dto.GetSpendingDynamicHistoryInput{}
	if version != 0 {
		input.Version = &version
	}

	spendingDynamic, err := r.Repo.GetSpendingDynamic(budgetID, unitID, input)
	if err != nil {
		var apiErr *errors.APIError
		if goerrors.As(err, &apiErr) {
			if apiErr.StatusCode != 404 {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic"))
			}
		}
	}

	accounts, err := r.Repo.GetAccountItems(nil)
	if err != nil {
		return nil, err
	}

	tree := buildSpendingDynamicTree(accounts.Data, spendingDynamic)

	return dto.Response{
		Status:  "success",
		Message: "Here's the data you asked for!",
		Items:   tree,
	}, nil
}

func (r *Resolver) SpendingDynamicHistoryOverview(params graphql.ResolveParams) (interface{}, error) {
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

	spendingDynamicHistory, err := r.Repo.GetSpendingDynamicHistory(budgetID, unitID)
	if err != nil {
		var apiErr *errors.APIError
		if goerrors.As(err, &apiErr) {
			if apiErr.StatusCode != 404 {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting spending dynamic"))
			}
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the data you asked for!",
		Items:   spendingDynamicHistory,
	}, nil
}

func buildAccountMap(accounts []*structs.AccountItem) map[int]*structs.AccountItem {
	accountMap := make(map[int]*structs.AccountItem)
	for _, account := range accounts {
		account := account // capture range variable
		accountMap[account.ID] = account
	}
	return accountMap
}

func buildAccountTree(accounts []*structs.AccountItem) map[int][]*structs.AccountItem {
	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].SerialNumber < accounts[j].SerialNumber
	})

	accountTree := make(map[int][]*structs.AccountItem)
	for _, account := range accounts {
		if account.ParentID != nil {
			parentID := *account.ParentID
			accountTree[parentID] = append(accountTree[parentID], account)
		} else {
			accountTree[0] = append(accountTree[0], account)
		}
	}
	return accountTree
}

func populateSpendingData(accounts map[int]*structs.AccountItem, spendingData []dto.SpendingDynamicDTO) map[int]*dto.SpendingDynamicDTO {
	spendingMap := make(map[int]*dto.SpendingDynamicDTO)
	for _, data := range spendingData {
		data := data // capture range variable
		spendingMap[data.AccountID] = &data
	}
	return spendingMap
}

func buildSpendingDynamicTree(accounts []*structs.AccountItem, spendingData []dto.SpendingDynamicDTO) []*dto.SpendingDynamicDTO {
	accountMap := buildAccountMap(accounts)
	accountTree := buildAccountTree(accounts)
	spendingMap := populateSpendingData(accountMap, spendingData)

	var roots []*dto.SpendingDynamicDTO

	for _, account := range accountTree[0] {
		root := &dto.SpendingDynamicDTO{
			AccountID:           account.ID,
			AccountSerialNumber: account.SerialNumber,
			AccountTitle:        account.Title,
		}

		if data, exists := spendingMap[account.ID]; exists {
			root = data
		}

		buildTreeRecursively(account.ID, root, accountTree, spendingMap)
		calculateSums(root)
		roots = append(roots, root)
	}

	return roots
}

func buildTreeRecursively(accountID int, parent *dto.SpendingDynamicDTO, accountTree map[int][]*structs.AccountItem, spendingMap map[int]*dto.SpendingDynamicDTO) {
	for _, childAccount := range accountTree[accountID] {
		child := &dto.SpendingDynamicDTO{
			AccountID:           childAccount.ID,
			AccountSerialNumber: childAccount.SerialNumber,
			AccountTitle:        childAccount.Title,
		}

		if data, exists := spendingMap[childAccount.ID]; exists {
			child.January = data.January
			child.February = data.January
			child.March = data.January
			child.April = data.January
			child.May = data.January
			child.June = data.January
			child.July = data.January
			child.August = data.January
			child.September = data.January
			child.October = data.January
			child.November = data.January
			child.December = data.January
		}

		buildTreeRecursively(childAccount.ID, child, accountTree, spendingMap)
		parent.Children = append(parent.Children, child)
	}
}

func calculateSums(node *dto.SpendingDynamicDTO) {
	for _, child := range node.Children {
		calculateSums(child)
		node.January.Value = node.January.Value.Add(child.January.Value)
		node.January.Savings = node.January.Savings.Add(child.January.Savings)
		node.February.Value = node.February.Value.Add(child.February.Value)
		node.February.Savings = node.February.Savings.Add(child.February.Savings)
		node.March.Value = node.March.Value.Add(child.March.Value)
		node.March.Savings = node.March.Savings.Add(child.March.Savings)
		node.April.Value = node.April.Value.Add(child.April.Value)
		node.April.Savings = node.April.Savings.Add(child.April.Savings)
		node.May.Value = node.May.Value.Add(child.May.Value)
		node.May.Savings = node.May.Savings.Add(child.May.Savings)
		node.June.Value = node.June.Value.Add(child.June.Value)
		node.June.Savings = node.June.Savings.Add(child.June.Savings)
		node.July.Value = node.July.Value.Add(child.July.Value)
		node.July.Savings = node.July.Savings.Add(child.July.Savings)
		node.August.Value = node.August.Value.Add(child.August.Value)
		node.August.Savings = node.August.Savings.Add(child.August.Savings)
		node.September.Value = node.September.Value.Add(child.September.Value)
		node.September.Savings = node.September.Savings.Add(child.September.Savings)
		node.October.Value = node.October.Value.Add(child.October.Value)
		node.October.Savings = node.October.Savings.Add(child.October.Savings)
		node.November.Value = node.November.Value.Add(child.November.Value)
		node.November.Savings = node.November.Savings.Add(child.November.Savings)
		node.December.Value = node.December.Value.Add(child.December.Value)
		node.December.Savings = node.December.Savings.Add(child.December.Savings)
	}
	node.TotalSavings = node.January.Savings.Add(node.February.Savings).Add(node.March.Savings).Add(node.April.Savings).
		Add(node.May.Savings).Add(node.June.Savings).Add(node.July.Savings).Add(node.August.Savings).
		Add(node.September.Savings).Add(node.October.Savings).Add(node.November.Savings).Add(node.December.Savings)
}
