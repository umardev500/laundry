package dto

import (
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// CreateTenantRequest defines the expected payload for creating a tenant.
type CreateTenantRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=100"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone,omitempty" validate:"omitempty,min=6,max=20"`
	Address string `json:"address,omitempty" validate:"omitempty,max=255"`
	Status  string `json:"status,omitempty" validate:"omitempty,oneof=active suspended"`
}

// ToDomain converts the DTO to a domain.Tenant (recommended path)
func (r *CreateTenantRequest) ToDomain() (*domain.Tenant, error) {
	return &domain.Tenant{
		Name:   r.Name,
		Email:  r.Email,
		Phone:  r.Phone,
		Status: types.TenantStatus(r.Status),
	}, nil
}
