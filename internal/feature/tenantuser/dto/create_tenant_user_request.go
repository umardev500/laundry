package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/tenantuser/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// CreateTenantUserRequest represents the payload to create a tenant user.
type CreateTenantUserRequest struct {
	UserID   uuid.UUID    `json:"user_id" validate:"required"`
	TenantID uuid.UUID    `json:"tenant_id" validate:"required"`
	Status   types.Status `json:"status" validate:"omitempty,oneof=active suspended deleted"`
}

// ToDomain converts the DTO to a domain model. Caller can set ID if needed.
func (r *CreateTenantUserRequest) ToDomain() *domain.TenantUser {
	status := r.Status
	if status == "" {
		status = types.StatusActive
	}
	return &domain.TenantUser{
		UserID:   r.UserID,
		TenantID: r.TenantID,
		Status:   status,
	}
}
