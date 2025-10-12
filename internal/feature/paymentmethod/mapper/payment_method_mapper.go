package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/domain"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

// FromEnt converts an Ent PaymentMethod to domain.PaymentMethod
func FromEnt(e *ent.PaymentMethod) *domain.PaymentMethod {
	if e == nil {
		return nil
	}
	return &domain.PaymentMethod{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}

// FromEntList converts a list of Ent PaymentMethods to domain.PaymentMethod slice
func FromEntList(list []*ent.PaymentMethod) []*domain.PaymentMethod {
	res := make([]*domain.PaymentMethod, len(list))
	for i, e := range list {
		res[i] = FromEnt(e)
	}
	return res
}

// ToResponse converts a domain.PaymentMethod to a DTO response
func ToResponse(d *domain.PaymentMethod) *dto.PaymentMethodResponse {
	if d == nil {
		return nil
	}
	return &dto.PaymentMethodResponse{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		DeletedAt:   d.DeletedAt,
	}
}

// ToResponseList converts a slice of domain.PaymentMethod to DTO slice
func ToResponseList(list []*domain.PaymentMethod) []*dto.PaymentMethodResponse {
	res := make([]*dto.PaymentMethodResponse, len(list))
	for i, d := range list {
		res[i] = ToResponse(d)
	}
	return res
}

// ToResponsePage converts paginated domain.PaymentMethod to paginated DTO response
func ToResponsePage(data *pagination.PageData[domain.PaymentMethod]) *pagination.PageData[dto.PaymentMethodResponse] {
	res := make([]*dto.PaymentMethodResponse, len(data.Data))
	for i, d := range data.Data {
		res[i] = ToResponse(d)
	}
	return &pagination.PageData[dto.PaymentMethodResponse]{
		Data:  res,
		Total: data.Total,
	}
}
