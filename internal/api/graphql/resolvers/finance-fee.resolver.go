package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) FeeInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Fee
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		return errors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	var item *structs.Fee

	if data.ID == 0 {
		item, err = r.Repo.CreateFee(params.Context, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdateFee(params.Context, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

	}

	singleItem, err := buildFeeResponseItem(*item, r)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	response.Item = *singleItem

	return response, nil
}

func (r *Resolver) FeeOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		fee, err := r.Repo.GetFee(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		feeResItem, err := buildFeeResponseItem(*fee, r)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.FeeResponseItem{feeResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetFeeListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["fee_subcategory_id"].(int); ok && value != 0 {
		input.FilterBySubcategoryID = &value
	}

	if value, ok := params.Args["fee_type_id"].(int); ok && value != 0 {
		input.FilterByFeeTypeID = &value
	}

	fees, total, err := r.Repo.GetFeeList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	feeResItem, err := buildFeeResponseItemList(fees, r)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   feeResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) FeeDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFee(params.Context, itemID)
	if err != nil {
		fmt.Printf("Deleting fee item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildFeeResponseItem(fee structs.Fee, r *Resolver) (*dto.FeeResponseItem, error) {
	status := dto.DropdownSimple{
		ID:    int(structs.UnpaidFeeStatus),
		Title: string(dto.FinancialFeeUnpaidFeeStatus),
	}
	switch fee.Status {
	case structs.PaidFeeStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.UnpaidFeeStatus),
			Title: string(dto.FinancialFeePaidFeeStatus),
		}
	case structs.UnpaidFeeStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.UnpaidFeeStatus),
			Title: string(dto.FinancialFeeUnpaidFeeStatus),
		}
	case structs.PartFeeStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.UnpaidFeeStatus),
			Title: string(dto.FinancialFeePartFeeStatus),
		}
	}

	feeType := dto.DropdownSimple{
		ID:    int(structs.LawsuitFeeType),
		Title: string(dto.LawsuitFeeType),
	}
	switch fee.FeeType {
	case structs.LawsuitFeeType:
		feeType = dto.DropdownSimple{
			ID:    int(structs.LawsuitFeeType),
			Title: string(dto.LawsuitFeeType),
		}
	case structs.JudgmentFeeType:
		feeType = dto.DropdownSimple{
			ID:    int(structs.JudgmentFeeType),
			Title: string(dto.JudgmentFeeType),
		}
	}

	feeSubcategory := dto.DropdownSimple{
		ID:    int(structs.CopyingFeeSubcategory),
		Title: string(dto.CopyingFeeSubcategory),
	}
	switch fee.FeeSubcategory {
	case structs.CopyingFeeSubcategory:
		feeSubcategory = dto.DropdownSimple{
			ID:    int(structs.CopyingFeeSubcategory),
			Title: string(dto.CopyingFeeSubcategory),
		}
	}

	response := dto.FeeResponseItem{
		ID:                     fee.ID,
		FeeType:                feeType,
		FeeSubcategory:         feeSubcategory,
		DecisionNumber:         fee.DecisionNumber,
		DecisionDate:           fee.DecisionDate,
		Subject:                fee.Subject,
		JMBG:                   fee.JMBG,
		Amount:                 fee.Amount,
		PaymentReferenceNumber: fee.PaymentReferenceNumber,
		DebitReferenceNumber:   fee.DebitReferenceNumber,
		ExecutionDate:          fee.ExecutionDate,
		PaymentDeadlineDate:    fee.PaymentDeadlineDate,
		Description:            fee.Description,
		Status:                 status,
		FeeDetails:             fee.FeeDetails,
		CreatedAt:              fee.CreatedAt,
		UpdatedAt:              fee.UpdatedAt,
	}

	if len(fee.File) > 0 {
		for _, fileID := range fee.File {
			file, err := r.Repo.GetFileByID(fileID)

			if err != nil {
				return nil, err
			}

			FileDropdown := dto.FileDropdownSimple{
				ID:   file.ID,
				Name: file.Name,
				Type: *file.Type,
			}
			response.File = append(response.File, FileDropdown)
		}
	}

	if fee.CourtAccountID != nil {
		courtAccount, err := r.Repo.GetAccountItemByID(*fee.CourtAccountID)

		if err != nil {
			return nil, err
		}

		courtAccountDropdown := &dto.DropdownSimple{
			ID:    courtAccount.ID,
			Title: courtAccount.Title,
		}
		response.CourtAccount = courtAccountDropdown
	}

	return &response, nil
}

func buildFeeResponseItemList(itemList []structs.Fee, r *Resolver) ([]*dto.FeeResponseItem, error) {
	var items []*dto.FeeResponseItem

	for _, item := range itemList {
		singleItem, err := buildFeeResponseItem(item, r)

		if err != nil {
			return nil, err
		}

		items = append(items, singleItem)

	}

	return items, nil
}
