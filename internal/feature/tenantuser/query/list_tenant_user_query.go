package query

import "github.com/umardev500/laundry/pkg/pagination"

type Order string

const (
	OrderCreatedAtAsc  Order = "created_at_asc"
	OrderCreatedAtDesc Order = "created_at_desc"
	OrderUpdatedAtAsc  Order = "updated_at_asc"
	OrderUpdatedAtDesc Order = "updated_at_desc"
)

// ListTenantUserQuery defines pagination and filters for tenant users.
type ListTenantUserQuery struct {
	pagination.Query
	Search         string `query:"search"`
	Status         string `query:"status"`
	IncludeDeleted bool   `query:"include_deleted"`
	Order          Order  `query:"order"`
}

func (q *ListTenantUserQuery) Normalize() {
	q.Query.Normalize(1, 10)
	if q.Order == "" {
		q.Order = OrderCreatedAtDesc
	}
}
