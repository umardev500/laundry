package query

import "github.com/umardev500/laundry/pkg/pagination"

type OrderBy string

const (
	OrderNameAsc       OrderBy = "name_asc"
	OrderNameDesc      OrderBy = "name_desc"
	OrderCreatedAtAsc  OrderBy = "created_at_asc"
	OrderCreatedAtDesc OrderBy = "created_at_desc"
)

type ListServiceCategoryQuery struct {
	pagination.Query
	Search         string  `query:"search"`
	Order          OrderBy `query:"order"`
	IncludeDeleted bool    `query:"include_deleted"`
}

func (q *ListServiceCategoryQuery) Normalize() {
	q.Query.Normalize(1, 10)
	if q.Order == "" {
		q.Order = OrderCreatedAtDesc
	}
}
