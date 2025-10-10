package query

import "github.com/umardev500/laundry/pkg/pagination"

type Order string

const (
	OrderNameAsc       Order = "name_asc"
	OrderNameDesc      Order = "name_desc"
	OrderCreatedAtAsc  Order = "created_at_asc"
	OrderCreatedAtDesc Order = "created_at_desc"
)

type ListServiceUnitQuery struct {
	pagination.Query
	Search string `query:"search"`
	Order  Order  `query:"order"`
}

func (q *ListServiceUnitQuery) Normalize() {
	q.Query.Normalize(1, 10)
	if q.Order == "" {
		q.Order = OrderCreatedAtDesc
	}
}
