package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type ServiceSeederSet seeder.Seeder

func NewServiceSeederSet(s *ServiceSeeder) []ServiceSeederSet {
	return []ServiceSeederSet{s}
}

var ProviderSet = wire.NewSet(
	NewServiceSeederSet,
	NewServiceSeeder,
)
