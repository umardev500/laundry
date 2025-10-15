package query

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// UpdateStatusQuery holds parameters for updating a tenant’s status.
type UpdateStatusQuery struct {
	Status string `json:"status" query:"status" form:"status"`
}

// ToDomain maps the query to a domain model.
func (q *UpdateStatusQuery) ToDomain(uid uuid.UUID) *domain.Tenant {
	return &domain.Tenant{
		ID:     uid,
		Status: types.TenantStatus(q.Status),
	}
}
