package query

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type UpdateFeatureStatusQuery struct {
	ID     string `params:"id"`
	Status string `params:"status"`
}

// UUID parses the tenant ID safely.
func (q *UpdateFeatureStatusQuery) UUID() (uuid.UUID, error) {
	if q.ID == "" {
		return uuid.Nil, fmt.Errorf("id is required")
	}
	return uuid.Parse(q.ID)
}

// Validate ensures ID and status fields are valid.
func (q *UpdateFeatureStatusQuery) Validate() error {
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
		return errors.New("invalid status value: must be 'active' or 'suspended'")
	}
}

func (q *UpdateFeatureStatusQuery) ToDomain() (*domain.Feature, error) {
	uid, err := q.UUID()
	if err != nil {
		return nil, err
	}

	return &domain.Feature{
		ID:     uid,
		Status: types.Status(q.Status),
	}, nil
}
