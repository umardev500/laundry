package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type MachineSeederSet seeder.Seeder

func NewMachineSeederSet(
	m *MachineSeeder,
) []MachineSeederSet {
	return []MachineSeederSet{m}
}

var ProviderSet = wire.NewSet(
	NewMachineSeederSet,
	NewMachineSeeder,
)
