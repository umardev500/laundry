package query

import "github.com/umardev500/laundry/pkg/pagination"

type RegencyOrder string

const (
	RegencyOrderNameAsc  RegencyOrder = "name_asc"
	RegencyOrderNameDesc RegencyOrder = "name_desc"
)

type ListRegencyQuery struct {
	pagination.Query
	Search string       `query:"search"`
	Order  RegencyOrder `query:"order"`
}

// Normalize ensures defaults are set.
func (q *ListRegencyQuery) Normalize() {
	q.Query.Normalize(1, 10) // default page=1, limit=10
	if q.Order == "" {
		q.Order = RegencyOrderNameAsc
	}
}
