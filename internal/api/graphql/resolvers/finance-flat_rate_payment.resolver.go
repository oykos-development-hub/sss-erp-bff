package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

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
		return errors.HandleAPPError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	var item *structs.FlatRatePayment

	if data.ID == 0 {
		item, err = r.Repo.CreateFlatRatePayment(params.Context, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdateFlatRatePayment(params.Context, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}
	}

	flatrateResItem, err := buildFlatRatePaymentResponseItem(*item)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	response.Item = flatrateResItem

	return response, nil
}

func (r *Resolver) FlatRatePaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		flatratePayment, err := r.Repo.GetFlatRatePayment(id)
		if err != nil {
			return errors.HandleAPPError(err)
		}
		flatratePaymentResItem, err := buildFlatRatePaymentResponseItem(*flatratePayment)
		if err != nil {
			return errors.HandleAPPError(err)
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
		return errors.HandleAPPError(err)
	}

	flatrateResItem, err := buildFlatRatePaymentResponseItemList(flatratePayments)
	if err != nil {
		return errors.HandleAPPError(err)
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

	err := r.Repo.DeleteFlatRatePayment(params.Context, itemID)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildFlatRatePaymentResponseItem(flatratePayment structs.FlatRatePayment) (*dto.FlatRatePaymentResponseItem, error) {

	status := dto.DropdownSimple{
		ID:    int(structs.PaidFlatRatePeymentStatus),
		Title: string(dto.FinancialFlatRatePaymentStatusPaid),
	}

	switch flatratePayment.Status {
	case structs.PaidFlatRatePeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PaidFlatRatePeymentStatus),
			Title: string(dto.FinancialFlatRatePaymentStatusPaid),
		}
	case structs.CancelledFlatRatePeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.CancelledFlatRatePeymentStatus),
			Title: string(dto.FinancialFlatRatePaymentStatusCanceled),
		}
	case structs.RetunedFlatRatePeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.RetunedFlatRatePeymentStatus),
			Title: string(dto.FinancialFlatRatePaymentStatusReturned),
		}
	}

	flatRatePaymentMethod := dto.DropdownSimple{
		ID:    int(structs.PaymentFlatRatePeymentMethod),
		Title: string(dto.FinancialFlatRatePaymentMethodPayment),
	}

	switch flatratePayment.PaymentMethod {
	case structs.PaymentFlatRatePeymentMethod:
		flatRatePaymentMethod = dto.DropdownSimple{
			ID:    int(structs.PaymentFlatRatePeymentMethod),
			Title: string(dto.FinancialFlatRatePaymentMethodPayment),
		}
	case structs.ForcedFlatRatePeymentMethod:
		flatRatePaymentMethod = dto.DropdownSimple{
			ID:    int(structs.ForcedFlatRatePeymentMethod),
			Title: string(dto.FinancialFlatRatePaymentMethodForced),
		}
	case structs.CourtCostsFlatRatePeymentMethod:
		flatRatePaymentMethod = dto.DropdownSimple{
			ID:    int(structs.CourtCostsFlatRatePeymentMethod),
			Title: string(dto.FinancialFlatRatePaymentMethodCourtCosts),
		}
	}

	response := dto.FlatRatePaymentResponseItem{
		ID:                     flatratePayment.ID,
		FlatRateID:             flatratePayment.FlatRateID,
		PaymentMethod:          flatRatePaymentMethod,
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
			return nil, errors.Wrap(err, "build flat rate payment response item")
		}

		items = append(items, singleItem)

	}

	return items, nil
}
