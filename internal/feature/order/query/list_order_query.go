package query

import "github.com/umardev500/laundry/pkg/pagination"

// OrderBy defines valid sorting options for orders.
type OrderBy string

const (
	OrderCreatedAtAsc    OrderBy = "created_at_asc"
	OrderCreatedAtDesc   OrderBy = "created_at_desc"
	OrderUpdatedAtAsc    OrderBy = "updated_at_asc"
	OrderUpdatedAtDesc   OrderBy = "updated_at_desc"
	OrderTotalAmountAsc  OrderBy = "total_asc"
	OrderTotalAmountDesc OrderBy = "total_desc"
)

// ListOrderQuery defines filtering, sorting, and pagination for order listing.
type ListOrderQuery struct {
	pagination.Query
	Search               string  `query:"search"`
	Status               *string `query:"status"`
	IncludeDeleted       bool    `query:"include_deleted"`
	IncludeItems         bool    `query:"include_items"`
	IncludePayment       bool    `query:"include_payment"`
	IncludePaymentMethod bool    `query:"include_payment_method"`
	Order                OrderBy `query:"order"`
}

// Normalize applies default pagination and sort values.
func (q *ListOrderQuery) Normalize() {
	q.Query.Normalize(1, 10)
	if q.Order == "" {
		q.Order = OrderCreatedAtDesc
	}
}
