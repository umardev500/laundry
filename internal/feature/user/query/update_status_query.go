package query

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/user/domain"
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

	switch domain.Status(q.Status) {
	case domain.StatusActive, domain.StatusSuspended:
		return nil
	default:
		return errors.New("invalid status value: must be one of 'active', or 'suspended'")
	}
}

func (q *UpdateStatusQuery) ToDomainUserWithID(uid uuid.UUID) *domain.User {
	return &domain.User{
		ID:     uid,
		Status: domain.Status(q.Status),
	}
}
