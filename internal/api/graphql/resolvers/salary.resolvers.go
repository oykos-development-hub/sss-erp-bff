package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) SalaryOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		Salary, err := r.Repo.GetSalaryByID(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		res, err := buildSalary(*Salary, r)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.SalaryResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.SalaryFilter{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["month"].(string); ok && value != "" {
		input.Month = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	} else {
		input.OrganizationUnitID, _ = params.Context.Value(config.OrganizationUnitIDKey).(*int)
	}

	if value, ok := params.Args["activity_id"].(int); ok && value != 0 {
		input.ActivityID = &value
	}

	if value, ok := params.Args["year"].(int); ok && value != 0 {
		input.Year = &value
	}

	items, total, err := r.Repo.GetSalaryList(input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var resItems []dto.SalaryResponse
	for _, item := range items {
		resItem, err := buildSalary(item, r)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		resItems = append(resItems, *resItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   resItems,
		Total:   total,
	}, nil
}

func (r *Resolver) SalaryInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Salary
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		return apierrors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if data.OrganizationUnitID == 0 {

		organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}

		data.OrganizationUnitID = *organizationUnitID

	}

	var item *structs.Salary

	if data.ID == 0 {
		item, err = r.Repo.CreateSalary(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdateSalary(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

	}

	singleItem, err := buildSalary(*item, r)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) SalaryDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteSalary(itemID)
	if err != nil {
		fmt.Printf("Deleting fixed deposit failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildSalary(item structs.Salary, r *Resolver) (*dto.SalaryResponse, error) {
	response := dto.SalaryResponse{
		ID:                item.ID,
		Month:             item.Month,
		DateOfCalculation: item.DateOfCalculation,
		Description:       item.Description,
		Status:            item.Status,
		GrossPrice:        item.GrossPrice,
		VatPrice:          item.VatPrice,
		NetPrice:          item.NetPrice,
		ObligationsPrice:  item.ObligationsPrice,
		NumberOfEmployees: item.NumberOfEmployees,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	if item.ActivityID != 0 {
		activity, err := r.Repo.GetActivity(item.ActivityID)
		if err != nil {
			return nil, err
		}

		response.Activity.ID = activity.ID
		response.Activity.Title = activity.Title
	}

	if item.OrganizationUnitID != 0 {
		orgUnit, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)
		if err != nil {
			return nil, err
		}

		response.OrganizationUnit.ID = orgUnit.ID
		response.OrganizationUnit.Title = orgUnit.Title
	}

	for _, additionalExpense := range item.SalaryAdditionalExpenses {
		data, err := buildSalaryAdditionalExpense(additionalExpense, r)

		if err != nil {
			return nil, err
		}

		response.SalaryAdditionalExpenses = append(response.SalaryAdditionalExpenses, *data)
	}

	return &response, nil
}

func buildSalaryAdditionalExpense(item structs.SalaryAdditionalExpense, r *Resolver) (*dto.SalaryAdditionalExpensesResponse, error) {
	response := dto.SalaryAdditionalExpensesResponse{
		ID:          item.ID,
		SalaryID:    item.SalaryID,
		Amount:      item.Amount,
		BankAccount: item.BankAccount,
		Status:      item.Status,
		Type:        item.Type,
		Title:       item.Title,
	}

	if item.OrganizationUnitID != 0 {
		orgUnit, err := r.Repo.GetOrganizationUnitByID(item.OrganizationUnitID)
		if err != nil {
			return nil, err
		}

		response.OrganizationUnit.ID = orgUnit.ID
		response.OrganizationUnit.Title = orgUnit.Title
	}

	if item.AccountID != 0 {
		account, err := r.Repo.GetAccountItemByID(item.AccountID)
		if err != nil {
			return nil, err
		}

		response.Account.ID = account.ID
		response.Account.Title = account.Title
	}

	if item.SubjectID != 0 {
		subject, err := r.Repo.GetSupplier(item.SubjectID)

		if err != nil {
			return nil, err
		}

		response.Subject.ID = subject.ID
		response.Subject.Title = subject.Title
	}

	if item.DebtorID != 0 {
		debtor, err := r.Repo.GetUserProfileByID(item.DebtorID)
		if err != nil {
			return nil, err
		}

		response.Debtor.ID = debtor.ID
		response.Debtor.Title = debtor.FirstName + " " + debtor.LastName
	}

	return &response, nil
}
