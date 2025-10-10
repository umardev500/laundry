package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/machine/domain"
)

type UpdateMachineRequest struct {
	Name        string `json:"name" validate:"omitempty,min=2,max=100"`
	Description string `json:"description,omitempty" validate:"omitempty,max=255"`
}

func (r *UpdateMachineRequest) ToDomain(id uuid.UUID) (*domain.Machine, error) {
	return &domain.Machine{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
	}, nil
}
