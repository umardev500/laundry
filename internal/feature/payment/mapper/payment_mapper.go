package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
	"github.com/umardev500/laundry/internal/feature/payment/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"

	paymentMethodMapper "github.com/umardev500/laundry/internal/feature/paymentmethod/mapper"
)

// FromEnt converts an Ent Payment to a domain Payment
func FromEnt(e *ent.Payment) *domain.Payment {
	if e == nil {
		return nil
	}

	ref := getRef(e)

	return &domain.Payment{
		ID:              e.ID,
		UserID:          e.UserID,
		TenantID:        e.TenantID,
		RefID:           e.RefID,
		RefType:         types.PaymentType(e.RefType),
		PaymentMethodID: e.PaymentMethodID,
		Amount:          e.Amount,
		ReceivedAmount:  &e.ReceivedAmount,
		ChangeAmount:    &e.ChangeAmount,
		Notes:           e.Notes,
		Status:          types.PaymentStatus(e.Status),
		PaymentMethod:   paymentMethodMapper.FromEnt(e.Edges.PaymentMethod),
		PaidAt:          e.PaidAt,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		DeletedAt:       e.DeletedAt,
		Ref:             ref,
	}
}

// FromEntList converts a slice of Ent Payments to domain Payments
func FromEntList(ents []*ent.Payment) []*domain.Payment {
	payments := make([]*domain.Payment, len(ents))
	for i, e := range ents {
		payments[i] = FromEnt(e)
	}
	return payments
}

// ToResponse converts a domain Payment to a PaymentResponse DTO
func ToResponse(d *domain.Payment, refMapper types.RefMapper) *dto.PaymentResponse {
	if d == nil {
		return nil
	}

	var ref any
	if refMapper != nil && d.Ref != nil {
		ref = refMapper(d.Ref)
	}

	return &dto.PaymentResponse{
		ID:              d.ID,
		UserID:          d.UserID,
		TenantID:        d.TenantID,
		RefID:           d.RefID,
		RefType:         d.RefType,
		PaymentMethodID: d.PaymentMethodID,
		Amount:          d.Amount,
		ReceivedAmount:  d.ReceivedAmount,
		ChangeAmount:    d.ChangeAmount,
		Notes:           d.Notes,
		Status:          d.Status,
		PaymentMethod:   paymentMethodMapper.ToResponse(d.PaymentMethod),
		PaidAt:          d.PaidAt,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
		DeletedAt:       d.DeletedAt,
		Ref:             ref,
	}
}

// ToResponseList converts a slice of domain Payments to PaymentResponse DTOs
func ToResponseList(payments []*domain.Payment, refMapper types.RefMapper) []*dto.PaymentResponse {
	res := make([]*dto.PaymentResponse, len(payments))
	for i, d := range payments {
		res[i] = ToResponse(d, refMapper)
	}
	return res
}

// ToResponsePage converts paginated domain Payments to DTO pagination
func ToResponsePage(data *pagination.PageData[domain.Payment], refMapper types.RefMapper) *pagination.PageData[dto.PaymentResponse] {
	if data == nil || len(data.Data) == 0 {
		return &pagination.PageData[dto.PaymentResponse]{
			Data:  []*dto.PaymentResponse{},
			Total: 0,
		}
	}

	payments := ToResponseList(data.Data, refMapper)

	return &pagination.PageData[dto.PaymentResponse]{
		Data:  payments,
		Total: data.Total,
	}
}

// getRef returns the ref of a payment
func getRef(e *ent.Payment) any {
	if e == nil {
		return nil
	}

	refHandlers := map[types.PaymentType]func() any{
		types.PaymentTypeOrder: func() any {
			return e.Edges.Order
		},
		types.PaymentTypeSubscription: func() any {
			return nil
		},
	}

	if handler, ok := refHandlers[types.PaymentType(e.RefType)]; ok {
		return handler()
	}

	return nil
}
