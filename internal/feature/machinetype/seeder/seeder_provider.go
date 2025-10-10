package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type MachineTypeSeederSet seeder.Seeder

func NewMachineTypeSeederSet(m *MachineTypeSeeder) []MachineTypeSeederSet {
	return []MachineTypeSeederSet{m}
}

var ProviderSet = wire.NewSet(
	NewMachineTypeSeederSet,
	NewMachineTypeSeeder,
)
