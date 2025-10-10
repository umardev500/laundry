package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machinetype/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type CreateMachineTypeRequest struct {
	TenantID    uuid.UUID `json:"tenant_id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=2,max=100"`
	Description string    `json:"description,omitempty" validate:"omitempty,max=255"`
	Capacity    int       `json:"capacity,omitempty" validate:"omitempty,min=1"`
}

// ToDomain converts req to domain, respecting context tenant when tenant id omitted.
func (r *CreateMachineTypeRequest) ToDomain(ctx *appctx.Context) (*domain.MachineType, error) {
	if r.TenantID == uuid.Nil && ctx.TenantID() == nil {
		return nil, types.ErrTenantIDRequired
	}
	if r.TenantID == uuid.Nil {
		r.TenantID = *ctx.TenantID()
	}

	return &domain.MachineType{
		TenantID:    r.TenantID,
		Name:        r.Name,
		Description: r.Description,
		Capacity:    r.Capacity,
	}, nil
}
