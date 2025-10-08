package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
)

type UpdatePermissionRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

func (r *UpdatePermissionRequest) ToDomain(id uuid.UUID) (*domain.Permission, error) {
	return &domain.Permission{
		ID:          id,
		Name:        r.Name,
		DisplayName: r.DisplayName,
		Description: r.Description,
	}, nil
}
