package query

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/machine/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type UpdateStatusQuery struct {
	Status string `json:"status" query:"status" form:"status"`
}

func (q *UpdateStatusQuery) ToDomain(id uuid.UUID) (*domain.Machine, error) {
	return &domain.Machine{
		ID:     id,
		Status: types.MachineStatus(q.Status),
	}, nil
}
