package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) PropBenConfPaymentInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.PropBenConfPayment
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

	var item *structs.PropBenConfPayment

	if data.ID == 0 {
		item, err = r.Repo.CreatePropBenConfPayment(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdatePropBenConfPayment(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

	}

	response.Item = *item

	return response, nil
}

func (r *Resolver) PropBenConfPaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		PropBenConfPayment, err := r.Repo.GetPropBenConfPayment(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		PropBenConfPaymentResItem, err := buildPropBenConfPaymentResponseItem(*PropBenConfPayment)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.PropBenConfPaymentResponseItem{PropBenConfPaymentResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetPropBenConfPaymentListInputMS{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["property_benefits_confiscation_id"].(int); ok && value != 0 {
		input.PropBenConfID = &value
	}

	PropBenConfPayments, total, err := r.Repo.GetPropBenConfPaymentList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	PropBenConfResItem, err := buildPropBenConfPaymentResponseItemList(PropBenConfPayments)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   PropBenConfResItem,
		Total:   total,
	}, nil
}

func (r *Resolver) PropBenConfPaymentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeletePropBenConfPayment(itemID)
	if err != nil {
		fmt.Printf("Deleting property benefit confiscation payment item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildPropBenConfPaymentResponseItem(PropBenConfPayment structs.PropBenConfPayment) (*dto.PropBenConfPaymentResponseItem, error) {
	status := dto.FinancialPropBenConfPaymentStatusPaid
	switch PropBenConfPayment.Status {
	case structs.PaidPropBenConfPeymentStatus:
		status = dto.FinancialPropBenConfPaymentStatusPaid
	case structs.CancelledPropBenConfPeymentStatus:
		status = dto.FinancialPropBenConfPaymentStatusCanceled
	case structs.RetunedPropBenConfPeymentStatus:
		status = dto.FinancialPropBenConfPaymentStatusReturned
	}

	PropBenConfPaymentMethod := dto.FinancialPropBenConfPaymentMethodPayment
	switch PropBenConfPayment.PaymentMethod {
	case structs.PaymentPropBenConfPeymentMethod:
		PropBenConfPaymentMethod = dto.FinancialPropBenConfPaymentMethodPayment
	case structs.ForcedPropBenConfPeymentMethod:
		PropBenConfPaymentMethod = dto.FinancialPropBenConfPaymentMethodForced
	case structs.CourtCostsPropBenConfPeymentMethod:
		PropBenConfPaymentMethod = dto.FinancialPropBenConfPaymentMethodCourtCosts
	}
	response := dto.PropBenConfPaymentResponseItem{
		ID:                     PropBenConfPayment.ID,
		PropBenConfID:          PropBenConfPayment.PropBenConfID,
		PaymentMethod:          PropBenConfPaymentMethod,
		Amount:                 PropBenConfPayment.Amount,
		PaymentDate:            PropBenConfPayment.PaymentDate,
		PaymentDueDate:         PropBenConfPayment.PaymentDueDate,
		ReceiptNumber:          PropBenConfPayment.ReceiptNumber,
		PaymentReferenceNumber: PropBenConfPayment.PaymentReferenceNumber,
		DebitReferenceNumber:   PropBenConfPayment.DebitReferenceNumber,
		Status:                 status,
		CreatedAt:              PropBenConfPayment.CreatedAt,
		UpdatedAt:              PropBenConfPayment.UpdatedAt,
	}

	return &response, nil
}

func buildPropBenConfPaymentResponseItemList(itemList []structs.PropBenConfPayment) ([]*dto.PropBenConfPaymentResponseItem, error) {
	var items []*dto.PropBenConfPaymentResponseItem

	for _, item := range itemList {
		singleItem, err := buildPropBenConfPaymentResponseItem(item)

		if err != nil {
			return nil, err
		}

		items = append(items, singleItem)

	}

	return items, nil
}
