package query

import "github.com/umardev500/laundry/pkg/pagination"

type Order string

const (
	OrderEmailAsc      Order = "email_asc"
	OrderEmailDesc     Order = "email_desc"
	OrderCreatedAtAsc  Order = "created_at_asc"
	OrderCreatedAtDesc Order = "created_at_desc"
)

type ListUserQuery struct {
	pagination.Query
	Search string `query:"search"`
	Status string `query:"status"`
	Order  Order  `query:"order"`
}

func (q *ListUserQuery) Normalize() {
	q.Query.Normalize(1, 10)

	if q.Order == "" {
		q.Order = OrderCreatedAtAsc
	}
}
