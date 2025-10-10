package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

// ServiceCategorySeederSet represents the seeder type for service categories.
type ServiceCategorySeederSet seeder.Seeder

// NewServiceCategorySeederSet groups all service category seeders.
func NewServiceCategorySeederSet(
	s *ServiceCategorySeeder,
) []ServiceCategorySeederSet {
	return []ServiceCategorySeederSet{s}
}

// ProviderSet registers all seeder dependencies for DI with Wire.
var ProviderSet = wire.NewSet(
	NewServiceCategorySeederSet,
	NewServiceCategorySeeder,
)
