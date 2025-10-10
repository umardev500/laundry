package query

import "github.com/umardev500/laundry/pkg/pagination"

type Order string

const (
	OrderNameAsc       Order = "name_asc"
	OrderNameDesc      Order = "name_desc"
	OrderPriceAsc      Order = "price_asc"
	OrderPriceDesc     Order = "price_desc"
	OrderCreatedAtAsc  Order = "created_at_asc"
	OrderCreatedAtDesc Order = "created_at_desc"
)

type ListServiceQuery struct {
	pagination.Query
	Search         string `query:"search"`
	IncludeDeleted bool   `query:"include_deleted"`
	Order          Order  `query:"order"`
}

func (q *ListServiceQuery) Normalize() {
	q.Query.Normalize(1, 10)
	if q.Order == "" {
		q.Order = OrderCreatedAtDesc
	}
}
