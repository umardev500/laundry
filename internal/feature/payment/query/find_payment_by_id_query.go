package query

type FindPaymentByIdQuery struct {
	IncludeDeleted bool `query:"include_deleted"`
	IncludeRef     bool `query:"include_ref"`
}
