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

	return &domain.Payment{
		ID:              e.ID,
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
func ToResponse(d *domain.Payment) *dto.PaymentResponse {
	if d == nil {
		return nil
	}

	return &dto.PaymentResponse{
		ID:              d.ID,
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
	}
}

// ToResponseList converts a slice of domain Payments to PaymentResponse DTOs
func ToResponseList(payments []*domain.Payment) []*dto.PaymentResponse {
	res := make([]*dto.PaymentResponse, len(payments))
	for i, d := range payments {
		res[i] = ToResponse(d)
	}
	return res
}

// ToResponsePage converts paginated domain Payments to DTO pagination
func ToResponsePage(data *pagination.PageData[domain.Payment]) *pagination.PageData[dto.PaymentResponse] {
	if data == nil || len(data.Data) == 0 {
		return &pagination.PageData[dto.PaymentResponse]{
			Data:  []*dto.PaymentResponse{},
			Total: 0,
		}
	}

	payments := ToResponseList(data.Data)

	return &pagination.PageData[dto.PaymentResponse]{
		Data:  payments,
		Total: data.Total,
	}
}
