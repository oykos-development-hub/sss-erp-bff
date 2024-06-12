package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) FeePaymentInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FeePayment
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

	var item *structs.FeePayment

	if data.ID == 0 {
		item, err = r.Repo.CreateFeePayment(params.Context, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdateFeePayment(params.Context, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

	}

	feePaymentResItem, err := buildFeePaymentResponseItem(*item)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	response.Item = feePaymentResItem

	return response, nil
}

func (r *Resolver) FeePaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		feePayment, err := r.Repo.GetFeePayment(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		feePaymentResItem, err := buildFeePaymentResponseItem(*feePayment)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.FeePaymentResponseItem{feePaymentResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetFeePaymentListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["fee_id"].(int); ok && value != 0 {
		input.FeeID = &value
	}

	feePayments, total, err := r.Repo.GetFeePaymentList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	feeResItem, err := buildFeePaymentResponseItemList(feePayments)
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

func (r *Resolver) FeePaymentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFeePayment(params.Context, itemID)
	if err != nil {
		fmt.Printf("Deleting fee payment item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildFeePaymentResponseItem(feePayment structs.FeePayment) (*dto.FeePaymentResponseItem, error) {
	status := dto.DropdownSimple{
		ID:    int(structs.PaidFeePeymentStatus),
		Title: string(dto.FinancialFeePaymentStatusPaid),
	}

	switch feePayment.Status {
	case structs.PaidFeePeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PaidFeePeymentStatus),
			Title: string(dto.FinancialFeePaymentStatusPaid),
		}
	case structs.CancelledFeePeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.CancelledFeePeymentStatus),
			Title: string(dto.FinancialFeePaymentStatusCanceled),
		}
	case structs.RetunedFeePeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.RetunedFeePeymentStatus),
			Title: string(dto.FinancialFeePaymentStatusReturned),
		}
	}

	feePaymentMethod := dto.DropdownSimple{
		ID:    int(structs.PaymentFeePeymentMethod),
		Title: string(dto.FinancialFeePaymentMethodPayment),
	}

	switch feePayment.PaymentMethod {
	case structs.PaymentFeePeymentMethod:
		feePaymentMethod = dto.DropdownSimple{
			ID:    int(structs.PaymentFeePeymentMethod),
			Title: string(dto.FinancialFeePaymentMethodPayment),
		}
	case structs.ForcedFeePeymentMethod:
		feePaymentMethod = dto.DropdownSimple{
			ID:    int(structs.ForcedFeePeymentMethod),
			Title: string(dto.FinancialFeePaymentMethodForced),
		}
	}

	response := dto.FeePaymentResponseItem{
		ID:                     feePayment.ID,
		FeeID:                  feePayment.FeeID,
		PaymentMethod:          feePaymentMethod,
		Amount:                 feePayment.Amount,
		PaymentDate:            feePayment.PaymentDate,
		PaymentDueDate:         feePayment.PaymentDueDate,
		ReceiptNumber:          feePayment.ReceiptNumber,
		PaymentReferenceNumber: feePayment.PaymentReferenceNumber,
		DebitReferenceNumber:   feePayment.DebitReferenceNumber,
		Status:                 status,
		CreatedAt:              feePayment.CreatedAt,
		UpdatedAt:              feePayment.UpdatedAt,
	}

	return &response, nil
}

func buildFeePaymentResponseItemList(itemList []structs.FeePayment) ([]*dto.FeePaymentResponseItem, error) {
	var items []*dto.FeePaymentResponseItem

	for _, item := range itemList {
		singleItem, err := buildFeePaymentResponseItem(item)

		if err != nil {
			return nil, err
		}

		items = append(items, singleItem)

	}

	return items, nil
}
