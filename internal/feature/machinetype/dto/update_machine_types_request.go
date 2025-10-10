package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/machinetype/domain"
)

type UpdateMachineTypeRequest struct {
	Name        string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Description string `json:"description,omitempty" validate:"omitempty,max=255"`
	Capacity    int    `json:"capacity,omitempty" validate:"omitempty,min=1"`
}

func (r *UpdateMachineTypeRequest) ToDomain(id uuid.UUID) (*domain.MachineType, error) {
	return &domain.MachineType{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
		Capacity:    r.Capacity,
	}, nil
}
