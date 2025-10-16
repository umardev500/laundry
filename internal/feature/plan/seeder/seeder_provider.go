package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type PlanSeederSet seeder.Seeder

func NewPlanSeederSet(s *PlanSeeder) []PlanSeederSet {
	return []PlanSeederSet{s}
}

var ProviderSet = wire.NewSet(
	NewPlanSeederSet,
	NewPlanSeeder,
)
