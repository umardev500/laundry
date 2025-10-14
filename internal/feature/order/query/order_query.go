package query

type OrderQuery struct {
	IncludeDeleted       bool          `query:"include_deleted"`
	IncludeItems         bool          `query:"include_items"`
	IncludePayment       bool          `query:"include_payment"`
	IncludePaymentMethod bool          `query:"include_payment_method"`
	IncludeStatuses      bool          `query:"include_statuses"`
	StatusOrder          StatusesOrder `query:"status_order"`
}

// Normalize applies default pagination and sort values.
func (q *OrderQuery) Normalize() {
	if q.StatusOrder == "" {
		q.StatusOrder = StatusesOrderAsc
	}
}
