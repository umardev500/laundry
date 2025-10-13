package query

import (
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

// OrderBy for payments
type PaymentOrder string

const (
	PaymentOrderCreatedAtAsc  PaymentOrder = "created_at_asc"
	PaymentOrderCreatedAtDesc PaymentOrder = "created_at_desc"
	PaymentOrderAmountAsc     PaymentOrder = "amount_asc"
	PaymentOrderAmountDesc    PaymentOrder = "amount_desc"
)

// ListPaymentQuery defines filters for listing payments
type ListPaymentQuery struct {
	pagination.Query
	Search         string              `query:"search"`          // Search by notes or other fields
	Order          PaymentOrder        `query:"order"`           // Order by created_at, amount, etc.
	Status         types.PaymentStatus `query:"status"`          // Filter by payment status (optional)
	RefType        types.PaymentType   `query:"ref_type"`        // Filter by payment type (optional)
	IncludeDeleted bool                `query:"include_deleted"` // Include soft-deleted payments
	IncludeRef     bool                `query:"include_ref"`
}

// Normalize sets default pagination and ordering values
func (q *ListPaymentQuery) Normalize() {
	q.Query.Normalize(1, 10) // Default page 1, 10 items per page
	if q.Order == "" {
		q.Order = PaymentOrderCreatedAtDesc
	}
}
