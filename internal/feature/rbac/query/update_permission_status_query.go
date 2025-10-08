package query

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// UpdatePermissionStatusQuery is used to update the status of a permission (e.g. activate/suspend).
type UpdatePermissionStatusQuery struct {
	ID     string `params:"id"`
	Status string `params:"status"`
}

// UUID parses the permission ID safely.
func (q *UpdatePermissionStatusQuery) UUID() (uuid.UUID, error) {
	if q.ID == "" {
		return uuid.Nil, fmt.Errorf("id is required")
	}
	return uuid.Parse(q.ID)
}

// Validate ensures ID and status fields are valid.
func (q *UpdatePermissionStatusQuery) Validate() error {
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

// ToDomain converts the query into a domain.Permission entity.
func (q *UpdatePermissionStatusQuery) ToDomain() (*domain.Permission, error) {
	uid, err := q.UUID()
	if err != nil {
		return nil, err
	}

	return &domain.Permission{
		ID:     uid,
		Status: types.Status(q.Status),
	}, nil
}
