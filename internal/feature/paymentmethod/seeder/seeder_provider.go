package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

// PaymentMethodSeederSet represents a generic seeder interface for payment methods
type PaymentMethodSeederSet seeder.Seeder

func NewPaymentMethodSeederSet(paymentMethod *PaymentMethodSeeder) []PaymentMethodSeederSet {
	return []PaymentMethodSeederSet{
		paymentMethod,
	}
}

var ProviderSet = wire.NewSet(
	NewPaymentMethodSeeder,
	NewPaymentMethodSeederSet,
)
