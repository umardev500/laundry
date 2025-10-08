package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
)

type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

// ToDomain converts the request DTO to a domain.Role.
func (r *CreateRoleRequest) ToDomain(tenantID *uuid.UUID) (*domain.Role, error) {
	role := &domain.Role{
		TenantID:    tenantID,
		Name:        r.Name,
		Description: r.Description,
	}
	return role, nil
}
