package query

import "github.com/umardev500/laundry/pkg/pagination"

type Order string

const (
	OrderCreatedAtAsc  Order = "created_at_asc"
	OrderCreatedAtDesc Order = "created_at_desc"
	OrderUpdatedAtAsc  Order = "updated_at_asc"
	OrderUpdatedAtDesc Order = "updated_at_desc"
)

type ListPlatformUserQuery struct {
	pagination.Query
	Search         string `json:"search,omitempty" query:"search"`
	Status         string `json:"status,omitempty" query:"status"`
	Order          Order  `json:"order,omitempty" query:"order"`
	IncludeDeleted bool   `json:"include_deleted,omitempty" query:"include_deleted"`
}

// Normalize sets default pagination and ordering values.
func (q *ListPlatformUserQuery) Normalize() {
	// Ensure default pagination
	q.Query.Normalize(1, 10)

	// Set default ordering if not specified
	if q.Order == "" {
		q.Order = OrderCreatedAtAsc
	}
}
