package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/serviceunit/domain"
)

type CreateServiceUnitRequest struct {
	TenantID uuid.UUID `json:"tenant_id" validate:"required"`
	Name     string    `json:"name" validate:"required,min=2,max=100"`
	Symbol   string    `json:"symbol,omitempty" validate:"omitempty,max=10"`
}

func (r *CreateServiceUnitRequest) ToDomain(ctx *appctx.Context) *domain.ServiceUnit {
	if ctx.TenantID() != nil && r.TenantID == uuid.Nil {
		r.TenantID = *ctx.TenantID()
	}
	return &domain.ServiceUnit{
		TenantID: r.TenantID,
		Name:     r.Name,
		Symbol:   r.Symbol,
	}
}
