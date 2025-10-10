package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machine/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type CreateMachineRequest struct {
	TenantID      uuid.UUID `json:"tenant_id" validate:"required"`
	MachineTypeID uuid.UUID `json:"machine_type_id" validate:"required"`
	Name          string    `json:"name" validate:"required,min=2,max=100"`
	Description   string    `json:"description,omitempty" validate:"omitempty,max=255"`
	Status        string    `json:"status,omitempty" validate:"omitempty,oneof=available in_use maintenance offline reserved"`
}

func (r *CreateMachineRequest) ToDomain(ctx *appctx.Context) (*domain.Machine, error) {
	if r.TenantID == uuid.Nil && ctx.TenantID() == nil {
		return nil, types.ErrTenantIDRequired
	}

	if r.TenantID == uuid.Nil {
		r.TenantID = *ctx.TenantID()
	}

	return &domain.Machine{
		TenantID: r.TenantID,
		MachineTypeID: func() *uuid.UUID {
			if r.MachineTypeID == uuid.Nil {
				return nil
			}
			return &r.MachineTypeID
		}(),
		Name:        r.Name,
		Description: r.Description,
		Status:      types.MachineStatus(r.Status),
	}, nil
}
