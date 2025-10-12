package domain

import "github.com/google/uuid"

// Fixed IDs for seeded payment methods
var PaymentMethodIDs = struct {
	Cash     uuid.UUID
	Card     uuid.UUID
	Transfer uuid.UUID
}{
	Cash:     uuid.MustParse("11111111-1111-1111-1111-111111111111"),
	Card:     uuid.MustParse("22222222-2222-2222-2222-222222222222"),
	Transfer: uuid.MustParse("33333333-3333-3333-3333-333333333333"),
}
