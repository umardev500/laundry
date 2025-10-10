package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/serviceunit/domain"
)

type UpdateServiceUnitRequest struct {
	Name   string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Symbol string `json:"symbol,omitempty" validate:"omitempty,max=10"`
}

func (r *UpdateServiceUnitRequest) ToDomain(id uuid.UUID) *domain.ServiceUnit {
	return &domain.ServiceUnit{
		ID:     id,
		Name:   r.Name,
		Symbol: r.Symbol,
	}
}
