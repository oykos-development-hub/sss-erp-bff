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
		item, err = r.Repo.CreatePropBenConfPayment(params.Context, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		item, err = r.Repo.UpdatePropBenConfPayment(params.Context, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	propBenConfPaymentResItem, err := buildPropBenConfPaymentResponseItem(*item)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	response.Item = propBenConfPaymentResItem

	return response, nil
}

func (r *Resolver) PropBenConfPaymentOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		propBenConfPayment, err := r.Repo.GetPropBenConfPayment(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		propBenConfPaymentResItem, err := buildPropBenConfPaymentResponseItem(*propBenConfPayment)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.PropBenConfPaymentResponseItem{propBenConfPaymentResItem},
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

	err := r.Repo.DeletePropBenConfPayment(params.Context, itemID)
	if err != nil {
		fmt.Printf("Deleting property benefit confiscation payment item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildPropBenConfPaymentResponseItem(propbenconfPayment structs.PropBenConfPayment) (*dto.PropBenConfPaymentResponseItem, error) {
	status := dto.DropdownSimple{
		ID:    int(structs.PaidPropBenConfPeymentStatus),
		Title: string(dto.FinancialPropBenConfPaymentStatusPaid),
	}

	switch propbenconfPayment.Status {
	case structs.PaidPropBenConfPeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.PaidPropBenConfPeymentStatus),
			Title: string(dto.FinancialPropBenConfPaymentStatusPaid),
		}
	case structs.CancelledPropBenConfPeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.CancelledPropBenConfPeymentStatus),
			Title: string(dto.FinancialPropBenConfPaymentStatusCanceled),
		}
	case structs.RetunedPropBenConfPeymentStatus:
		status = dto.DropdownSimple{
			ID:    int(structs.RetunedPropBenConfPeymentStatus),
			Title: string(dto.FinancialPropBenConfPaymentStatusReturned),
		}
	}

	propbenconfPaymentMethod := dto.DropdownSimple{
		ID:    int(structs.PaymentPropBenConfPeymentMethod),
		Title: string(dto.FinancialPropBenConfPaymentMethodPayment),
	}

	switch propbenconfPayment.PaymentMethod {
	case structs.PaymentPropBenConfPeymentMethod:
		propbenconfPaymentMethod = dto.DropdownSimple{
			ID:    int(structs.PaymentPropBenConfPeymentMethod),
			Title: string(dto.FinancialPropBenConfPaymentMethodPayment),
		}
	case structs.ForcedPropBenConfPeymentMethod:
		propbenconfPaymentMethod = dto.DropdownSimple{
			ID:    int(structs.ForcedPropBenConfPeymentMethod),
			Title: string(dto.FinancialPropBenConfPaymentMethodForced),
		}
	case structs.CourtCostsPropBenConfPeymentMethod:
		propbenconfPaymentMethod = dto.DropdownSimple{
			ID:    int(structs.CourtCostsPropBenConfPeymentMethod),
			Title: string(dto.FinancialPropBenConfPaymentMethodCourtCosts),
		}
	}

	response := dto.PropBenConfPaymentResponseItem{
		ID:                     propbenconfPayment.ID,
		PropBenConfID:          propbenconfPayment.PropBenConfID,
		PaymentMethod:          propbenconfPaymentMethod,
		Amount:                 propbenconfPayment.Amount,
		PaymentDate:            propbenconfPayment.PaymentDate,
		PaymentDueDate:         propbenconfPayment.PaymentDueDate,
		ReceiptNumber:          propbenconfPayment.ReceiptNumber,
		PaymentReferenceNumber: propbenconfPayment.PaymentReferenceNumber,
		DebitReferenceNumber:   propbenconfPayment.DebitReferenceNumber,
		Status:                 status,
		CreatedAt:              propbenconfPayment.CreatedAt,
		UpdatedAt:              propbenconfPayment.UpdatedAt,
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
