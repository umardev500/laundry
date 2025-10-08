package query

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type UpdateStatusQuery struct {
	ID     string `params:"id"`
	Status string `params:"status"`
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

	switch types.Status(q.Status) {
	case types.StatusActive, types.StatusSuspended:
		return nil
	default:
		return errors.New("invalid status value: must be one of 'active', or 'suspended'")
	}
}

func (q *UpdateStatusQuery) ToDomainPlatformUserWithID(uid uuid.UUID) *domain.PlatformUser {
	return &domain.PlatformUser{
		ID:     uid,
		Status: types.Status(q.Status),
	}
}
