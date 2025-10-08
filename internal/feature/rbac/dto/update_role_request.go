package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
)

// UpdateRoleRequest defines the input data for updating an existing role.
type UpdateRoleRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

// ToDomain converts the request to a domain.Role object.
func (r *UpdateRoleRequest) ToDomain(roleID uuid.UUID, tenantID *uuid.UUID) (*domain.Role, error) {
	role := &domain.Role{
		ID:          roleID,
		TenantID:    tenantID,
		Name:        r.Name,
		Description: r.Description,
	}
	return role, nil
}
