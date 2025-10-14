package query

type StatusHistoryByIDQuery struct {
	IncludeOrder    bool `query:"include_order"` // Include associated order
	IncludeOrderRef bool `query:"include_order_ref"`
}
