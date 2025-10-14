package query

type OrderQuery struct {
	IncludeDeleted       bool `query:"include_deleted"`
	IncludeItems         bool `query:"include_items"`
	IncludePayment       bool `query:"include_payment"`
	IncludePaymentMethod bool `query:"include_payment_method"`
	IncludeStatuses      bool `query:"include_statuses"`
}
