package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) FinePaymentInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.FinePayment
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var item *structs.FinePayment

	if data.ID == 0 {
		item, err = r.Repo.CreateFinePayment(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		item, err = r.Repo.UpdateFinePayment(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	}

	finePaymentResItem, err := buildFinePaymentResponseItem(*item)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Item = finePaymentResItem

	return response, nil
}

func (r *Resolver) FinePaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		finePayment, err := r.Repo.GetFinePayment(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		finePaymentResItem, err := buildFinePaymentResponseItem(*finePayment)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.FinePaymentResponseItem{finePaymentResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetFinePaymentListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["fine_id"].(int); ok && value != 0 {
		input.FineID = &value
	}

	finePayments, total, err := r.Repo.GetFinePaymentList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	fineResItem, err := buildFinePaymentResponseItemList(finePayments)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   fineResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) FinePaymentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteFinePayment(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildFinePaymentResponseItem(finePayment structs.FinePayment) (*dto.FinePaymentResponseItem, error) {
	status := dto.DropdownSimple{
		ID:    int(structs.PaidFinePeymentStatus),
		Title: string(dto.FinancialFinePaymentStatusPaid),
	}

	switch finePayment.Status {
	case structs.PaidFinePeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PaidFinePeymentStatus),
			Title: string(dto.FinancialFinePaymentStatusPaid),
		}
	case structs.CancelledFinePeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.CancelledFinePeymentStatus),
			Title: string(dto.FinancialFinePaymentStatusCanceled),
		}
	case structs.RetunedFinePeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.RetunedFinePeymentStatus),
			Title: string(dto.FinancialFinePaymentStatusReturned),
		}
	}

	finePaymentMethod := dto.DropdownSimple{
		ID:    int(structs.PaymentFinePeymentMethod),
		Title: string(dto.FinancialFinePaymentMethodPayment),
	}

	switch finePayment.PaymentMethod {
	case structs.PaymentFinePeymentMethod:
		finePaymentMethod = dto.DropdownSimple{
			ID:    int(structs.PaymentFinePeymentMethod),
			Title: string(dto.FinancialFinePaymentMethodPayment),
		}
	case structs.ForcedFinePeymentMethod:
		finePaymentMethod = dto.DropdownSimple{
			ID:    int(structs.ForcedFinePeymentMethod),
			Title: string(dto.FinancialFinePaymentMethodForced),
		}
	case structs.CourtCostsFinePeymentMethod:
		finePaymentMethod = dto.DropdownSimple{
			ID:    int(structs.CourtCostsFinePeymentMethod),
			Title: string(dto.FinancialFinePaymentMethodCourtCosts),
		}
	}
	response := dto.FinePaymentResponseItem{
		ID:                     finePayment.ID,
		FineID:                 finePayment.FineID,
		PaymentMethod:          finePaymentMethod,
		Amount:                 finePayment.Amount,
		PaymentDate:            finePayment.PaymentDate,
		PaymentDueDate:         finePayment.PaymentDueDate,
		ReceiptNumber:          finePayment.ReceiptNumber,
		PaymentReferenceNumber: finePayment.PaymentReferenceNumber,
		DebitReferenceNumber:   finePayment.DebitReferenceNumber,
		Status:                 status,
		CreatedAt:              finePayment.CreatedAt,
		UpdatedAt:              finePayment.UpdatedAt,
	}

	return &response, nil
}

func buildFinePaymentResponseItemList(itemList []structs.FinePayment) ([]*dto.FinePaymentResponseItem, error) {
	var items []*dto.FinePaymentResponseItem

	for _, item := range itemList {
		singleItem, err := buildFinePaymentResponseItem(item)

		if err != nil {
			return nil, errors.Wrap(err, "build fine payment response item")
		}

		items = append(items, singleItem)

	}

	return items, nil
}
