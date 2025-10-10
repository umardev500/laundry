package query

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/machine/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type UpdateStatusQuery struct {
	ID     string `params:"id"`
	Status string `json:"status" query:"status" form:"status"`
}

func (q *UpdateStatusQuery) UUID() (uuid.UUID, error) {
	if q.ID == "" {
		return uuid.Nil, fmt.Errorf("id is required")
	}
	return uuid.Parse(q.ID)
}

func (q *UpdateStatusQuery) Validate() error {
	if q.ID == "" {
		return fmt.Errorf("id is required")
	}
	if q.Status == "" {
		return fmt.Errorf("status is required")
	}

	switch types.MachineStatus(q.Status) {
	case types.MachineStatusAvailable, types.MachineStatusInUse, types.MachineStatusMaintenance, types.MachineStatusOffline, types.MachineStatusReserved:
		return nil
	default:
		return errors.New("invalid status value")
	}
}

func (q *UpdateStatusQuery) ToDomain() (*domain.Machine, error) {
	uid, err := q.UUID()
	if err != nil {
		return nil, err
	}

	return &domain.Machine{
		ID:     uid,
		Status: types.MachineStatus(q.Status),
	}, nil
}
