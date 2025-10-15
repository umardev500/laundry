package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenantuser/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// CreateTenantUserRequest represents the payload to create a tenant user.
type CreateTenantUserRequest struct {
	UserID   uuid.UUID              `json:"user_id" validate:"required"`
	TenantID *uuid.UUID             `json:"tenant_id,omitempty"`
	Status   types.TenantUserStatus `json:"status" validate:"omitempty,oneof=active suspended deleted"`
}

// ToDomain converts the DTO to a domain model. Caller can set ID if needed.
func (r *CreateTenantUserRequest) ToDomain(ctx *appctx.Context) *domain.TenantUser {
	status := r.Status
	if status == "" {
		status = types.TenantUserStatusActive
	}

	tenantID := ctx.TenantID()
	if r.TenantID != nil {
		tenantID = r.TenantID
	}

	return &domain.TenantUser{
		UserID:   r.UserID,
		TenantID: *tenantID,
		Status:   status,
	}
}
