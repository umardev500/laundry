package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/servicecategory/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type CreateServiceCategoryRequest struct {
	TenantID    uuid.UUID `json:"tenant_id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
}

func (r *CreateServiceCategoryRequest) ToDomain(ctx *appctx.Context) (*domain.ServiceCategory, error) {
	// 1️⃣ If context has TenantID and request doesn’t — inherit it
	if ctx.TenantID() != nil && r.TenantID == uuid.Nil {
		r.TenantID = *ctx.TenantID()
	}

	// 2️⃣ If still empty after that — reject
	if r.TenantID == uuid.Nil {
		return nil, types.ErrTenantIDRequired
	}

	// 3️⃣ Build domain model
	return &domain.ServiceCategory{
		TenantID:    r.TenantID,
		Name:        r.Name,
		Description: r.Description,
	}, nil
}
