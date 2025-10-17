package query

import "github.com/umardev500/laundry/pkg/pagination"

type ProviceOrder string

const (
	ProviceOrderNameAsc  DistrictOrder = "name_asc"
	ProviceOrderNameDesc DistrictOrder = "name_desc"
)

type ListProvinceQuery struct {
	pagination.Query
	Search string        `query:"search"`
	Order  DistrictOrder `query:"order"`
}

// Normalize ensures defaults are set.
func (q *ListProvinceQuery) Normalize() {
	q.Query.Normalize(1, 10) // default page=1, limit=10
	if q.Order == "" {
		q.Order = ProviceOrderNameAsc
	}
}
