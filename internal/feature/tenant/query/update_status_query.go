package query

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// UpdateStatusQuery holds parameters for updating a tenant’s status.
type UpdateStatusQuery struct {
	ID     string `params:"id"`
	Status string `json:"status" query:"status" form:"status"`
}

// UUID parses the tenant ID safely.
func (q *UpdateStatusQuery) UUID() (uuid.UUID, error) {
	if q.ID == "" {
		return uuid.Nil, fmt.Errorf("id is required")
	}
	return uuid.Parse(q.ID)
}

// Validate ensures ID and status fields are valid.
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
		return errors.New("invalid status value: must be 'active' or 'suspended'")
	}
}

// ToDomainTenantWithID maps the query to a domain model.
func (q *UpdateStatusQuery) ToDomainTenantWithID(uid uuid.UUID) *domain.Tenant {
	return &domain.Tenant{
		ID:     uid,
		Status: types.Status(q.Status),
	}
}
