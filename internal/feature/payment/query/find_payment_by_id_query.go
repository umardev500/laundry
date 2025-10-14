package query

import "github.com/google/uuid"

type FindPaymentByIdQuery struct {
	ID         uuid.UUID
	IncludeRef bool `query:"include_ref"`
}
