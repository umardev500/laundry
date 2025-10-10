package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/service/domain"
	"github.com/umardev500/laundry/pkg/utils"
)

type CreateServiceRequest struct {
	TenantID          uuid.UUID `json:"tenant_id" validate:"required"`
	ServiceUnitID     uuid.UUID `json:"service_unit_id,omitempty"`
	ServiceCategoryID uuid.UUID `json:"service_category_id,omitempty"`
	Name              string    `json:"name" validate:"required,min=2,max=200"`
	Price             float64   `json:"price,omitempty"`
	Description       string    `json:"description,omitempty" validate:"omitempty,max=1024"`
}

func (r *CreateServiceRequest) ToDomain(ctx *appctx.Context) *domain.Service {
	if ctx.TenantID() != nil && r.TenantID == uuid.Nil {
		r.TenantID = *ctx.TenantID()
	}

	return &domain.Service{
		TenantID:          r.TenantID,
		ServiceUnitID:     utils.NilIfUUIDZero(r.ServiceUnitID),
		ServiceCategoryID: utils.NilIfUUIDZero(r.ServiceCategoryID),
		Name:              r.Name,
		Price:             r.Price,
		Description:       r.Description,
	}
}
