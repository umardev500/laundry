package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

// ServiceUnitSeederSet represents a typed seeder slice for dependency injection.
type ServiceUnitSeederSet seeder.Seeder

// NewServiceUnitSeederSet creates a slice of Seeder instances for service unit data.
func NewServiceUnitSeederSet(
	s *ServiceUnitSeeder,
) []ServiceUnitSeederSet {
	return []ServiceUnitSeederSet{s}
}

// ProviderSet wires the service unit seeders.
var ProviderSet = wire.NewSet(
	NewServiceUnitSeederSet,
	NewServiceUnitSeeder,
)
