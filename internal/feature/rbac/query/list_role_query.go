package query

import "github.com/umardev500/laundry/pkg/pagination"

// Order defines sorting options for roles.
type Order string

const (
	OrderNameAsc       FeatureOrder = "name_asc"
	OrderNameDesc      FeatureOrder = "name_desc"
	OrderCreatedAtAsc  FeatureOrder = "created_at_asc"
	OrderCreatedAtDesc FeatureOrder = "created_at_desc"
	OrderUpdatedAtAsc  FeatureOrder = "updated_at_asc"
	OrderUpdatedAtDesc FeatureOrder = "updated_at_desc"
)

// ListRoleQuery holds pagination and filtering options for roles.
type ListRoleQuery struct {
	pagination.Query
	TenantID       string       `query:"tenant_id"`
	Search         string       `query:"search"`
	IncludeDeleted bool         `query:"include_deleted"`
	Order          FeatureOrder `query:"order"`
}

// Normalize ensures default values for pagination and order.
func (q *ListRoleQuery) Normalize() {
	q.Query.Normalize(1, 10)

	if q.Order == "" {
		q.Order = OrderCreatedAtDesc
	}
}
