package query

import "github.com/umardev500/laundry/pkg/pagination"

// Order represents sorting options for machines.
type Order string

const (
	OrderNameAsc       Order = "name_asc"
	OrderNameDesc      Order = "name_desc"
	OrderCreatedAtAsc  Order = "created_at_asc"
	OrderCreatedAtDesc Order = "created_at_desc"
	OrderUpdatedAtAsc  Order = "updated_at_asc"
	OrderUpdatedAtDesc Order = "updated_at_desc"
)

type ListMachineQuery struct {
	pagination.Query
	Search         string `query:"search"`
	Status         string `query:"status"`
	IncludeDeleted bool   `query:"include_deleted"`
	Order          Order  `query:"order"`
}

func (q *ListMachineQuery) Normalize() {
	q.Query.Normalize(1, 10)
	if q.Order == "" {
		q.Order = OrderCreatedAtDesc
	}
}
