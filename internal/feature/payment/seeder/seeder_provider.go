package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type PaymentSeederSet seeder.Seeder

func NewPaymentSeederSet(
	payment *PaymentSeeder,
) []PaymentSeederSet {
	return []PaymentSeederSet{
		payment,
	}
}

var ProviderSet = wire.NewSet(
	NewPaymentSeeder,
	NewPaymentSeederSet,
)
