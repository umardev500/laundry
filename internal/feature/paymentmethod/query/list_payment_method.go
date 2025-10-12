package query

import "github.com/umardev500/laundry/pkg/pagination"

type Order string

const (
	OrderNameAsc       Order = "name_asc"
	OrderNameDesc      Order = "name_desc"
	OrderCreatedAtAsc  Order = "created_at_asc"
	OrderCreatedAtDesc Order = "created_at_desc"
)

// ListPaymentMethodQuery defines filters and pagination for listing payment methods
type ListPaymentMethodQuery struct {
	pagination.Query
	Search         string `query:"search"`
	IncludeDeleted bool   `query:"include_deleted"`
	Order          Order  `query:"order"`
}

// Normalize sets default pagination and ordering
func (q *ListPaymentMethodQuery) Normalize() {
	q.Query.Normalize(1, 10) // default page=1, limit=10
	if q.Order == "" {
		q.Order = OrderCreatedAtDesc
	}
}
