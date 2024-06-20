package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ProcedureCostPaymentInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.ProcedureCostPayment
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

	var item *structs.ProcedureCostPayment

	if data.ID == 0 {
		item, err = r.Repo.CreateProcedureCostPayment(params.Context, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdateProcedureCostPayment(params.Context, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}
	}
	procedurecostPaymentResItem, err := buildProcedureCostPaymentResponseItem(*item)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	response.Item = procedurecostPaymentResItem

	return response, nil
}

func (r *Resolver) ProcedureCostPaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		procedurecostPayment, err := r.Repo.GetProcedureCostPayment(id)
		if err != nil {
			return errors.HandleAPPError(err)
		}
		procedurecostPaymentResItem, err := buildProcedureCostPaymentResponseItem(*procedurecostPayment)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.ProcedureCostPaymentResponseItem{procedurecostPaymentResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetProcedureCostPaymentListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["procedure_cost_id"].(int); ok && value != 0 {
		input.ProcedureCostID = &value
	}

	procedurecostPayments, total, err := r.Repo.GetProcedureCostPaymentList(&input)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	procedurecostResItem, err := buildProcedureCostPaymentResponseItemList(procedurecostPayments)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   procedurecostResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) ProcedureCostPaymentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteProcedureCostPayment(params.Context, itemID)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildProcedureCostPaymentResponseItem(procedurecostPayment structs.ProcedureCostPayment) (*dto.ProcedureCostPaymentResponseItem, error) {
	status := dto.DropdownSimple{
		ID:    int(structs.PaidProcedureCostPeymentStatus),
		Title: string(dto.FinancialProcedureCostPaymentStatusPaid),
	}

	switch procedurecostPayment.Status {
	case structs.PaidProcedureCostPeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PaidProcedureCostPeymentStatus),
			Title: string(dto.FinancialProcedureCostPaymentStatusPaid),
		}
	case structs.CancelledProcedureCostPeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.CancelledProcedureCostPeymentStatus),
			Title: string(dto.FinancialProcedureCostPaymentStatusCanceled),
		}
	case structs.RetunedProcedureCostPeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.RetunedProcedureCostPeymentStatus),
			Title: string(dto.FinancialProcedureCostPaymentStatusReturned),
		}
	}

	procedurecostPaymentMethod := dto.DropdownSimple{
		ID:    int(structs.PaymentProcedureCostPeymentMethod),
		Title: string(dto.FinancialProcedureCostPaymentMethodPayment),
	}

	switch procedurecostPayment.PaymentMethod {
	case structs.PaymentProcedureCostPeymentMethod:
		procedurecostPaymentMethod = dto.DropdownSimple{
			ID:    int(structs.PaymentProcedureCostPeymentMethod),
			Title: string(dto.FinancialProcedureCostPaymentMethodPayment),
		}
	case structs.ForcedProcedureCostPeymentMethod:
		procedurecostPaymentMethod = dto.DropdownSimple{
			ID:    int(structs.ForcedProcedureCostPeymentMethod),
			Title: string(dto.FinancialProcedureCostPaymentMethodForced),
		}
	case structs.CourtCostsProcedureCostPeymentMethod:
		procedurecostPaymentMethod = dto.DropdownSimple{
			ID:    int(structs.CourtCostsProcedureCostPeymentMethod),
			Title: string(dto.FinancialProcedureCostPaymentMethodCourtCosts),
		}
	}

	response := dto.ProcedureCostPaymentResponseItem{
		ID:                     procedurecostPayment.ID,
		ProcedureCostID:        procedurecostPayment.ProcedureCostID,
		PaymentMethod:          procedurecostPaymentMethod,
		Amount:                 procedurecostPayment.Amount,
		PaymentDate:            procedurecostPayment.PaymentDate,
		PaymentDueDate:         procedurecostPayment.PaymentDueDate,
		ReceiptNumber:          procedurecostPayment.ReceiptNumber,
		PaymentReferenceNumber: procedurecostPayment.PaymentReferenceNumber,
		DebitReferenceNumber:   procedurecostPayment.DebitReferenceNumber,
		Status:                 status,
		CreatedAt:              procedurecostPayment.CreatedAt,
		UpdatedAt:              procedurecostPayment.UpdatedAt,
	}

	return &response, nil
}

func buildProcedureCostPaymentResponseItemList(itemList []structs.ProcedureCostPayment) ([]*dto.ProcedureCostPaymentResponseItem, error) {
	var items []*dto.ProcedureCostPaymentResponseItem

	for _, item := range itemList {
		singleItem, err := buildProcedureCostPaymentResponseItem(item)

		if err != nil {
			return nil, errors.Wrap(err, "build procedure cost payment response item")
		}

		items = append(items, singleItem)

	}

	return items, nil
}
