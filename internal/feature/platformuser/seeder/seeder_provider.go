package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type PlatformUserSeederSet seeder.Seeder

func NewUserSeederSet(
	user *PlatformUserSeeder,
) []PlatformUserSeederSet {
	return []PlatformUserSeederSet{
		user,
	}
}

var ProviderSet = wire.NewSet(
	NewUserSeederSet,
	NewPlatformUserSeeder,
)
