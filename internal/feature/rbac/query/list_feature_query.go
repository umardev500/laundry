package query

import "github.com/umardev500/laundry/pkg/pagination"

type FeatureOrder string

const (
	FeatureOrderNameAsc       FeatureOrder = "name_asc"
	FeatureOrderNameDesc      FeatureOrder = "name_desc"
	FeatureOrderCreatedAtAsc  FeatureOrder = "created_at_asc"
	FeatureOrderCreatedAtDesc FeatureOrder = "created_at_desc"
	FeatureOrderUpdatedAtAsc  FeatureOrder = "updated_at_asc"
	FeatureOrderUpdatedAtDesc FeatureOrder = "updated_at_desc"
)

type ListFeatureQuery struct {
	pagination.Query
	Search         string       `query:"search"`
	IncludeDeleted bool         `query:"include_deleted"`
	Order          FeatureOrder `query:"order"`
}

func (q *ListFeatureQuery) Normalize() {
	q.Query.Normalize(1, 10)
	if q.Order == "" {
		q.Order = OrderNameAsc
	}
}
