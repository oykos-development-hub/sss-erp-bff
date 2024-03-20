package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) FlatRatePaymentInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FlatRatePayment
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

	var item *structs.FlatRatePayment

	if data.ID == 0 {
		item, err = r.Repo.CreateFlatRatePayment(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdateFlatRatePayment(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

	}

	response.Item = *item

	return response, nil
}

func (r *Resolver) FlatRatePaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		flatratePayment, err := r.Repo.GetFlatRatePayment(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		flatratePaymentResItem, err := buildFlatRatePaymentResponseItem(*flatratePayment)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.FlatRatePaymentResponseItem{flatratePaymentResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetFlatRatePaymentListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["flat_rate_id"].(int); ok && value != 0 {
		input.FlatRateID = &value
	}

	flatratePayments, total, err := r.Repo.GetFlatRatePaymentList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	flatrateResItem, err := buildFlatRatePaymentResponseItemList(flatratePayments)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   flatrateResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) FlatRatePaymentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFlatRatePayment(itemID)
	if err != nil {
		fmt.Printf("Deleting flatrate payment item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildFlatRatePaymentResponseItem(flatratePayment structs.FlatRatePayment) (*dto.FlatRatePaymentResponseItem, error) {
	status := dto.FinancialFlatRatePaymentStatusPaid
	switch flatratePayment.Status {
	case structs.PaidFlatRatePeymentStatus:
		status = dto.FinancialFlatRatePaymentStatusPaid
	case structs.CancelledFlatRatePeymentStatus:
		status = dto.FinancialFlatRatePaymentStatusCanceled
	case structs.RetunedFlatRatePeymentStatus:
		status = dto.FinancialFlatRatePaymentStatusReturned
	}

	flatratePaymentMethod := dto.FinancialFlatRatePaymentMethodPayment
	switch flatratePayment.PaymentMethod {
	case structs.PaymentFlatRatePeymentMethod:
		flatratePaymentMethod = dto.FinancialFlatRatePaymentMethodPayment
	case structs.ForcedFlatRatePeymentMethod:
		flatratePaymentMethod = dto.FinancialFlatRatePaymentMethodForced
	case structs.CourtCostsFlatRatePeymentMethod:
		flatratePaymentMethod = dto.FinancialFlatRatePaymentMethodCourtCosts
	}
	response := dto.FlatRatePaymentResponseItem{
		ID:                     flatratePayment.ID,
		FlatRateID:             flatratePayment.FlatRateID,
		PaymentMethod:          flatratePaymentMethod,
		Amount:                 flatratePayment.Amount,
		PaymentDate:            flatratePayment.PaymentDate,
		PaymentDueDate:         flatratePayment.PaymentDueDate,
		ReceiptNumber:          flatratePayment.ReceiptNumber,
		PaymentReferenceNumber: flatratePayment.PaymentReferenceNumber,
		DebitReferenceNumber:   flatratePayment.DebitReferenceNumber,
		Status:                 status,
		CreatedAt:              flatratePayment.CreatedAt,
		UpdatedAt:              flatratePayment.UpdatedAt,
	}

	return &response, nil
}

func buildFlatRatePaymentResponseItemList(itemList []structs.FlatRatePayment) ([]*dto.FlatRatePaymentResponseItem, error) {
	var items []*dto.FlatRatePaymentResponseItem

	for _, item := range itemList {
		singleItem, err := buildFlatRatePaymentResponseItem(item)

		if err != nil {
			return nil, err
		}

		items = append(items, singleItem)

	}

	return items, nil
}
