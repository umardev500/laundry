package query

import "github.com/umardev500/laundry/pkg/pagination"

type DistrictOrder string

const (
	DistrictOrderNameAsc  DistrictOrder = "name_asc"
	DistrictOrderNameDesc DistrictOrder = "name_desc"
)

type ListDistrictQuery struct {
	pagination.Query
	Search string        `query:"search"`
	Order  DistrictOrder `query:"order"`
}

// Normalize ensures defaults are set.
func (q *ListDistrictQuery) Normalize() {
	q.Query.Normalize(1, 10) // default page=1, limit=10
	if q.Order == "" {
		q.Order = DistrictOrderNameAsc
	}
}
