package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type OrderStatusHistorySeederSet seeder.Seeder

func NewOrderStatusHistorySeederSet(
	orderStatusHistory *OrderStatusHistorySeeder,
) []OrderStatusHistorySeederSet {
	return []OrderStatusHistorySeederSet{
		orderStatusHistory,
	}
}

var ProviderSet = wire.NewSet(
	NewOrderStatusHistorySeeder,
	NewOrderStatusHistorySeederSet,
)
