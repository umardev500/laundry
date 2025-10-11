package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type OrderSeederSet seeder.Seeder

func NewOrderSeederSet(
	order *OrderSeeder,
	orderItem *OrderItemSeeder,
) []OrderSeederSet {
	return []OrderSeederSet{
		order,
		orderItem,
	}
}

var ProviderSet = wire.NewSet(
	NewOrderSeeder,
	NewOrderItemSeeder,
	NewOrderSeederSet,
)
