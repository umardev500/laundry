package query

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

type OrderStatusBy string

const (
	StatusCreatedAtAsc  OrderStatusBy = "created_at_asc"
	StatusCreatedAtDesc OrderStatusBy = "created_at_desc"
)

// -------------------------
// OrderStatusHistoryListQuery
// -------------------------

type OrderStatusHistoryListQuery struct {
	pagination.Query
	OrderID         *uuid.UUID        `query:"order_id"`
	Status          types.OrderStatus `query:"status"`
	IncludeNotes    bool              `query:"include_notes"`
	Order           OrderStatusBy     `query:"order"`
	IncludeOrder    bool              `query:"include_order"`
	IncludeOrderRef bool              `query:"include_order_ref"`
}

// Methods for OrderStatusHistoryListQuery
func (q *OrderStatusHistoryListQuery) Normalize() {
	q.Query.Normalize(1, 10)
	if q.Order == "" {
		q.Order = StatusCreatedAtDesc
	}
}
