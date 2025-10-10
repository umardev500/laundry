package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/servicecategory/domain"
)

type UpdateServiceCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (r *UpdateServiceCategoryRequest) ToDomain(id uuid.UUID) *domain.ServiceCategory {
	return &domain.ServiceCategory{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
	}
}
