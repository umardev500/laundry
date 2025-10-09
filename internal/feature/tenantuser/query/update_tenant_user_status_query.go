package query

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/tenantuser/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// UpdateStatusQuery represents query params for updating a tenant user's status.
type UpdateStatusQuery struct {
	ID     string `params:"id"`
	Status string `params:"status"`
}

// UUID safely parses the TenantUser ID.
func (q *UpdateStatusQuery) UUID() (uuid.UUID, error) {
	if q.ID == "" {
		return uuid.Nil, fmt.Errorf("id is required")
	}
	return uuid.Parse(q.ID)
}

// Validate ensures ID and Status are valid.
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
		return errors.New("invalid status value: must be one of 'active' or 'suspended'")
	}
}

// ToDomainTenantUserWithID maps query data to a domain.TenantUser with the parsed ID.
func (q *UpdateStatusQuery) ToDomain() (*domain.TenantUser, error) {
	uid, err := q.UUID()
	if err != nil {
		return nil, err
	}

	return &domain.TenantUser{
		ID:     uid,
		Status: types.Status(q.Status),
	}, nil
}
