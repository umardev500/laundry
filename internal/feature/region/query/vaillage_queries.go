package query

import "github.com/umardev500/laundry/pkg/pagination"

type VillageOrder string

const (
	VillageOrderNameAsc  VillageOrder = "name_asc"
	VillageOrderNameDesc VillageOrder = "name_desc"
)

type ListVillageQuery struct {
	pagination.Query
	Search string       `query:"search"`
	Order  VillageOrder `query:"order"`
}

// Normalize ensures defaults are set.
func (q *ListVillageQuery) Normalize() {
	q.Query.Normalize(1, 10) // default page=1, limit=10
	if q.Order == "" {
		q.Order = VillageOrderNameAsc
	}
}
