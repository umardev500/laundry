package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

// SubscriptionSeederSet defines a common interface for running all subscription seeders.
type SubscriptionSeederSet seeder.Seeder

// NewSubscriptionSeederSet groups all subscription-related seeders into a slice.
func NewSubscriptionSeederSet(s *SubscriptionSeeder) []SubscriptionSeederSet {
	return []SubscriptionSeederSet{s}
}

// ProviderSet exposes the SubscriptionSeeder dependencies to Wire.
var ProviderSet = wire.NewSet(
	NewSubscriptionSeederSet,
	NewSubscriptionSeeder,
)
